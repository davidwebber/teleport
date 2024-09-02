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

package userintegrationtasks_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	headerv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	userintegrationtasksv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/userintegrationtasks/v1"
	"github.com/gravitational/teleport/api/types/userintegrationtasks"
)

func TestValidateUserIntegrationTask(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		task    *userintegrationtasksv1.UserIntegrationTask
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "NilUserIntegrationTask",
			task:    nil,
			wantErr: require.Error,
		},
		{
			name: "ValidUserIntegrationTask",
			task: &userintegrationtasksv1.UserIntegrationTask{
				Kind:    "user_integration_task",
				Version: "v1",
				Metadata: &headerv1.Metadata{
					Name: "test",
				},
				Spec: &userintegrationtasksv1.UserIntegrationTaskSpec{
					Integration: "my-integration",
					TaskType:    "discover-ec2",
					IssueType:   "failed to enroll ec2 instances",
					DiscoverEc2: &userintegrationtasksv1.DiscoverEC2{},
				},
			},
			wantErr: require.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userintegrationtasks.ValidateUserIntegrationTask(tt.task)
			tt.wantErr(t, err)
		})
	}
}
