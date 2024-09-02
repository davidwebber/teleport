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

package local

import (
	"context"

	"github.com/gravitational/trace"

	userintegrationtasksv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/userintegrationtasks/v1"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/backend"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/teleport/lib/services/local/generic"
)

type UserIntegrationTasksService struct {
	service *generic.ServiceWrapper[*userintegrationtasksv1.UserIntegrationTask]
}

const userIntegrationTasksKey = "user_integration_tasks"

// NewUserIntegrationTasksService creates a new UserIntegrationTasksService.
func NewUserIntegrationTasksService(backend backend.Backend) (*UserIntegrationTasksService, error) {
	service, err := generic.NewServiceWrapper(
		generic.ServiceWrapperConfig[*userintegrationtasksv1.UserIntegrationTask]{
			Backend:       backend,
			ResourceKind:  types.KindAccessGraphSecretAuthorizedKey,
			BackendPrefix: authorizedKeysPrefix,
			MarshalFunc:   services.MarshalProtoResource[*userintegrationtasksv1.UserIntegrationTask],
			UnmarshalFunc: services.UnmarshalProtoResource[*userintegrationtasksv1.UserIntegrationTask],
		})
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &UserIntegrationTasksService{service: service}, nil
}

func (s *UserIntegrationTasksService) ListUserIntegrationTasks(ctx context.Context, pagesize int64, lastKey string) ([]*userintegrationtasksv1.UserIntegrationTask, string, error) {
	r, nextToken, err := s.service.ListResources(ctx, int(pagesize), lastKey)
	return r, nextToken, trace.Wrap(err)
}

func (s *UserIntegrationTasksService) GetUserIntegrationTask(ctx context.Context, name string) (*userintegrationtasksv1.UserIntegrationTask, error) {
	r, err := s.service.GetResource(ctx, name)
	return r, trace.Wrap(err)
}

func (s *UserIntegrationTasksService) CreateUserIntegrationTask(ctx context.Context, userIntegrationTask *userintegrationtasksv1.UserIntegrationTask) (*userintegrationtasksv1.UserIntegrationTask, error) {
	r, err := s.service.CreateResource(ctx, userIntegrationTask)
	return r, trace.Wrap(err)
}

func (s *UserIntegrationTasksService) UpdateUserIntegrationTask(ctx context.Context, userIntegrationTask *userintegrationtasksv1.UserIntegrationTask) (*userintegrationtasksv1.UserIntegrationTask, error) {
	r, err := s.service.ConditionalUpdateResource(ctx, userIntegrationTask)
	return r, trace.Wrap(err)
}

func (s *UserIntegrationTasksService) UpsertUserIntegrationTask(ctx context.Context, userIntegrationTask *userintegrationtasksv1.UserIntegrationTask) (*userintegrationtasksv1.UserIntegrationTask, error) {
	r, err := s.service.UpsertResource(ctx, userIntegrationTask)
	return r, trace.Wrap(err)
}

func (s *UserIntegrationTasksService) DeleteUserIntegrationTask(ctx context.Context, name string) error {
	err := s.service.DeleteResource(ctx, name)
	return trace.Wrap(err)
}

func (s *UserIntegrationTasksService) DeleteAllUserIntegrationTasks(ctx context.Context) error {
	err := s.service.DeleteAllResources(ctx)
	return trace.Wrap(err)
}
