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

package services

import (
	"context"

	userintegrationtasksv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/userintegrationtasks/v1"
)

// UserIntegrationTasks is the interface for managing user integration tasks resources.
type UserIntegrationTasks interface {
	// CreateUserIntegrationTask creates a new user integration tasks resource.
	CreateUserIntegrationTask(context.Context, *userintegrationtasksv1.UserIntegrationTask) (*userintegrationtasksv1.UserIntegrationTask, error)
	// UpsertUserIntegrationTask creates or updates the user integration tasks resource.
	UpsertUserIntegrationTask(context.Context, *userintegrationtasksv1.UserIntegrationTask) (*userintegrationtasksv1.UserIntegrationTask, error)
	// GetUserIntegrationTask returns the user integration tasks resource by name.
	GetUserIntegrationTask(ctx context.Context, name string) (*userintegrationtasksv1.UserIntegrationTask, error)
	// ListUserIntegrationTasks returns the user integration tasks resources.
	ListUserIntegrationTasks(ctx context.Context, pageSize int64, nextToken string) ([]*userintegrationtasksv1.UserIntegrationTask, string, error)
	// UpdateUserIntegrationTask updates the user integration tasks resource.
	UpdateUserIntegrationTask(context.Context, *userintegrationtasksv1.UserIntegrationTask) (*userintegrationtasksv1.UserIntegrationTask, error)
	// DeleteUserIntegrationTask deletes the user integration tasks resource by name.
	DeleteUserIntegrationTask(context.Context, string) error
	// DeleteAllUserIntegrationTasks deletes all User Integration Tasks.
	DeleteAllUserIntegrationTasks(context.Context) error
}

// MarshalUserIntegrationTask marshals the UserIntegrationTask object into a JSON byte array.
func MarshalUserIntegrationTask(object *userintegrationtasksv1.UserIntegrationTask, opts ...MarshalOption) ([]byte, error) {
	return MarshalProtoResource(object, opts...)
}

// UnmarshalUserIntegrationTask unmarshals the UserIntegrationTask object from a JSON byte array.
func UnmarshalUserIntegrationTask(data []byte, opts ...MarshalOption) (*userintegrationtasksv1.UserIntegrationTask, error) {
	return UnmarshalProtoResource[*userintegrationtasksv1.UserIntegrationTask](data, opts...)
}
