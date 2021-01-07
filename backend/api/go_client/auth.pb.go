// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend/api/auth.proto

package go_client // import "github.com/kubeflow/pipelines/backend/api/go_client"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthorizeRequest_Resources int32

const (
	AuthorizeRequest_UNASSIGNED_RESOURCES AuthorizeRequest_Resources = 0
	AuthorizeRequest_VIEWERS              AuthorizeRequest_Resources = 1
)

var AuthorizeRequest_Resources_name = map[int32]string{
	0: "UNASSIGNED_RESOURCES",
	1: "VIEWERS",
}
var AuthorizeRequest_Resources_value = map[string]int32{
	"UNASSIGNED_RESOURCES": 0,
	"VIEWERS":              1,
}

func (x AuthorizeRequest_Resources) String() string {
	return proto.EnumName(AuthorizeRequest_Resources_name, int32(x))
}
func (AuthorizeRequest_Resources) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_auth_b463ef3269931e86, []int{0, 0}
}

type AuthorizeRequest_Verb int32

const (
	AuthorizeRequest_UNASSIGNED_VERB AuthorizeRequest_Verb = 0
	AuthorizeRequest_CREATE          AuthorizeRequest_Verb = 1
	AuthorizeRequest_GET             AuthorizeRequest_Verb = 2
	AuthorizeRequest_DELETE          AuthorizeRequest_Verb = 3
)

var AuthorizeRequest_Verb_name = map[int32]string{
	0: "UNASSIGNED_VERB",
	1: "CREATE",
	2: "GET",
	3: "DELETE",
}
var AuthorizeRequest_Verb_value = map[string]int32{
	"UNASSIGNED_VERB": 0,
	"CREATE":          1,
	"GET":             2,
	"DELETE":          3,
}

func (x AuthorizeRequest_Verb) String() string {
	return proto.EnumName(AuthorizeRequest_Verb_name, int32(x))
}
func (AuthorizeRequest_Verb) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_auth_b463ef3269931e86, []int{0, 1}
}

