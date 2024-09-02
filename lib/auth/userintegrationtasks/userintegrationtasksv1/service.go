/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package userintegrationtasksv1

import (
	"context"

	"github.com/gravitational/trace"
	"google.golang.org/protobuf/types/known/emptypb"

	userintegrationtasksv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/userintegrationtasks/v1"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/api/types/userintegrationtasks"
	"github.com/gravitational/teleport/lib/authz"
	"github.com/gravitational/teleport/lib/services"
)

// ServiceConfig holds configuration options for the UserIntegrationTask gRPC service.
type ServiceConfig struct {
	// Authorizer is the authorizer to use.
	Authorizer authz.Authorizer

	// Backend is the backend for storing UserIntegrationTask.
	Backend services.UserIntegrationTasks

	// Cache is the cache for storing UserIntegrationTask.
	Cache Reader
}

// CheckAndSetDefaults checks the ServiceConfig fields and returns an error if
// a required param is not provided.
// Authorizer, Cache and Backend are required params
func (s *ServiceConfig) CheckAndSetDefaults() error {
	if s.Authorizer == nil {
		return trace.BadParameter("authorizer is required")
	}
	if s.Backend == nil {
		return trace.BadParameter("backend is required")
	}
	if s.Cache == nil {
		return trace.BadParameter("cache is required")
	}

	return nil
}

// Reader contains the methods defined for cache access.
type Reader interface {
	ListUserIntegrationTasks(ctx context.Context, pageSize int64, nextToken string) ([]*userintegrationtasksv1.UserIntegrationTask, string, error)
	GetUserIntegrationTask(ctx context.Context, name string) (*userintegrationtasksv1.UserIntegrationTask, error)
}

// Service implements the teleport.UserIntegrationTask.v1.UserIntegrationTaskService RPC service.
type Service struct {
	userintegrationtasksv1.UnimplementedUserIntegrationTaskServiceServer

	authorizer authz.Authorizer
	backend    services.UserIntegrationTasks
	cache      Reader
}

// NewService returns a new UserIntegrationTask gRPC service.
func NewService(cfg ServiceConfig) (*Service, error) {
	if err := cfg.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	return &Service{
		authorizer: cfg.Authorizer,
		backend:    cfg.Backend,
		cache:      cfg.Cache,
	}, nil
}

// CreateUserIntegrationTask creates user integration task resource.
func (s *Service) CreateUserIntegrationTask(ctx context.Context, req *userintegrationtasksv1.CreateUserIntegrationTaskRequest) (*userintegrationtasksv1.UserIntegrationTask, error) {
	authCtx, err := s.authorizer.Authorize(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if err := authCtx.CheckAccessToKind(types.KindUserIntegrationTask, types.VerbCreate); err != nil {
		return nil, trace.Wrap(err)
	}

	if err := userintegrationtasks.ValidateUserIntegrationTask(req.UserIntegrationTask); err != nil {
		return nil, trace.Wrap(err)
	}

	rsp, err := s.backend.CreateUserIntegrationTask(ctx, req.UserIntegrationTask)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return rsp, nil
}

// ListUserIntegrationTasks returns a list of user integration tasks.
func (s *Service) ListUserIntegrationTasks(ctx context.Context, req *userintegrationtasksv1.ListUserIntegrationTasksRequest) (*userintegrationtasksv1.ListUserIntegrationTasksResponse, error) {
	authCtx, err := s.authorizer.Authorize(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if err := authCtx.CheckAccessToKind(types.KindUserIntegrationTask, types.VerbRead, types.VerbList); err != nil {
		return nil, trace.Wrap(err)
	}

	rsp, nextToken, err := s.cache.ListUserIntegrationTasks(ctx, req.PageSize, req.PageToken)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return &userintegrationtasksv1.ListUserIntegrationTasksResponse{
		UserIntegrationTasks: rsp,
		NextPageToken:        nextToken,
	}, nil
}

// GetUserIntegrationTask returns user integration task resource.
func (s *Service) GetUserIntegrationTask(ctx context.Context, req *userintegrationtasksv1.GetUserIntegrationTaskRequest) (*userintegrationtasksv1.UserIntegrationTask, error) {
	authCtx, err := s.authorizer.Authorize(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if err := authCtx.CheckAccessToKind(types.KindUserIntegrationTask, types.VerbRead); err != nil {
		return nil, trace.Wrap(err)
	}

	rsp, err := s.cache.GetUserIntegrationTask(ctx, req.GetName())
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return rsp, nil

}

// UpdateUserIntegrationTask updates user integration task resource.
func (s *Service) UpdateUserIntegrationTask(ctx context.Context, req *userintegrationtasksv1.UpdateUserIntegrationTaskRequest) (*userintegrationtasksv1.UserIntegrationTask, error) {
	authCtx, err := s.authorizer.Authorize(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if err := authCtx.CheckAccessToKind(types.KindUserIntegrationTask, types.VerbUpdate); err != nil {
		return nil, trace.Wrap(err)
	}

	if err := userintegrationtasks.ValidateUserIntegrationTask(req.UserIntegrationTask); err != nil {
		return nil, trace.Wrap(err)
	}

	rsp, err := s.backend.UpdateUserIntegrationTask(ctx, req.UserIntegrationTask)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return rsp, nil
}

// UpsertUserIntegrationTask upserts user integration task resource.
func (s *Service) UpsertUserIntegrationTask(ctx context.Context, req *userintegrationtasksv1.UpsertUserIntegrationTaskRequest) (*userintegrationtasksv1.UserIntegrationTask, error) {
	authCtx, err := s.authorizer.Authorize(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if err := authCtx.CheckAccessToKind(types.KindUserIntegrationTask, types.VerbUpdate, types.VerbCreate); err != nil {
		return nil, trace.Wrap(err)
	}

	if err := userintegrationtasks.ValidateUserIntegrationTask(req.UserIntegrationTask); err != nil {
		return nil, trace.Wrap(err)
	}

	rsp, err := s.backend.UpsertUserIntegrationTask(ctx, req.UserIntegrationTask)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return rsp, nil

}

// DeleteUserIntegrationTask deletes user integration task resource.
func (s *Service) DeleteUserIntegrationTask(ctx context.Context, req *userintegrationtasksv1.DeleteUserIntegrationTaskRequest) (*emptypb.Empty, error) {
	authCtx, err := s.authorizer.Authorize(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if err := authCtx.CheckAccessToKind(types.KindUserIntegrationTask, types.VerbDelete); err != nil {
		return nil, trace.Wrap(err)
	}

	if err := s.backend.DeleteUserIntegrationTask(ctx, req.GetName()); err != nil {
		return nil, trace.Wrap(err)
	}

	return &emptypb.Empty{}, nil
}
