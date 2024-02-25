// Copyright 2024 Stacklok, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package reminder sends reminders to the minder server to process entities in background.
package reminder

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	reminderconfig "github.com/stacklok/minder/internal/config/reminder"
	"github.com/stacklok/minder/internal/db"
	minderv1 "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
	pb "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/util/sets"
)

// Reminder sends reminders to the minder server to process entities in background.
type Reminder struct {
	cfg      *reminderconfig.Config
	stop     chan struct{}
	stopOnce sync.Once

	projectListCursor string
	repoListCursor    map[projectProviderPair]string

	ticker *time.Ticker
	logger zerolog.Logger

	// TODO: Create a client connection, check authn/z stuff

	// A new connection is established for each interval to prevent the connection from being
	// closed by the server due to inactivity. No 'KeepAlive' here.
	conn *grpc.ClientConn
}

type projectProviderPair struct {
	projectId uuid.UUID
	provider  string
}

// NewReminder creates a new Reminder instance
func NewReminder(config *reminderconfig.Config) *Reminder {
	return &Reminder{
		cfg:            config,
		stop:           make(chan struct{}),
		repoListCursor: make(map[projectProviderPair]string),
		logger:         zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

// Start starts the reminder by sending reminders at regular intervals
func (r *Reminder) Start(ctx context.Context) error {
	select {
	case <-r.stop:
		return errors.New("reminder stopped, cannot start again")
	default:
	}

	interval, err := time.ParseDuration(r.cfg.RecurrenceConfig.Interval)
	if err != nil {
		return err
	}

	if interval <= 0 {
		return fmt.Errorf("invalid interval: %s", r.cfg.RecurrenceConfig.Interval)
	}

	r.ticker = time.NewTicker(interval)
	defer r.ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			r.logger.Info().Msg("reminder stopped")
			return nil
		case <-r.stop:
			r.logger.Info().Msg("reminder stopped")
			return nil
		case <-r.ticker.C:
			// In-case sending reminders i.e. iterating over entities consumes more time than the
			// interval, the ticker will adjust the time interval or drop ticks to make up for
			// slow receivers.
			if errs := r.sendReminders(ctx); errs != nil {
				for _, err := range errs {
					r.logger.Error().Err(err).Msg("reconciliation request unsuccessful")
				}
			}
		}
	}
}

// Stop stops the reminder by stopping the ticker and closing the stop channel
func (r *Reminder) Stop() {
	if r.ticker != nil {
		defer r.ticker.Stop()
	}
	r.stopOnce.Do(func() {
		close(r.stop)
	})
}

func (r *Reminder) sendReminders(ctx context.Context) []error {
	projectClient := minderv1.NewProjectServiceClient(r.conn)
	listProjectsResp, err := projectClient.ListProjects(ctx, &pb.ListProjectsRequest{
		Limit:  int64(r.cfg.RecurrenceConfig.MinProjectFetchLimit),
		Cursor: r.projectListCursor,
	})
	if errors.Is(err, sql.ErrNoRows) {
		r.logger.Debug().Msgf("no projects found, current cursor: %s", r.projectListCursor)
		return nil
	} else if err != nil {
		return []error{fmt.Errorf("error listing projects: %w", err)}
	}

	// Update the cursor for the next iteration
	r.updateProjectListCursor(listProjectsResp.Cursor)

	projects, err := dbProjectsFromPB(listProjectsResp.Results)
	if err != nil {
		return []error{fmt.Errorf("error converting projects: %w", err)}
	}

	repos, err := r.getRepositoryBatch(ctx, projects)
	if err != nil {
		return []error{fmt.Errorf("error listing projects: %w", err)}
	}

	errorSlice := make([]error, 0)

	reconciliationServiceClient := minderv1.NewReconciliationServiceClient(r.conn)
	for _, repo := range repos {
		projID := repo.ProjectID.String()
		_, err := reconciliationServiceClient.CreateRepositoryReconciliationTask(ctx, &minderv1.CreateRepositoryReconciliationTaskRequest{
			Context: &minderv1.Context{
				Project:  &projID,
				Provider: &repo.Provider,
			},
			RepositoryId: int64(repo.RepoID),
		})
		if err != nil {
			errorSlice = append(errorSlice, err)
		}
	}
	return errorSlice
}

func (r *Reminder) getRepositoryBatch(ctx context.Context, projects []*db.Project) ([]*db.Repository, error) {
	repos, err := r.getReposForReconciliation(ctx, projects, r.cfg.RecurrenceConfig.MaxPerProject)
	if err != nil {
		return nil, fmt.Errorf("error getting repos for reconciliation: %w", err)
	}

	if len(repos) < r.cfg.RecurrenceConfig.BatchSize {
		remCap := r.cfg.RecurrenceConfig.BatchSize - len(repos)
		additionalRepos, err := r.getAdditionalReposForReconciliation(ctx, remCap)
		if err != nil {
			return nil, fmt.Errorf("error getting additional repos for reconciliation: %w", err)
		}

		repos = append(repos, additionalRepos...)
	}
	return repos, nil
}

