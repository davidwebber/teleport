/*
Copyright 2024 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

// Auto Discover EC2 issues identifiers which are used to provide better error messages to the user.
const (
	// AutoDiscoverEC2IssueEICEFailedToCreateNode is used when the EICE flow fails to auto enroll an EC2 instance
	// as an EICE node.
	AutoDiscoverEC2IssueEICEFailedToCreateNode = "ec2-eice-creation"
	// AutoDiscoverEC2IssueScriptSSMAgentNotRunning is used when the SSM Agent is not present in the instance.
	// This can also happen when the SSM was not able to connect to AWS Systems Manager.
	AutoDiscoverEC2IssueScriptSSMAgentNotRunning = "ec2-ssm-agent-not-running"
)
