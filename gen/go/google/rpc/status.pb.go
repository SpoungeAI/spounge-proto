// Copyright 2025 Google LLC
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
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: google/rpc/status.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The `Status` type defines a logical error model that is suitable for
// different programming environments, including REST APIs and RPC APIs. It is
// used by [gRPC](https://github.com/grpc). Each `Status` message contains
// three pieces of data: error code, error message, and error details.
//
// You can find out more about this error model and how to work with it in the
// [API Design Guide](https://cloud.google.com/apis/design/errors).
type Status struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	// by the client.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A list of messages that carry the error details.  There is a common set of
	// message types for APIs to use.
	Details       []*anypb.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_google_rpc_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_google_rpc_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_google_rpc_status_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_google_rpc_status_proto protoreflect.FileDescriptor

const file_google_rpc_status_proto_rawDesc = "" +
	"\n" +
	"\x17google/rpc/status.proto\x12\n" +
	"google.rpc\x1a\x19google/protobuf/any.proto\"f\n" +
	"\x06Status\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12.\n" +
	"\adetails\x18\x03 \x03(\v2\x14.google.protobuf.AnyR\adetailsB\xa0\x01\n" +
	"\x0ecom.google.rpcB\vStatusProtoP\x01Z5github.com/spounge-ai/spounge-proto/gen/go/google/rpc\xf8\x01\x01\xa2\x02\x03GRX\xaa\x02\n" +
	"Google.Rpc\xca\x02\n" +
	"Google\\Rpc\xe2\x02\x16Google\\Rpc\\GPBMetadata\xea\x02\vGoogle::Rpcb\x06proto3"

var (
	file_google_rpc_status_proto_rawDescOnce sync.Once
	file_google_rpc_status_proto_rawDescData []byte
)

func file_google_rpc_status_proto_rawDescGZIP() []byte {
	file_google_rpc_status_proto_rawDescOnce.Do(func() {
		file_google_rpc_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_rpc_status_proto_rawDesc), len(file_google_rpc_status_proto_rawDesc)))
	})
	return file_google_rpc_status_proto_rawDescData
}

var file_google_rpc_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_rpc_status_proto_goTypes = []any{
	(*Status)(nil),    // 0: google.rpc.Status
	(*anypb.Any)(nil), // 1: google.protobuf.Any
}
var file_google_rpc_status_proto_depIdxs = []int32{
	1, // 0: google.rpc.Status.details:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_rpc_status_proto_init() }
func file_google_rpc_status_proto_init() {
	if File_google_rpc_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_rpc_status_proto_rawDesc), len(file_google_rpc_status_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_rpc_status_proto_goTypes,
		DependencyIndexes: file_google_rpc_status_proto_depIdxs,
		MessageInfos:      file_google_rpc_status_proto_msgTypes,
	}.Build()
	File_google_rpc_status_proto = out.File
	file_google_rpc_status_proto_goTypes = nil
	file_google_rpc_status_proto_depIdxs = nil
}
