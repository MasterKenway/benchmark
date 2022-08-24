// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: hello_grpc.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserProviderClient is the client API for UserProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProviderClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetUserByName(ctx context.Context, in *GetUserByNameRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
}

type userProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProviderClient(cc grpc.ClientConnInterface) UserProviderClient {
	return &userProviderClient{cc}
}

func (c *userProviderClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/provider.UserProvider/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProviderClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/provider.UserProvider/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProviderClient) GetUserByName(ctx context.Context, in *GetUserByNameRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/provider.UserProvider/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProviderServer is the server API for UserProvider service.
// All implementations must embed UnimplementedUserProviderServer
// for forward compatibility
type UserProviderServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUsersResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	GetUserByName(context.Context, *GetUserByNameRequest) (*GetUsersResponse, error)
	mustEmbedUnimplementedUserProviderServer()
}

// UnimplementedUserProviderServer must be embedded to have forward compatible implementations.
type UnimplementedUserProviderServer struct {
}

func (UnimplementedUserProviderServer) GetUser(context.Context, *GetUserRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserProviderServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserProviderServer) GetUserByName(context.Context, *GetUserByNameRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedUserProviderServer) mustEmbedUnimplementedUserProviderServer() {}

// UnsafeUserProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProviderServer will
// result in compilation errors.
type UnsafeUserProviderServer interface {
	mustEmbedUnimplementedUserProviderServer()
}

func RegisterUserProviderServer(s grpc.ServiceRegistrar, srv UserProviderServer) {
	s.RegisterService(&UserProvider_ServiceDesc, srv)
}

func _UserProvider_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProviderServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.UserProvider/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProviderServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProvider_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProviderServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.UserProvider/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProviderServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProvider_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProviderServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.UserProvider/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProviderServer).GetUserByName(ctx, req.(*GetUserByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserProvider_ServiceDesc is the grpc.ServiceDesc for UserProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "provider.UserProvider",
	HandlerType: (*UserProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserProvider_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserProvider_GetUsers_Handler,
		},
		{
			MethodName: "GetUserByName",
			Handler:    _UserProvider_GetUserByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_grpc.proto",
}
