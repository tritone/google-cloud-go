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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/shopping/merchant/datasources/v1beta/datasourcetypes.proto

package datasourcespb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Data Source Channel.
//
// Channel is used to distinguish between data sources for different product
// verticals.
type PrimaryProductDataSource_Channel int32

const (
	// Not specified.
	PrimaryProductDataSource_CHANNEL_UNSPECIFIED PrimaryProductDataSource_Channel = 0
	// Online product.
	PrimaryProductDataSource_ONLINE_PRODUCTS PrimaryProductDataSource_Channel = 1
	// Local product.
	PrimaryProductDataSource_LOCAL_PRODUCTS PrimaryProductDataSource_Channel = 2
	// Unified data source for both local and online products.
	PrimaryProductDataSource_PRODUCTS PrimaryProductDataSource_Channel = 3
)

// Enum value maps for PrimaryProductDataSource_Channel.
var (
	PrimaryProductDataSource_Channel_name = map[int32]string{
		0: "CHANNEL_UNSPECIFIED",
		1: "ONLINE_PRODUCTS",
		2: "LOCAL_PRODUCTS",
		3: "PRODUCTS",
	}
	PrimaryProductDataSource_Channel_value = map[string]int32{
		"CHANNEL_UNSPECIFIED": 0,
		"ONLINE_PRODUCTS":     1,
		"LOCAL_PRODUCTS":      2,
		"PRODUCTS":            3,
	}
)

func (x PrimaryProductDataSource_Channel) Enum() *PrimaryProductDataSource_Channel {
	p := new(PrimaryProductDataSource_Channel)
	*p = x
	return p
}

func (x PrimaryProductDataSource_Channel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrimaryProductDataSource_Channel) Descriptor() protoreflect.EnumDescriptor {
	return file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_enumTypes[0].Descriptor()
}

func (PrimaryProductDataSource_Channel) Type() protoreflect.EnumType {
	return &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_enumTypes[0]
}

func (x PrimaryProductDataSource_Channel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrimaryProductDataSource_Channel.Descriptor instead.
func (PrimaryProductDataSource_Channel) EnumDescriptor() ([]byte, []int) {
	return file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescGZIP(), []int{0, 0}
}

// The primary data source for local and online products.
type PrimaryProductDataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. Specifies the type of data source channel.
	Channel PrimaryProductDataSource_Channel `protobuf:"varint,3,opt,name=channel,proto3,enum=google.shopping.merchant.datasources.v1beta.PrimaryProductDataSource_Channel" json:"channel,omitempty"`
	// Optional. Immutable. The feed label that is specified on the data source
	// level.
	//
	// Must be less than or equal to 20 uppercase letters (A-Z), numbers (0-9),
	// and dashes (-).
	//
	// See also [migration to feed
	// labels](https://developers.google.com/shopping-content/guides/products/feed-labels).
	//
	// `feedLabel` and `contentLanguage` must be either both set or unset for data
	// sources with product content type.
	// They must be set for data sources with a file input.
	//
	// If set, the data source will only accept products matching this
	// combination. If unset, the data source will accept products without that
	// restriction.
	FeedLabel *string `protobuf:"bytes,4,opt,name=feed_label,json=feedLabel,proto3,oneof" json:"feed_label,omitempty"`
	// Optional. Immutable. The two-letter ISO 639-1 language of the items in the
	// data source.
	//
	// `feedLabel` and `contentLanguage` must be either both set or unset.
	// The fields can only be unset for data sources without file input.
	//
	// If set, the data source will only accept products matching this
	// combination. If unset, the data source will accept products without that
	// restriction.
	ContentLanguage *string `protobuf:"bytes,5,opt,name=content_language,json=contentLanguage,proto3,oneof" json:"content_language,omitempty"`
	// Optional. The countries where the items may be displayed. Represented as a
	// [CLDR territory
	// code](https://github.com/unicode-org/cldr/blob/latest/common/main/en.xml).
	Countries []string `protobuf:"bytes,6,rep,name=countries,proto3" json:"countries,omitempty"`
}

func (x *PrimaryProductDataSource) Reset() {
	*x = PrimaryProductDataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryProductDataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryProductDataSource) ProtoMessage() {}

func (x *PrimaryProductDataSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryProductDataSource.ProtoReflect.Descriptor instead.
func (*PrimaryProductDataSource) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescGZIP(), []int{0}
}

func (x *PrimaryProductDataSource) GetChannel() PrimaryProductDataSource_Channel {
	if x != nil {
		return x.Channel
	}
	return PrimaryProductDataSource_CHANNEL_UNSPECIFIED
}

