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

// [START phishingprotection_v1beta1_generated_PhishingProtectionServiceV1Beta1_ReportPhishing_sync]

package main

import (
	"context"

	phishingprotection "cloud.google.com/go/phishingprotection/apiv1beta1"
	phishingprotectionpb "google.golang.org/genproto/googleapis/cloud/phishingprotection/v1beta1"
)

func main() {
	ctx := context.Background()
	c, err := phishingprotection.NewPhishingProtectionServiceV1Beta1Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &phishingprotectionpb.ReportPhishingRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ReportPhishing(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END phishingprotection_v1beta1_generated_PhishingProtectionServiceV1Beta1_ReportPhishing_sync]
