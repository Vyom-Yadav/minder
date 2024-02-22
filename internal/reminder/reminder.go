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
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	reminderconfig "github.com/stacklok/minder/internal/config/reminder"
	"github.com/stacklok/minder/internal/db"
)

type Reminder struct {
	store db.Store
	cfg   *reminderconfig.Config

	projectListCursor string
	repoListCursor    map[projectProviderPair]string

	ticker *time.Ticker
}

type projectProviderPair struct {
	projectId uuid.UUID
	provider  string
}

// NewReminder creates a new Reminder instance
func NewReminder(store db.Store, config *reminderconfig.Config) *Reminder {
	return &Reminder{
		store:          store,
		cfg:            config,
		repoListCursor: make(map[projectProviderPair]string),
	}
}

func (r *Reminder) Start(ctx context.Context) error {
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
			log.Println("Shutting down reminder process...")
			return nil
		case <-r.ticker.C:
			if err := r.sendReminders(ctx); err != nil {
				// Non-fatal error, log and continue
				// TODO: Use zerolog
				log.Printf("Error sending reminders: %v", err)
			}
		}
	}
}

func (r *Reminder) Stop() {
	if r.ticker != nil {
		defer r.ticker.Stop()
	}
}

func (r *Reminder) sendReminders(ctx context.Context) error {
	listProjectsResp, err := r.listProjects(ctx, listProjectsRequest{
		cursor: r.projectListCursor,
		limit:  r.cfg.RecurrenceConfig.MinProjectFetchLimit,
	})
	if err != nil {
		return fmt.Errorf("error listing projects: %w", err)
	}

	// Update the cursor for the next iteration
	r.projectListCursor = listProjectsResp.cursor

	repos, err := r.getRepositoryBatch(ctx, listProjectsResp)
	if err != nil {
		return err
	}

	batch := make([]any, 0, r.cfg.RecurrenceConfig.BatchSize)
	batch = append(batch, repos...)

	// TODO: Send the batch for eval with r.cfg.RecurrenceConfig.MaxRetries
	// dial GRPC server and send batch
	return nil
}

func (r *Reminder) getRepositoryBatch(ctx context.Context, listProjectsResp *listProjectsResponse) ([]any, error) {
	repos, err := r.getReposForReconciliation(ctx, listProjectsResp.projects, r.cfg.RecurrenceConfig.MaxPerProject)
	if err != nil {
		return nil, fmt.Errorf("error getting repos for reconciliation: %w", err)
	}

	if len(repos) < r.cfg.RecurrenceConfig.BatchSize {
		remCap := r.cfg.RecurrenceConfig.BatchSize - len(repos)
		additionalRepos, err := r.getAdditionalReposForReconciliation(ctx, int32(remCap))
		if err != nil {
			return nil, fmt.Errorf("error getting additional repos for reconciliation: %w", err)
		}

		repos = append(repos, additionalRepos...)
	}
	return repos, nil
}

func (r *Reminder) getReposForReconciliation(ctx context.Context, projects []*db.Project, maxPerProject int32) ([]any, error) {
	repos := make([]any, 0, r.cfg.RecurrenceConfig.BatchSize)

	for _, project := range projects {
		listRepoResp, err := r.listRepositories(ctx, listRepoRequest{
			projectId: project.ID,
			provider:  string(db.ProviderTypeGithub),
			limit:     maxPerProject,
			// Use the cursor from the last iteration, if it is empty, either it is the first
			// iteration or the cursor is exhausted. In either case, we can fetch from the start.
			cursor: r.repoListCursor[projectProviderPair{
				projectId: project.ID,
				provider:  string(db.ProviderTypeGithub),
			}],
		})

		if err != nil {
			return nil, fmt.Errorf("error listing repositories: %w", err)
		}

		// Update the cursor for the next iteration
		r.repoListCursor[projectProviderPair{
			projectId: project.ID,
			provider:  string(db.ProviderTypeGithub),
		}] = listRepoResp.cursor

		for _, repo := range listRepoResp.results {
			repos = append(repos, repo)
		}
	}
	return repos, nil
}

func (r *Reminder) getAdditionalReposForReconciliation(ctx context.Context, additionalSpaces int32) ([]any, error) {
	additionalRepos := make([]any, 0, additionalSpaces)
	for additionalSpaces > 0 {
		listProjectsResp, err := r.listProjects(ctx, listProjectsRequest{
			cursor: r.projectListCursor,
			limit:  1,
		})
		if err != nil {
			return nil, fmt.Errorf("error listing projects: %w", err)
		}

		// Update the cursor for the next iteration
		r.projectListCursor = listProjectsResp.cursor

		repos, err := r.getReposForReconciliation(ctx, listProjectsResp.projects, additionalSpaces)
		if err != nil {
			return nil, fmt.Errorf("error getting repos for reconciliation: %w", err)
		}

		additionalSpaces -= int32(len(repos))
		additionalRepos = append(additionalRepos, repos...)
	}

	return additionalRepos, nil
}
