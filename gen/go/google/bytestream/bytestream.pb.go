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
// source: google/bytestream/bytestream.proto

package bytestream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Request object for ByteStream.Read.
type ReadRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the resource to read.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The offset for the first byte to return in the read, relative to the start
	// of the resource.
	//
	// A `read_offset` that is negative or greater than the size of the resource
	// will cause an `OUT_OF_RANGE` error.
	ReadOffset int64 `protobuf:"varint,2,opt,name=read_offset,json=readOffset,proto3" json:"read_offset,omitempty"`
	// The maximum number of `data` bytes the server is allowed to return in the
	// sum of all `ReadResponse` messages. A `read_limit` of zero indicates that
	// there is no limit, and a negative `read_limit` will cause an error.
	//
	// If the stream returns fewer bytes than allowed by the `read_limit` and no
	// error occurred, the stream includes all data from the `read_offset` to the
	// end of the resource.
	ReadLimit     int64 `protobuf:"varint,3,opt,name=read_limit,json=readLimit,proto3" json:"read_limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	mi := &file_google_bytestream_bytestream_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_bytestream_bytestream_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_google_bytestream_bytestream_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ReadRequest) GetReadOffset() int64 {
	if x != nil {
		return x.ReadOffset
	}
	return 0
}

func (x *ReadRequest) GetReadLimit() int64 {
	if x != nil {
		return x.ReadLimit
	}
	return 0
}

// Response object for ByteStream.Read.
type ReadResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A portion of the data for the resource. The service **may** leave `data`
	// empty for any given `ReadResponse`. This enables the service to inform the
	// client that the request is still live while it is running an operation to
	// generate more data.
	Data          []byte `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	mi := &file_google_bytestream_bytestream_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_bytestream_bytestream_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_google_bytestream_bytestream_proto_rawDescGZIP(), []int{1}
}

func (x *ReadResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Request object for ByteStream.Write.
type WriteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the resource to write. This **must** be set on the first
	// `WriteRequest` of each `Write()` action. If it is set on subsequent calls,
	// it **must** match the value of the first request.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The offset from the beginning of the resource at which the data should be
	// written. It is required on all `WriteRequest`s.
	//
	// In the first `WriteRequest` of a `Write()` action, it indicates
	// the initial offset for the `Write()` call. The value **must** be equal to
	// the `committed_size` that a call to `QueryWriteStatus()` would return.
	//
	// On subsequent calls, this value **must** be set and **must** be equal to
	// the sum of the first `write_offset` and the sizes of all `data` bundles
	// sent previously on this stream.
	//
	// An incorrect value will cause an error.
	WriteOffset int64 `protobuf:"varint,2,opt,name=write_offset,json=writeOffset,proto3" json:"write_offset,omitempty"`
	// If `true`, this indicates that the write is complete. Sending any
	// `WriteRequest`s subsequent to one in which `finish_write` is `true` will
	// cause an error.
	FinishWrite bool `protobuf:"varint,3,opt,name=finish_write,json=finishWrite,proto3" json:"finish_write,omitempty"`
	// A portion of the data for the resource. The client **may** leave `data`
	// empty for any given `WriteRequest`. This enables the client to inform the
	// service that the request is still live while it is running an operation to
	// generate more data.
	Data          []byte `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	mi := &file_google_bytestream_bytestream_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_bytestream_bytestream_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_google_bytestream_bytestream_proto_rawDescGZIP(), []int{2}
}

func (x *WriteRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *WriteRequest) GetWriteOffset() int64 {
	if x != nil {
		return x.WriteOffset
	}
	return 0
}

func (x *WriteRequest) GetFinishWrite() bool {
	if x != nil {
		return x.FinishWrite
	}
	return false
}

func (x *WriteRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Response object for ByteStream.Write.
type WriteResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The number of bytes that have been processed for the given resource.
	CommittedSize int64 `protobuf:"varint,1,opt,name=committed_size,json=committedSize,proto3" json:"committed_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	mi := &file_google_bytestream_bytestream_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_bytestream_bytestream_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_google_bytestream_bytestream_proto_rawDescGZIP(), []int{3}
}

func (x *WriteResponse) GetCommittedSize() int64 {
	if x != nil {
		return x.CommittedSize
	}
	return 0
}

// Request object for ByteStream.QueryWriteStatus.
type QueryWriteStatusRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the resource whose write status is being requested.
	ResourceName  string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryWriteStatusRequest) Reset() {
	*x = QueryWriteStatusRequest{}
	mi := &file_google_bytestream_bytestream_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryWriteStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryWriteStatusRequest) ProtoMessage() {}

