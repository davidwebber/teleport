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

package userintegrationtasks

import (
	"github.com/gravitational/trace"

	headerv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	userintegrationtasksv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/userintegrationtasks/v1"
	"github.com/gravitational/teleport/api/types"
)

// NewUserIntegrationTask creates a new UserIntegrationTask object.
// It validates the object before returning it.
func NewUserIntegrationTask(name string, spec *userintegrationtasksv1.UserIntegrationTaskSpec) (*userintegrationtasksv1.UserIntegrationTask, error) {
	cj := &userintegrationtasksv1.UserIntegrationTask{
		Kind:    types.KindUserIntegrationTask,
		Version: types.V1,
		Metadata: &headerv1.Metadata{
			Name: name,
		},
		Spec: spec,
	}

	if err := ValidateUserIntegrationTask(cj); err != nil {
		return nil, trace.Wrap(err)
	}

	return cj, nil
}

const (
	// TaskTypeDiscoverEC2 identifies a User Integration Tasks that is created
	// when an auto-enrollment of an EC2 instance fails.
	// UserIntegrationTasks that have this Task Type must include the DiscoverEC2 field.
	TaskTypeDiscoverEC2 = "discover-ec2"
)

// ValidateUserIntegrationTask validates the UserIntegrationTask object without modifying it.
func ValidateUserIntegrationTask(uit *userintegrationtasksv1.UserIntegrationTask) error {
	switch {
	case uit.GetKind() != types.KindUserIntegrationTask:
		return trace.BadParameter("invalid kind")
	case uit.GetVersion() != types.V1:
		return trace.BadParameter("invalid version")
	case uit.GetSubKind() != "":
		return trace.BadParameter("invalid sub kind, must be empty")
	case uit.GetMetadata() == nil:
		return trace.BadParameter("user integration task metadata is nil")
	case uit.Metadata.GetName() == "":
		return trace.BadParameter("user integration task name is empty")
	case uit.GetSpec() == nil:
		return trace.BadParameter("user integration task spec is nil")
	case uit.GetSpec().Integration == "":
		return trace.BadParameter("integration is required")
	}

	switch uit.Spec.TaskType {
	case TaskTypeDiscoverEC2:
		if err := validateDiscoverEC2TaskType(uit); err != nil {
			return trace.Wrap(err)
		}
	default:
		return trace.BadParameter("task type %q is not valid", uit.Spec.TaskType)
	}

	return nil
}

func validateDiscoverEC2TaskType(uit *userintegrationtasksv1.UserIntegrationTask) error {
	if uit.Spec.DiscoverEc2 == nil {
		return trace.BadParameter("%s requires the discover_ec2 field", TaskTypeDiscoverEC2)
	}
	if uit.Spec.IssueType == "" {
		return trace.BadParameter("issue type is required")
	}

	return nil
}
