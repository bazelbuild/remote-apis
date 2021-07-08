// Copyright 2020 The Bazel Authors.
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

// Log Stream API

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: build/bazel/remote/logstream/v1/remote_logstream.proto

package remotelogstream

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Contains all information necessary to create a new LogStream resource.
type CreateLogStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource of the created LogStream.
	// The list of valid types of parent resources of LogStreams is up to the
	// implementing server.
	// Example: projects/123
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *CreateLogStreamRequest) Reset() {
	*x = CreateLogStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_bazel_remote_logstream_v1_remote_logstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogStreamRequest) ProtoMessage() {}

func (x *CreateLogStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_build_bazel_remote_logstream_v1_remote_logstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogStreamRequest.ProtoReflect.Descriptor instead.
func (*CreateLogStreamRequest) Descriptor() ([]byte, []int) {
	return file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLogStreamRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

// A handle to a log (an ordered sequence of bytes).
type LogStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Structured name of the resource in the format:
	//   {parent=**}/logstreams/{logstream_id}
	//   Example: projects/123/logstreams/456-def
	// Attempting to call the Byte Stream API's `Write` RPC with a LogStream's
	//   `name` as the value for `ByteStream.Write.resource_name` is an error.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Resource name to pass to `ByteStream.Write` in the format:
	//   {parent=**}/logstreams/{logstream_id}/{write_token}
	//   Example: projects/123/logstreams/456-def/789-ghi
	// Attempting to call the Byte Stream API's `Read` RPC with a LogStream's
	//   `write_resource_name` as the value for `ByteStream.Write.resource_name`
	//   is an error.
	//
	// `write_resource_name` is separate from `name` to ensure that only the
	// intended writers can write to a given LogStream. Writers must address write
	// operations to the `write_resource_name`, not the `name`, and must have
	// permission to write LogStreams. `write_resource_name` embeds a secret token
	// and should be protected accordingly; a mishandled `write_resource_name` can
	// result in unintended writers corrupting the LogStream. Therefore, the field
	// should be excluded from calls to any calls which retrieve LogStream
	// metadata (i.e.: `GetLogStream`).
	//
	// Bytes written to this resource must to be readable when `ByteStream.Read`
	// is called with the `name` resource.
	// Reading a write_resource_name must return an INVALID_ARGUMENT error.
	WriteResourceName string `protobuf:"bytes,2,opt,name=write_resource_name,json=writeResourceName,proto3" json:"write_resource_name,omitempty"`
}

func (x *LogStream) Reset() {
	*x = LogStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_bazel_remote_logstream_v1_remote_logstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogStream) ProtoMessage() {}

func (x *LogStream) ProtoReflect() protoreflect.Message {
	mi := &file_build_bazel_remote_logstream_v1_remote_logstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogStream.ProtoReflect.Descriptor instead.
func (*LogStream) Descriptor() ([]byte, []int) {
	return file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDescGZIP(), []int{1}
}

func (x *LogStream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LogStream) GetWriteResourceName() string {
	if x != nil {
		return x.WriteResourceName
	}
	return ""
}

var File_build_bazel_remote_logstream_v1_remote_logstream_proto protoreflect.FileDescriptor

var file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDesc = []byte{
	0x0a, 0x36, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x6c, 0x6f, 0x67,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0x30, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x4f, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x8c, 0x01, 0x0a,
	0x10, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x78, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x37, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a,
	0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x00, 0x42, 0x71, 0x0a, 0x1f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x14,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x6c, 0x6f,
	0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xa2, 0x02, 0x02, 0x52, 0x4c, 0xaa, 0x02, 0x1f, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDescOnce sync.Once
	file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDescData = file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDesc
)

func file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDescGZIP() []byte {
	file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDescOnce.Do(func() {
		file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDescData)
	})
	return file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDescData
}

var file_build_bazel_remote_logstream_v1_remote_logstream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_build_bazel_remote_logstream_v1_remote_logstream_proto_goTypes = []interface{}{
	(*CreateLogStreamRequest)(nil), // 0: build.bazel.remote.logstream.v1.CreateLogStreamRequest
	(*LogStream)(nil),              // 1: build.bazel.remote.logstream.v1.LogStream
}
var file_build_bazel_remote_logstream_v1_remote_logstream_proto_depIdxs = []int32{
	0, // 0: build.bazel.remote.logstream.v1.LogStreamService.CreateLogStream:input_type -> build.bazel.remote.logstream.v1.CreateLogStreamRequest
	1, // 1: build.bazel.remote.logstream.v1.LogStreamService.CreateLogStream:output_type -> build.bazel.remote.logstream.v1.LogStream
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_build_bazel_remote_logstream_v1_remote_logstream_proto_init() }
func file_build_bazel_remote_logstream_v1_remote_logstream_proto_init() {
	if File_build_bazel_remote_logstream_v1_remote_logstream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_build_bazel_remote_logstream_v1_remote_logstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogStreamRequest); i {
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
		file_build_bazel_remote_logstream_v1_remote_logstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogStream); i {
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
			RawDescriptor: file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_build_bazel_remote_logstream_v1_remote_logstream_proto_goTypes,
		DependencyIndexes: file_build_bazel_remote_logstream_v1_remote_logstream_proto_depIdxs,
		MessageInfos:      file_build_bazel_remote_logstream_v1_remote_logstream_proto_msgTypes,
	}.Build()
	File_build_bazel_remote_logstream_v1_remote_logstream_proto = out.File
	file_build_bazel_remote_logstream_v1_remote_logstream_proto_rawDesc = nil
	file_build_bazel_remote_logstream_v1_remote_logstream_proto_goTypes = nil
	file_build_bazel_remote_logstream_v1_remote_logstream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogStreamServiceClient is the client API for LogStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogStreamServiceClient interface {
	// Create a LogStream which may be written to.
	//
	// The returned LogStream resource name will include a `write_resource_name`
	// which is the resource to use when writing to the LogStream.
	// Callers of CreateLogStream are expected to NOT publish the
	// `write_resource_name`.
	CreateLogStream(ctx context.Context, in *CreateLogStreamRequest, opts ...grpc.CallOption) (*LogStream, error)
}

type logStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogStreamServiceClient(cc grpc.ClientConnInterface) LogStreamServiceClient {
	return &logStreamServiceClient{cc}
}

func (c *logStreamServiceClient) CreateLogStream(ctx context.Context, in *CreateLogStreamRequest, opts ...grpc.CallOption) (*LogStream, error) {
	out := new(LogStream)
	err := c.cc.Invoke(ctx, "/build.bazel.remote.logstream.v1.LogStreamService/CreateLogStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogStreamServiceServer is the server API for LogStreamService service.
type LogStreamServiceServer interface {
	// Create a LogStream which may be written to.
	//
	// The returned LogStream resource name will include a `write_resource_name`
	// which is the resource to use when writing to the LogStream.
	// Callers of CreateLogStream are expected to NOT publish the
	// `write_resource_name`.
	CreateLogStream(context.Context, *CreateLogStreamRequest) (*LogStream, error)
}

// UnimplementedLogStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogStreamServiceServer struct {
}

func (*UnimplementedLogStreamServiceServer) CreateLogStream(context.Context, *CreateLogStreamRequest) (*LogStream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLogStream not implemented")
}

func RegisterLogStreamServiceServer(s *grpc.Server, srv LogStreamServiceServer) {
	s.RegisterService(&_LogStreamService_serviceDesc, srv)
}

func _LogStreamService_CreateLogStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLogStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogStreamServiceServer).CreateLogStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build.bazel.remote.logstream.v1.LogStreamService/CreateLogStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogStreamServiceServer).CreateLogStream(ctx, req.(*CreateLogStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "build.bazel.remote.logstream.v1.LogStreamService",
	HandlerType: (*LogStreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLogStream",
			Handler:    _LogStreamService_CreateLogStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "build/bazel/remote/logstream/v1/remote_logstream.proto",
}
