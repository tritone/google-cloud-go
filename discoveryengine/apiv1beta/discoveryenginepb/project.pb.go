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
// source: google/cloud/discoveryengine/v1beta/project.proto

package discoveryenginepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The agreement states this terms of service.
type Project_ServiceTerms_State int32

const (
	// The default value of the enum. This value is not actually used.
	Project_ServiceTerms_STATE_UNSPECIFIED Project_ServiceTerms_State = 0
	// The project has given consent to the terms of service.
	Project_ServiceTerms_TERMS_ACCEPTED Project_ServiceTerms_State = 1
	// The project is pending to review and accept the terms of service.
	Project_ServiceTerms_TERMS_PENDING Project_ServiceTerms_State = 2
	// The project has declined or revoked the agreement to terms of service.
	Project_ServiceTerms_TERMS_DECLINED Project_ServiceTerms_State = 3
)

// Enum value maps for Project_ServiceTerms_State.
var (
	Project_ServiceTerms_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "TERMS_ACCEPTED",
		2: "TERMS_PENDING",
		3: "TERMS_DECLINED",
	}
	Project_ServiceTerms_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"TERMS_ACCEPTED":    1,
		"TERMS_PENDING":     2,
		"TERMS_DECLINED":    3,
	}
)

func (x Project_ServiceTerms_State) Enum() *Project_ServiceTerms_State {
	p := new(Project_ServiceTerms_State)
	*p = x
	return p
}

func (x Project_ServiceTerms_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Project_ServiceTerms_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_discoveryengine_v1beta_project_proto_enumTypes[0].Descriptor()
}

func (Project_ServiceTerms_State) Type() protoreflect.EnumType {
	return &file_google_cloud_discoveryengine_v1beta_project_proto_enumTypes[0]
}

func (x Project_ServiceTerms_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Project_ServiceTerms_State.Descriptor instead.
func (Project_ServiceTerms_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_project_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Metadata and configurations for a Google Cloud project in the service.
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Full resource name of the project, for example
	// `projects/{project_number}`.
	// Note that when making requests, project number and project id are both
	// acceptable, but the server will always respond in project number.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The timestamp when this project is created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The timestamp when this project is successfully provisioned.
	// Empty value means this project is still provisioning and is not ready for
	// use.
	ProvisionCompletionTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=provision_completion_time,json=provisionCompletionTime,proto3" json:"provision_completion_time,omitempty"`
	// Output only. A map of terms of services. The key is the `id` of
	// [ServiceTerms][google.cloud.discoveryengine.v1beta.Project.ServiceTerms].
	ServiceTermsMap map[string]*Project_ServiceTerms `protobuf:"bytes,4,rep,name=service_terms_map,json=serviceTermsMap,proto3" json:"service_terms_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Project) GetProvisionCompletionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ProvisionCompletionTime
	}
	return nil
}

func (x *Project) GetServiceTermsMap() map[string]*Project_ServiceTerms {
	if x != nil {
		return x.ServiceTermsMap
	}
	return nil
}

// Metadata about the terms of service.
type Project_ServiceTerms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of this terms of service.
	// Available terms:
	//
	// * `GA_DATA_USE_TERMS`: [Terms for data
	// use](https://cloud.google.com/retail/data-use-terms). When using this as
	// `id`, the acceptable
	// [version][google.cloud.discoveryengine.v1beta.Project.ServiceTerms.version]
	// to provide is `2022-11-23`.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The version string of the terms of service.
	// For acceptable values, see the comments for
	// [id][google.cloud.discoveryengine.v1beta.Project.ServiceTerms.id] above.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Whether the project has accepted/rejected the service terms or it is
	// still pending.
	State Project_ServiceTerms_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.discoveryengine.v1beta.Project_ServiceTerms_State" json:"state,omitempty"`
	// The last time when the project agreed to the terms of service.
	AcceptTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=accept_time,json=acceptTime,proto3" json:"accept_time,omitempty"`
	// The last time when the project declined or revoked the agreement to terms
	// of service.
	DeclineTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=decline_time,json=declineTime,proto3" json:"decline_time,omitempty"`
}

func (x *Project_ServiceTerms) Reset() {
	*x = Project_ServiceTerms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project_ServiceTerms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project_ServiceTerms) ProtoMessage() {}

func (x *Project_ServiceTerms) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project_ServiceTerms.ProtoReflect.Descriptor instead.
func (*Project_ServiceTerms) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_project_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Project_ServiceTerms) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project_ServiceTerms) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Project_ServiceTerms) GetState() Project_ServiceTerms_State {
	if x != nil {
		return x.State
	}
	return Project_ServiceTerms_STATE_UNSPECIFIED
}

