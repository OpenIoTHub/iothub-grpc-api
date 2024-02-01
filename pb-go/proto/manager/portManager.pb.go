// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: portManager.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 手动添加的局域网主机
type PortInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortInfoList []*PortInfo `protobuf:"bytes,1,rep,name=PortInfoList,proto3" json:"PortInfoList,omitempty"`
}

func (x *PortInfoList) Reset() {
	*x = PortInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portManager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortInfoList) ProtoMessage() {}

func (x *PortInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_portManager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortInfoList.ProtoReflect.Descriptor instead.
func (*PortInfoList) Descriptor() ([]byte, []int) {
	return file_portManager_proto_rawDescGZIP(), []int{0}
}

func (x *PortInfoList) GetPortInfoList() []*PortInfo {
	if x != nil {
		return x.PortInfoList
	}
	return nil
}

type HttpInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpInfoList []*HttpInfo `protobuf:"bytes,1,rep,name=HttpInfoList,proto3" json:"HttpInfoList,omitempty"`
}

func (x *HttpInfoList) Reset() {
	*x = HttpInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portManager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpInfoList) ProtoMessage() {}

func (x *HttpInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_portManager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpInfoList.ProtoReflect.Descriptor instead.
func (*HttpInfoList) Descriptor() ([]byte, []int) {
	return file_portManager_proto_rawDescGZIP(), []int{1}
}

func (x *HttpInfoList) GetHttpInfoList() []*HttpInfo {
	if x != nil {
		return x.HttpInfoList
	}
	return nil
}

type PortInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID                string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Name                string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description         string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Domain              string `protobuf:"bytes,4,opt,name=Domain,proto3" json:"Domain,omitempty"`
	Username            string `protobuf:"bytes,5,opt,name=Username,proto3" json:"Username,omitempty"`
	Password            string `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
	Port                int32  `protobuf:"varint,7,opt,name=Port,proto3" json:"Port,omitempty"`
	LocalPort           int32  `protobuf:"varint,8,opt,name=LocalPort,proto3" json:"LocalPort,omitempty"`
	BindAllAddr         bool   `protobuf:"varint,9,opt,name=BindAllAddr,proto3" json:"BindAllAddr,omitempty"`
	NetworkProtocol     string `protobuf:"bytes,10,opt,name=NetworkProtocol,proto3" json:"NetworkProtocol,omitempty"`
	ApplicationProtocol string `protobuf:"bytes,11,opt,name=ApplicationProtocol,proto3" json:"ApplicationProtocol,omitempty"`
	HostUUID            string `protobuf:"bytes,12,opt,name=HostUUID,proto3" json:"HostUUID,omitempty"`
}

func (x *PortInfo) Reset() {
	*x = PortInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portManager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortInfo) ProtoMessage() {}

func (x *PortInfo) ProtoReflect() protoreflect.Message {
	mi := &file_portManager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortInfo.ProtoReflect.Descriptor instead.
func (*PortInfo) Descriptor() ([]byte, []int) {
	return file_portManager_proto_rawDescGZIP(), []int{2}
}

func (x *PortInfo) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *PortInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PortInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PortInfo) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *PortInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PortInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PortInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PortInfo) GetLocalPort() int32 {
	if x != nil {
		return x.LocalPort
	}
	return 0
}

func (x *PortInfo) GetBindAllAddr() bool {
	if x != nil {
		return x.BindAllAddr
	}
	return false
}

func (x *PortInfo) GetNetworkProtocol() string {
	if x != nil {
		return x.NetworkProtocol
	}
	return ""
}

func (x *PortInfo) GetApplicationProtocol() string {
	if x != nil {
		return x.ApplicationProtocol
	}
	return ""
}

func (x *PortInfo) GetHostUUID() string {
	if x != nil {
		return x.HostUUID
	}
	return ""
}

type HttpInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID                string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Name                string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description         string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Domain              string `protobuf:"bytes,4,opt,name=Domain,proto3" json:"Domain,omitempty"`
	Username            string `protobuf:"bytes,5,opt,name=Username,proto3" json:"Username,omitempty"`
	Password            string `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
	RemoteAddr          string `protobuf:"bytes,7,opt,name=RemoteAddr,proto3" json:"RemoteAddr,omitempty"`
	RemotePort          int32  `protobuf:"varint,8,opt,name=RemotePort,proto3" json:"RemotePort,omitempty"`
	NetworkProtocol     string `protobuf:"bytes,9,opt,name=NetworkProtocol,proto3" json:"NetworkProtocol,omitempty"`
	ApplicationProtocol string `protobuf:"bytes,10,opt,name=ApplicationProtocol,proto3" json:"ApplicationProtocol,omitempty"`
	GatewayUuid         string `protobuf:"bytes,11,opt,name=GatewayUuid,proto3" json:"GatewayUuid,omitempty"`
}

