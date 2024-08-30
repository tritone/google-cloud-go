// Copyright 2024 Google LLC
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

package accounts

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"

	accountspb "cloud.google.com/go/shopping/merchant/accounts/apiv1beta/accountspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newBusinessIdentityClientHook clientHook

// BusinessIdentityCallOptions contains the retry settings for each method of BusinessIdentityClient.
type BusinessIdentityCallOptions struct {
	GetBusinessIdentity    []gax.CallOption
	UpdateBusinessIdentity []gax.CallOption
}

func defaultBusinessIdentityGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("merchantapi.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("merchantapi.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("merchantapi.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBusinessIdentityCallOptions() *BusinessIdentityCallOptions {
	return &BusinessIdentityCallOptions{
		GetBusinessIdentity:    []gax.CallOption{},
		UpdateBusinessIdentity: []gax.CallOption{},
	}
}

func defaultBusinessIdentityRESTCallOptions() *BusinessIdentityCallOptions {
	return &BusinessIdentityCallOptions{
		GetBusinessIdentity:    []gax.CallOption{},
		UpdateBusinessIdentity: []gax.CallOption{},
	}
}

// internalBusinessIdentityClient is an interface that defines the methods available from Merchant API.
type internalBusinessIdentityClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetBusinessIdentity(context.Context, *accountspb.GetBusinessIdentityRequest, ...gax.CallOption) (*accountspb.BusinessIdentity, error)
	UpdateBusinessIdentity(context.Context, *accountspb.UpdateBusinessIdentityRequest, ...gax.CallOption) (*accountspb.BusinessIdentity, error)
}

// BusinessIdentityClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to support business
// identity (at https://support.google.com/merchants/answer/12564247) API.
type BusinessIdentityClient struct {
	// The internal transport-dependent client.
	internalClient internalBusinessIdentityClient

	// The call options for this service.
	CallOptions *BusinessIdentityCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BusinessIdentityClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BusinessIdentityClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BusinessIdentityClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetBusinessIdentity retrieves the business identity of an account.
func (c *BusinessIdentityClient) GetBusinessIdentity(ctx context.Context, req *accountspb.GetBusinessIdentityRequest, opts ...gax.CallOption) (*accountspb.BusinessIdentity, error) {
	return c.internalClient.GetBusinessIdentity(ctx, req, opts...)
}

// UpdateBusinessIdentity updates the business identity of an account. Executing this method requires
// admin access.
func (c *BusinessIdentityClient) UpdateBusinessIdentity(ctx context.Context, req *accountspb.UpdateBusinessIdentityRequest, opts ...gax.CallOption) (*accountspb.BusinessIdentity, error) {
	return c.internalClient.UpdateBusinessIdentity(ctx, req, opts...)
}

// businessIdentityGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type businessIdentityGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing BusinessIdentityClient
	CallOptions **BusinessIdentityCallOptions

	// The gRPC API client.
	businessIdentityClient accountspb.BusinessIdentityServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewBusinessIdentityClient creates a new business identity service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to support business
// identity (at https://support.google.com/merchants/answer/12564247) API.
func NewBusinessIdentityClient(ctx context.Context, opts ...option.ClientOption) (*BusinessIdentityClient, error) {
	clientOpts := defaultBusinessIdentityGRPCClientOptions()
	if newBusinessIdentityClientHook != nil {
		hookOpts, err := newBusinessIdentityClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BusinessIdentityClient{CallOptions: defaultBusinessIdentityCallOptions()}

	c := &businessIdentityGRPCClient{
		connPool:               connPool,
		businessIdentityClient: accountspb.NewBusinessIdentityServiceClient(connPool),
		CallOptions:            &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *businessIdentityGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *businessIdentityGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *businessIdentityGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type businessIdentityRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing BusinessIdentityClient
	CallOptions **BusinessIdentityCallOptions
}

// NewBusinessIdentityRESTClient creates a new business identity service rest client.
//
// Service to support business
// identity (at https://support.google.com/merchants/answer/12564247) API.
func NewBusinessIdentityRESTClient(ctx context.Context, opts ...option.ClientOption) (*BusinessIdentityClient, error) {
	clientOpts := append(defaultBusinessIdentityRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultBusinessIdentityRESTCallOptions()
	c := &businessIdentityRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &BusinessIdentityClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultBusinessIdentityRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://merchantapi.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://merchantapi.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://merchantapi.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *businessIdentityRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *businessIdentityRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *businessIdentityRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *businessIdentityGRPCClient) GetBusinessIdentity(ctx context.Context, req *accountspb.GetBusinessIdentityRequest, opts ...gax.CallOption) (*accountspb.BusinessIdentity, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetBusinessIdentity[0:len((*c.CallOptions).GetBusinessIdentity):len((*c.CallOptions).GetBusinessIdentity)], opts...)
	var resp *accountspb.BusinessIdentity
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.businessIdentityClient.GetBusinessIdentity(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *businessIdentityGRPCClient) UpdateBusinessIdentity(ctx context.Context, req *accountspb.UpdateBusinessIdentityRequest, opts ...gax.CallOption) (*accountspb.BusinessIdentity, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "business_identity.name", url.QueryEscape(req.GetBusinessIdentity().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateBusinessIdentity[0:len((*c.CallOptions).UpdateBusinessIdentity):len((*c.CallOptions).UpdateBusinessIdentity)], opts...)
	var resp *accountspb.BusinessIdentity
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.businessIdentityClient.UpdateBusinessIdentity(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBusinessIdentity retrieves the business identity of an account.
func (c *businessIdentityRESTClient) GetBusinessIdentity(ctx context.Context, req *accountspb.GetBusinessIdentityRequest, opts ...gax.CallOption) (*accountspb.BusinessIdentity, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetBusinessIdentity[0:len((*c.CallOptions).GetBusinessIdentity):len((*c.CallOptions).GetBusinessIdentity)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.BusinessIdentity{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// UpdateBusinessIdentity updates the business identity of an account. Executing this method requires
// admin access.
func (c *businessIdentityRESTClient) UpdateBusinessIdentity(ctx context.Context, req *accountspb.UpdateBusinessIdentityRequest, opts ...gax.CallOption) (*accountspb.BusinessIdentity, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetBusinessIdentity()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v", req.GetBusinessIdentity().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		field, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(field[1:len(field)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "business_identity.name", url.QueryEscape(req.GetBusinessIdentity().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateBusinessIdentity[0:len((*c.CallOptions).UpdateBusinessIdentity):len((*c.CallOptions).UpdateBusinessIdentity)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.BusinessIdentity{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