func (x *PrimaryProductDataSource) GetFeedLabel() string {
	if x != nil && x.FeedLabel != nil {
		return *x.FeedLabel
	}
	return ""
}

func (x *PrimaryProductDataSource) GetContentLanguage() string {
	if x != nil && x.ContentLanguage != nil {
		return *x.ContentLanguage
	}
	return ""
}

func (x *PrimaryProductDataSource) GetCountries() []string {
	if x != nil {
		return x.Countries
	}
	return nil
}

// The supplemental data source for local and online products.
type SupplementalProductDataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Immutable. The feed label that is specified on the data source
	// level.
	//
	// Must be less than or equal to 20 uppercase letters (A-Z), numbers (0-9),
	// and dashes (-).
	//
	// See also [migration to feed
	// labels](https://developers.google.com/shopping-content/guides/products/feed-labels).
	//
	// `feedLabel` and `contentLanguage` must be either both set or unset for data
	// sources with product content type.
	// They must be set for data sources with a file input.
	//
	// If set, the data source will only accept products matching this
	// combination. If unset, the data source will accept produts without that
	// restriction.
	FeedLabel *string `protobuf:"bytes,4,opt,name=feed_label,json=feedLabel,proto3,oneof" json:"feed_label,omitempty"`
	// Optional. Immutable. The two-letter ISO 639-1 language of the items in the
	// data source.
	//
	// `feedLabel` and `contentLanguage` must be either both set or unset.
	// The fields can only be unset for data sources without file input.
	//
	// If set, the data source will only accept products matching this
	// combination. If unset, the data source will accept produts without that
	// restriction.
	ContentLanguage *string `protobuf:"bytes,5,opt,name=content_language,json=contentLanguage,proto3,oneof" json:"content_language,omitempty"`
}

func (x *SupplementalProductDataSource) Reset() {
	*x = SupplementalProductDataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplementalProductDataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplementalProductDataSource) ProtoMessage() {}

func (x *SupplementalProductDataSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplementalProductDataSource.ProtoReflect.Descriptor instead.
func (*SupplementalProductDataSource) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescGZIP(), []int{1}
}

func (x *SupplementalProductDataSource) GetFeedLabel() string {
	if x != nil && x.FeedLabel != nil {
		return *x.FeedLabel
	}
	return ""
}

func (x *SupplementalProductDataSource) GetContentLanguage() string {
	if x != nil && x.ContentLanguage != nil {
		return *x.ContentLanguage
	}
	return ""
}

// The local inventory data source.
type LocalInventoryDataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The feed label of the offers to which the local
	// inventory is provided.
	//
	// Must be less than or equal to 20 uppercase letters (A-Z), numbers (0-9),
	// and dashes (-).
	//
	// See also [migration to feed
	// labels](https://developers.google.com/shopping-content/guides/products/feed-labels).
	FeedLabel string `protobuf:"bytes,4,opt,name=feed_label,json=feedLabel,proto3" json:"feed_label,omitempty"`
	// Required. Immutable. The two-letter ISO 639-1 language of the items to
	// which the local inventory is provided.
	ContentLanguage string `protobuf:"bytes,5,opt,name=content_language,json=contentLanguage,proto3" json:"content_language,omitempty"`
}

func (x *LocalInventoryDataSource) Reset() {
	*x = LocalInventoryDataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalInventoryDataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalInventoryDataSource) ProtoMessage() {}

func (x *LocalInventoryDataSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalInventoryDataSource.ProtoReflect.Descriptor instead.
func (*LocalInventoryDataSource) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescGZIP(), []int{2}
}

func (x *LocalInventoryDataSource) GetFeedLabel() string {
	if x != nil {
		return x.FeedLabel
	}
	return ""
}

func (x *LocalInventoryDataSource) GetContentLanguage() string {
	if x != nil {
		return x.ContentLanguage
	}
	return ""
}

// The regional inventory data source.
type RegionalInventoryDataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The feed label of the offers to which the regional
	// inventory is provided.
	//
	// Must be less than or equal to 20 uppercase letters (A-Z), numbers (0-9),
	// and dashes (-).
	//
	// See also [migration to feed
	// labels](https://developers.google.com/shopping-content/guides/products/feed-labels).
	FeedLabel string `protobuf:"bytes,4,opt,name=feed_label,json=feedLabel,proto3" json:"feed_label,omitempty"`
	// Required. Immutable. The two-letter ISO 639-1 language of the items to
	// which the regional inventory is provided.
	ContentLanguage string `protobuf:"bytes,5,opt,name=content_language,json=contentLanguage,proto3" json:"content_language,omitempty"`
}

func (x *RegionalInventoryDataSource) Reset() {
	*x = RegionalInventoryDataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionalInventoryDataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionalInventoryDataSource) ProtoMessage() {}

