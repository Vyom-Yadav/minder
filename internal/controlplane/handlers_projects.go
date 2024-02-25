// Copyright 2023 Stacklok, Inc
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

package controlplane

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/stacklok/minder/internal/db"
	"github.com/stacklok/minder/internal/util"
	cursorutil "github.com/stacklok/minder/internal/util/cursor"
	minderv1 "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
	pb "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// maxFetchLimit is the maximum number of projects that can be fetched from the database in one call
const maxProjectFetchLimit = 50

func (s *Server) ListProjects(ctx context.Context, in *pb.ListProjectsRequest) (*pb.ListProjectResponse, error) {
	cursor, err := cursorutil.NewProjectCursor(in.Cursor)
	if err != nil {
		return nil, err
	}

	if in.Limit > maxProjectFetchLimit {
		return nil, util.UserVisibleError(codes.InvalidArgument, "limit too high, max is %d", maxProjectFetchLimit)
	}

	limit := sql.NullInt64{
		// Add 1 to the limit to check if there are more results
		Int64: in.Limit + 1,
		Valid: in.Limit > 0,
	}
	projects, err := s.store.ListProjects(ctx, db.ListProjectsParams{
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
	if limit.Valid && int64(len(projects)) == limit.Int64 {
		nextCursor = cursorutil.ProjectCursor{
			CreatedAt: projects[in.Limit].CreatedAt,
			Id:        projects[in.Limit].ID,
		}

		// remove the (req.limit + 1)th element from the results
		projects = projects[:in.Limit]
	}

	results := make([]*minderv1.Project, len(projects))
	for i := range projects {
		results[i] = &minderv1.Project{
			ProjectId: projects[i].ID.String(),
			Name:      projects[i].Name,
			CreatedAt: timestamppb.New(projects[i].CreatedAt),
			UpdatedAt: timestamppb.New(projects[i].UpdatedAt),
		}
	}

	return &pb.ListProjectResponse{
		Results: results,
		Cursor:  nextCursor.String(),
	}, nil
}
