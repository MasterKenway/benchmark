//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: samples_api.proto

package api

import (
	reflect "reflect"
	sync "sync"
)

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samples_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_samples_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_samples_api_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x71, 0xf1, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x1a, 0x17, 0x17, 0x54, 0x4d, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71,
	0x62, 0x3a, 0x4c, 0x11, 0x8c, 0x6b, 0xe4, 0xc9, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a,
	0x24, 0x64, 0xc7, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0xd6, 0x25, 0x24, 0xa1, 0x87, 0xe4, 0x02, 0x64,
	0xcb, 0xa4, 0xc4, 0xb0, 0xc8, 0x14, 0xe4, 0x54, 0x2a, 0x31, 0x38, 0x99, 0x71, 0x49, 0x67, 0xe6,
	0xeb, 0xa5, 0x17, 0x15, 0x24, 0xeb, 0xa5, 0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4, 0x16, 0x23, 0xa9,
	0x75, 0xe2, 0x07, 0x2b, 0x0e, 0x07, 0xb1, 0x03, 0x40, 0x5e, 0x0a, 0x60, 0x5c, 0xc4, 0xc4, 0xec,
	0xe1, 0x13, 0x9e, 0xc4, 0x06, 0xf6, 0xa1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x0a, 0xdc,
	0xe8, 0xf5, 0x00, 0x00, 0x00,
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The response message containing the greetings
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samples_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_samples_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_samples_api_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_samples_api_proto protoreflect.FileDescriptor

var file_samples_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x64, 0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x32, 0xd7, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x12, 0x60, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x2c, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f,
	0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f,
	0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_samples_api_proto_rawDescOnce sync.Once
	file_samples_api_proto_rawDescData = file_samples_api_proto_rawDesc
)

func file_samples_api_proto_rawDescGZIP() []byte {
	file_samples_api_proto_rawDescOnce.Do(func() {
		file_samples_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_samples_api_proto_rawDescData)
	})
	return file_samples_api_proto_rawDescData
}

var file_samples_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_samples_api_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: org.apache.dubbogo.samples.api.HelloRequest
	(*User)(nil),         // 1: org.apache.dubbogo.samples.api.User
}
var file_samples_api_proto_depIdxs = []int32{
	0, // 0: org.apache.dubbogo.samples.api.Greeter.SayHello:input_type -> org.apache.dubbogo.samples.api.HelloRequest
	0, // 1: org.apache.dubbogo.samples.api.Greeter.SayHelloStream:input_type -> org.apache.dubbogo.samples.api.HelloRequest
	1, // 2: org.apache.dubbogo.samples.api.Greeter.SayHello:output_type -> org.apache.dubbogo.samples.api.User
	1, // 3: org.apache.dubbogo.samples.api.Greeter.SayHelloStream:output_type -> org.apache.dubbogo.samples.api.User
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_samples_api_proto_init() }
func file_samples_api_proto_init() {
	if File_samples_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samples_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_samples_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_samples_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_samples_api_proto_goTypes,
		DependencyIndexes: file_samples_api_proto_depIdxs,
		MessageInfos:      file_samples_api_proto_msgTypes,
	}.Build()
	File_samples_api_proto = out.File
	file_samples_api_proto_rawDesc = nil
	file_samples_api_proto_goTypes = nil
	file_samples_api_proto_depIdxs = nil
}