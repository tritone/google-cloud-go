// Copyright 2021 Google LLC
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

// Code generated by cloud.google.com/go/internal/gapicgen/gensnippets. DO NOT EDIT.

// [START workflows_v1beta_generated_Workflows_CreateWorkflow_sync]

package main

import (
	"context"

	workflows "cloud.google.com/go/workflows/apiv1beta"
	workflowspb "google.golang.org/genproto/googleapis/cloud/workflows/v1beta"
)

func main() {
	ctx := context.Background()
	c, err := workflows.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &workflowspb.CreateWorkflowRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/workflows/v1beta#CreateWorkflowRequest.
	}
	op, err := c.CreateWorkflow(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END workflows_v1beta_generated_Workflows_CreateWorkflow_sync]
