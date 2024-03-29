// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: pkg/model/kindling_p9x.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type P9XRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *P9XRequest) Reset() {
	*x = P9XRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_kindling_p9x_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P9XRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P9XRequest) ProtoMessage() {}

func (x *P9XRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_kindling_p9x_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P9XRequest.ProtoReflect.Descriptor instead.
func (*P9XRequest) Descriptor() ([]byte, []int) {
	return file_pkg_model_kindling_p9x_proto_rawDescGZIP(), []int{0}
}

func (x *P9XRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type P9XResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas []*P9XData `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty"`
}

func (x *P9XResponse) Reset() {
	*x = P9XResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_kindling_p9x_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P9XResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P9XResponse) ProtoMessage() {}

func (x *P9XResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_kindling_p9x_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P9XResponse.ProtoReflect.Descriptor instead.
func (*P9XResponse) Descriptor() ([]byte, []int) {
	return file_pkg_model_kindling_p9x_proto_rawDescGZIP(), []int{1}
}

func (x *P9XResponse) GetDatas() []*P9XData {
	if x != nil {
		return x.Datas
	}
	return nil
}

type P9XData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string  `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	ContainerId string  `protobuf:"bytes,2,opt,name=containerId,proto3" json:"containerId,omitempty"`
	Value       float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *P9XData) Reset() {
	*x = P9XData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_kindling_p9x_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P9XData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P9XData) ProtoMessage() {}

func (x *P9XData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_kindling_p9x_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P9XData.ProtoReflect.Descriptor instead.
func (*P9XData) Descriptor() ([]byte, []int) {
	return file_pkg_model_kindling_p9x_proto_rawDescGZIP(), []int{2}
}

func (x *P9XData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *P9XData) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *P9XData) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_pkg_model_kindling_p9x_proto protoreflect.FileDescriptor

var file_pkg_model_kindling_p9x_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6b, 0x69, 0x6e, 0x64,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x39, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x6b, 0x69, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x1c, 0x0a, 0x0a, 0x50, 0x39, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x36, 0x0a, 0x0b, 0x50, 0x39, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6b, 0x69, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x50, 0x39, 0x78, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x22, 0x53,
	0x0a, 0x07, 0x50, 0x39, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x47, 0x0a, 0x0a, 0x50, 0x39, 0x58, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x39, 0x78, 0x12, 0x14, 0x2e,
	0x6b, 0x69, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x39, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6b, 0x69, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50,
	0x39, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_model_kindling_p9x_proto_rawDescOnce sync.Once
	file_pkg_model_kindling_p9x_proto_rawDescData = file_pkg_model_kindling_p9x_proto_rawDesc
)

func file_pkg_model_kindling_p9x_proto_rawDescGZIP() []byte {
	file_pkg_model_kindling_p9x_proto_rawDescOnce.Do(func() {
		file_pkg_model_kindling_p9x_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_model_kindling_p9x_proto_rawDescData)
	})
	return file_pkg_model_kindling_p9x_proto_rawDescData
}

var file_pkg_model_kindling_p9x_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_model_kindling_p9x_proto_goTypes = []interface{}{
	(*P9XRequest)(nil),  // 0: kindling.P9xRequest
	(*P9XResponse)(nil), // 1: kindling.P9xResponse
	(*P9XData)(nil),     // 2: kindling.P9xData
}
var file_pkg_model_kindling_p9x_proto_depIdxs = []int32{
	2, // 0: kindling.P9xResponse.datas:type_name -> kindling.P9xData
	0, // 1: kindling.P9XService.QueryP9x:input_type -> kindling.P9xRequest
	1, // 2: kindling.P9XService.QueryP9x:output_type -> kindling.P9xResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_model_kindling_p9x_proto_init() }
func file_pkg_model_kindling_p9x_proto_init() {
	if File_pkg_model_kindling_p9x_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_model_kindling_p9x_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P9XRequest); i {
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
		file_pkg_model_kindling_p9x_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P9XResponse); i {
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
		file_pkg_model_kindling_p9x_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P9XData); i {
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
			RawDescriptor: file_pkg_model_kindling_p9x_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_model_kindling_p9x_proto_goTypes,
		DependencyIndexes: file_pkg_model_kindling_p9x_proto_depIdxs,
		MessageInfos:      file_pkg_model_kindling_p9x_proto_msgTypes,
	}.Build()
	File_pkg_model_kindling_p9x_proto = out.File
	file_pkg_model_kindling_p9x_proto_rawDesc = nil
	file_pkg_model_kindling_p9x_proto_goTypes = nil
	file_pkg_model_kindling_p9x_proto_depIdxs = nil
}
