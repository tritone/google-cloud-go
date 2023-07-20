// Copyright 2023 Google LLC
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

package storage

import (
	"context"
	"fmt"
	"github.com/googleapis/gax-go/v2/callctx"
	"google.golang.org/api/option"
	"io"
	"net/http"
	"strings"
	"testing"
)

// Tests that sending a custom header via the context works as expected
// for a few different ops.
func TestCustomHeaders(t *testing.T) {
	var cases = []struct {
		name string
		call func(ctx context.Context, client *Client) error
		resp *http.Response
		opts []option.ClientOption
	}{
		{
			name: "Metadata op",
			call: func(ctx context.Context, client *Client) error {
				_, err := client.Bucket("my-bucket").Attrs(ctx)
				return err
			},
			resp: &http.Response{
				StatusCode: 200,
				Body:       bodyReader(`{"name":"my-bucket"}`),
			},
		},
		{
			name: "Writer",
			call: func(ctx context.Context, client *Client) error {
				w := client.Bucket("my-bucket").Object("my-object").NewWriter(ctx)
				if _, err := io.Copy(w, strings.NewReader("object data")); err != nil {
					return fmt.Errorf("io.Copy: %w", err)
				}
				return w.Close()
			},
			resp: &http.Response{
				StatusCode: 200,
				Body:       bodyReader(`{"name": "my-object"}`),
			},
		},
		{
			name: "Reader XML",
			call: func(ctx context.Context, client *Client) error {
				r, err := client.Bucket("my-bucket").Object("my-object").NewReader(ctx)
				if err != nil {
					return fmt.Errorf("NewReader: %w", err)
				}
				if _, err := io.Copy(io.Discard, r); err != nil {
					return fmt.Errorf("io.Copy: %w", err)
				}
				return r.Close()
			},
			resp: &http.Response{
				StatusCode: 200,
				Body:       bodyReader("object data"),
			},
			opts: []option.ClientOption{WithXMLReads()},
		},
		{
			name: "Reader JSON",
			call: func(ctx context.Context, client *Client) error {
				r, err := client.Bucket("my-bucket").Object("my-object").NewReader(ctx)
				if err != nil {
					return fmt.Errorf("NewReader: %w", err)
				}
				if _, err := io.Copy(io.Discard, r); err != nil {
					return fmt.Errorf("io.Copy: %w", err)
				}
				return r.Close()
			},
			resp: &http.Response{
				StatusCode: 200,
				Body:       bodyReader("object data"),
			},
			opts: []option.ClientOption{WithJSONReads()},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(r *testing.T) {
			mt := &mockTransport{}
			client := mockClient(t, mt, c.opts...)

			key := "x-goog-custom-audit-foo"
			value := "bar"
			ctx := callctx.SetHeaders(context.Background(), key, value)

			mt.addResult(c.resp, nil)
			if err := c.call(ctx, client); err != nil {
				r.Fatal(err)
			}
			if gotValue := mt.gotReq.Header.Get(key); gotValue != value {
				r.Errorf("Got headers %v, want to contain %q: %q", mt.gotReq.Header, key, value)
			}
		})
	}

}
