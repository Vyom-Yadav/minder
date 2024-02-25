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
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	mockdb "github.com/stacklok/minder/database/mock"
	"github.com/stacklok/minder/internal/db"
	cursorutil "github.com/stacklok/minder/internal/util/cursor"
)

func Test_listProjects(t *testing.T) {
	t.Parallel()

	type want struct {
		projects []db.Project
		cursorId uuid.UUID
	}

	tests := []struct {
		name      string
		req       listProjectsRequest
		want      want
		buildStub func(store *mockdb.MockStore, req listProjectsRequest, proj []db.Project)
		err       string
	}{
		{
			name: "extra project as cursor",
			req: listProjectsRequest{
				cursor: "",
				limit:  3,
			},
			want: want{
				projects: []db.Project{
					{
						Name: "project1",
						ID:   generateUUIDFromName(t, "project1"),
					},
					{
						Name: "project2",
						ID:   generateUUIDFromName(t, "project2"),
					},
					{
						Name: "project3",
						ID:   generateUUIDFromName(t, "project3"),
					},
				},
				cursorId: generateUUIDFromName(t, "project4"),
			},
			buildStub: func(store *mockdb.MockStore, req listProjectsRequest, proj []db.Project) {
				extraResp := []db.Project{
					{
						Name: "project4",
						ID:   generateUUIDFromName(t, "project4"),
					},
				}

				returnedResp := append(proj, extraResp...)

				store.EXPECT().ListProjects(gomock.Any(), db.ListProjectsParams{
					CreatedAt: sql.NullTime{
						Time:  time.Time{},
						Valid: false,
					},
					ID: uuid.NullUUID{
						UUID:  uuid.Nil,
						Valid: false,
					},
					Limit: sql.NullInt64{
						Int64: int64(req.limit + 1),
						Valid: true,
					},
				}).Return(returnedResp, nil)
			},
			err: "",
		},
		{
			name: "error with wrong cursor",
			req: listProjectsRequest{
				cursor: "wrong-cursor",
				limit:  3,
			},
			want: want{},
			buildStub: func(store *mockdb.MockStore, req listProjectsRequest, proj []db.Project) {
			},
			err: "error decoding cursor: illegal base64 data at input byte 5",
		},
		{
			name: "error with store",
			req: listProjectsRequest{
				cursor: "",
				limit:  3,
			},
			want: want{},
			buildStub: func(store *mockdb.MockStore, req listProjectsRequest, proj []db.Project) {
				store.EXPECT().ListProjects(gomock.Any(), gomock.Any()).
					Return(nil, sql.ErrConnDone)
			},
			err: sql.ErrConnDone.Error(),
		},
		{
			name: "no extra project as cursor",
			req: listProjectsRequest{
				cursor: "",
				limit:  1,
			},
			want: want{
				projects: []db.Project{
					{
						Name: "project1",
						ID:   generateUUIDFromName(t, "project1"),
					},
				},
				cursorId: uuid.Nil,
			},
			buildStub: func(store *mockdb.MockStore, req listProjectsRequest, proj []db.Project) {
				store.EXPECT().ListProjects(gomock.Any(), gomock.Any()).
					Return(proj, nil)
			},
			err: "",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			test.buildStub(store, test.req, test.want.projects)

			r := &Reminder{
				store: store,
			}

			ctx := context.Background()
			got, err := r.listProjects(ctx, test.req)
			if test.err != "" {
				require.EqualError(t, err, test.err)
				return
			}
			require.NoError(t, err)

			wanted := make([]*db.Project, len(test.want.projects))
			for i := range test.want.projects {
				wanted[i] = &test.want.projects[i]
			}

			require.ElementsMatch(t, wanted, got.projects)

			cursor, err := cursorutil.NewProjectCursor(got.cursor)
			require.NoError(t, err)

			require.Equal(t, test.want.cursorId, cursor.Id)
		})
	}
}

