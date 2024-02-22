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
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/stacklok/minder/internal/db"
	cursorutil "github.com/stacklok/minder/internal/util/cursor"
)

type listProjectsResponse struct {
	projects []*db.Project
	cursor   string
}

type listProjectsRequest struct {
	cursor string
	limit  int32
}

func (r *Reminder) listProjects(ctx context.Context, req listProjectsRequest) (*listProjectsResponse, error) {
	cursor, err := cursorutil.NewProjectCursor(req.cursor)
	if err != nil {
		return nil, err
	}

	limit := sql.NullInt32{
		// Add 1 to the limit to check if there are more results
		Int32: req.limit + 1,
		Valid: req.limit > 0,
	}
	projects, err := r.store.ListProjects(ctx, db.ListProjectsParams{
		CreatedAt: sql.NullTime{
			Time:  cursor.CreatedAt,
			Valid: !cursor.CreatedAt.IsZero(),
		},
		ID: uuid.NullUUID{
			UUID:  cursor.Id,
			Valid: cursor.Id != uuid.Nil,
		},
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}

	var nextCursor cursorutil.ProjectCursor
	if limit.Valid && len(projects) == int(limit.Int32) {
		nextCursor = cursorutil.ProjectCursor{
			CreatedAt: projects[req.limit].CreatedAt,
			Id:        projects[req.limit].ID,
		}

		// remove the (req.limit + 1)th element from the results
		projects = projects[:req.limit]
	}

	results := make([]*db.Project, len(projects))
	for i, p := range projects {
		p := p
		results[i] = &p
	}

	return &listProjectsResponse{
		projects: results,
		cursor:   nextCursor.String(),
	}, nil
}

type listRepoRequest struct {
	projectId uuid.UUID
	provider  string
	cursor    string
	limit     int32
}

type listRepoResponse struct {
	results []*db.Repository
	cursor  string
}

func (r *Reminder) listRepositories(ctx context.Context, req listRepoRequest) (*listRepoResponse, error) {
	reqRepoCursor, err := cursorutil.NewRepoCursor(req.cursor)
	if err != nil {
		return nil, err
	}

	repoId := sql.NullInt32{
		Valid: reqRepoCursor.ProjectId == req.projectId.String() && reqRepoCursor.Provider == req.provider,
		Int32: reqRepoCursor.RepoId,
	}

	limit := sql.NullInt32{
		Valid: req.limit > 0,
		Int32: req.limit + 1,
	}

	repos, err := r.store.ListRepositoriesByProjectID(ctx, db.ListRepositoriesByProjectIDParams{
		Provider:  req.provider,
		ProjectID: req.projectId,
		RepoID:    repoId,
		Limit:     limit,
	})

	if err != nil {
		return nil, err
	}

	var results []*db.Repository

	for i, r := range repos {
		r := r
		results[i] = &r
	}

	var respRepoCursor *cursorutil.RepoCursor
	if limit.Valid && len(repos) == int(limit.Int32) {
		respRepoCursor = &cursorutil.RepoCursor{
			ProjectId: req.projectId.String(),
			Provider:  req.provider,
			RepoId:    repos[req.limit].RepoID,
		}

		// remove the (req.limit + 1)th element from the results
		results = results[:req.limit]
	}

	return &listRepoResponse{
		results: results,
		cursor:  respRepoCursor.String(),
	}, nil
}
