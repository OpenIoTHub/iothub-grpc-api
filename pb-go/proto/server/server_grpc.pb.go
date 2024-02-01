// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/server/server.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HttpManagerClient is the client API for HttpManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HttpManagerClient interface {
	CreateOneHTTP(ctx context.Context, in *HTTPConfig, opts ...grpc.CallOption) (*HTTPConfig, error)
	UpdateOneHTTP(ctx context.Context, in *HTTPConfig, opts ...grpc.CallOption) (*HTTPConfig, error)
	DeleteOneHTTP(ctx context.Context, in *HTTPConfig, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetOneHTTP(ctx context.Context, in *HTTPConfig, opts ...grpc.CallOption) (*HTTPConfig, error)
	GetAllHTTP(ctx context.Context, in *Device, opts ...grpc.CallOption) (*HTTPList, error)
}

type httpManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewHttpManagerClient(cc grpc.ClientConnInterface) HttpManagerClient {
	return &httpManagerClient{cc}
}

func (c *httpManagerClient) CreateOneHTTP(ctx context.Context, in *HTTPConfig, opts ...grpc.CallOption) (*HTTPConfig, error) {
	out := new(HTTPConfig)
	err := c.cc.Invoke(ctx, "/pb.HttpManager/CreateOneHTTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpManagerClient) UpdateOneHTTP(ctx context.Context, in *HTTPConfig, opts ...grpc.CallOption) (*HTTPConfig, error) {
	out := new(HTTPConfig)
	err := c.cc.Invoke(ctx, "/pb.HttpManager/UpdateOneHTTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpManagerClient) DeleteOneHTTP(ctx context.Context, in *HTTPConfig, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.HttpManager/DeleteOneHTTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpManagerClient) GetOneHTTP(ctx context.Context, in *HTTPConfig, opts ...grpc.CallOption) (*HTTPConfig, error) {
	out := new(HTTPConfig)
	err := c.cc.Invoke(ctx, "/pb.HttpManager/GetOneHTTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpManagerClient) GetAllHTTP(ctx context.Context, in *Device, opts ...grpc.CallOption) (*HTTPList, error) {
	out := new(HTTPList)
	err := c.cc.Invoke(ctx, "/pb.HttpManager/GetAllHTTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpManagerServer is the server API for HttpManager service.
// All implementations must embed UnimplementedHttpManagerServer
// for forward compatibility
type HttpManagerServer interface {
	CreateOneHTTP(context.Context, *HTTPConfig) (*HTTPConfig, error)
	UpdateOneHTTP(context.Context, *HTTPConfig) (*HTTPConfig, error)
	DeleteOneHTTP(context.Context, *HTTPConfig) (*emptypb.Empty, error)
	GetOneHTTP(context.Context, *HTTPConfig) (*HTTPConfig, error)
	GetAllHTTP(context.Context, *Device) (*HTTPList, error)
	mustEmbedUnimplementedHttpManagerServer()
}

// UnimplementedHttpManagerServer must be embedded to have forward compatible implementations.
type UnimplementedHttpManagerServer struct {
}

func (UnimplementedHttpManagerServer) CreateOneHTTP(context.Context, *HTTPConfig) (*HTTPConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOneHTTP not implemented")
}
func (UnimplementedHttpManagerServer) UpdateOneHTTP(context.Context, *HTTPConfig) (*HTTPConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOneHTTP not implemented")
}
func (UnimplementedHttpManagerServer) DeleteOneHTTP(context.Context, *HTTPConfig) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOneHTTP not implemented")
}
func (UnimplementedHttpManagerServer) GetOneHTTP(context.Context, *HTTPConfig) (*HTTPConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneHTTP not implemented")
}
func (UnimplementedHttpManagerServer) GetAllHTTP(context.Context, *Device) (*HTTPList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllHTTP not implemented")
}
func (UnimplementedHttpManagerServer) mustEmbedUnimplementedHttpManagerServer() {}

// UnsafeHttpManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HttpManagerServer will
// result in compilation errors.
type UnsafeHttpManagerServer interface {
	mustEmbedUnimplementedHttpManagerServer()
}

func RegisterHttpManagerServer(s grpc.ServiceRegistrar, srv HttpManagerServer) {
	s.RegisterService(&HttpManager_ServiceDesc, srv)
}

func _HttpManager_CreateOneHTTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpManagerServer).CreateOneHTTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HttpManager/CreateOneHTTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpManagerServer).CreateOneHTTP(ctx, req.(*HTTPConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpManager_UpdateOneHTTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpManagerServer).UpdateOneHTTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HttpManager/UpdateOneHTTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpManagerServer).UpdateOneHTTP(ctx, req.(*HTTPConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpManager_DeleteOneHTTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpManagerServer).DeleteOneHTTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HttpManager/DeleteOneHTTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpManagerServer).DeleteOneHTTP(ctx, req.(*HTTPConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpManager_GetOneHTTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpManagerServer).GetOneHTTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HttpManager/GetOneHTTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpManagerServer).GetOneHTTP(ctx, req.(*HTTPConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpManager_GetAllHTTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpManagerServer).GetAllHTTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HttpManager/GetAllHTTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpManagerServer).GetAllHTTP(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

// HttpManager_ServiceDesc is the grpc.ServiceDesc for HttpManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HttpManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HttpManager",
	HandlerType: (*HttpManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOneHTTP",
			Handler:    _HttpManager_CreateOneHTTP_Handler,
		},
		{
			MethodName: "UpdateOneHTTP",
			Handler:    _HttpManager_UpdateOneHTTP_Handler,
		},
		{
			MethodName: "DeleteOneHTTP",
			Handler:    _HttpManager_DeleteOneHTTP_Handler,
		},
		{
			MethodName: "GetOneHTTP",
			Handler:    _HttpManager_GetOneHTTP_Handler,
		},
		{
			MethodName: "GetAllHTTP",
			Handler:    _HttpManager_GetAllHTTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server/server.proto",
}
