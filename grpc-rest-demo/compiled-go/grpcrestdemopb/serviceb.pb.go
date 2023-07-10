// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: services/serviceb.proto

package grpcrestdemopb

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

type PongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PongRequest) Reset() {
	*x = PongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_serviceb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongRequest) ProtoMessage() {}

func (x *PongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_serviceb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongRequest.ProtoReflect.Descriptor instead.
func (*PongRequest) Descriptor() ([]byte, []int) {
	return file_services_serviceb_proto_rawDescGZIP(), []int{0}
}

func (x *PongRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Out string `protobuf:"bytes,1,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *PongResponse) Reset() {
	*x = PongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_serviceb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongResponse) ProtoMessage() {}

func (x *PongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_serviceb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongResponse.ProtoReflect.Descriptor instead.
func (*PongResponse) Descriptor() ([]byte, []int) {
	return file_services_serviceb_proto_rawDescGZIP(), []int{1}
}

func (x *PongResponse) GetOut() string {
	if x != nil {
		return x.Out
	}
	return ""
}

var File_services_serviceb_proto protoreflect.FileDescriptor

var file_services_serviceb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x64, 0x65, 0x76, 0x2e, 0x6a,
	0x70, 0x6f, 0x6e, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6f, 0x75, 0x74, 0x32, 0x71, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x12, 0x65, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x2c, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x6a, 0x70, 0x6f, 0x6e, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6a, 0x70,
	0x6f, 0x6e, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x72, 0x65, 0x73, 0x74, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_serviceb_proto_rawDescOnce sync.Once
	file_services_serviceb_proto_rawDescData = file_services_serviceb_proto_rawDesc
)

func file_services_serviceb_proto_rawDescGZIP() []byte {
	file_services_serviceb_proto_rawDescOnce.Do(func() {
		file_services_serviceb_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_serviceb_proto_rawDescData)
	})
	return file_services_serviceb_proto_rawDescData
}

var file_services_serviceb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_serviceb_proto_goTypes = []interface{}{
	(*PongRequest)(nil),  // 0: dev.jponc.grpc_rest.serviceb.v1.PongRequest
	(*PongResponse)(nil), // 1: dev.jponc.grpc_rest.serviceb.v1.PongResponse
}
var file_services_serviceb_proto_depIdxs = []int32{
	0, // 0: dev.jponc.grpc_rest.serviceb.v1.ServiceB.Pong:input_type -> dev.jponc.grpc_rest.serviceb.v1.PongRequest
	1, // 1: dev.jponc.grpc_rest.serviceb.v1.ServiceB.Pong:output_type -> dev.jponc.grpc_rest.serviceb.v1.PongResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_serviceb_proto_init() }
func file_services_serviceb_proto_init() {
	if File_services_serviceb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_serviceb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongRequest); i {
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
		file_services_serviceb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongResponse); i {
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
			RawDescriptor: file_services_serviceb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_serviceb_proto_goTypes,
		DependencyIndexes: file_services_serviceb_proto_depIdxs,
		MessageInfos:      file_services_serviceb_proto_msgTypes,
	}.Build()
	File_services_serviceb_proto = out.File
	file_services_serviceb_proto_rawDesc = nil
	file_services_serviceb_proto_goTypes = nil
	file_services_serviceb_proto_depIdxs = nil
}
