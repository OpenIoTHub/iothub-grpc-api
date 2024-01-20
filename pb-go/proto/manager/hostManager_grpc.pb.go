// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: hostManager.proto

package pb

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

// HostManagerClient is the client API for HostManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostManagerClient interface {
	// Host
	GetAllHosts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HostInfoList, error)
	AddHost(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*OperationResponse, error)
	UpdateHost(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*OperationResponse, error)
	DelHost(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*OperationResponse, error)
	// 设置主机的MAC物理地址(用于WoL远程唤醒)
	SetDeviceMac(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*OperationResponse, error)
}

type hostManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewHostManagerClient(cc grpc.ClientConnInterface) HostManagerClient {
	return &hostManagerClient{cc}
}

func (c *hostManagerClient) GetAllHosts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HostInfoList, error) {
	out := new(HostInfoList)
	err := c.cc.Invoke(ctx, "/pb.HostManager/GetAllHosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostManagerClient) AddHost(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.HostManager/AddHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostManagerClient) UpdateHost(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.HostManager/UpdateHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostManagerClient) DelHost(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.HostManager/DelHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostManagerClient) SetDeviceMac(ctx context.Context, in *HostInfo, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.HostManager/SetDeviceMac", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostManagerServer is the server API for HostManager service.
// All implementations must embed UnimplementedHostManagerServer
// for forward compatibility
type HostManagerServer interface {
	// Host
	GetAllHosts(context.Context, *Empty) (*HostInfoList, error)
	AddHost(context.Context, *HostInfo) (*OperationResponse, error)
	UpdateHost(context.Context, *HostInfo) (*OperationResponse, error)
	DelHost(context.Context, *HostInfo) (*OperationResponse, error)
	// 设置主机的MAC物理地址(用于WoL远程唤醒)
	SetDeviceMac(context.Context, *HostInfo) (*OperationResponse, error)
	mustEmbedUnimplementedHostManagerServer()
}

// UnimplementedHostManagerServer must be embedded to have forward compatible implementations.
type UnimplementedHostManagerServer struct {
}

func (UnimplementedHostManagerServer) GetAllHosts(context.Context, *Empty) (*HostInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllHosts not implemented")
}
func (UnimplementedHostManagerServer) AddHost(context.Context, *HostInfo) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHost not implemented")
}
func (UnimplementedHostManagerServer) UpdateHost(context.Context, *HostInfo) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHost not implemented")
}
func (UnimplementedHostManagerServer) DelHost(context.Context, *HostInfo) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelHost not implemented")
}
func (UnimplementedHostManagerServer) SetDeviceMac(context.Context, *HostInfo) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceMac not implemented")
}
func (UnimplementedHostManagerServer) mustEmbedUnimplementedHostManagerServer() {}

// UnsafeHostManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostManagerServer will
// result in compilation errors.
type UnsafeHostManagerServer interface {
	mustEmbedUnimplementedHostManagerServer()
}

func RegisterHostManagerServer(s grpc.ServiceRegistrar, srv HostManagerServer) {
	s.RegisterService(&HostManager_ServiceDesc, srv)
}

func _HostManager_GetAllHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostManagerServer).GetAllHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HostManager/GetAllHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostManagerServer).GetAllHosts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostManager_AddHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostManagerServer).AddHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HostManager/AddHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostManagerServer).AddHost(ctx, req.(*HostInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostManager_UpdateHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostManagerServer).UpdateHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HostManager/UpdateHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostManagerServer).UpdateHost(ctx, req.(*HostInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostManager_DelHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostManagerServer).DelHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HostManager/DelHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostManagerServer).DelHost(ctx, req.(*HostInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostManager_SetDeviceMac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostManagerServer).SetDeviceMac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HostManager/SetDeviceMac",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostManagerServer).SetDeviceMac(ctx, req.(*HostInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// HostManager_ServiceDesc is the grpc.ServiceDesc for HostManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HostManager",
	HandlerType: (*HostManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllHosts",
			Handler:    _HostManager_GetAllHosts_Handler,
		},
		{
			MethodName: "AddHost",
			Handler:    _HostManager_AddHost_Handler,
		},
		{
			MethodName: "UpdateHost",
			Handler:    _HostManager_UpdateHost_Handler,
		},
		{
			MethodName: "DelHost",
			Handler:    _HostManager_DelHost_Handler,
		},
		{
			MethodName: "SetDeviceMac",
			Handler:    _HostManager_SetDeviceMac_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hostManager.proto",
}