func (x *Project_ServiceTerms) GetAcceptTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AcceptTime
	}
	return nil
}

func (x *Project_ServiceTerms) GetDeclineTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeclineTime
	}
	return nil
}

var File_google_cloud_discoveryengine_v1beta_project_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1beta_project_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x06, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x19,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x17, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x72, 0x0a, 0x11, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4d, 0x61, 0x70, 0x1a, 0xe6, 0x02,
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x72,
	0x6d, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x3b, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x64, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54,
	0x45, 0x52, 0x4d, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x54, 0x45, 0x52, 0x4d, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x45, 0x52, 0x4d, 0x53, 0x5f, 0x44, 0x45, 0x43, 0x4c,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x7d, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x4f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x3f, 0xea, 0x41, 0x3c, 0x0a, 0x26, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x42, 0x93, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52,
	0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1beta_project_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1beta_project_proto_rawDescData = file_google_cloud_discoveryengine_v1beta_project_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1beta_project_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1beta_project_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1beta_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1beta_project_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1beta_project_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1beta_project_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_discoveryengine_v1beta_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_discoveryengine_v1beta_project_proto_goTypes = []any{
	(Project_ServiceTerms_State)(0), // 0: google.cloud.discoveryengine.v1beta.Project.ServiceTerms.State
	(*Project)(nil),                 // 1: google.cloud.discoveryengine.v1beta.Project
	(*Project_ServiceTerms)(nil),    // 2: google.cloud.discoveryengine.v1beta.Project.ServiceTerms
	nil,                             // 3: google.cloud.discoveryengine.v1beta.Project.ServiceTermsMapEntry
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
}
var file_google_cloud_discoveryengine_v1beta_project_proto_depIdxs = []int32{
	4, // 0: google.cloud.discoveryengine.v1beta.Project.create_time:type_name -> google.protobuf.Timestamp
	4, // 1: google.cloud.discoveryengine.v1beta.Project.provision_completion_time:type_name -> google.protobuf.Timestamp
	3, // 2: google.cloud.discoveryengine.v1beta.Project.service_terms_map:type_name -> google.cloud.discoveryengine.v1beta.Project.ServiceTermsMapEntry
	0, // 3: google.cloud.discoveryengine.v1beta.Project.ServiceTerms.state:type_name -> google.cloud.discoveryengine.v1beta.Project.ServiceTerms.State
	4, // 4: google.cloud.discoveryengine.v1beta.Project.ServiceTerms.accept_time:type_name -> google.protobuf.Timestamp
	4, // 5: google.cloud.discoveryengine.v1beta.Project.ServiceTerms.decline_time:type_name -> google.protobuf.Timestamp
	2, // 6: google.cloud.discoveryengine.v1beta.Project.ServiceTermsMapEntry.value:type_name -> google.cloud.discoveryengine.v1beta.Project.ServiceTerms
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1beta_project_proto_init() }
func file_google_cloud_discoveryengine_v1beta_project_proto_init() {
	if File_google_cloud_discoveryengine_v1beta_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1beta_project_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Project); i {
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
		file_google_cloud_discoveryengine_v1beta_project_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Project_ServiceTerms); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1beta_project_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1beta_project_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1beta_project_proto_depIdxs,
		EnumInfos:         file_google_cloud_discoveryengine_v1beta_project_proto_enumTypes,
		MessageInfos:      file_google_cloud_discoveryengine_v1beta_project_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1beta_project_proto = out.File
	file_google_cloud_discoveryengine_v1beta_project_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1beta_project_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1beta_project_proto_depIdxs = nil
}
