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

package channel_test

import (
	"context"

	channel "cloud.google.com/go/channel/apiv1"
	"google.golang.org/api/iterator"
	channelpb "google.golang.org/genproto/googleapis/cloud/channel/v1"
)

func ExampleNewCloudChannelClient() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleCloudChannelClient_ListCustomers() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListCustomersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListCustomersRequest.
	}
	it := c.ListCustomers(ctx, req)
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

func ExampleCloudChannelClient_GetCustomer() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.GetCustomerRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#GetCustomerRequest.
	}
	resp, err := c.GetCustomer(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_CheckCloudIdentityAccountsExist() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.CheckCloudIdentityAccountsExistRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#CheckCloudIdentityAccountsExistRequest.
	}
	resp, err := c.CheckCloudIdentityAccountsExist(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_CreateCustomer() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.CreateCustomerRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#CreateCustomerRequest.
	}
	resp, err := c.CreateCustomer(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_UpdateCustomer() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.UpdateCustomerRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#UpdateCustomerRequest.
	}
	resp, err := c.UpdateCustomer(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_DeleteCustomer() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.DeleteCustomerRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#DeleteCustomerRequest.
	}
	err = c.DeleteCustomer(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCloudChannelClient_ProvisionCloudIdentity() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ProvisionCloudIdentityRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ProvisionCloudIdentityRequest.
	}
	op, err := c.ProvisionCloudIdentity(ctx, req)
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

func ExampleCloudChannelClient_ListEntitlements() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListEntitlementsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListEntitlementsRequest.
	}
	it := c.ListEntitlements(ctx, req)
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

func ExampleCloudChannelClient_ListTransferableSkus() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListTransferableSkusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListTransferableSkusRequest.
	}
	it := c.ListTransferableSkus(ctx, req)
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

func ExampleCloudChannelClient_ListTransferableOffers() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListTransferableOffersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListTransferableOffersRequest.
	}
	it := c.ListTransferableOffers(ctx, req)
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

func ExampleCloudChannelClient_GetEntitlement() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.GetEntitlementRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#GetEntitlementRequest.
	}
	resp, err := c.GetEntitlement(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_CreateEntitlement() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.CreateEntitlementRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#CreateEntitlementRequest.
	}
	op, err := c.CreateEntitlement(ctx, req)
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

func ExampleCloudChannelClient_ChangeParameters() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ChangeParametersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ChangeParametersRequest.
	}
	op, err := c.ChangeParameters(ctx, req)
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

func ExampleCloudChannelClient_ChangeRenewalSettings() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ChangeRenewalSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ChangeRenewalSettingsRequest.
	}
	op, err := c.ChangeRenewalSettings(ctx, req)
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

func ExampleCloudChannelClient_ChangeOffer() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ChangeOfferRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ChangeOfferRequest.
	}
	op, err := c.ChangeOffer(ctx, req)
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

func ExampleCloudChannelClient_StartPaidService() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.StartPaidServiceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#StartPaidServiceRequest.
	}
	op, err := c.StartPaidService(ctx, req)
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

func ExampleCloudChannelClient_SuspendEntitlement() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.SuspendEntitlementRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#SuspendEntitlementRequest.
	}
	op, err := c.SuspendEntitlement(ctx, req)
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

func ExampleCloudChannelClient_CancelEntitlement() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.CancelEntitlementRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#CancelEntitlementRequest.
	}
	op, err := c.CancelEntitlement(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCloudChannelClient_ActivateEntitlement() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ActivateEntitlementRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ActivateEntitlementRequest.
	}
	op, err := c.ActivateEntitlement(ctx, req)
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

func ExampleCloudChannelClient_TransferEntitlements() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.TransferEntitlementsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#TransferEntitlementsRequest.
	}
	op, err := c.TransferEntitlements(ctx, req)
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

func ExampleCloudChannelClient_TransferEntitlementsToGoogle() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.TransferEntitlementsToGoogleRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#TransferEntitlementsToGoogleRequest.
	}
	op, err := c.TransferEntitlementsToGoogle(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCloudChannelClient_ListChannelPartnerLinks() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListChannelPartnerLinksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListChannelPartnerLinksRequest.
	}
	it := c.ListChannelPartnerLinks(ctx, req)
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

func ExampleCloudChannelClient_GetChannelPartnerLink() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.GetChannelPartnerLinkRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#GetChannelPartnerLinkRequest.
	}
	resp, err := c.GetChannelPartnerLink(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_CreateChannelPartnerLink() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.CreateChannelPartnerLinkRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#CreateChannelPartnerLinkRequest.
	}
	resp, err := c.CreateChannelPartnerLink(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_UpdateChannelPartnerLink() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.UpdateChannelPartnerLinkRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#UpdateChannelPartnerLinkRequest.
	}
	resp, err := c.UpdateChannelPartnerLink(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_LookupOffer() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.LookupOfferRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#LookupOfferRequest.
	}
	resp, err := c.LookupOffer(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_ListProducts() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListProductsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListProductsRequest.
	}
	it := c.ListProducts(ctx, req)
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

func ExampleCloudChannelClient_ListSkus() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListSkusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListSkusRequest.
	}
	it := c.ListSkus(ctx, req)
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

func ExampleCloudChannelClient_ListOffers() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListOffersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListOffersRequest.
	}
	it := c.ListOffers(ctx, req)
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

func ExampleCloudChannelClient_ListPurchasableSkus() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListPurchasableSkusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListPurchasableSkusRequest.
	}
	it := c.ListPurchasableSkus(ctx, req)
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

func ExampleCloudChannelClient_ListPurchasableOffers() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListPurchasableOffersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListPurchasableOffersRequest.
	}
	it := c.ListPurchasableOffers(ctx, req)
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

func ExampleCloudChannelClient_RegisterSubscriber() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.RegisterSubscriberRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#RegisterSubscriberRequest.
	}
	resp, err := c.RegisterSubscriber(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_UnregisterSubscriber() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.UnregisterSubscriberRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#UnregisterSubscriberRequest.
	}
	resp, err := c.UnregisterSubscriber(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudChannelClient_ListSubscribers() {
	ctx := context.Background()
	c, err := channel.NewCloudChannelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &channelpb.ListSubscribersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/channel/v1#ListSubscribersRequest.
	}
	it := c.ListSubscribers(ctx, req)
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