func (r *Reminder) getReposForReconciliation(ctx context.Context, projects []*db.Project, fetchLimit int) ([]*db.Repository, error) {
	repos := make([]*db.Repository, 0)

	// Instead of querying which providers are registered for a project, we can simply test all
	// providers we support. Performance shouldn't be an issue as ListRepositories endpoint is
	// indexed using provider.
	githubProvider := string(db.ProviderTypeGithub)

	for _, project := range projects {
		cursorKey := projectProviderPair{
			projectId: project.ID,
			provider:  githubProvider,
		}

		client := minderv1.NewRepositoryServiceClient(r.conn)
		projectIdStr := project.ID.String()
		listRepoResp, err := client.ListRepositories(ctx, &pb.ListRepositoriesRequest{
			Limit: int64(fetchLimit),
			Context: &pb.Context{
				Provider: &githubProvider,
				Project:  &projectIdStr,
			},
			// Use the cursor from the last iteration. If the cursor is empty, it will fetch
			// the first page.
			Cursor: r.repoListCursor[cursorKey],
		})

		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Debug().Msgf("no repositories found for project: %s", project.ID)
			continue
		} else if err != nil {
			return nil, fmt.Errorf("error listing repositories: %w", err)
		}

		// Update the cursor for the next iteration
		r.repoListCursor[cursorKey] = listRepoResp.Cursor
		if listRepoResp.Cursor == "" {
			// Remove the cursor from the map if it's empty. This keeps the map size in check.
			// Default empty cursor is used to fetch the first page in the subsequent iterations.
			delete(r.repoListCursor, cursorKey)
		}

		resultRepos, err := dbRepositoriesFromPB(listRepoResp.Results)
		if err != nil {
			return nil, fmt.Errorf("error converting repositories: %w", err)
		}

		repos, err = r.getEligibleRepos(ctx, resultRepos)
		if err != nil {
			return nil, fmt.Errorf("error getting eligible repos: %w", err)
		}
	}
	return repos, nil
}

func (r *Reminder) getAdditionalReposForReconciliation(ctx context.Context, additionalSpaces int) ([]*db.Repository, error) {
	additionalRepos := make([]*db.Repository, 0, additionalSpaces)

	// TODO: Is this optimization necessary?
	// minProjectsToFetch := math.Ceil(float64(additionalSpaces) / float64(r.cfg.RecurrenceConfig.MaxPerProject))

	for additionalSpaces > 0 {
		client := minderv1.NewProjectServiceClient(r.conn)
		// Fetch the next project (one-by-one) to get the repos
		listProjectsResp, err := client.ListProjects(ctx, &pb.ListProjectsRequest{
			Cursor: r.projectListCursor,
			Limit:  1,
		})
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Debug().Msgf("no more projects found, current cursor: %s", r.projectListCursor)
			break
		} else if err != nil {
			return nil, fmt.Errorf("error listing projects: %w", err)
		}

		fetchLimit := additionalSpaces
		if additionalSpaces >= r.cfg.RecurrenceConfig.MaxPerProject {
			fetchLimit = r.cfg.RecurrenceConfig.MaxPerProject

			// Update the cursor for the next iteration. If additionalSpaces < MaxPerProject, then
			// we will fetch the same project again in the next iteration. This is done to prevent
			// evaluating only small number of entities from a project in an iteration.
			r.updateProjectListCursor(listProjectsResp.Cursor)
		}

		resProjects, err := dbProjectsFromPB(listProjectsResp.Results)
		repos, err := r.getReposForReconciliation(ctx, resProjects, fetchLimit)
		if err != nil {
			return nil, fmt.Errorf("error getting repos for reconciliation: %w", err)
		}

		additionalSpaces -= len(repos)
		additionalRepos = append(additionalRepos, repos...)
	}

	return additionalRepos, nil
}

func (r *Reminder) updateProjectListCursor(newCursor string) {
	r.projectListCursor = newCursor
}

func (r *Reminder) getEligibleRepos(ctx context.Context, repos []*db.Repository) ([]*db.Repository, error) {
	repoIds := make([]string, len(repos))
	for i, repo := range repos {
		repoIds[i] = repo.ID.String()
	}

	eligibleRepoIds, err := r.getEligibleRepoIds(ctx, repoIds)
	if err != nil {
		return nil, err
	}

	eligibleRepos := make([]*db.Repository, 0, len(repos))
	for _, repo := range repos {
		if eligibleRepoIds.Has(repo.ID) {
			eligibleRepos = append(eligibleRepos, repo)
		}
	}

	return eligibleRepos, nil
}

func (r *Reminder) getEligibleRepoIds(ctx context.Context, repoId []string) (sets.Set[uuid.UUID], error) {
	minElapsed, err := time.ParseDuration(r.cfg.RecurrenceConfig.MinElapsed)
	if err != nil {
		return nil, err
	}

	client := minderv1.NewRepositoryServiceClient(r.conn)
	oldestRuleEvaluations, err := client.ListOldestRuleEvaluationByIds(ctx, &pb.ListOldestRuleEvaluationByIdsRequest{
		RepositoryIds: repoId,
	})
	if err != nil {
		return nil, err
	}

	eligibleRepoIds := sets.New[uuid.UUID]()

	for _, ruleEvaluation := range oldestRuleEvaluations.RepositoryRuleEvaluations {
		lastUpdated := ruleEvaluation.OldestRuleEvaluation.AsTime()

		if lastUpdated.Add(minElapsed).Before(time.Now()) {
			parsedUUID, err := uuid.Parse(ruleEvaluation.RepositoryId)
			if err != nil {
				return nil, err
			}
			eligibleRepoIds.Insert(parsedUUID)
		}
	}
	return eligibleRepoIds, nil
}
