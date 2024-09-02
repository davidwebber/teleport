// Copyright 2024 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package userintegrationtask

import (
	"context"

	"github.com/gravitational/trace"

	userintegrationtaskv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/userintegrationtasks/v1"
)

// Client is a client for the User Integration Task API.
type Client struct {
	grpcClient userintegrationtaskv1.UserIntegrationTaskServiceClient
}

// NewClient creates a new User Integration Task client.
func NewClient(grpcClient userintegrationtaskv1.UserIntegrationTaskServiceClient) *Client {
	return &Client{
		grpcClient: grpcClient,
	}
}

// ListUserIntegrationTasks returns a list of User Integration Tasks.
func (c *Client) ListUserIntegrationTasks(ctx context.Context, pageSize int64, nextToken string) ([]*userintegrationtaskv1.UserIntegrationTask, string, error) {
	resp, err := c.grpcClient.ListUserIntegrationTasks(ctx, &userintegrationtaskv1.ListUserIntegrationTasksRequest{
		PageSize:  pageSize,
		PageToken: nextToken,
	})
	if err != nil {
		return nil, "", trace.Wrap(err)
	}

	return resp.UserIntegrationTasks, resp.NextPageToken, nil
}

// CreateUserIntegrationTask creates a new User Integration Task.
func (c *Client) CreateUserIntegrationTask(ctx context.Context, req *userintegrationtaskv1.UserIntegrationTask) (*userintegrationtaskv1.UserIntegrationTask, error) {
	rsp, err := c.grpcClient.CreateUserIntegrationTask(ctx, &userintegrationtaskv1.CreateUserIntegrationTaskRequest{
		UserIntegrationTask: req,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return rsp, nil
}

// GetUserIntegrationTask returns a User Integration Task by name.
func (c *Client) GetUserIntegrationTask(ctx context.Context, name string) (*userintegrationtaskv1.UserIntegrationTask, error) {
	rsp, err := c.grpcClient.GetUserIntegrationTask(ctx, &userintegrationtaskv1.GetUserIntegrationTaskRequest{
		Name: name,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return rsp, nil
}

// UpdateUserIntegrationTask updates an existing User Integration Task.
func (c *Client) UpdateUserIntegrationTask(ctx context.Context, req *userintegrationtaskv1.UserIntegrationTask) (*userintegrationtaskv1.UserIntegrationTask, error) {
	rsp, err := c.grpcClient.UpdateUserIntegrationTask(ctx, &userintegrationtaskv1.UpdateUserIntegrationTaskRequest{
		UserIntegrationTask: req,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return rsp, nil
}

// UpsertUserIntegrationTask upserts a User Integration Task.
func (c *Client) UpsertUserIntegrationTask(ctx context.Context, req *userintegrationtaskv1.UserIntegrationTask) (*userintegrationtaskv1.UserIntegrationTask, error) {
	rsp, err := c.grpcClient.UpsertUserIntegrationTask(ctx, &userintegrationtaskv1.UpsertUserIntegrationTaskRequest{
		UserIntegrationTask: req,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return rsp, nil
}

// DeleteUserIntegrationTask deletes a User Integration Task.
func (c *Client) DeleteUserIntegrationTask(ctx context.Context, name string) error {
	_, err := c.grpcClient.DeleteUserIntegrationTask(ctx, &userintegrationtaskv1.DeleteUserIntegrationTaskRequest{
		Name: name,
	})
	return trace.Wrap(err)
}

// DeleteAllUserIntegrationTasks deletes all User Integration Tasks.
// Not implemented. Added to satisfy the interface.
func (c *Client) DeleteAllUserIntegrationTasks(_ context.Context) error {
	return trace.NotImplemented("DeleteAllUserIntegrationTasks is not implemented")
}