func Test_listRepositories(t *testing.T) {
	t.Parallel()

	type want struct {
		repositories []db.Repository
		cursor       cursorutil.RepoCursor
	}

	tests := []struct {
		name      string
		req       listRepoRequest
		want      want
		buildStub func(store *mockdb.MockStore, req listRepoRequest, repos []db.Repository)
		err       string
	}{
		{
			name: "extra repository as cursor",
			req: listRepoRequest{
				projectId: generateUUIDFromName(t, "project1"),
				provider:  "github",
				cursor:    "",
				limit:     3,
			},
			want: want{
				repositories: []db.Repository{
					{RepoID: 1},
					{RepoID: 2},
					{RepoID: 3},
				},
				cursor: cursorutil.RepoCursor{
					ProjectId: generateUUIDFromName(t, "project1").String(),
					Provider:  "github",
					RepoId:    4,
				},
			},
			buildStub: func(store *mockdb.MockStore, req listRepoRequest, repos []db.Repository) {
				extraResp := []db.Repository{
					{RepoID: 4},
				}

				returnedResp := append(repos, extraResp...)

				store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
					Provider:  req.provider,
					ProjectID: req.projectId,
					RepoID:    sql.NullInt32{Valid: false},
					Limit:     sql.NullInt64{Int64: int64(req.limit + 1), Valid: true},
				}).Return(returnedResp, nil)
			},
			err: "",
		},
		{
			name: "error with wrong cursor",
			req: listRepoRequest{
				projectId: uuid.New(),
				provider:  "github",
				cursor:    "wrong-cursor",
				limit:     3,
			},
			want: want{},
			buildStub: func(store *mockdb.MockStore, req listRepoRequest, repos []db.Repository) {
			},
			err: "error decoding cursor: illegal base64 data at input byte 5",
		},
		{
			name: "error with store",
			req: listRepoRequest{
				projectId: uuid.New(),
				provider:  "github",
				cursor:    "",
				limit:     3,
			},
			want: want{},
			buildStub: func(store *mockdb.MockStore, req listRepoRequest, repos []db.Repository) {
				store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
					Provider:  req.provider,
					ProjectID: req.projectId,
					RepoID:    sql.NullInt32{Valid: false},
					Limit:     sql.NullInt64{Int64: int64(req.limit + 1), Valid: true},
				}).Return(nil, sql.ErrConnDone)
			},
			err: sql.ErrConnDone.Error(),
		},
		{
			name: "no extra repository as cursor",
			req: listRepoRequest{
				projectId: uuid.New(),
				provider:  "github",
				cursor:    "",
				limit:     1,
			},
			want: want{
				repositories: []db.Repository{
					{RepoID: 1},
				},
				cursor: cursorutil.RepoCursor{},
			},
			buildStub: func(store *mockdb.MockStore, req listRepoRequest, repos []db.Repository) {
				store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
					Provider:  req.provider,
					ProjectID: req.projectId,
					RepoID:    sql.NullInt32{Valid: false},
					Limit:     sql.NullInt64{Int64: int64(req.limit + 1), Valid: true},
				}).Return(repos, nil)
			},
			err: "",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			test.buildStub(store, test.req, test.want.repositories)

			r := &Reminder{
				store: store,
			}

			ctx := context.Background()
			got, err := r.listRepositories(ctx, test.req)
			if test.err != "" {
				require.EqualError(t, err, test.err)
				return
			}
			require.NoError(t, err)

			wanted := make([]*db.Repository, len(test.want.repositories))
			for i := range test.want.repositories {
				wanted[i] = &test.want.repositories[i]
			}

			require.ElementsMatch(t, wanted, got.results)

			cursor, err := cursorutil.NewRepoCursor(got.cursor)
			require.NoError(t, err)

			require.Equal(t, test.want.cursor, *cursor)
		})
	}
}

func Test_getOldestRepoRuleEvaluation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		want      time.Time
		buildStub func(store *mockdb.MockStore, want time.Time)
		err       string
	}{
		{
			name: "success",
			want: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			buildStub: func(store *mockdb.MockStore, want time.Time) {
				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
					Return(want, nil)
			},
			err: "",
		},
		{
			name: "error with store",
			want: time.Time{},
			buildStub: func(store *mockdb.MockStore, want time.Time) {
				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
					Return(nil, sql.ErrConnDone)
			},
			err: sql.ErrConnDone.Error(),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			test.buildStub(store, test.want)

			r := &Reminder{
				store: store,
			}

			ctx := context.Background()
			got, err := r.getOldestRepoRuleEvaluation(ctx, uuid.New())
			if test.err != "" {
				require.EqualError(t, err, test.err)
				return
			}
			require.NoError(t, err)

			require.Equal(t, test.want, got)
		})
	}
}