func (x *QueryWriteStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_bytestream_bytestream_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryWriteStatusRequest.ProtoReflect.Descriptor instead.
func (*QueryWriteStatusRequest) Descriptor() ([]byte, []int) {
	return file_google_bytestream_bytestream_proto_rawDescGZIP(), []int{4}
}

func (x *QueryWriteStatusRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Response object for ByteStream.QueryWriteStatus.
type QueryWriteStatusResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The number of bytes that have been processed for the given resource.
	CommittedSize int64 `protobuf:"varint,1,opt,name=committed_size,json=committedSize,proto3" json:"committed_size,omitempty"`
	// `complete` is `true` only if the client has sent a `WriteRequest` with
	// `finish_write` set to true, and the server has processed that request.
	Complete      bool `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryWriteStatusResponse) Reset() {
	*x = QueryWriteStatusResponse{}
	mi := &file_google_bytestream_bytestream_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryWriteStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryWriteStatusResponse) ProtoMessage() {}

func (x *QueryWriteStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_bytestream_bytestream_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryWriteStatusResponse.ProtoReflect.Descriptor instead.
func (*QueryWriteStatusResponse) Descriptor() ([]byte, []int) {
	return file_google_bytestream_bytestream_proto_rawDescGZIP(), []int{5}
}

func (x *QueryWriteStatusResponse) GetCommittedSize() int64 {
	if x != nil {
		return x.CommittedSize
	}
	return 0
}

func (x *QueryWriteStatusResponse) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

var File_google_bytestream_bytestream_proto protoreflect.FileDescriptor

const file_google_bytestream_bytestream_proto_rawDesc = "" +
	"\n" +
	"\"google/bytestream/bytestream.proto\x12\x11google.bytestream\"r\n" +
	"\vReadRequest\x12#\n" +
	"\rresource_name\x18\x01 \x01(\tR\fresourceName\x12\x1f\n" +
	"\vread_offset\x18\x02 \x01(\x03R\n" +
	"readOffset\x12\x1d\n" +
	"\n" +
	"read_limit\x18\x03 \x01(\x03R\treadLimit\"\"\n" +
	"\fReadResponse\x12\x12\n" +
	"\x04data\x18\n" +
	" \x01(\fR\x04data\"\x8d\x01\n" +
	"\fWriteRequest\x12#\n" +
	"\rresource_name\x18\x01 \x01(\tR\fresourceName\x12!\n" +
	"\fwrite_offset\x18\x02 \x01(\x03R\vwriteOffset\x12!\n" +
	"\ffinish_write\x18\x03 \x01(\bR\vfinishWrite\x12\x12\n" +
	"\x04data\x18\n" +
	" \x01(\fR\x04data\"6\n" +
	"\rWriteResponse\x12%\n" +
	"\x0ecommitted_size\x18\x01 \x01(\x03R\rcommittedSize\">\n" +
	"\x17QueryWriteStatusRequest\x12#\n" +
	"\rresource_name\x18\x01 \x01(\tR\fresourceName\"]\n" +
	"\x18QueryWriteStatusResponse\x12%\n" +
	"\x0ecommitted_size\x18\x01 \x01(\x03R\rcommittedSize\x12\x1a\n" +
	"\bcomplete\x18\x02 \x01(\bR\bcomplete2\x92\x02\n" +
	"\n" +
	"ByteStream\x12I\n" +
	"\x04Read\x12\x1e.google.bytestream.ReadRequest\x1a\x1f.google.bytestream.ReadResponse0\x01\x12L\n" +
	"\x05Write\x12\x1f.google.bytestream.WriteRequest\x1a .google.bytestream.WriteResponse(\x01\x12k\n" +
	"\x10QueryWriteStatus\x12*.google.bytestream.QueryWriteStatusRequest\x1a+.google.bytestream.QueryWriteStatusResponseB\xcb\x01\n" +
	"\x15com.google.bytestreamB\x0fBytestreamProtoP\x01Z<github.com/spounge-ai/spounge-proto/gen/go/google/bytestream\xa2\x02\x03GBX\xaa\x02\x11Google.Bytestream\xca\x02\x11Google\\Bytestream\xe2\x02\x1dGoogle\\Bytestream\\GPBMetadata\xea\x02\x12Google::Bytestreamb\x06proto3"

var (
	file_google_bytestream_bytestream_proto_rawDescOnce sync.Once
	file_google_bytestream_bytestream_proto_rawDescData []byte
)

func file_google_bytestream_bytestream_proto_rawDescGZIP() []byte {
	file_google_bytestream_bytestream_proto_rawDescOnce.Do(func() {
		file_google_bytestream_bytestream_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_bytestream_bytestream_proto_rawDesc), len(file_google_bytestream_bytestream_proto_rawDesc)))
	})
	return file_google_bytestream_bytestream_proto_rawDescData
}

var file_google_bytestream_bytestream_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_bytestream_bytestream_proto_goTypes = []any{
	(*ReadRequest)(nil),              // 0: google.bytestream.ReadRequest
	(*ReadResponse)(nil),             // 1: google.bytestream.ReadResponse
	(*WriteRequest)(nil),             // 2: google.bytestream.WriteRequest
	(*WriteResponse)(nil),            // 3: google.bytestream.WriteResponse
	(*QueryWriteStatusRequest)(nil),  // 4: google.bytestream.QueryWriteStatusRequest
	(*QueryWriteStatusResponse)(nil), // 5: google.bytestream.QueryWriteStatusResponse
}
var file_google_bytestream_bytestream_proto_depIdxs = []int32{
	0, // 0: google.bytestream.ByteStream.Read:input_type -> google.bytestream.ReadRequest
	2, // 1: google.bytestream.ByteStream.Write:input_type -> google.bytestream.WriteRequest
	4, // 2: google.bytestream.ByteStream.QueryWriteStatus:input_type -> google.bytestream.QueryWriteStatusRequest
	1, // 3: google.bytestream.ByteStream.Read:output_type -> google.bytestream.ReadResponse
	3, // 4: google.bytestream.ByteStream.Write:output_type -> google.bytestream.WriteResponse
	5, // 5: google.bytestream.ByteStream.QueryWriteStatus:output_type -> google.bytestream.QueryWriteStatusResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_bytestream_bytestream_proto_init() }
func file_google_bytestream_bytestream_proto_init() {
	if File_google_bytestream_bytestream_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_bytestream_bytestream_proto_rawDesc), len(file_google_bytestream_bytestream_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_bytestream_bytestream_proto_goTypes,
		DependencyIndexes: file_google_bytestream_bytestream_proto_depIdxs,
		MessageInfos:      file_google_bytestream_bytestream_proto_msgTypes,
	}.Build()
	File_google_bytestream_bytestream_proto = out.File
	file_google_bytestream_bytestream_proto_goTypes = nil
	file_google_bytestream_bytestream_proto_depIdxs = nil
}
