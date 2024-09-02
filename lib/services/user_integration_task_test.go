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
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	headerv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	userintegrationtasksv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/userintegrationtasks/v1"
	"github.com/gravitational/teleport/lib/utils"
)

func TestMarshalUserIntegrationTaskRoundTrip(t *testing.T) {
	t.Parallel()

	obj := &userintegrationtasksv1.UserIntegrationTask{
		Version: "v1",
		Kind:    "user_integration_task",
		Metadata: &headerv1.Metadata{
			Name: "example-user-integration-task",
			Labels: map[string]string{
				"env": "example",
			},
		},
		Spec: &userintegrationtasksv1.UserIntegrationTaskSpec{
			Integration: "my-integration",
			TaskType:    "discover-ec2",
			IssueType:   "SSM_AGENT_MISSING",
			DiscoverEc2: &userintegrationtasksv1.DiscoverEC2{Instances: map[string]*userintegrationtasksv1.DiscoverEC2Instance{
				"i-1234567890": {
					State:           "OPEN",
					Name:            "instance-name",
					Region:          "us-east-1",
					InvocationUrl:   "https://example.com/",
					DiscoveryConfig: "config",
					DiscoveryGroup:  "group",
					SyncTime:        timestamppb.Now(),
				},
			}},
		},
	}

	out, err := MarshalUserIntegrationTask(obj)
	require.NoError(t, err)
	newObj, err := UnmarshalUserIntegrationTask(out)
	require.NoError(t, err)
	require.True(t, proto.Equal(obj, newObj), "messages are not equal")
}

func TestUnmarshalUserIntegrationTask(t *testing.T) {
	t.Parallel()

	syncTime := timestamppb.Now()
	syncTimeString := syncTime.AsTime().Format(time.RFC3339Nano)

	correctUserIntegrationTaskYAML := fmt.Sprintf(`
version: v1
kind: user_integration_task
metadata:
  name: example-user-integration-task
  labels:
    env: example
spec:
  integration: my-integration
  task_type: discover-ec2
  issue_type: SSM_AGENT_MISSING
  discover_ec2:
    instances:
      i-1234567890:
        state: OPEN
        name: instance-name
        region: us-east-1
        invocation_url: https://example.com/
        discovery_config: config
        discovery_group: group
        sync_time: "%s"
`, syncTimeString)

	data, err := utils.ToJSON([]byte(correctUserIntegrationTaskYAML))
	require.NoError(t, err)

	expected := &userintegrationtasksv1.UserIntegrationTask{
		Version: "v1",
		Kind:    "user_integration_task",
		Metadata: &headerv1.Metadata{
			Name: "example-user-integration-task",
			Labels: map[string]string{
				"env": "example",
			},
		},
		Spec: &userintegrationtasksv1.UserIntegrationTaskSpec{
			Integration: "my-integration",
			TaskType:    "discover-ec2",
			IssueType:   "SSM_AGENT_MISSING",
			DiscoverEc2: &userintegrationtasksv1.DiscoverEC2{Instances: map[string]*userintegrationtasksv1.DiscoverEC2Instance{
				"i-1234567890": {
					State:           "OPEN",
					Name:            "instance-name",
					Region:          "us-east-1",
					InvocationUrl:   "https://example.com/",
					DiscoveryConfig: "config",
					DiscoveryGroup:  "group",
					SyncTime:        syncTime,
				},
			}},
		},
	}

	obj, err := UnmarshalUserIntegrationTask(data)
	require.NoError(t, err)
	require.True(t, proto.Equal(expected, obj), "UserIntegrationTask objects are not equal")
}
