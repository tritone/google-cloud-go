// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/retail/v2/prediction_service.proto

package retailpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for Predict method.
type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Full resource name of the format:
	// `{placement=projects/*/locations/global/catalogs/default_catalog/servingConfigs/*}`
	// or
	// `{placement=projects/*/locations/global/catalogs/default_catalog/placements/*}`.
	// We recommend using the `servingConfigs` resource. `placements` is a legacy
	// resource.
	// The ID of the Recommendations AI serving config or placement.
	// Before you can request predictions from your model, you must create at
	// least one serving config or placement for it. For more information, see
	// [Manage serving configs]
	// (https://cloud.google.com/retail/docs/manage-configs).
	//
	// The full list of available serving configs can be seen at
	// https://console.cloud.google.com/ai/retail/catalogs/default_catalog/configs
	Placement string `protobuf:"bytes,1,opt,name=placement,proto3" json:"placement,omitempty"`
	// Required. Context about the user, what they are looking at and what action
	// they took to trigger the predict request. Note that this user event detail
	// won't be ingested to userEvent logs. Thus, a separate userEvent write
	// request is required for event logging.
	//
	// Don't set
	// [UserEvent.visitor_id][google.cloud.retail.v2.UserEvent.visitor_id] or
	// [UserInfo.user_id][google.cloud.retail.v2.UserInfo.user_id] to the same
	// fixed ID for different users. If you are trying to receive non-personalized
	// recommendations (not recommended; this can negatively impact model
	// performance), instead set
	// [UserEvent.visitor_id][google.cloud.retail.v2.UserEvent.visitor_id] to a
	// random unique ID and leave
	// [UserInfo.user_id][google.cloud.retail.v2.UserInfo.user_id] unset.
	UserEvent *UserEvent `protobuf:"bytes,2,opt,name=user_event,json=userEvent,proto3" json:"user_event,omitempty"`
	// Maximum number of results to return. Set this property to the number of
	// prediction results needed. If zero, the service will choose a reasonable
	// default. The maximum allowed value is 100. Values above 100 will be coerced
	// to 100.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// This field is not used; leave it unset.
	//
	// Deprecated: Marked as deprecated in google/cloud/retail/v2/prediction_service.proto.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Filter for restricting prediction results with a length limit of 5,000
	// characters. Accepts values for tags and the `filterOutOfStockItems` flag.
	//
	//   - Tag expressions. Restricts predictions to products that match all of the
	//     specified tags. Boolean operators `OR` and `NOT` are supported if the
	//     expression is enclosed in parentheses, and must be separated from the
	//     tag values by a space. `-"tagA"` is also supported and is equivalent to
	//     `NOT "tagA"`. Tag values must be double quoted UTF-8 encoded strings
	//     with a size limit of 1,000 characters.
	//
	//     Note: "Recently viewed" models don't support tag filtering at the
	//     moment.
	//
	//   - filterOutOfStockItems. Restricts predictions to products that do not
	//     have a
	//     stockState value of OUT_OF_STOCK.
	//
	// Examples:
	//
	//   - tag=("Red" OR "Blue") tag="New-Arrival" tag=(NOT "promotional")
	//   - filterOutOfStockItems  tag=(-"promotional")
	//   - filterOutOfStockItems
	//
	// If your filter blocks all prediction results, the API will return *no*
	// results. If instead you want empty result sets to return generic
	// (unfiltered) popular products, set `strictFiltering` to False in
	// `PredictRequest.params`. Note that the API will never return items with
	// storageStatus of "EXPIRED" or "DELETED" regardless of filter choices.
	//
	// If `filterSyntaxV2` is set to true under the `params` field, then
	// attribute-based expressions are expected instead of the above described
	// tag-based syntax. Examples:
	//
	//   - (colors: ANY("Red", "Blue")) AND NOT (categories: ANY("Phones"))
	//   - (availability: ANY("IN_STOCK")) AND
	//     (colors: ANY("Red") OR categories: ANY("Phones"))
	//
	// For more information, see
	// [Filter recommendations](https://cloud.google.com/retail/docs/filter-recs).
	Filter string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	// Use validate only mode for this prediction query. If set to true, a
	// dummy model will be used that returns arbitrary products.
	// Note that the validate only mode should only be used for testing the API,
	// or if the model is not ready.
	ValidateOnly bool `protobuf:"varint,6,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// Additional domain specific parameters for the predictions.
	//
	// Allowed values:
	//
	//   - `returnProduct`: Boolean. If set to true, the associated product
	//     object will be returned in the `results.metadata` field in the
	//     prediction response.
	//   - `returnScore`: Boolean. If set to true, the prediction 'score'
	//     corresponding to each returned product will be set in the
	//     `results.metadata` field in the prediction response. The given
	//     'score' indicates the probability of a product being clicked/purchased
	//     given the user's context and history.
	//   - `strictFiltering`: Boolean. True by default. If set to false, the service
	//     will return generic (unfiltered) popular products instead of empty if
	//     your filter blocks all prediction results.
	//   - `priceRerankLevel`: String. Default empty. If set to be non-empty, then
	//     it needs to be one of {'no-price-reranking', 'low-price-reranking',
	//     'medium-price-reranking', 'high-price-reranking'}. This gives
	//     request-level control and adjusts prediction results based on product
	//     price.
	//   - `diversityLevel`: String. Default empty. If set to be non-empty, then
	//     it needs to be one of {'no-diversity', 'low-diversity',
	//     'medium-diversity', 'high-diversity', 'auto-diversity'}. This gives
	//     request-level control and adjusts prediction results based on product
	//     category.
	//   - `filterSyntaxV2`: Boolean. False by default. If set to true, the `filter`
	//     field is interpreteted according to the new, attribute-based syntax.
	Params map[string]*structpb.Value `protobuf:"bytes,7,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The labels applied to a resource must meet the following requirements:
	//
	//   - Each resource can have multiple labels, up to a maximum of 64.
	//   - Each label must be a key-value pair.
	//   - Keys have a minimum length of 1 character and a maximum length of 63
	//     characters and cannot be empty. Values can be empty and have a maximum
	//     length of 63 characters.
	//   - Keys and values can contain only lowercase letters, numeric characters,
	//     underscores, and dashes. All characters must use UTF-8 encoding, and
	//     international characters are allowed.
	//   - The key portion of a label must be unique. However, you can use the same
	//     key with multiple resources.
	//   - Keys must start with a lowercase letter or international character.
	//
	// See [Google Cloud
	// Document](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements)
	// for more details.
	Labels map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_prediction_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_prediction_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_prediction_service_proto_rawDescGZIP(), []int{0}
}

