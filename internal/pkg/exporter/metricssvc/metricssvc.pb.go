//
//Copyright (c) Advanced Micro Devices, Inc. All rights reserved.
//
//Licensed under the Apache License, Version 2.0 (the \"License\");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an \"AS IS\" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

//----------------------------------------------------------------------------
///
/// \file
/// AMD Metrics Services
///
//----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: metricssvc.proto

package metricssvc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GPUHealth int32

const (
	GPUHealth_UNKNOWN   GPUHealth = 0
	GPUHealth_HEALTHY   GPUHealth = 1
	GPUHealth_UNHEALTHY GPUHealth = 2
)

// Enum value maps for GPUHealth.
var (
	GPUHealth_name = map[int32]string{
		0: "UNKNOWN",
		1: "HEALTHY",
		2: "UNHEALTHY",
	}
	GPUHealth_value = map[string]int32{
		"UNKNOWN":   0,
		"HEALTHY":   1,
		"UNHEALTHY": 2,
	}
)

func (x GPUHealth) Enum() *GPUHealth {
	p := new(GPUHealth)
	*p = x
	return p
}

func (x GPUHealth) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GPUHealth) Descriptor() protoreflect.EnumDescriptor {
	return file_metricssvc_proto_enumTypes[0].Descriptor()
}

func (GPUHealth) Type() protoreflect.EnumType {
	return &file_metricssvc_proto_enumTypes[0]
}

func (x GPUHealth) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GPUHealth.Descriptor instead.
func (GPUHealth) EnumDescriptor() ([]byte, []int) {
	return file_metricssvc_proto_rawDescGZIP(), []int{0}
}

type GPUState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the GPU
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// uuid of the GPU
	UUID string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	// health state string value of GPUHealth enum
	Health string `protobuf:"bytes,3,opt,name=Health,proto3" json:"Health,omitempty"`
	// Workload associated with GPU
	AssociatedWorkload []string `protobuf:"bytes,4,rep,name=AssociatedWorkload,proto3" json:"AssociatedWorkload,omitempty"`
	// PCIe Bus ID refers to device ID in amd device plugin
	Device string `protobuf:"bytes,5,opt,name=Device,proto3" json:"Device,omitempty"`
}

func (x *GPUState) Reset() {
	*x = GPUState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricssvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPUState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPUState) ProtoMessage() {}

func (x *GPUState) ProtoReflect() protoreflect.Message {
	mi := &file_metricssvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPUState.ProtoReflect.Descriptor instead.
func (*GPUState) Descriptor() ([]byte, []int) {
	return file_metricssvc_proto_rawDescGZIP(), []int{0}
}

func (x *GPUState) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *GPUState) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *GPUState) GetHealth() string {
	if x != nil {
		return x.Health
	}
	return ""
}

func (x *GPUState) GetAssociatedWorkload() []string {
	if x != nil {
		return x.AssociatedWorkload
	}
	return nil
}

func (x *GPUState) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

type GPUGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of id of the GPU
	ID []string `protobuf:"bytes,1,rep,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GPUGetRequest) Reset() {
	*x = GPUGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricssvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPUGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPUGetRequest) ProtoMessage() {}

func (x *GPUGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metricssvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPUGetRequest.ProtoReflect.Descriptor instead.
func (*GPUGetRequest) Descriptor() ([]byte, []int) {
	return file_metricssvc_proto_rawDescGZIP(), []int{1}
}

func (x *GPUGetRequest) GetID() []string {
	if x != nil {
		return x.ID
	}
	return nil
}

type GPUUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of id of the GPU
	ID []string `protobuf:"bytes,1,rep,name=ID,proto3" json:"ID,omitempty"`
	// set health state string value of GPUHealth enum
	Health []string `protobuf:"bytes,2,rep,name=Health,proto3" json:"Health,omitempty"`
}

func (x *GPUUpdateRequest) Reset() {
	*x = GPUUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricssvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPUUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPUUpdateRequest) ProtoMessage() {}

func (x *GPUUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metricssvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPUUpdateRequest.ProtoReflect.Descriptor instead.
func (*GPUUpdateRequest) Descriptor() ([]byte, []int) {
	return file_metricssvc_proto_rawDescGZIP(), []int{2}
}

func (x *GPUUpdateRequest) GetID() []string {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *GPUUpdateRequest) GetHealth() []string {
	if x != nil {
		return x.Health
	}
	return nil
}

type GPUStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of requested GPU States
	GPUState []*GPUState `protobuf:"bytes,1,rep,name=GPUState,proto3" json:"GPUState,omitempty"`
}

func (x *GPUStateResponse) Reset() {
	*x = GPUStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricssvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPUStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPUStateResponse) ProtoMessage() {}

func (x *GPUStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metricssvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPUStateResponse.ProtoReflect.Descriptor instead.
func (*GPUStateResponse) Descriptor() ([]byte, []int) {
	return file_metricssvc_proto_rawDescGZIP(), []int{3}
}

func (x *GPUStateResponse) GetGPUState() []*GPUState {
	if x != nil {
		return x.GPUState
	}
	return nil
}

var File_metricssvc_proto protoreflect.FileDescriptor

var file_metricssvc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x76, 0x63, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x08,
	0x47, 0x50, 0x55, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x12, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74,
	0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x12, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x1f, 0x0a, 0x0d,
	0x47, 0x50, 0x55, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x3a, 0x0a,
	0x10, 0x47, 0x50, 0x55, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x50, 0x55,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x08, 0x47, 0x50, 0x55, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x50, 0x55,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x08, 0x47, 0x50, 0x55, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2a,
	0x34, 0x0a, 0x09, 0x47, 0x50, 0x55, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41,
	0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x48, 0x45, 0x41, 0x4c,
	0x54, 0x48, 0x59, 0x10, 0x02, 0x32, 0x9a, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47,
	0x50, 0x55, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x50, 0x55, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x76, 0x63, 0x2e,
	0x47, 0x50, 0x55, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x76, 0x63, 0x2e,
	0x47, 0x50, 0x55, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metricssvc_proto_rawDescOnce sync.Once
	file_metricssvc_proto_rawDescData = file_metricssvc_proto_rawDesc
)

func file_metricssvc_proto_rawDescGZIP() []byte {
	file_metricssvc_proto_rawDescOnce.Do(func() {
		file_metricssvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_metricssvc_proto_rawDescData)
	})
	return file_metricssvc_proto_rawDescData
}

var file_metricssvc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metricssvc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_metricssvc_proto_goTypes = []any{
	(GPUHealth)(0),           // 0: metricssvc.GPUHealth
	(*GPUState)(nil),         // 1: metricssvc.GPUState
	(*GPUGetRequest)(nil),    // 2: metricssvc.GPUGetRequest
	(*GPUUpdateRequest)(nil), // 3: metricssvc.GPUUpdateRequest
	(*GPUStateResponse)(nil), // 4: metricssvc.GPUStateResponse
	(*emptypb.Empty)(nil),    // 5: google.protobuf.Empty
}
var file_metricssvc_proto_depIdxs = []int32{
	1, // 0: metricssvc.GPUStateResponse.GPUState:type_name -> metricssvc.GPUState
	2, // 1: metricssvc.MetricsService.GetGPUState:input_type -> metricssvc.GPUGetRequest
	5, // 2: metricssvc.MetricsService.List:input_type -> google.protobuf.Empty
	4, // 3: metricssvc.MetricsService.GetGPUState:output_type -> metricssvc.GPUStateResponse
	4, // 4: metricssvc.MetricsService.List:output_type -> metricssvc.GPUStateResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metricssvc_proto_init() }
func file_metricssvc_proto_init() {
	if File_metricssvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metricssvc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GPUState); i {
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
		file_metricssvc_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GPUGetRequest); i {
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
		file_metricssvc_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GPUUpdateRequest); i {
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
		file_metricssvc_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GPUStateResponse); i {
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
			RawDescriptor: file_metricssvc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metricssvc_proto_goTypes,
		DependencyIndexes: file_metricssvc_proto_depIdxs,
		EnumInfos:         file_metricssvc_proto_enumTypes,
		MessageInfos:      file_metricssvc_proto_msgTypes,
	}.Build()
	File_metricssvc_proto = out.File
	file_metricssvc_proto_rawDesc = nil
	file_metricssvc_proto_goTypes = nil
	file_metricssvc_proto_depIdxs = nil
}