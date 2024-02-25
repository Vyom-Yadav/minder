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

package reminder

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/stacklok/minder/internal/db"
	minderv1 "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
)

func dbProjectFromPB(project *minderv1.Project) (*db.Project, error) {
	projectUUID, err := uuid.Parse(project.ProjectId)
	if err != nil {
		return nil, err
	}
	return &db.Project{
		ID:        projectUUID,
		Name:      project.Name,
		CreatedAt: project.CreatedAt.AsTime(),
		UpdatedAt: project.UpdatedAt.AsTime(),
	}, nil
}

func dbProjectsFromPB(projects []*minderv1.Project) ([]*db.Project, error) {
	dbProjects := make([]*db.Project, 0, len(projects))
	for _, project := range projects {
		dbProject, err := dbProjectFromPB(project)
		if err != nil {
			return nil, err
		}
		dbProjects = append(dbProjects, dbProject)
	}
	return dbProjects, nil
}

func dbRepositoryFromPB(repo *minderv1.Repository) (*db.Repository, error) {
	if repo.Id == nil {
		return nil, errors.New("repository id is required")
	}

	repoUUID, err := uuid.Parse(*repo.Id)
	if err != nil {
		return nil, err
	}

	if repo.Context == nil {
		return nil, errors.New("repository context is required")
	}

	if repo.Context.Project == nil {
		return nil, errors.New("repository context project is required")
	}

	if repo.Context.Provider == nil {
		return nil, errors.New("repository provider is required")
	}

	projectUUID, err := uuid.Parse(*repo.Context.Project)
	if err != nil {
		return nil, err
	}

	return &db.Repository{
		ID:        repoUUID,
		Provider:  *repo.Context.Provider,
		ProjectID: projectUUID,
		RepoOwner: repo.Owner,
		RepoName:  repo.Name,
		RepoID:    repo.RepoId,
		IsPrivate: repo.IsPrivate,
		IsFork:    repo.IsFork,
		WebhookID: sql.NullInt32{
			Valid: repo.HookId != 0,
			// TODO: Create an issue for this inconsistency
			Int32: int32(repo.HookId),
		},
		WebhookUrl: repo.HookUrl,
		DeployUrl:  repo.DeployUrl,
		CloneUrl:   repo.CloneUrl,
		CreatedAt:  repo.CreatedAt.AsTime(),
		UpdatedAt:  repo.UpdatedAt.AsTime(),
		DefaultBranch: sql.NullString{
			String: repo.DefaultBranch,
			Valid:  repo.DefaultBranch != "",
		},
	}, nil
}

func dbRepositoriesFromPB(repos []*minderv1.Repository) ([]*db.Repository, error) {
	dbRepos := make([]*db.Repository, 0, len(repos))
	for _, repo := range repos {
		dbRepo, err := dbRepositoryFromPB(repo)
		if err != nil {
			return nil, err
		}
		dbRepos = append(dbRepos, dbRepo)
	}
	return dbRepos, nil
}
