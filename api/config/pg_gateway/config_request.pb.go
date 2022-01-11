// -*- mode: protobuf; indent-tabs-mode: t; c-basic-offset: 8; tab-width: 8 -*-

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: config/pg_gateway/config_request.proto

package pg_gateway

import (
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V1 *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pg_gateway_config_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_pg_gateway_config_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_pg_gateway_config_request_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if x != nil {
		return x.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys *ConfigRequest_V1_System `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
}

func (x *ConfigRequest_V1) Reset() {
	*x = ConfigRequest_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pg_gateway_config_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1) ProtoMessage() {}

func (x *ConfigRequest_V1) ProtoReflect() protoreflect.Message {
	mi := &file_config_pg_gateway_config_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return file_config_pg_gateway_config_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if x != nil {
		return x.Sys
	}
	return nil
}

type ConfigRequest_V1_System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mlsa      *shared.Mlsa                       `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Service   *ConfigRequest_V1_System_Service   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Tls       *shared.TLSCredentials             `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Log       *shared.Log                        `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	Timeouts  *ConfigRequest_V1_System_Timeouts  `protobuf:"bytes,5,opt,name=timeouts,proto3" json:"timeouts,omitempty" toml:"timeouts,omitempty" mapstructure:"timeouts,omitempty"`
	Resolvers *ConfigRequest_V1_System_Resolvers `protobuf:"bytes,6,opt,name=resolvers,proto3" json:"resolvers,omitempty" toml:"resolvers,omitempty" mapstructure:"resolvers,omitempty"`
}

func (x *ConfigRequest_V1_System) Reset() {
	*x = ConfigRequest_V1_System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pg_gateway_config_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System) ProtoMessage() {}

func (x *ConfigRequest_V1_System) ProtoReflect() protoreflect.Message {
	mi := &file_config_pg_gateway_config_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return file_config_pg_gateway_config_request_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if x != nil {
		return x.Mlsa
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetLog() *shared.Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetTimeouts() *ConfigRequest_V1_System_Timeouts {
	if x != nil {
		return x.Timeouts
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetResolvers() *ConfigRequest_V1_System_Resolvers {
	if x != nil {
		return x.Resolvers
	}
	return nil
}

type ConfigRequest_V1_System_Resolvers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nameservers             []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=nameservers,proto3" json:"nameservers,omitempty" toml:"nameservers,omitempty" mapstructure:"nameservers,omitempty"`
	EnableSystemNameservers *wrapperspb.BoolValue     `protobuf:"bytes,2,opt,name=enable_system_nameservers,json=enableSystemNameservers,proto3" json:"enable_system_nameservers,omitempty" toml:"enable_system_nameservers,omitempty" mapstructure:"enable_system_nameservers,omitempty"`
}

func (x *ConfigRequest_V1_System_Resolvers) Reset() {
	*x = ConfigRequest_V1_System_Resolvers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pg_gateway_config_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Resolvers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Resolvers) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Resolvers) ProtoReflect() protoreflect.Message {
	mi := &file_config_pg_gateway_config_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Resolvers.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Resolvers) Descriptor() ([]byte, []int) {
	return file_config_pg_gateway_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *ConfigRequest_V1_System_Resolvers) GetNameservers() []*wrapperspb.StringValue {
	if x != nil {
		return x.Nameservers
	}
	return nil
}

func (x *ConfigRequest_V1_System_Resolvers) GetEnableSystemNameservers() *wrapperspb.BoolValue {
	if x != nil {
		return x.EnableSystemNameservers
	}
	return nil
}

type ConfigRequest_V1_System_Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDomain *wrapperspb.BoolValue   `protobuf:"bytes,1,opt,name=is_domain,json=isDomain,proto3" json:"is_domain,omitempty" toml:"is_domain,omitempty" mapstructure:"is_domain,omitempty"`
	Address  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" toml:"address,omitempty" mapstructure:"address,omitempty"`
	Port     *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
}

func (x *ConfigRequest_V1_System_Endpoint) Reset() {
	*x = ConfigRequest_V1_System_Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pg_gateway_config_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Endpoint) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_config_pg_gateway_config_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Endpoint.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Endpoint) Descriptor() ([]byte, []int) {
	return file_config_pg_gateway_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

func (x *ConfigRequest_V1_System_Endpoint) GetIsDomain() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsDomain
	}
	return nil
}

