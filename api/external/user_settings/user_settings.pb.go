// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: external/user_settings/user_settings.proto

package user_settings

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Connector string `protobuf:"bytes,2,opt,name=connector,proto3" json:"connector,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_user_settings_user_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_external_user_settings_user_settings_proto_msgTypes[0]
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
	return file_external_user_settings_user_settings_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetConnector() string {
	if x != nil {
		return x.Connector
	}
	return ""
}

type GetUserSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user to get settings for
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserSettingsRequest) Reset() {
	*x = GetUserSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_user_settings_user_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSettingsRequest) ProtoMessage() {}

func (x *GetUserSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_user_settings_user_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetUserSettingsRequest) Descriptor() ([]byte, []int) {
	return file_external_user_settings_user_settings_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserSettingsRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type PutUserSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the user. Cannot be changed. Used to sign in.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The user settings to persist.
	Settings map[string]*UserSettingValue `protobuf:"bytes,2,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PutUserSettingsRequest) Reset() {
	*x = PutUserSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_user_settings_user_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutUserSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutUserSettingsRequest) ProtoMessage() {}

func (x *PutUserSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_user_settings_user_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutUserSettingsRequest.ProtoReflect.Descriptor instead.
func (*PutUserSettingsRequest) Descriptor() ([]byte, []int) {
	return file_external_user_settings_user_settings_proto_rawDescGZIP(), []int{2}
}

func (x *PutUserSettingsRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PutUserSettingsRequest) GetSettings() map[string]*UserSettingValue {
	if x != nil {
		return x.Settings
	}
	return nil
}

type DeleteUserSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the user.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DeleteUserSettingsRequest) Reset() {
	*x = DeleteUserSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_user_settings_user_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserSettingsRequest) ProtoMessage() {}

func (x *DeleteUserSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_user_settings_user_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserSettingsRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserSettingsRequest) Descriptor() ([]byte, []int) {
	return file_external_user_settings_user_settings_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteUserSettingsRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *User                        `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Settings map[string]*UserSettingValue `protobuf:"bytes,2,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetUserSettingsResponse) Reset() {
	*x = GetUserSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_user_settings_user_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSettingsResponse) ProtoMessage() {}

func (x *GetUserSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_user_settings_user_settings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetUserSettingsResponse) Descriptor() ([]byte, []int) {
	return file_external_user_settings_user_settings_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserSettingsResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetUserSettingsResponse) GetSettings() map[string]*UserSettingValue {
	if x != nil {
		return x.Settings
	}
	return nil
}

type PutUserSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *PutUserSettingsResponse) Reset() {
	*x = PutUserSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_user_settings_user_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutUserSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutUserSettingsResponse) ProtoMessage() {}

