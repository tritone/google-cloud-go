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

// [START dataproc_v1_generated_AutoscalingPolicyService_DeleteAutoscalingPolicy_sync]

package main

import (
	"context"

	dataproc "cloud.google.com/go/dataproc/apiv1"
	dataprocpb "google.golang.org/genproto/googleapis/cloud/dataproc/v1"
)

func main() {
	ctx := context.Background()
	c, err := dataproc.NewAutoscalingPolicyClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &dataprocpb.DeleteAutoscalingPolicyRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteAutoscalingPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END dataproc_v1_generated_AutoscalingPolicyService_DeleteAutoscalingPolicy_sync]