func (x *ConfigRequest_V1_System_Endpoint) GetAddress() *wrapperspb.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ConfigRequest_V1_System_Endpoint) GetPort() *wrapperspb.StringValue {
	if x != nil {
		return x.Port
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	Host               *wrapperspb.StringValue             `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"` // The listen host is no longer setable(localhost only)
	Port               *wrapperspb.Int32Value              `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	ExternalPostgresql *shared.External_Postgresql         `protobuf:"bytes,4,opt,name=external_postgresql,json=externalPostgresql,proto3" json:"external_postgresql,omitempty" toml:"external_postgresql,omitempty" mapstructure:"external_postgresql,omitempty"`
	ParsedNodes        []*ConfigRequest_V1_System_Endpoint `protobuf:"bytes,5,rep,name=parsed_nodes,json=parsedNodes,proto3" json:"parsed_nodes,omitempty" toml:"parsed_nodes,omitempty" mapstructure:"parsed_nodes,omitempty"`
}

func (x *ConfigRequest_V1_System_Service) Reset() {
	*x = ConfigRequest_V1_System_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pg_gateway_config_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Service) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_pg_gateway_config_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Service.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return file_config_pg_gateway_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 2}
}

// Deprecated: Do not use.
func (x *ConfigRequest_V1_System_Service) GetHost() *wrapperspb.StringValue {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetPort() *wrapperspb.Int32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetExternalPostgresql() *shared.External_Postgresql {
	if x != nil {
		return x.ExternalPostgresql
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetParsedNodes() []*ConfigRequest_V1_System_Endpoint {
	if x != nil {
		return x.ParsedNodes
	}
	return nil
}

type ConfigRequest_V1_System_Timeouts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connect *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=connect,proto3" json:"connect,omitempty" toml:"connect,omitempty" mapstructure:"connect,omitempty"`
	Idle    *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=idle,proto3" json:"idle,omitempty" toml:"idle,omitempty" mapstructure:"idle,omitempty"`
}

func (x *ConfigRequest_V1_System_Timeouts) Reset() {
	*x = ConfigRequest_V1_System_Timeouts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pg_gateway_config_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Timeouts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Timeouts) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Timeouts) ProtoReflect() protoreflect.Message {
	mi := &file_config_pg_gateway_config_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Timeouts.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Timeouts) Descriptor() ([]byte, []int) {
	return file_config_pg_gateway_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 3}
}

func (x *ConfigRequest_V1_System_Timeouts) GetConnect() *wrapperspb.Int32Value {
	if x != nil {
		return x.Connect
	}
	return nil
}

func (x *ConfigRequest_V1_System_Timeouts) GetIdle() *wrapperspb.Int32Value {
	if x != nil {
		return x.Idle
	}
	return nil
}

var File_config_pg_gateway_config_request_proto protoreflect.FileDescriptor

var file_config_pg_gateway_config_request_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x67,
	0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x61, 0x32, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x32, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab,
	0x0b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x40, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x70, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x52, 0x02,
	0x76, 0x31, 0x1a, 0xbc, 0x0a, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x49, 0x0a, 0x03, 0x73, 0x79, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x67, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x03, 0x73, 0x79, 0x73, 0x1a, 0xea, 0x09, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x34, 0x0a, 0x04, 0x6d, 0x6c, 0x73, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6c, 0x73, 0x61, 0x52,
	0x04, 0x6d, 0x6c, 0x73, 0x61, 0x12, 0x59, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x67, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x31,
	0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f,
	0x67, 0x12, 0x5c, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x67, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x73, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x12,
	0x5f, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73,
	0x1a, 0xa3, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x12, 0x3e,
	0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x56,
	0x0a, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x1a, 0xad, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x69, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x36, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0xd4, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x15, 0xc2, 0xf3, 0x18, 0x11, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x10, 0xa1, 0x4f, 0x1a, 0x03, 0x74, 0x63, 0x70, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x60, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x52, 0x12,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x71, 0x6c, 0x12, 0x63, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70,
	0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x1a, 0x72, 0x0a,
	0x08, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x12, 0x2f, 0x0a, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x69, 0x64, 0x6c,
	0x65, 0x3a, 0x19, 0xc2, 0xf3, 0x18, 0x15, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2d, 0x70, 0x67, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x70, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_pg_gateway_config_request_proto_rawDescOnce sync.Once
	file_config_pg_gateway_config_request_proto_rawDescData = file_config_pg_gateway_config_request_proto_rawDesc
)

func file_config_pg_gateway_config_request_proto_rawDescGZIP() []byte {
	file_config_pg_gateway_config_request_proto_rawDescOnce.Do(func() {
		file_config_pg_gateway_config_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_pg_gateway_config_request_proto_rawDescData)
	})
	return file_config_pg_gateway_config_request_proto_rawDescData
}

