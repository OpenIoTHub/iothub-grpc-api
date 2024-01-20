// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: configManager.proto

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

// ConfigManagerClient is the client API for ConfigManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigManagerClient interface {
	// 用户配置
	// 普通配置一次性操作多个
	GetUserConfigByKey(ctx context.Context, in *StringValue, opts ...grpc.CallOption) (*StringValue, error)
	GetAllUserConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserConfigMap, error)
	// 创建或者更新
	SetUserConfigByKey(ctx context.Context, in *UserConfigMap, opts ...grpc.CallOption) (*OperationResponse, error)
	// 删除
	DelAllUserConfig(ctx context.Context, in *UserConfigMap, opts ...grpc.CallOption) (*OperationResponse, error)
	DelUserConfigByKey(ctx context.Context, in *StringValue, opts ...grpc.CallOption) (*OperationResponse, error)
}

type configManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigManagerClient(cc grpc.ClientConnInterface) ConfigManagerClient {
	return &configManagerClient{cc}
}

func (c *configManagerClient) GetUserConfigByKey(ctx context.Context, in *StringValue, opts ...grpc.CallOption) (*StringValue, error) {
	out := new(StringValue)
	err := c.cc.Invoke(ctx, "/pb.ConfigManager/GetUserConfigByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configManagerClient) GetAllUserConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserConfigMap, error) {
	out := new(UserConfigMap)
	err := c.cc.Invoke(ctx, "/pb.ConfigManager/GetAllUserConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configManagerClient) SetUserConfigByKey(ctx context.Context, in *UserConfigMap, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.ConfigManager/SetUserConfigByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configManagerClient) DelAllUserConfig(ctx context.Context, in *UserConfigMap, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.ConfigManager/DelAllUserConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configManagerClient) DelUserConfigByKey(ctx context.Context, in *StringValue, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/pb.ConfigManager/DelUserConfigByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigManagerServer is the server API for ConfigManager service.
// All implementations must embed UnimplementedConfigManagerServer
// for forward compatibility
type ConfigManagerServer interface {
	// 用户配置
	// 普通配置一次性操作多个
	GetUserConfigByKey(context.Context, *StringValue) (*StringValue, error)
	GetAllUserConfig(context.Context, *Empty) (*UserConfigMap, error)
	// 创建或者更新
	SetUserConfigByKey(context.Context, *UserConfigMap) (*OperationResponse, error)
	// 删除
	DelAllUserConfig(context.Context, *UserConfigMap) (*OperationResponse, error)
	DelUserConfigByKey(context.Context, *StringValue) (*OperationResponse, error)
	mustEmbedUnimplementedConfigManagerServer()
}

// UnimplementedConfigManagerServer must be embedded to have forward compatible implementations.
type UnimplementedConfigManagerServer struct {
}

func (UnimplementedConfigManagerServer) GetUserConfigByKey(context.Context, *StringValue) (*StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserConfigByKey not implemented")
}
func (UnimplementedConfigManagerServer) GetAllUserConfig(context.Context, *Empty) (*UserConfigMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserConfig not implemented")
}
func (UnimplementedConfigManagerServer) SetUserConfigByKey(context.Context, *UserConfigMap) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserConfigByKey not implemented")
}
func (UnimplementedConfigManagerServer) DelAllUserConfig(context.Context, *UserConfigMap) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAllUserConfig not implemented")
}
func (UnimplementedConfigManagerServer) DelUserConfigByKey(context.Context, *StringValue) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUserConfigByKey not implemented")
}
func (UnimplementedConfigManagerServer) mustEmbedUnimplementedConfigManagerServer() {}

// UnsafeConfigManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigManagerServer will
// result in compilation errors.
type UnsafeConfigManagerServer interface {
	mustEmbedUnimplementedConfigManagerServer()
}

func RegisterConfigManagerServer(s grpc.ServiceRegistrar, srv ConfigManagerServer) {
	s.RegisterService(&ConfigManager_ServiceDesc, srv)
}

func _ConfigManager_GetUserConfigByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManagerServer).GetUserConfigByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ConfigManager/GetUserConfigByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManagerServer).GetUserConfigByKey(ctx, req.(*StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigManager_GetAllUserConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManagerServer).GetAllUserConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ConfigManager/GetAllUserConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManagerServer).GetAllUserConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigManager_SetUserConfigByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConfigMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManagerServer).SetUserConfigByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ConfigManager/SetUserConfigByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManagerServer).SetUserConfigByKey(ctx, req.(*UserConfigMap))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigManager_DelAllUserConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConfigMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManagerServer).DelAllUserConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ConfigManager/DelAllUserConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManagerServer).DelAllUserConfig(ctx, req.(*UserConfigMap))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigManager_DelUserConfigByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManagerServer).DelUserConfigByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ConfigManager/DelUserConfigByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManagerServer).DelUserConfigByKey(ctx, req.(*StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigManager_ServiceDesc is the grpc.ServiceDesc for ConfigManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ConfigManager",
	HandlerType: (*ConfigManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserConfigByKey",
			Handler:    _ConfigManager_GetUserConfigByKey_Handler,
		},
		{
			MethodName: "GetAllUserConfig",
			Handler:    _ConfigManager_GetAllUserConfig_Handler,
		},
		{
			MethodName: "SetUserConfigByKey",
			Handler:    _ConfigManager_SetUserConfigByKey_Handler,
		},
		{
			MethodName: "DelAllUserConfig",
			Handler:    _ConfigManager_DelAllUserConfig_Handler,
		},
		{
			MethodName: "DelUserConfigByKey",
			Handler:    _ConfigManager_DelUserConfigByKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configManager.proto",
}
