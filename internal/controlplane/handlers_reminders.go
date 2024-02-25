// Copyright 2024 Stacklok, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controlplane

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/stacklok/minder/internal/db"
	"github.com/stacklok/minder/internal/logger"
	"github.com/stacklok/minder/internal/reconcilers"
	pb "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRepositoryReconciliationTask(ctx context.Context,
	in *pb.CreateRepositoryReconciliationTaskRequest) (
	*pb.CreateRepositoryReconciliationTaskResponse, error,
) {

	projectIdString := in.GetContext().Project
	if projectIdString == nil {
		return nil, status.Errorf(codes.InvalidArgument, "projectId is required")
	}

	projectId, err := uuid.Parse(*projectIdString)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "projectId is invalid")
	}

	provider, err := getProviderFromRequestOrDefault(ctx, s.store, in, projectId)
	if err != nil {
		return nil, providerError(err)
	}

	// Check if the repo exists
	repo, err := s.store.GetRepositoryByIDAndProject(ctx, db.GetRepositoryByIDAndProjectParams{
		Provider:  provider.Name,
		RepoID:    int32(in.GetRepositoryId()),
		ProjectID: projectId,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "repository not found: %v", err)
	}

	msg, err := reconcilers.NewRepoReconcilerMessage(repo.Provider, repo.RepoID, repo.ProjectID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating reconciler message: %v", err)
	}

	// This is a non-fatal error, so we'll just log it and continue with the next ones
	if err := s.evt.Publish(reconcilers.InternalReconcilerEventTopic, msg); err != nil {
		log.Printf("error publishing reconciler event: %v", err)
	}

	// Telemetry logging
	logger.BusinessRecord(ctx).Provider = repo.Provider
	logger.BusinessRecord(ctx).Project = repo.ProjectID
	logger.BusinessRecord(ctx).Repository = repo.ID

	return &pb.CreateRepositoryReconciliationTaskResponse{
		TaskId: msg.UUID,
	}, nil
}
