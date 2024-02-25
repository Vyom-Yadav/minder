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

//import (
//	"context"
//	"database/sql"
//	"fmt"
//	"testing"
//	"time"
//
//	"github.com/golang/mock/gomock"
//	"github.com/google/uuid"
//	mockdb "github.com/stacklok/minder/database/mock"
//	reminderconfig "github.com/stacklok/minder/internal/config/reminder"
//	"github.com/stacklok/minder/internal/db"
//	cursorutil "github.com/stacklok/minder/internal/util/cursor"
//	"github.com/stretchr/testify/require"
//)
//
//func Test_getReposForReconciliation(t *testing.T) {
//	t.Parallel()
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	type want struct {
//		repos          []*db.Repository
//		repoListCursor map[projectProviderPair]cursorutil.RepoCursor
//	}
//
//	tests := []struct {
//		name          string
//		projects      map[*db.Project][]*db.Repository
//		fetchLimit    int
//		recurrenceCfg reminderconfig.RecurrenceConfig
//		buildStub     func(store *mockdb.MockStore, projects map[*db.Project][]*db.Repository)
//		want          want
//		err           string
//	}{
//		{
//			name: "get repos for reconciliation",
//			projects: map[*db.Project][]*db.Repository{
//				{
//					ID:   generateUUIDFromName(t, "project1"),
//					Name: "project1",
//				}: getRepoTillId(t, 3),
//			},
//			fetchLimit: 3,
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed: "1h",
//				BatchSize:  3,
//			},
//			buildStub: func(store *mockdb.MockStore, projects map[*db.Project][]*db.Repository) {
//				for project, repos := range projects {
//					returnedRepos := make([]db.Repository, len(repos))
//					for i := range repos {
//						returnedRepos[i] = *repos[i]
//					}
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 4, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//				store.EXPECT().ListOld(gomock.Any(), gomock.Any()).
//					Return(time.Now().Add(-2*time.Hour), nil).AnyTimes()
//			},
//			want: want{
//				repos: getRepoTillId(t, 3),
//				// Exhausted cursors are deleted to keep size in check
//				repoListCursor: map[projectProviderPair]cursorutil.RepoCursor{},
//			},
//			err: "",
//		},
//		{
//			name: "get repos for reconciliation with additional repos",
//			projects: map[*db.Project][]*db.Repository{
//				{
//					ID:   generateUUIDFromName(t, "project1"),
//					Name: "project1",
//				}: getRepoTillId(t, 2),
//			},
//			fetchLimit: 2,
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed: "1h",
//				BatchSize:  2,
//			},
//			buildStub: func(store *mockdb.MockStore, projects map[*db.Project][]*db.Repository) {
//				for project, repos := range projects {
//					returnedRepos := make([]db.Repository, len(repos))
//					for i := range repos {
//						returnedRepos[i] = *repos[i]
//					}
//
//					returnedRepos = append(returnedRepos, db.Repository{RepoID: 3})
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 3, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
//					Return(time.Now().Add(-2*time.Hour), nil).AnyTimes()
//			},
//			want: want{
//				repos: getRepoTillId(t, 2),
//				repoListCursor: map[projectProviderPair]cursorutil.RepoCursor{
//					{
//						projectId: generateUUIDFromName(t, "project1"),
//						provider:  "github",
//					}: {
//						ProjectId: generateUUIDFromName(t, "project1").String(),
//						Provider:  "github",
//						RepoId:    3,
//					},
//				},
//			},
//			err: "",
//		},
//		{
//			name: "error listing repositories",
//			projects: map[*db.Project][]*db.Repository{
//				{
//					ID:   generateUUIDFromName(t, "project1"),
//					Name: "project1",
//				}: {
//					{RepoID: 1},
//				},
//			},
//			fetchLimit: 3,
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed: "1h",
//				BatchSize:  3,
//			},
//			buildStub: func(store *mockdb.MockStore, projects map[*db.Project][]*db.Repository) {
//				store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), gomock.Any()).
//					Return(nil, fmt.Errorf("some error"))
//			},
//			want: want{},
//			err:  "error listing repositories: some error",
//		},
//		{
//			name: "get repos from multiple projects with additional repos",
//			projects: map[*db.Project][]*db.Repository{
//				{
//					ID:   generateUUIDFromName(t, "project1"),
//					Name: "project1",
//				}: {
//					{RepoID: 1},
//				},
//				{
//					ID:   generateUUIDFromName(t, "project2"),
//					Name: "project2",
//				}: {
//					{RepoID: 2},
//				},
//			},
//			fetchLimit: 1,
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed: "1h",
//				BatchSize:  2,
//			},
//			buildStub: func(store *mockdb.MockStore, projects map[*db.Project][]*db.Repository) {
//				for project, repos := range projects {
//					returnedRepos := make([]db.Repository, len(repos))
//					for i := range repos {
//						returnedRepos[i] = *repos[i]
//					}
//
//					// Add an additional repo for each project, it doesn't matter if RepoID is same
//					returnedRepos = append(returnedRepos, db.Repository{RepoID: 3})
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 2, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
//					Return(time.Now().Add(-2*time.Hour), nil).AnyTimes()
//			},
//			want: want{
//				repos: getRepoTillId(t, 2),
//				repoListCursor: map[projectProviderPair]cursorutil.RepoCursor{
//					{
//						projectId: generateUUIDFromName(t, "project1"),
//						provider:  "github",
//					}: {
//						ProjectId: generateUUIDFromName(t, "project1").String(),
//						Provider:  "github",
//						RepoId:    3,
//					},
//					{
//						projectId: generateUUIDFromName(t, "project2"),
//						provider:  "github",
//					}: {
//						ProjectId: generateUUIDFromName(t, "project2").String(),
//						Provider:  "github",
//						RepoId:    3,
//					},
//				},
//			},
//		},
//		{
//			name: "invalid min elapsed time",
//			projects: map[*db.Project][]*db.Repository{
//				{
//					ID:   generateUUIDFromName(t, "project1"),
//					Name: "project1",
//				}: {
//					{RepoID: 1},
//				},
//			},
//			fetchLimit: 1,
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed: "invalid",
//				BatchSize:  2,
//			},
//			buildStub: func(store *mockdb.MockStore, projects map[*db.Project][]*db.Repository) {
//				for project, repos := range projects {
//					returnedRepos := make([]db.Repository, len(repos))
//					for i := range repos {
//						returnedRepos[i] = *repos[i]
//					}
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 2, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
//					Return(time.Now().Add(-2*time.Hour), nil).AnyTimes()
//			},
//			want: want{},
//			err:  "error checking min elapsed time: time: invalid duration \"invalid\"",
//		},
//		{
//			name: "no repos found for project",
//			projects: map[*db.Project][]*db.Repository{
//				{
//					ID:   generateUUIDFromName(t, "project1"),
//					Name: "project1",
//				}: {},
//			},
//			fetchLimit: 1,
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed: "1h",
//				BatchSize:  2,
//			},
//			buildStub: func(store *mockdb.MockStore, projects map[*db.Project][]*db.Repository) {
//				store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), gomock.Any()).
//					Return(nil, sql.ErrNoRows)
//			},
//			want: want{},
//			err:  "",
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//
//		t.Run(test.name, func(t *testing.T) {
//			t.Parallel()
//
//			store := mockdb.NewMockStore(ctrl)
//			test.buildStub(store, test.projects)
//			cfg := &reminderconfig.Config{
//				RecurrenceConfig: test.recurrenceCfg,
//			}
//
//			r := NewReminder(store, cfg)
//
//			projectsSlice := make([]*db.Project, 0, len(test.projects))
//			for project := range test.projects {
//				projectsSlice = append(projectsSlice, project)
//			}
//
//			got, err := r.getReposForReconciliation(context.Background(), projectsSlice, test.fetchLimit)
//			if test.err != "" {
//				require.EqualError(t, err, test.err)
//				return
//			} else {
//				require.NoError(t, err)
//			}
//			require.ElementsMatch(t, got, test.want.repos)
//			require.Equal(t, len(test.want.repoListCursor), len(r.repoListCursor))
//			for k, v := range test.want.repoListCursor {
//				require.Equal(t, v.String(), r.repoListCursor[k])
//			}
//		})
//	}
//}
//
//func Test_getRepositoryBatch(t *testing.T) {
//	t.Parallel()
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	type want struct {
//		repos          []*db.Repository
//		repoListCursor map[projectProviderPair]cursorutil.RepoCursor
//		projectCursor  cursorutil.ProjectCursor
//	}
//
//	type projectAndRepos struct {
//		project *db.Project
//		repos   []*db.Repository
//	}
//
//	tests := []struct {
//		name                         string
//		projectsAndRepos             []projectAndRepos
//		additionalProjectsAndRepos   []projectAndRepos
//		additionalNonQueriedProjects []*db.Project
//		recurrenceCfg                reminderconfig.RecurrenceConfig
//		buildStub                    func(store *mockdb.MockStore, projectsAndRepos, additionalProjectsAndRepos []projectAndRepos, additionalNonQueriedProjects []*db.Project)
//		want                         want
//		err                          string
//	}{
//		{
//			name: "get repository batch no additional repos",
//			projectsAndRepos: []projectAndRepos{
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-1"),
//						Name:      "project-1",
//						CreatedAt: getCreatedAtFromNum(t, 1),
//					},
//					repos: getRepoTillId(t, 4),
//				},
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-2"),
//						Name:      "project-2",
//						CreatedAt: getCreatedAtFromNum(t, 2),
//					},
//					repos: getRepoTillId(t, 4),
//				},
//			},
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed:    "1h",
//				BatchSize:     6,
//				MaxPerProject: 3,
//			},
//			buildStub: func(store *mockdb.MockStore, projectsAndRepos, additionalProjectsAndRepos []projectAndRepos, additionalNonQueriedProjects []*db.Project) {
//				for _, projAndRepos := range projectsAndRepos {
//					returnedRepos := make([]db.Repository, len(projAndRepos.repos))
//					for i := range projAndRepos.repos {
//						returnedRepos[i] = *projAndRepos.repos[i]
//					}
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: projAndRepos.project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 4, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
//					Return(time.Now().Add(-2*time.Hour), nil).AnyTimes()
//			},
//			want: want{
//				repos: append(getRepoTillId(t, 3), getRepoTillId(t, 3)...),
//				repoListCursor: map[projectProviderPair]cursorutil.RepoCursor{
//					{
//						projectId: generateUUIDFromName(t, "project-1"),
//						provider:  "github",
//					}: {
//						ProjectId: generateUUIDFromName(t, "project-1").String(),
//						Provider:  "github",
//						RepoId:    4,
//					},
//					{
//						projectId: generateUUIDFromName(t, "project-2"),
//						provider:  "github",
//					}: {
//						ProjectId: generateUUIDFromName(t, "project-2").String(),
//						Provider:  "github",
//						RepoId:    4,
//					},
//				},
//			},
//		},
//		{
//			name: "get repository batch with additional repos with profile cursor update",
//			projectsAndRepos: []projectAndRepos{
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-1"),
//						Name:      "project-1",
//						CreatedAt: getCreatedAtFromNum(t, 1),
//					},
//					repos: getRepoTillId(t, 3),
//				},
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-2"),
//						Name:      "project-2",
//						CreatedAt: getCreatedAtFromNum(t, 2),
//					},
//					repos: getRepoTillId(t, 3),
//				},
//			},
//			additionalProjectsAndRepos: []projectAndRepos{
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-3"),
//						Name:      "project-3",
//						CreatedAt: getCreatedAtFromNum(t, 3),
//					},
//					repos: getRepoTillId(t, 3),
//				},
//			},
//			additionalNonQueriedProjects: []*db.Project{
//				{
//					ID:        generateUUIDFromName(t, "project-4"),
//					Name:      "project-4",
//					CreatedAt: getCreatedAtFromNum(t, 4),
//				},
//			},
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed:    "1h",
//				BatchSize:     6,
//				MaxPerProject: 2,
//			},
//			buildStub: func(store *mockdb.MockStore, projectsAndRepos, additionalProjectsAndRepos []projectAndRepos, additionalNonQueriedProjects []*db.Project) {
//				for _, projAndRepos := range projectsAndRepos {
//					returnedRepos := make([]db.Repository, len(projAndRepos.repos))
//					for i := range projAndRepos.repos {
//						returnedRepos[i] = *projAndRepos.repos[i]
//					}
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: projAndRepos.project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 3, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//
//				for _, projAndRepos := range additionalProjectsAndRepos {
//					returnedRepos := make([]db.Repository, len(projAndRepos.repos))
//					for i := range projAndRepos.repos {
//						returnedRepos[i] = *projAndRepos.repos[i]
//					}
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: projAndRepos.project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 3, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//
//				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
//					Return(time.Now().Add(-2*time.Hour), nil).AnyTimes()
//
//				projects := make([]db.Project, len(additionalProjectsAndRepos)+len(additionalNonQueriedProjects))
//				for i, projAndRepos := range additionalProjectsAndRepos {
//					projects[i] = *projAndRepos.project
//				}
//
//				for i, proj := range additionalNonQueriedProjects {
//					projects[i+len(additionalProjectsAndRepos)] = *proj
//				}
//
//				store.EXPECT().ListProjects(gomock.Any(), gomock.Any()).
//					Return(projects, nil)
//			},
//			want: want{
//				repos: append(getRepoTillId(t, 2), append(getRepoTillId(t, 2), getRepoTillId(t, 2)...)...),
//				repoListCursor: map[projectProviderPair]cursorutil.RepoCursor{
//					{
//						projectId: generateUUIDFromName(t, "project-1"),
//						provider:  "github",
//					}: {
//						ProjectId: generateUUIDFromName(t, "project-1").String(),
//						Provider:  "github",
//						RepoId:    3,
//					},
//					{
//						projectId: generateUUIDFromName(t, "project-2"),
//						provider:  "github",
//					}: {
//						ProjectId: generateUUIDFromName(t, "project-2").String(),
//						Provider:  "github",
//						RepoId:    3,
//					},
//					{
//						projectId: generateUUIDFromName(t, "project-3"),
//						provider:  "github",
//					}: {
//						ProjectId: generateUUIDFromName(t, "project-3").String(),
//						Provider:  "github",
//						RepoId:    3,
//					},
//				},
//				projectCursor: cursorutil.ProjectCursor{
//					CreatedAt: getCreatedAtFromNum(t, 4),
//					Id:        generateUUIDFromName(t, "project-4"),
//				},
//			},
//		},
//		{
//			name: "no additional projects found",
//			projectsAndRepos: []projectAndRepos{
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-1"),
//						Name:      "project-1",
//						CreatedAt: getCreatedAtFromNum(t, 1),
//					},
//					repos: getRepoTillId(t, 3),
//				},
//			},
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed:    "1h",
//				BatchSize:     6,
//				MaxPerProject: 3,
//			},
//			buildStub: func(store *mockdb.MockStore, projectsAndRepos, additionalProjectsAndRepos []projectAndRepos, additionalNonQueriedProjects []*db.Project) {
//				for _, projAndRepos := range projectsAndRepos {
//					returnedRepos := make([]db.Repository, len(projAndRepos.repos))
//					for i := range projAndRepos.repos {
//						returnedRepos[i] = *projAndRepos.repos[i]
//					}
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: projAndRepos.project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 4, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
//					Return(time.Now().Add(-2*time.Hour), nil).AnyTimes()
//
//				store.EXPECT().ListProjects(gomock.Any(), gomock.Any()).
//					Return(nil, sql.ErrNoRows)
//			},
//			want: want{
//				repos: getRepoTillId(t, 3),
//				// Exhausted cursors are deleted to keep size in check
//				repoListCursor: map[projectProviderPair]cursorutil.RepoCursor{},
//			},
//		},
//		{
//			name: "error listing additional projects",
//			projectsAndRepos: []projectAndRepos{
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-1"),
//						Name:      "project-1",
//						CreatedAt: getCreatedAtFromNum(t, 1),
//					},
//					repos: getRepoTillId(t, 3),
//				},
//			},
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed:    "1h",
//				BatchSize:     6,
//				MaxPerProject: 3,
//			},
//			buildStub: func(store *mockdb.MockStore, projectsAndRepos, additionalProjectsAndRepos []projectAndRepos, additionalNonQueriedProjects []*db.Project) {
//				for _, projAndRepos := range projectsAndRepos {
//					returnedRepos := make([]db.Repository, len(projAndRepos.repos))
//					for i := range projAndRepos.repos {
//						returnedRepos[i] = *projAndRepos.repos[i]
//					}
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: projAndRepos.project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 4, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
//					Return(time.Now().Add(-2*time.Hour), nil).AnyTimes()
//
//				store.EXPECT().ListProjects(gomock.Any(), gomock.Any()).
//					Return(nil, sql.ErrConnDone)
//			},
//			want: want{},
//			err:  fmt.Sprintf("error getting additional repos for reconciliation: error listing projects: %s", sql.ErrConnDone),
//		},
//		{
//			name: "error listing additional repositories",
//			projectsAndRepos: []projectAndRepos{
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-1"),
//						Name:      "project-1",
//						CreatedAt: getCreatedAtFromNum(t, 1),
//					},
//					repos: getRepoTillId(t, 3),
//				},
//			},
//			additionalProjectsAndRepos: []projectAndRepos{
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-2"),
//						Name:      "project-2",
//						CreatedAt: getCreatedAtFromNum(t, 2),
//					},
//					repos: getRepoTillId(t, 3),
//				},
//			},
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed:    "1h",
//				BatchSize:     6,
//				MaxPerProject: 3,
//			},
//			buildStub: func(store *mockdb.MockStore, projectsAndRepos, additionalProjectsAndRepos []projectAndRepos, additionalNonQueriedProjects []*db.Project) {
//				for _, projAndRepos := range projectsAndRepos {
//					returnedRepos := make([]db.Repository, len(projAndRepos.repos))
//					for i := range projAndRepos.repos {
//						returnedRepos[i] = *projAndRepos.repos[i]
//					}
//
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: projAndRepos.project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 4, Valid: true},
//					}).Return(returnedRepos, nil)
//				}
//
//				for _, additionalProjectAndRepos := range additionalProjectsAndRepos {
//					store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), db.ListRepositoriesByProjectIDParams{
//						Provider:  "github",
//						ProjectID: additionalProjectAndRepos.project.ID,
//						RepoID:    sql.NullInt32{Valid: false},
//						Limit:     sql.NullInt64{Int64: 4, Valid: true},
//					}).Return(nil, sql.ErrConnDone)
//				}
//
//				projects := make([]db.Project, len(additionalProjectsAndRepos))
//				for i, additionalProjectAndRepos := range additionalProjectsAndRepos {
//					projects[i] = *additionalProjectAndRepos.project
//				}
//
//				store.EXPECT().ListProjects(gomock.Any(), gomock.Any()).
//					Return(projects, nil)
//
//				store.EXPECT().GetOldestRuleEvaluationByRepositoryId(gomock.Any(), gomock.Any()).
//					Return(time.Now().Add(-2*time.Hour), nil).AnyTimes()
//			},
//			want: want{},
//			err:  fmt.Sprintf("error getting additional repos for reconciliation: error getting repos for reconciliation: error listing repositories: %s", sql.ErrConnDone),
//		},
//		{
//			name: "error getting repos for reconciliation",
//			projectsAndRepos: []projectAndRepos{
//				{
//					project: &db.Project{
//						ID:        generateUUIDFromName(t, "project-1"),
//						Name:      "project-1",
//						CreatedAt: getCreatedAtFromNum(t, 1),
//					},
//					repos: getRepoTillId(t, 3),
//				},
//			},
//			recurrenceCfg: reminderconfig.RecurrenceConfig{
//				MinElapsed:    "1h",
//				BatchSize:     6,
//				MaxPerProject: 3,
//			},
//			buildStub: func(store *mockdb.MockStore, projectsAndRepos, additionalProjectsAndRepos []projectAndRepos, additionalNonQueriedProjects []*db.Project) {
//				store.EXPECT().ListRepositoriesByProjectID(gomock.Any(), gomock.Any()).
//					Return(nil, sql.ErrTxDone)
//			},
//			want: want{},
//			err:  fmt.Sprintf("error getting repos for reconciliation: error listing repositories: %s", sql.ErrTxDone),
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//
//		t.Run(test.name, func(t *testing.T) {
//			t.Parallel()
//
//			store := mockdb.NewMockStore(ctrl)
//			test.buildStub(store, test.projectsAndRepos, test.additionalProjectsAndRepos, test.additionalNonQueriedProjects)
//			cfg := &reminderconfig.Config{
//				RecurrenceConfig: test.recurrenceCfg,
//			}
//
//			r := NewReminder(store, cfg)
//			projects := make([]*db.Project, len(test.projectsAndRepos))
//			for i, projAndRepos := range test.projectsAndRepos {
//				projects[i] = projAndRepos.project
//			}
//
//			got, err := r.getRepositoryBatch(context.Background(), projects)
//			if test.err != "" {
//				require.EqualError(t, err, test.err)
//				return
//			} else {
//				require.NoError(t, err)
//			}
//			require.ElementsMatch(t, got, test.want.repos)
//			require.Equal(t, len(test.want.repoListCursor), len(r.repoListCursor))
//			for k, v := range test.want.repoListCursor {
//				require.Equal(t, v.String(), r.repoListCursor[k])
//			}
//
//			cursor, err := cursorutil.NewProjectCursor(r.projectListCursor)
//			require.NoError(t, err)
//
//			// We do not cover null cursor, that is tested separately
//			require.NotNil(t, cursor)
//			require.Equal(t, test.want.projectCursor, *cursor)
//		})
//	}
//}
//
//func getCreatedAtFromNum(t *testing.T, numFromZeroToSixty int) time.Time {
//	t.Helper()
//
//	creationTime := time.Date(2023, time.February, 23, 10, 0, numFromZeroToSixty, 0, time.UTC)
//	return creationTime
//}
//
//func getRepoTillId(t *testing.T, tillId int32) []*db.Repository {
//	t.Helper()
//
//	repos := make([]*db.Repository, 0, tillId)
//	for i := int32(1); i <= tillId; i++ {
//		repos = append(repos, &db.Repository{RepoID: i})
//	}
//	return repos
//}
//
//func generateUUIDFromName(t *testing.T, name string) uuid.UUID {
//	t.Helper()
//	return uuid.NewSHA1(uuid.Nil, []byte(name))
//}
