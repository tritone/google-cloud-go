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

package privateca_test

import (
	"context"

	privateca "cloud.google.com/go/security/privateca/apiv1"
	"google.golang.org/api/iterator"
	privatecapb "google.golang.org/genproto/googleapis/cloud/security/privateca/v1"
)

func ExampleNewCertificateAuthorityClient() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleCertificateAuthorityClient_CreateCertificate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.CreateCertificateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#CreateCertificateRequest.
	}
	resp, err := c.CreateCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetCertificate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.GetCertificateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#GetCertificateRequest.
	}
	resp, err := c.GetCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificates() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ListCertificatesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#ListCertificatesRequest.
	}
	it := c.ListCertificates(ctx, req)
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

func ExampleCertificateAuthorityClient_RevokeCertificate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.RevokeCertificateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#RevokeCertificateRequest.
	}
	resp, err := c.RevokeCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_UpdateCertificate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.UpdateCertificateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#UpdateCertificateRequest.
	}
	resp, err := c.UpdateCertificate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ActivateCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ActivateCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#ActivateCertificateAuthorityRequest.
	}
	op, err := c.ActivateCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_CreateCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.CreateCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#CreateCertificateAuthorityRequest.
	}
	op, err := c.CreateCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_DisableCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.DisableCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#DisableCertificateAuthorityRequest.
	}
	op, err := c.DisableCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_EnableCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.EnableCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#EnableCertificateAuthorityRequest.
	}
	op, err := c.EnableCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_FetchCertificateAuthorityCsr() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.FetchCertificateAuthorityCsrRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#FetchCertificateAuthorityCsrRequest.
	}
	resp, err := c.FetchCertificateAuthorityCsr(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.GetCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#GetCertificateAuthorityRequest.
	}
	resp, err := c.GetCertificateAuthority(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificateAuthorities() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ListCertificateAuthoritiesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#ListCertificateAuthoritiesRequest.
	}
	it := c.ListCertificateAuthorities(ctx, req)
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

func ExampleCertificateAuthorityClient_UndeleteCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.UndeleteCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#UndeleteCertificateAuthorityRequest.
	}
	op, err := c.UndeleteCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_DeleteCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.DeleteCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#DeleteCertificateAuthorityRequest.
	}
	op, err := c.DeleteCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_UpdateCertificateAuthority() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.UpdateCertificateAuthorityRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#UpdateCertificateAuthorityRequest.
	}
	op, err := c.UpdateCertificateAuthority(ctx, req)
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

func ExampleCertificateAuthorityClient_CreateCaPool() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.CreateCaPoolRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#CreateCaPoolRequest.
	}
	op, err := c.CreateCaPool(ctx, req)
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

func ExampleCertificateAuthorityClient_UpdateCaPool() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.UpdateCaPoolRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#UpdateCaPoolRequest.
	}
	op, err := c.UpdateCaPool(ctx, req)
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

func ExampleCertificateAuthorityClient_GetCaPool() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.GetCaPoolRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#GetCaPoolRequest.
	}
	resp, err := c.GetCaPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCaPools() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ListCaPoolsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#ListCaPoolsRequest.
	}
	it := c.ListCaPools(ctx, req)
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

func ExampleCertificateAuthorityClient_DeleteCaPool() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.DeleteCaPoolRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#DeleteCaPoolRequest.
	}
	op, err := c.DeleteCaPool(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCertificateAuthorityClient_FetchCaCerts() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.FetchCaCertsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#FetchCaCertsRequest.
	}
	resp, err := c.FetchCaCerts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_GetCertificateRevocationList() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.GetCertificateRevocationListRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#GetCertificateRevocationListRequest.
	}
	resp, err := c.GetCertificateRevocationList(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificateRevocationLists() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ListCertificateRevocationListsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#ListCertificateRevocationListsRequest.
	}
	it := c.ListCertificateRevocationLists(ctx, req)
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

func ExampleCertificateAuthorityClient_UpdateCertificateRevocationList() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.UpdateCertificateRevocationListRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#UpdateCertificateRevocationListRequest.
	}
	op, err := c.UpdateCertificateRevocationList(ctx, req)
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

func ExampleCertificateAuthorityClient_CreateCertificateTemplate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.CreateCertificateTemplateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#CreateCertificateTemplateRequest.
	}
	op, err := c.CreateCertificateTemplate(ctx, req)
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

func ExampleCertificateAuthorityClient_DeleteCertificateTemplate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.DeleteCertificateTemplateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#DeleteCertificateTemplateRequest.
	}
	op, err := c.DeleteCertificateTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCertificateAuthorityClient_GetCertificateTemplate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.GetCertificateTemplateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#GetCertificateTemplateRequest.
	}
	resp, err := c.GetCertificateTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCertificateAuthorityClient_ListCertificateTemplates() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.ListCertificateTemplatesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#ListCertificateTemplatesRequest.
	}
	it := c.ListCertificateTemplates(ctx, req)
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

func ExampleCertificateAuthorityClient_UpdateCertificateTemplate() {
	ctx := context.Background()
	c, err := privateca.NewCertificateAuthorityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &privatecapb.UpdateCertificateTemplateRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#UpdateCertificateTemplateRequest.
	}
	op, err := c.UpdateCertificateTemplate(ctx, req)
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