func (x *PutUserSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_user_settings_user_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutUserSettingsResponse.ProtoReflect.Descriptor instead.
func (*PutUserSettingsResponse) Descriptor() ([]byte, []int) {
	return file_external_user_settings_user_settings_proto_rawDescGZIP(), []int{5}
}

func (x *PutUserSettingsResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserSettingValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Default value for this setting.
	DefaultValue string `protobuf:"bytes,1,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// Value for this setting.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Enabled
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Valid values for this setting.
	ValidValues []string `protobuf:"bytes,4,rep,name=valid_values,json=validValues,proto3" json:"valid_values,omitempty"`
}

func (x *UserSettingValue) Reset() {
	*x = UserSettingValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_user_settings_user_settings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSettingValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSettingValue) ProtoMessage() {}

func (x *UserSettingValue) ProtoReflect() protoreflect.Message {
	mi := &file_external_user_settings_user_settings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSettingValue.ProtoReflect.Descriptor instead.
func (*UserSettingValue) Descriptor() ([]byte, []int) {
	return file_external_user_settings_user_settings_proto_rawDescGZIP(), []int{6}
}

func (x *UserSettingValue) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

func (x *UserSettingValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *UserSettingValue) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *UserSettingValue) GetValidValues() []string {
	if x != nil {
		return x.ValidValues
	}
	return nil
}

var File_external_user_settings_user_settings_proto protoreflect.FileDescriptor

var file_external_user_settings_user_settings_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x53,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0xa6, 0x02, 0x0a, 0x16, 0x50, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x50, 0x75,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x6e, 0x0a, 0x0d,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x47, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x56, 0x0a, 0x19,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0xa8, 0x02, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x62, 0x0a, 0x08, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a,
	0x6e, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x47, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x54, 0x0a, 0x17, 0x50, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x32, 0xf6, 0x05, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xef, 0x01, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x37,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x69, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x30, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x7d, 0x8a, 0xb5, 0x18,
	0x10, 0x0a, 0x0e, 0x69, 0x61, 0x6d, 0x3a, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x8a, 0xb5, 0x18, 0x17, 0x12, 0x15, 0x69, 0x61, 0x6d, 0x3a, 0x69, 0x6e, 0x74, 0x72, 0x6f,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x3a, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0xf2, 0x01, 0x0a,
	0x0f, 0x50, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x50, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x50, 0x75, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x1a, 0x32, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x30, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x7d, 0x3a,
	0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x10, 0x0a, 0x0e, 0x69, 0x61, 0x6d, 0x3a, 0x69, 0x6e, 0x74, 0x72,
	0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x8a, 0xb5, 0x18, 0x17, 0x12, 0x15, 0x69, 0x61, 0x6d, 0x3a,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x3a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0xf7, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x8c, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x34, 0x2a, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x7d, 0x8a, 0xb5, 0x18, 0x4e, 0x0a, 0x30, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x3a, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x7d, 0x12,
	0x1a, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x3a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x3a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_user_settings_user_settings_proto_rawDescOnce sync.Once
	file_external_user_settings_user_settings_proto_rawDescData = file_external_user_settings_user_settings_proto_rawDesc
)

func file_external_user_settings_user_settings_proto_rawDescGZIP() []byte {
	file_external_user_settings_user_settings_proto_rawDescOnce.Do(func() {
		file_external_user_settings_user_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_user_settings_user_settings_proto_rawDescData)
	})
	return file_external_user_settings_user_settings_proto_rawDescData
}

var file_external_user_settings_user_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_external_user_settings_user_settings_proto_goTypes = []interface{}{
	(*User)(nil),                      // 0: chef.automate.api.user_settings.User
	(*GetUserSettingsRequest)(nil),    // 1: chef.automate.api.user_settings.GetUserSettingsRequest
	(*PutUserSettingsRequest)(nil),    // 2: chef.automate.api.user_settings.PutUserSettingsRequest
	(*DeleteUserSettingsRequest)(nil), // 3: chef.automate.api.user_settings.DeleteUserSettingsRequest
	(*GetUserSettingsResponse)(nil),   // 4: chef.automate.api.user_settings.GetUserSettingsResponse
	(*PutUserSettingsResponse)(nil),   // 5: chef.automate.api.user_settings.PutUserSettingsResponse
	(*UserSettingValue)(nil),          // 6: chef.automate.api.user_settings.UserSettingValue
	nil,                               // 7: chef.automate.api.user_settings.PutUserSettingsRequest.SettingsEntry
	nil,                               // 8: chef.automate.api.user_settings.GetUserSettingsResponse.SettingsEntry
	(*emptypb.Empty)(nil),             // 9: google.protobuf.Empty
}
var file_external_user_settings_user_settings_proto_depIdxs = []int32{
	0,  // 0: chef.automate.api.user_settings.GetUserSettingsRequest.user:type_name -> chef.automate.api.user_settings.User
	0,  // 1: chef.automate.api.user_settings.PutUserSettingsRequest.user:type_name -> chef.automate.api.user_settings.User
	7,  // 2: chef.automate.api.user_settings.PutUserSettingsRequest.settings:type_name -> chef.automate.api.user_settings.PutUserSettingsRequest.SettingsEntry
	0,  // 3: chef.automate.api.user_settings.DeleteUserSettingsRequest.user:type_name -> chef.automate.api.user_settings.User
	0,  // 4: chef.automate.api.user_settings.GetUserSettingsResponse.user:type_name -> chef.automate.api.user_settings.User
	8,  // 5: chef.automate.api.user_settings.GetUserSettingsResponse.settings:type_name -> chef.automate.api.user_settings.GetUserSettingsResponse.SettingsEntry
	0,  // 6: chef.automate.api.user_settings.PutUserSettingsResponse.user:type_name -> chef.automate.api.user_settings.User
	6,  // 7: chef.automate.api.user_settings.PutUserSettingsRequest.SettingsEntry.value:type_name -> chef.automate.api.user_settings.UserSettingValue
	6,  // 8: chef.automate.api.user_settings.GetUserSettingsResponse.SettingsEntry.value:type_name -> chef.automate.api.user_settings.UserSettingValue
	1,  // 9: chef.automate.api.user_settings.UserSettingsService.GetUserSettings:input_type -> chef.automate.api.user_settings.GetUserSettingsRequest
	2,  // 10: chef.automate.api.user_settings.UserSettingsService.PutUserSettings:input_type -> chef.automate.api.user_settings.PutUserSettingsRequest
	3,  // 11: chef.automate.api.user_settings.UserSettingsService.DeleteUserSettings:input_type -> chef.automate.api.user_settings.DeleteUserSettingsRequest
	4,  // 12: chef.automate.api.user_settings.UserSettingsService.GetUserSettings:output_type -> chef.automate.api.user_settings.GetUserSettingsResponse
	5,  // 13: chef.automate.api.user_settings.UserSettingsService.PutUserSettings:output_type -> chef.automate.api.user_settings.PutUserSettingsResponse
	9,  // 14: chef.automate.api.user_settings.UserSettingsService.DeleteUserSettings:output_type -> google.protobuf.Empty
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_external_user_settings_user_settings_proto_init() }
func file_external_user_settings_user_settings_proto_init() {
	if File_external_user_settings_user_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_user_settings_user_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_external_user_settings_user_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSettingsRequest); i {
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
		file_external_user_settings_user_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutUserSettingsRequest); i {
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
		file_external_user_settings_user_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserSettingsRequest); i {
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
		file_external_user_settings_user_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSettingsResponse); i {
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
		file_external_user_settings_user_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutUserSettingsResponse); i {
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
		file_external_user_settings_user_settings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSettingValue); i {
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
			RawDescriptor: file_external_user_settings_user_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_user_settings_user_settings_proto_goTypes,
		DependencyIndexes: file_external_user_settings_user_settings_proto_depIdxs,
		MessageInfos:      file_external_user_settings_user_settings_proto_msgTypes,
	}.Build()
	File_external_user_settings_user_settings_proto = out.File
	file_external_user_settings_user_settings_proto_rawDesc = nil
	file_external_user_settings_user_settings_proto_goTypes = nil
	file_external_user_settings_user_settings_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserSettingsServiceClient is the client API for UserSettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSettingsServiceClient interface {
	// GetUserSettings returns all of the preferences for a given user
	GetUserSettings(ctx context.Context, in *GetUserSettingsRequest, opts ...grpc.CallOption) (*GetUserSettingsResponse, error)
	// PutUserSettings upserts all of the preferences for a given user
	PutUserSettings(ctx context.Context, in *PutUserSettingsRequest, opts ...grpc.CallOption) (*PutUserSettingsResponse, error)
	// DeleteUserSettings deletes all settings for a given user
	DeleteUserSettings(ctx context.Context, in *DeleteUserSettingsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userSettingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSettingsServiceClient(cc grpc.ClientConnInterface) UserSettingsServiceClient {
	return &userSettingsServiceClient{cc}
}

func (c *userSettingsServiceClient) GetUserSettings(ctx context.Context, in *GetUserSettingsRequest, opts ...grpc.CallOption) (*GetUserSettingsResponse, error) {
	out := new(GetUserSettingsResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.user_settings.UserSettingsService/GetUserSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSettingsServiceClient) PutUserSettings(ctx context.Context, in *PutUserSettingsRequest, opts ...grpc.CallOption) (*PutUserSettingsResponse, error) {
	out := new(PutUserSettingsResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.user_settings.UserSettingsService/PutUserSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSettingsServiceClient) DeleteUserSettings(ctx context.Context, in *DeleteUserSettingsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chef.automate.api.user_settings.UserSettingsService/DeleteUserSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSettingsServiceServer is the server API for UserSettingsService service.
type UserSettingsServiceServer interface {
	// GetUserSettings returns all of the preferences for a given user
	GetUserSettings(context.Context, *GetUserSettingsRequest) (*GetUserSettingsResponse, error)
	// PutUserSettings upserts all of the preferences for a given user
	PutUserSettings(context.Context, *PutUserSettingsRequest) (*PutUserSettingsResponse, error)
	// DeleteUserSettings deletes all settings for a given user
	DeleteUserSettings(context.Context, *DeleteUserSettingsRequest) (*emptypb.Empty, error)
}

// UnimplementedUserSettingsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserSettingsServiceServer struct {
}

func (*UnimplementedUserSettingsServiceServer) GetUserSettings(context.Context, *GetUserSettingsRequest) (*GetUserSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSettings not implemented")
}
func (*UnimplementedUserSettingsServiceServer) PutUserSettings(context.Context, *PutUserSettingsRequest) (*PutUserSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutUserSettings not implemented")
}
func (*UnimplementedUserSettingsServiceServer) DeleteUserSettings(context.Context, *DeleteUserSettingsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserSettings not implemented")
}

func RegisterUserSettingsServiceServer(s *grpc.Server, srv UserSettingsServiceServer) {
	s.RegisterService(&_UserSettingsService_serviceDesc, srv)
}

func _UserSettingsService_GetUserSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSettingsServiceServer).GetUserSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.user_settings.UserSettingsService/GetUserSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSettingsServiceServer).GetUserSettings(ctx, req.(*GetUserSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSettingsService_PutUserSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutUserSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSettingsServiceServer).PutUserSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.user_settings.UserSettingsService/PutUserSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSettingsServiceServer).PutUserSettings(ctx, req.(*PutUserSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSettingsService_DeleteUserSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSettingsServiceServer).DeleteUserSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.user_settings.UserSettingsService/DeleteUserSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSettingsServiceServer).DeleteUserSettings(ctx, req.(*DeleteUserSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSettingsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.user_settings.UserSettingsService",
	HandlerType: (*UserSettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserSettings",
			Handler:    _UserSettingsService_GetUserSettings_Handler,
		},
		{
			MethodName: "PutUserSettings",
			Handler:    _UserSettingsService_PutUserSettings_Handler,
		},
		{
			MethodName: "DeleteUserSettings",
			Handler:    _UserSettingsService_DeleteUserSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/user_settings/user_settings.proto",
}
