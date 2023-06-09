// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: userServer.proto

package UserServer

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

const (
	UserServer_CreateUser_FullMethodName = "/UserServer.UserServer/CreateUser"
	UserServer_GetUser_FullMethodName    = "/UserServer.UserServer/GetUser"
)

// UserServerClient is the client API for UserServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServerClient interface {
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
}

type userServerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServerClient(cc grpc.ClientConnInterface) UserServerClient {
	return &userServerClient{cc}
}

func (c *userServerClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error) {
	out := new(CreateUserResp)
	err := c.cc.Invoke(ctx, UserServer_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServerClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, UserServer_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServerServer is the server API for UserServer service.
// All implementations must embed UnimplementedUserServerServer
// for forward compatibility
type UserServerServer interface {
	CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error)
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
	mustEmbedUnimplementedUserServerServer()
}

// UnimplementedUserServerServer must be embedded to have forward compatible implementations.
type UnimplementedUserServerServer struct {
}

func (UnimplementedUserServerServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServerServer) GetUser(context.Context, *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServerServer) mustEmbedUnimplementedUserServerServer() {}

// UnsafeUserServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServerServer will
// result in compilation errors.
type UnsafeUserServerServer interface {
	mustEmbedUnimplementedUserServerServer()
}

func RegisterUserServerServer(s grpc.ServiceRegistrar, srv UserServerServer) {
	s.RegisterService(&UserServer_ServiceDesc, srv)
}

func _UserServer_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServer_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServer_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServerServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServer_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServerServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServer_ServiceDesc is the grpc.ServiceDesc for UserServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserServer.UserServer",
	HandlerType: (*UserServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserServer_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserServer_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userServer.proto",
}
