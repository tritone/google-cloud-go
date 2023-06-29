// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google/maps/routing/v2/transit_preferences.proto

package routingpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A set of values used to specify the mode of transit.
type TransitPreferences_TransitTravelMode int32

const (
	// No transit travel mode specified.
	TransitPreferences_TRANSIT_TRAVEL_MODE_UNSPECIFIED TransitPreferences_TransitTravelMode = 0
	// Travel by bus.
	TransitPreferences_BUS TransitPreferences_TransitTravelMode = 1
	// Travel by subway.
	TransitPreferences_SUBWAY TransitPreferences_TransitTravelMode = 2
	// Travel by train.
	TransitPreferences_TRAIN TransitPreferences_TransitTravelMode = 3
	// Travel by light rail or tram.
	TransitPreferences_LIGHT_RAIL TransitPreferences_TransitTravelMode = 4
	// Travel by rail. This is equivalent to a combination of `SUBWAY`, `TRAIN`,
	// and `LIGHT_RAIL`.
	TransitPreferences_RAIL TransitPreferences_TransitTravelMode = 5
)

// Enum value maps for TransitPreferences_TransitTravelMode.
var (
	TransitPreferences_TransitTravelMode_name = map[int32]string{
		0: "TRANSIT_TRAVEL_MODE_UNSPECIFIED",
		1: "BUS",
		2: "SUBWAY",
		3: "TRAIN",
		4: "LIGHT_RAIL",
		5: "RAIL",
	}
	TransitPreferences_TransitTravelMode_value = map[string]int32{
		"TRANSIT_TRAVEL_MODE_UNSPECIFIED": 0,
		"BUS":                             1,
		"SUBWAY":                          2,
		"TRAIN":                           3,
		"LIGHT_RAIL":                      4,
		"RAIL":                            5,
	}
)

func (x TransitPreferences_TransitTravelMode) Enum() *TransitPreferences_TransitTravelMode {
	p := new(TransitPreferences_TransitTravelMode)
	*p = x
	return p
}

func (x TransitPreferences_TransitTravelMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransitPreferences_TransitTravelMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routing_v2_transit_preferences_proto_enumTypes[0].Descriptor()
}

func (TransitPreferences_TransitTravelMode) Type() protoreflect.EnumType {
	return &file_google_maps_routing_v2_transit_preferences_proto_enumTypes[0]
}

func (x TransitPreferences_TransitTravelMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransitPreferences_TransitTravelMode.Descriptor instead.
func (TransitPreferences_TransitTravelMode) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_transit_preferences_proto_rawDescGZIP(), []int{0, 0}
}

// Specifies routing preferences for transit routes.
type TransitPreferences_TransitRoutingPreference int32

const (
	// No preference specified.
	TransitPreferences_TRANSIT_ROUTING_PREFERENCE_UNSPECIFIED TransitPreferences_TransitRoutingPreference = 0
	// Indicates that the calculated route should prefer limited amounts of
	// walking.
	TransitPreferences_LESS_WALKING TransitPreferences_TransitRoutingPreference = 1
	// Indicates that the calculated route should prefer a limited number of
	// transfers.
	TransitPreferences_FEWER_TRANSFERS TransitPreferences_TransitRoutingPreference = 2
)

// Enum value maps for TransitPreferences_TransitRoutingPreference.
var (
	TransitPreferences_TransitRoutingPreference_name = map[int32]string{
		0: "TRANSIT_ROUTING_PREFERENCE_UNSPECIFIED",
		1: "LESS_WALKING",
		2: "FEWER_TRANSFERS",
	}
	TransitPreferences_TransitRoutingPreference_value = map[string]int32{
		"TRANSIT_ROUTING_PREFERENCE_UNSPECIFIED": 0,
		"LESS_WALKING":                           1,
		"FEWER_TRANSFERS":                        2,
	}
)

func (x TransitPreferences_TransitRoutingPreference) Enum() *TransitPreferences_TransitRoutingPreference {
	p := new(TransitPreferences_TransitRoutingPreference)
	*p = x
	return p
}

func (x TransitPreferences_TransitRoutingPreference) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransitPreferences_TransitRoutingPreference) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routing_v2_transit_preferences_proto_enumTypes[1].Descriptor()
}

func (TransitPreferences_TransitRoutingPreference) Type() protoreflect.EnumType {
	return &file_google_maps_routing_v2_transit_preferences_proto_enumTypes[1]
}

func (x TransitPreferences_TransitRoutingPreference) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransitPreferences_TransitRoutingPreference.Descriptor instead.
func (TransitPreferences_TransitRoutingPreference) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_transit_preferences_proto_rawDescGZIP(), []int{0, 1}
}

