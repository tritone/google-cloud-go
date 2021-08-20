// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package memcache_test

import (
	"context"

	memcache "cloud.google.com/go/memcache/apiv1beta2"
	"google.golang.org/api/iterator"
	memcachepb "google.golang.org/genproto/googleapis/cloud/memcache/v1beta2"
)

func ExampleNewCloudMemcacheClient() {
	ctx := context.Background()
	c, err := memcache.NewCloudMemcacheClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleCloudMemcacheClient_ListInstances() {
	ctx := context.Background()
	c, err := memcache.NewCloudMemcacheClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &memcachepb.ListInstancesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/memcache/v1beta2#ListInstancesRequest.
	}
	it := c.ListInstances(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudMemcacheClient_GetInstance() {
	ctx := context.Background()
	c, err := memcache.NewCloudMemcacheClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &memcachepb.GetInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/memcache/v1beta2#GetInstanceRequest.
	}
	resp, err := c.GetInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudMemcacheClient_CreateInstance() {
	ctx := context.Background()
	c, err := memcache.NewCloudMemcacheClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &memcachepb.CreateInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/memcache/v1beta2#CreateInstanceRequest.
	}
	op, err := c.CreateInstance(ctx, req)
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

func ExampleCloudMemcacheClient_UpdateInstance() {
	ctx := context.Background()
	c, err := memcache.NewCloudMemcacheClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &memcachepb.UpdateInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/memcache/v1beta2#UpdateInstanceRequest.
	}
	op, err := c.UpdateInstance(ctx, req)
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

func ExampleCloudMemcacheClient_UpdateParameters() {
	ctx := context.Background()
	c, err := memcache.NewCloudMemcacheClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &memcachepb.UpdateParametersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/memcache/v1beta2#UpdateParametersRequest.
	}
	op, err := c.UpdateParameters(ctx, req)
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

func ExampleCloudMemcacheClient_DeleteInstance() {
	ctx := context.Background()
	c, err := memcache.NewCloudMemcacheClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &memcachepb.DeleteInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/memcache/v1beta2#DeleteInstanceRequest.
	}
	op, err := c.DeleteInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCloudMemcacheClient_ApplyParameters() {
	ctx := context.Background()
	c, err := memcache.NewCloudMemcacheClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &memcachepb.ApplyParametersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/memcache/v1beta2#ApplyParametersRequest.
	}
	op, err := c.ApplyParameters(ctx, req)
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

func ExampleCloudMemcacheClient_ApplySoftwareUpdate() {
	ctx := context.Background()
	c, err := memcache.NewCloudMemcacheClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &memcachepb.ApplySoftwareUpdateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/memcache/v1beta2#ApplySoftwareUpdateRequest.
	}
	op, err := c.ApplySoftwareUpdate(ctx, req)
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