func (x *HttpInfo) Reset() {
	*x = HttpInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portManager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpInfo) ProtoMessage() {}

func (x *HttpInfo) ProtoReflect() protoreflect.Message {
	mi := &file_portManager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpInfo.ProtoReflect.Descriptor instead.
func (*HttpInfo) Descriptor() ([]byte, []int) {
	return file_portManager_proto_rawDescGZIP(), []int{3}
}

func (x *HttpInfo) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *HttpInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HttpInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HttpInfo) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *HttpInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *HttpInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *HttpInfo) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *HttpInfo) GetRemotePort() int32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

func (x *HttpInfo) GetNetworkProtocol() string {
	if x != nil {
		return x.NetworkProtocol
	}
	return ""
}

func (x *HttpInfo) GetApplicationProtocol() string {
	if x != nil {
		return x.ApplicationProtocol
	}
	return ""
}

func (x *HttpInfo) GetGatewayUuid() string {
	if x != nil {
		return x.GatewayUuid
	}
	return ""
}

var File_portManager_proto protoreflect.FileDescriptor

var file_portManager_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0c, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x0c, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xf0, 0x02, 0x0a, 0x08, 0x50, 0x6f, 0x72, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x28, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x30, 0x0a, 0x13, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x48, 0x6f, 0x73, 0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x48, 0x6f, 0x73, 0x74, 0x55, 0x55, 0x49, 0x44, 0x22, 0xe2, 0x02, 0x0a, 0x08, 0x48, 0x74,
	0x74, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x30, 0x0a, 0x13, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x55, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x55, 0x75, 0x69, 0x64, 0x32, 0xfd,
	0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x39,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x72, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x74, 0x74, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portManager_proto_rawDescOnce sync.Once
	file_portManager_proto_rawDescData = file_portManager_proto_rawDesc
)

func file_portManager_proto_rawDescGZIP() []byte {
	file_portManager_proto_rawDescOnce.Do(func() {
		file_portManager_proto_rawDescData = protoimpl.X.CompressGZIP(file_portManager_proto_rawDescData)
	})
	return file_portManager_proto_rawDescData
}

var file_portManager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_portManager_proto_goTypes = []interface{}{
	(*PortInfoList)(nil),      // 0: pb.PortInfoList
	(*HttpInfoList)(nil),      // 1: pb.HttpInfoList
	(*PortInfo)(nil),          // 2: pb.PortInfo
	(*HttpInfo)(nil),          // 3: pb.HttpInfo
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
	(*OperationResponse)(nil), // 5: pb.OperationResponse
}
var file_portManager_proto_depIdxs = []int32{
	2, // 0: pb.PortInfoList.PortInfoList:type_name -> pb.PortInfo
	3, // 1: pb.HttpInfoList.HttpInfoList:type_name -> pb.HttpInfo
	4, // 2: pb.PortManager.GetAllPorts:input_type -> google.protobuf.Empty
	2, // 3: pb.PortManager.AddPort:input_type -> pb.PortInfo
	2, // 4: pb.PortManager.UpdatePort:input_type -> pb.PortInfo
	2, // 5: pb.PortManager.DelPort:input_type -> pb.PortInfo
	4, // 6: pb.PortManager.GetAllPortInfoListByServerUuid:input_type -> google.protobuf.Empty
	4, // 7: pb.PortManager.GetAllHttpInfoListByServerUuid:input_type -> google.protobuf.Empty
	0, // 8: pb.PortManager.GetAllPorts:output_type -> pb.PortInfoList
	5, // 9: pb.PortManager.AddPort:output_type -> pb.OperationResponse
	5, // 10: pb.PortManager.UpdatePort:output_type -> pb.OperationResponse
	5, // 11: pb.PortManager.DelPort:output_type -> pb.OperationResponse
	0, // 12: pb.PortManager.GetAllPortInfoListByServerUuid:output_type -> pb.PortInfoList
	1, // 13: pb.PortManager.GetAllHttpInfoListByServerUuid:output_type -> pb.HttpInfoList
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_portManager_proto_init() }
func file_portManager_proto_init() {
	if File_portManager_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_portManager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortInfoList); i {
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
		file_portManager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpInfoList); i {
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
		file_portManager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortInfo); i {
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
		file_portManager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpInfo); i {
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
			RawDescriptor: file_portManager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portManager_proto_goTypes,
		DependencyIndexes: file_portManager_proto_depIdxs,
		MessageInfos:      file_portManager_proto_msgTypes,
	}.Build()
	File_portManager_proto = out.File
	file_portManager_proto_rawDesc = nil
	file_portManager_proto_goTypes = nil
	file_portManager_proto_depIdxs = nil
}