// Preferences for `TRANSIT` based routes that influence the route that is
// returned.
type TransitPreferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A set of travel modes to use when getting a `TRANSIT` route. Defaults to
	// all supported modes of travel.
	AllowedTravelModes []TransitPreferences_TransitTravelMode `protobuf:"varint,1,rep,packed,name=allowed_travel_modes,json=allowedTravelModes,proto3,enum=google.maps.routing.v2.TransitPreferences_TransitTravelMode" json:"allowed_travel_modes,omitempty"`
	// A routing preference that, when specified, influences the `TRANSIT` route
	// returned.
	RoutingPreference TransitPreferences_TransitRoutingPreference `protobuf:"varint,2,opt,name=routing_preference,json=routingPreference,proto3,enum=google.maps.routing.v2.TransitPreferences_TransitRoutingPreference" json:"routing_preference,omitempty"`
}

func (x *TransitPreferences) Reset() {
	*x = TransitPreferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routing_v2_transit_preferences_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitPreferences) ProtoMessage() {}

func (x *TransitPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_transit_preferences_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitPreferences.ProtoReflect.Descriptor instead.
func (*TransitPreferences) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_transit_preferences_proto_rawDescGZIP(), []int{0}
}

func (x *TransitPreferences) GetAllowedTravelModes() []TransitPreferences_TransitTravelMode {
	if x != nil {
		return x.AllowedTravelModes
	}
	return nil
}

func (x *TransitPreferences) GetRoutingPreference() TransitPreferences_TransitRoutingPreference {
	if x != nil {
		return x.RoutingPreference
	}
	return TransitPreferences_TRANSIT_ROUTING_PREFERENCE_UNSPECIFIED
}

var File_google_maps_routing_v2_transit_preferences_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_transit_preferences_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x22, 0xdb, 0x03, 0x0a, 0x12, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x6e, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x74, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x12, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x72, 0x0a, 0x12, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x72, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x49, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x42, 0x55, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x42, 0x57,
	0x41, 0x59, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x10, 0x03, 0x12,
	0x0e, 0x0a, 0x0a, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x52, 0x41, 0x49, 0x4c, 0x10, 0x04, 0x12,
	0x08, 0x0a, 0x04, 0x52, 0x41, 0x49, 0x4c, 0x10, 0x05, 0x22, 0x6d, 0x0a, 0x18, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x54,
	0x5f, 0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45,
	0x4e, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x57, 0x41, 0x4c, 0x4b, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x45, 0x57, 0x45, 0x52, 0x5f, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x46, 0x45, 0x52, 0x53, 0x10, 0x02, 0x42, 0xcc, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x70, 0x62, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xf8, 0x01,
	0x01, 0xa2, 0x02, 0x05, 0x47, 0x4d, 0x52, 0x56, 0x32, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73,
	0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_transit_preferences_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_transit_preferences_proto_rawDescData = file_google_maps_routing_v2_transit_preferences_proto_rawDesc
)

func file_google_maps_routing_v2_transit_preferences_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_transit_preferences_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_transit_preferences_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_transit_preferences_proto_rawDescData)
	})
	return file_google_maps_routing_v2_transit_preferences_proto_rawDescData
}

var file_google_maps_routing_v2_transit_preferences_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_maps_routing_v2_transit_preferences_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_routing_v2_transit_preferences_proto_goTypes = []interface{}{
	(TransitPreferences_TransitTravelMode)(0),        // 0: google.maps.routing.v2.TransitPreferences.TransitTravelMode
	(TransitPreferences_TransitRoutingPreference)(0), // 1: google.maps.routing.v2.TransitPreferences.TransitRoutingPreference
	(*TransitPreferences)(nil),                       // 2: google.maps.routing.v2.TransitPreferences
}
var file_google_maps_routing_v2_transit_preferences_proto_depIdxs = []int32{
	0, // 0: google.maps.routing.v2.TransitPreferences.allowed_travel_modes:type_name -> google.maps.routing.v2.TransitPreferences.TransitTravelMode
	1, // 1: google.maps.routing.v2.TransitPreferences.routing_preference:type_name -> google.maps.routing.v2.TransitPreferences.TransitRoutingPreference
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_transit_preferences_proto_init() }
func file_google_maps_routing_v2_transit_preferences_proto_init() {
	if File_google_maps_routing_v2_transit_preferences_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routing_v2_transit_preferences_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitPreferences); i {
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
			RawDescriptor: file_google_maps_routing_v2_transit_preferences_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_transit_preferences_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_transit_preferences_proto_depIdxs,
		EnumInfos:         file_google_maps_routing_v2_transit_preferences_proto_enumTypes,
		MessageInfos:      file_google_maps_routing_v2_transit_preferences_proto_msgTypes,
	}.Build()
	File_google_maps_routing_v2_transit_preferences_proto = out.File
	file_google_maps_routing_v2_transit_preferences_proto_rawDesc = nil
	file_google_maps_routing_v2_transit_preferences_proto_goTypes = nil
	file_google_maps_routing_v2_transit_preferences_proto_depIdxs = nil
}