type AuthorizeRequest struct {
	Namespace            string                     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Resources            AuthorizeRequest_Resources `protobuf:"varint,2,opt,name=resources,proto3,enum=api.AuthorizeRequest_Resources" json:"resources,omitempty"`
	Verb                 AuthorizeRequest_Verb      `protobuf:"varint,3,opt,name=verb,proto3,enum=api.AuthorizeRequest_Verb" json:"verb,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AuthorizeRequest) Reset()         { *m = AuthorizeRequest{} }
func (m *AuthorizeRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizeRequest) ProtoMessage()    {}
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_b463ef3269931e86, []int{0}
}
func (m *AuthorizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeRequest.Unmarshal(m, b)
}
func (m *AuthorizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeRequest.Marshal(b, m, deterministic)
}
func (dst *AuthorizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeRequest.Merge(dst, src)
}
func (m *AuthorizeRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizeRequest.Size(m)
}
func (m *AuthorizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeRequest proto.InternalMessageInfo

func (m *AuthorizeRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *AuthorizeRequest) GetResources() AuthorizeRequest_Resources {
	if m != nil {
		return m.Resources
	}
	return AuthorizeRequest_UNASSIGNED_RESOURCES
}

func (m *AuthorizeRequest) GetVerb() AuthorizeRequest_Verb {
	if m != nil {
		return m.Verb
	}
	return AuthorizeRequest_UNASSIGNED_VERB
}

func init() {
	proto.RegisterType((*AuthorizeRequest)(nil), "api.AuthorizeRequest")
	proto.RegisterEnum("api.AuthorizeRequest_Resources", AuthorizeRequest_Resources_name, AuthorizeRequest_Resources_value)
	proto.RegisterEnum("api.AuthorizeRequest_Verb", AuthorizeRequest_Verb_name, AuthorizeRequest_Verb_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.AuthService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Authorize(context.Context, *AuthorizeRequest) (*empty.Empty, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _AuthService_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/api/auth.proto",
}

func init() { proto.RegisterFile("backend/api/auth.proto", fileDescriptor_auth_b463ef3269931e86) }

var fileDescriptor_auth_b463ef3269931e86 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0x8c, 0x9d, 0x2a, 0xc1, 0x2f, 0x94, 0x9a, 0x6d, 0x29, 0x91, 0x09, 0x6a, 0x94, 0x53, 0x0f,
	0xd4, 0x56, 0xd3, 0x2b, 0x1c, 0x92, 0x76, 0x55, 0x55, 0x82, 0x22, 0xad, 0xd3, 0x20, 0xf5, 0x52,
	0xad, 0xdd, 0x17, 0x67, 0x55, 0xc7, 0x6b, 0xd6, 0xeb, 0x54, 0x70, 0x44, 0xe2, 0x03, 0x80, 0x4f,
	0xe3, 0x17, 0xf8, 0x10, 0xe4, 0x4d, 0x9a, 0x46, 0x90, 0xd3, 0x6a, 0xdf, 0xcc, 0x9b, 0x19, 0x69,
	0x1e, 0xec, 0x47, 0x3c, 0xbe, 0xc3, 0xec, 0x36, 0xe0, 0xb9, 0x08, 0x78, 0xa9, 0xa7, 0x7e, 0xae,
	0xa4, 0x96, 0xa4, 0xce, 0x73, 0xe1, 0x75, 0x12, 0x29, 0x93, 0x14, 0x17, 0x58, 0x96, 0x49, 0xcd,
	0xb5, 0x90, 0x59, 0xb1, 0xa0, 0x78, 0xaf, 0x96, 0xa8, 0xf9, 0x45, 0xe5, 0x24, 0xc0, 0x59, 0xae,
	0xbf, 0x2c, 0xc1, 0x97, 0xeb, 0xba, 0xa8, 0x94, 0x54, 0x4b, 0xe0, 0x8d, 0x79, 0xe2, 0xa3, 0x04,
	0xb3, 0xa3, 0xe2, 0x9e, 0x27, 0x09, 0xaa, 0x40, 0xe6, 0x46, 0xf7, 0x7f, 0x8f, 0xde, 0x0f, 0x1b,
	0xdc, 0x41, 0xa9, 0xa7, 0x52, 0x89, 0xaf, 0xc8, 0xf0, 0x73, 0x89, 0x85, 0x26, 0x1d, 0x70, 0x32,
	0x3e, 0xc3, 0x22, 0xe7, 0x31, 0xb6, 0xad, 0xae, 0x75, 0xe8, 0xb0, 0xc7, 0x01, 0x79, 0x07, 0x8e,
	0xc2, 0x42, 0x96, 0x2a, 0xc6, 0xa2, 0x6d, 0x77, 0xad, 0xc3, 0x67, 0xfd, 0x03, 0x9f, 0xe7, 0xc2,
	0xff, 0x57, 0xc7, 0x67, 0x0f, 0x34, 0xf6, 0xb8, 0x41, 0x7c, 0xd8, 0x9a, 0xa3, 0x8a, 0xda, 0x75,
	0xb3, 0xe9, 0x6d, 0xde, 0x1c, 0xa3, 0x8a, 0x98, 0xe1, 0xf5, 0xfa, 0xe0, 0xac, 0x74, 0x48, 0x1b,
	0xf6, 0xae, 0x2e, 0x07, 0x61, 0x78, 0x71, 0x7e, 0x49, 0xcf, 0x6e, 0x18, 0x0d, 0x3f, 0x5e, 0xb1,
	0x53, 0x1a, 0xba, 0x35, 0xd2, 0x82, 0xe6, 0xf8, 0x82, 0x7e, 0xa2, 0x2c, 0x74, 0xad, 0xde, 0x5b,
	0xd8, 0xaa, 0x14, 0xc8, 0x2e, 0xec, 0xac, 0xd1, 0xc7, 0x94, 0x0d, 0xdd, 0x1a, 0x01, 0x68, 0x9c,
	0x32, 0x3a, 0x18, 0x51, 0xd7, 0x22, 0x4d, 0xa8, 0x9f, 0xd3, 0x91, 0x6b, 0x57, 0xc3, 0x33, 0xfa,
	0x9e, 0x8e, 0xa8, 0x5b, 0xef, 0x23, 0xb4, 0xaa, 0x40, 0x21, 0xaa, 0xb9, 0x88, 0x91, 0x8c, 0xc1,
	0x59, 0xe5, 0x23, 0x2f, 0x36, 0xe6, 0xf5, 0xf6, 0xfd, 0x45, 0x57, 0xfe, 0x43, 0x57, 0x3e, 0xad,
	0xba, 0xea, 0x79, 0xdf, 0x7e, 0xff, 0xf9, 0x65, 0xef, 0x11, 0x52, 0xd5, 0x54, 0x04, 0xf3, 0xe3,
	0x08, 0x35, 0x3f, 0x36, 0x77, 0x30, 0xfc, 0x6e, 0xfd, 0x1c, 0x7c, 0x60, 0x1d, 0x68, 0xde, 0xe2,
	0x84, 0x97, 0xa9, 0x26, 0xcf, 0xc9, 0x0e, 0x6c, 0x7b, 0x2d, 0xe3, 0x10, 0x6a, 0xae, 0xcb, 0xe2,
	0xfa, 0x00, 0x5e, 0x43, 0x63, 0x88, 0x5c, 0xa1, 0x22, 0xbb, 0x4f, 0x6c, 0x6f, 0x9b, 0x2f, 0x9d,
	0x4d, 0x89, 0x5d, 0x3b, 0x7a, 0x0a, 0xb0, 0x22, 0xd4, 0xae, 0x4f, 0x12, 0xa1, 0xa7, 0x65, 0xe4,
	0xc7, 0x72, 0x16, 0xdc, 0x95, 0x11, 0x4e, 0x52, 0x79, 0x1f, 0xe4, 0x22, 0xc7, 0x54, 0x64, 0x58,
	0x04, 0xeb, 0x27, 0x93, 0xc8, 0x9b, 0x38, 0x15, 0x98, 0xe9, 0xa8, 0x61, 0x32, 0x9f, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0x47, 0x2d, 0x41, 0xaa, 0x02, 0x00, 0x00,
}
