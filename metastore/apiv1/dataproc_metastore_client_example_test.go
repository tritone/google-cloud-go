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

package metastore_test

import (
	"context"

	metastore "cloud.google.com/go/metastore/apiv1"
	"google.golang.org/api/iterator"
	metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1"
)

func ExampleNewDataprocMetastoreClient() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleDataprocMetastoreClient_ListServices() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.ListServicesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListServices(ctx, req)
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

func ExampleDataprocMetastoreClient_GetService() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.GetServiceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataprocMetastoreClient_CreateService() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.CreateServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateService(ctx, req)
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

func ExampleDataprocMetastoreClient_UpdateService() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.UpdateServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateService(ctx, req)
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

func ExampleDataprocMetastoreClient_DeleteService() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.DeleteServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDataprocMetastoreClient_ListMetadataImports() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.ListMetadataImportsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListMetadataImports(ctx, req)
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

func ExampleDataprocMetastoreClient_GetMetadataImport() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.GetMetadataImportRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetMetadataImport(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataprocMetastoreClient_CreateMetadataImport() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.CreateMetadataImportRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateMetadataImport(ctx, req)
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

func ExampleDataprocMetastoreClient_UpdateMetadataImport() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.UpdateMetadataImportRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateMetadataImport(ctx, req)
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

func ExampleDataprocMetastoreClient_ExportMetadata() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.ExportMetadataRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ExportMetadata(ctx, req)
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

func ExampleDataprocMetastoreClient_RestoreService() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.RestoreServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RestoreService(ctx, req)
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

func ExampleDataprocMetastoreClient_ListBackups() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.ListBackupsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListBackups(ctx, req)
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

func ExampleDataprocMetastoreClient_GetBackup() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.GetBackupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetBackup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataprocMetastoreClient_CreateBackup() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.CreateBackupRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateBackup(ctx, req)
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

func ExampleDataprocMetastoreClient_DeleteBackup() {
	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metastorepb.DeleteBackupRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteBackup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