func (x *PredictRequest) GetPlacement() string {
	if x != nil {
		return x.Placement
	}
	return ""
}

func (x *PredictRequest) GetUserEvent() *UserEvent {
	if x != nil {
		return x.UserEvent
	}
	return nil
}

func (x *PredictRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Deprecated: Marked as deprecated in google/cloud/retail/v2/prediction_service.proto.
func (x *PredictRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *PredictRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *PredictRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *PredictRequest) GetParams() map[string]*structpb.Value {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *PredictRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Response message for predict method.
type PredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of recommended products. The order represents the ranking (from the
	// most relevant product to the least).
	Results []*PredictResponse_PredictionResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// A unique attribution token. This should be included in the
	// [UserEvent][google.cloud.retail.v2.UserEvent] logs resulting from this
	// recommendation, which enables accurate attribution of recommendation model
	// performance.
	AttributionToken string `protobuf:"bytes,2,opt,name=attribution_token,json=attributionToken,proto3" json:"attribution_token,omitempty"`
	// IDs of products in the request that were missing from the inventory.
	MissingIds []string `protobuf:"bytes,3,rep,name=missing_ids,json=missingIds,proto3" json:"missing_ids,omitempty"`
	// True if the validateOnly property was set in the request.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *PredictResponse) Reset() {
	*x = PredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_prediction_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse) ProtoMessage() {}

func (x *PredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_prediction_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse.ProtoReflect.Descriptor instead.
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_prediction_service_proto_rawDescGZIP(), []int{1}
}

func (x *PredictResponse) GetResults() []*PredictResponse_PredictionResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PredictResponse) GetAttributionToken() string {
	if x != nil {
		return x.AttributionToken
	}
	return ""
}

func (x *PredictResponse) GetMissingIds() []string {
	if x != nil {
		return x.MissingIds
	}
	return nil
}

func (x *PredictResponse) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// PredictionResult represents the recommendation prediction results.
type PredictResponse_PredictionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the recommended product
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Additional product metadata / annotations.
	//
	// Possible values:
	//
	//   - `product`: JSON representation of the product. Is set if
	//     `returnProduct` is set to true in `PredictRequest.params`.
	//   - `score`: Prediction score in double value. Is set if
	//     `returnScore` is set to true in `PredictRequest.params`.
	Metadata map[string]*structpb.Value `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PredictResponse_PredictionResult) Reset() {
	*x = PredictResponse_PredictionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_prediction_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse_PredictionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse_PredictionResult) ProtoMessage() {}

func (x *PredictResponse_PredictionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_prediction_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse_PredictionResult.ProtoReflect.Descriptor instead.
func (*PredictResponse_PredictionResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_prediction_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PredictResponse_PredictionResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PredictResponse_PredictionResult) GetMetadata() map[string]*structpb.Value {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_google_cloud_retail_v2_prediction_service_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2_prediction_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x04, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x4a, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4a, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x51, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xb6, 0x03, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x1a, 0xdb,
	0x01, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x62, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x53, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xe1, 0x02, 0x0a,
	0x11, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x80, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x26,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xa3, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x9c, 0x01, 0x3a, 0x01, 0x2a, 0x5a, 0x4f, 0x3a, 0x01,
	0x2a, 0x22, 0x4a, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x22, 0x46, 0x2f,
	0x76, 0x32, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x2f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x1a, 0x49, 0xca, 0x41, 0x15, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x42, 0xc1, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x42,
	0x16, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0xa2, 0x02, 0x06,
	0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0xca,
	0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2_prediction_service_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2_prediction_service_proto_rawDescData = file_google_cloud_retail_v2_prediction_service_proto_rawDesc
)

func file_google_cloud_retail_v2_prediction_service_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2_prediction_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2_prediction_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2_prediction_service_proto_rawDescData)
	})
	return file_google_cloud_retail_v2_prediction_service_proto_rawDescData
}

var file_google_cloud_retail_v2_prediction_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_retail_v2_prediction_service_proto_goTypes = []any{
	(*PredictRequest)(nil),                   // 0: google.cloud.retail.v2.PredictRequest
	(*PredictResponse)(nil),                  // 1: google.cloud.retail.v2.PredictResponse
	nil,                                      // 2: google.cloud.retail.v2.PredictRequest.ParamsEntry
	nil,                                      // 3: google.cloud.retail.v2.PredictRequest.LabelsEntry
	(*PredictResponse_PredictionResult)(nil), // 4: google.cloud.retail.v2.PredictResponse.PredictionResult
	nil,                                      // 5: google.cloud.retail.v2.PredictResponse.PredictionResult.MetadataEntry
	(*UserEvent)(nil),                        // 6: google.cloud.retail.v2.UserEvent
	(*structpb.Value)(nil),                   // 7: google.protobuf.Value
}
var file_google_cloud_retail_v2_prediction_service_proto_depIdxs = []int32{
	6, // 0: google.cloud.retail.v2.PredictRequest.user_event:type_name -> google.cloud.retail.v2.UserEvent
	2, // 1: google.cloud.retail.v2.PredictRequest.params:type_name -> google.cloud.retail.v2.PredictRequest.ParamsEntry
	3, // 2: google.cloud.retail.v2.PredictRequest.labels:type_name -> google.cloud.retail.v2.PredictRequest.LabelsEntry
	4, // 3: google.cloud.retail.v2.PredictResponse.results:type_name -> google.cloud.retail.v2.PredictResponse.PredictionResult
	7, // 4: google.cloud.retail.v2.PredictRequest.ParamsEntry.value:type_name -> google.protobuf.Value
	5, // 5: google.cloud.retail.v2.PredictResponse.PredictionResult.metadata:type_name -> google.cloud.retail.v2.PredictResponse.PredictionResult.MetadataEntry
	7, // 6: google.cloud.retail.v2.PredictResponse.PredictionResult.MetadataEntry.value:type_name -> google.protobuf.Value
	0, // 7: google.cloud.retail.v2.PredictionService.Predict:input_type -> google.cloud.retail.v2.PredictRequest
	1, // 8: google.cloud.retail.v2.PredictionService.Predict:output_type -> google.cloud.retail.v2.PredictResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2_prediction_service_proto_init() }
func file_google_cloud_retail_v2_prediction_service_proto_init() {
	if File_google_cloud_retail_v2_prediction_service_proto != nil {
		return
	}
	file_google_cloud_retail_v2_user_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2_prediction_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PredictRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_retail_v2_prediction_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PredictResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_retail_v2_prediction_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PredictResponse_PredictionResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_retail_v2_prediction_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_retail_v2_prediction_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2_prediction_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2_prediction_service_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2_prediction_service_proto = out.File
	file_google_cloud_retail_v2_prediction_service_proto_rawDesc = nil
	file_google_cloud_retail_v2_prediction_service_proto_goTypes = nil
	file_google_cloud_retail_v2_prediction_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Makes a recommendation prediction.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
}

type predictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictionServiceClient(cc grpc.ClientConnInterface) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Makes a recommendation prediction.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
}

// UnimplementedPredictionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPredictionServiceServer struct {
}

func (*UnimplementedPredictionServiceServer) Predict(context.Context, *PredictRequest) (*PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.retail.v2.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/retail/v2/prediction_service.proto",
}