func (x *RegionalInventoryDataSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionalInventoryDataSource.ProtoReflect.Descriptor instead.
func (*RegionalInventoryDataSource) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescGZIP(), []int{3}
}

func (x *RegionalInventoryDataSource) GetFeedLabel() string {
	if x != nil {
		return x.FeedLabel
	}
	return ""
}

func (x *RegionalInventoryDataSource) GetContentLanguage() string {
	if x != nil {
		return x.ContentLanguage
	}
	return ""
}

// The promotion data source.
type PromotionDataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The target country used as part of the unique
	// identifier. Represented as a [CLDR territory
	// code](https://github.com/unicode-org/cldr/blob/latest/common/main/en.xml).
	//
	// Promotions are only available in selected
	// [countries](https://support.google.com/merchants/answer/4588460).
	TargetCountry string `protobuf:"bytes,1,opt,name=target_country,json=targetCountry,proto3" json:"target_country,omitempty"`
	// Required. Immutable. The two-letter ISO 639-1 language of the items in the
	// data source.
	ContentLanguage string `protobuf:"bytes,2,opt,name=content_language,json=contentLanguage,proto3" json:"content_language,omitempty"`
}

func (x *PromotionDataSource) Reset() {
	*x = PromotionDataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionDataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionDataSource) ProtoMessage() {}

func (x *PromotionDataSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionDataSource.ProtoReflect.Descriptor instead.
func (*PromotionDataSource) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescGZIP(), []int{4}
}

func (x *PromotionDataSource) GetTargetCountry() string {
	if x != nil {
		return x.TargetCountry
	}
	return ""
}

func (x *PromotionDataSource) GetContentLanguage() string {
	if x != nil {
		return x.ContentLanguage
	}
	return ""
}

var File_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto protoreflect.FileDescriptor

var file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x91, 0x03, 0x0a, 0x18, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x6f,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x06,
	0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x2a, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x01, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x09, 0x66,
	0x65, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x10, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x01, 0xe0, 0x41, 0x05, 0x48, 0x01, 0x52,
	0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x59, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x4e,
	0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x53, 0x10, 0x01, 0x12,
	0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54,
	0x53, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x53, 0x10,
	0x03, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x1d, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x01,
	0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0,
	0x41, 0x01, 0xe0, 0x41, 0x05, 0x48, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22,
	0x74, 0x0a, 0x18, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0a, 0x66,
	0x65, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x31, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41,
	0x02, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x77, 0x0a, 0x1b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05,
	0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x31, 0x0a, 0x10, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x77,
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0,
	0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0xb0, 0x02, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x14, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x57, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0x3b, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x2b, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x2b, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5c, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescOnce sync.Once
	file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescData = file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDesc
)

func file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescGZIP() []byte {
	file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescOnce.Do(func() {
		file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescData)
	})
	return file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDescData
}

var file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_goTypes = []interface{}{
	(PrimaryProductDataSource_Channel)(0), // 0: google.shopping.merchant.datasources.v1beta.PrimaryProductDataSource.Channel
	(*PrimaryProductDataSource)(nil),      // 1: google.shopping.merchant.datasources.v1beta.PrimaryProductDataSource
	(*SupplementalProductDataSource)(nil), // 2: google.shopping.merchant.datasources.v1beta.SupplementalProductDataSource
	(*LocalInventoryDataSource)(nil),      // 3: google.shopping.merchant.datasources.v1beta.LocalInventoryDataSource
	(*RegionalInventoryDataSource)(nil),   // 4: google.shopping.merchant.datasources.v1beta.RegionalInventoryDataSource
	(*PromotionDataSource)(nil),           // 5: google.shopping.merchant.datasources.v1beta.PromotionDataSource
}
var file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_depIdxs = []int32{
	0, // 0: google.shopping.merchant.datasources.v1beta.PrimaryProductDataSource.channel:type_name -> google.shopping.merchant.datasources.v1beta.PrimaryProductDataSource.Channel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_init() }
func file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_init() {
	if File_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryProductDataSource); i {
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
		file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplementalProductDataSource); i {
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
		file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalInventoryDataSource); i {
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
		file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionalInventoryDataSource); i {
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
		file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionDataSource); i {
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
	file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_goTypes,
		DependencyIndexes: file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_depIdxs,
		EnumInfos:         file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_enumTypes,
		MessageInfos:      file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_msgTypes,
	}.Build()
	File_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto = out.File
	file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_rawDesc = nil
	file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_goTypes = nil
	file_google_shopping_merchant_datasources_v1beta_datasourcetypes_proto_depIdxs = nil
}