var file_config_pg_gateway_config_request_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_config_pg_gateway_config_request_proto_goTypes = []interface{}{
	(*ConfigRequest)(nil),                     // 0: chef.automate.infra.pg_gateway.ConfigRequest
	(*ConfigRequest_V1)(nil),                  // 1: chef.automate.infra.pg_gateway.ConfigRequest.V1
	(*ConfigRequest_V1_System)(nil),           // 2: chef.automate.infra.pg_gateway.ConfigRequest.V1.System
	(*ConfigRequest_V1_System_Resolvers)(nil), // 3: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Resolvers
	(*ConfigRequest_V1_System_Endpoint)(nil),  // 4: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Endpoint
	(*ConfigRequest_V1_System_Service)(nil),   // 5: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Service
	(*ConfigRequest_V1_System_Timeouts)(nil),  // 6: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Timeouts
	(*shared.Mlsa)(nil),                       // 7: chef.automate.infra.config.Mlsa
	(*shared.TLSCredentials)(nil),             // 8: chef.automate.infra.config.TLSCredentials
	(*shared.Log)(nil),                        // 9: chef.automate.infra.config.Log
	(*wrapperspb.StringValue)(nil),            // 10: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),              // 11: google.protobuf.BoolValue
	(*wrapperspb.Int32Value)(nil),             // 12: google.protobuf.Int32Value
	(*shared.External_Postgresql)(nil),        // 13: chef.automate.infra.config.External.Postgresql
}
var file_config_pg_gateway_config_request_proto_depIdxs = []int32{
	1,  // 0: chef.automate.infra.pg_gateway.ConfigRequest.v1:type_name -> chef.automate.infra.pg_gateway.ConfigRequest.V1
	2,  // 1: chef.automate.infra.pg_gateway.ConfigRequest.V1.sys:type_name -> chef.automate.infra.pg_gateway.ConfigRequest.V1.System
	7,  // 2: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.mlsa:type_name -> chef.automate.infra.config.Mlsa
	5,  // 3: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.service:type_name -> chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Service
	8,  // 4: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.tls:type_name -> chef.automate.infra.config.TLSCredentials
	9,  // 5: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.log:type_name -> chef.automate.infra.config.Log
	6,  // 6: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.timeouts:type_name -> chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Timeouts
	3,  // 7: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.resolvers:type_name -> chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Resolvers
	10, // 8: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Resolvers.nameservers:type_name -> google.protobuf.StringValue
	11, // 9: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Resolvers.enable_system_nameservers:type_name -> google.protobuf.BoolValue
	11, // 10: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Endpoint.is_domain:type_name -> google.protobuf.BoolValue
	10, // 11: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Endpoint.address:type_name -> google.protobuf.StringValue
	10, // 12: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Endpoint.port:type_name -> google.protobuf.StringValue
	10, // 13: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Service.host:type_name -> google.protobuf.StringValue
	12, // 14: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Service.port:type_name -> google.protobuf.Int32Value
	13, // 15: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Service.external_postgresql:type_name -> chef.automate.infra.config.External.Postgresql
	4,  // 16: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Service.parsed_nodes:type_name -> chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Endpoint
	12, // 17: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Timeouts.connect:type_name -> google.protobuf.Int32Value
	12, // 18: chef.automate.infra.pg_gateway.ConfigRequest.V1.System.Timeouts.idle:type_name -> google.protobuf.Int32Value
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_config_pg_gateway_config_request_proto_init() }
func file_config_pg_gateway_config_request_proto_init() {
	if File_config_pg_gateway_config_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_pg_gateway_config_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_config_pg_gateway_config_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1); i {
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
		file_config_pg_gateway_config_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System); i {
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
		file_config_pg_gateway_config_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Resolvers); i {
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
		file_config_pg_gateway_config_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Endpoint); i {
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
		file_config_pg_gateway_config_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Service); i {
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
		file_config_pg_gateway_config_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Timeouts); i {
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
			RawDescriptor: file_config_pg_gateway_config_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_pg_gateway_config_request_proto_goTypes,
		DependencyIndexes: file_config_pg_gateway_config_request_proto_depIdxs,
		MessageInfos:      file_config_pg_gateway_config_request_proto_msgTypes,
	}.Build()
	File_config_pg_gateway_config_request_proto = out.File
	file_config_pg_gateway_config_request_proto_rawDesc = nil
	file_config_pg_gateway_config_request_proto_goTypes = nil
	file_config_pg_gateway_config_request_proto_depIdxs = nil
}
