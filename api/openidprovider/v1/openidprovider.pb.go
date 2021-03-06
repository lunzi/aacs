// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.1
// source: api/openidprovider/v1/openidprovider.proto

package openidproviderv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNo     string `protobuf:"bytes,4,opt,name=phone_no,json=phoneNo,proto3" json:"phone_no,omitempty"`
	Source      string `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	App         string `protobuf:"bytes,6,opt,name=app,proto3" json:"app,omitempty"`
	Retired     bool   `protobuf:"varint,7,opt,name=retired,proto3" json:"retired,omitempty"`
	Gender      string `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_api_openidprovider_v1_openidprovider_proto_rawDescGZIP(), []int{0}
}

func (x *Subject) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Subject) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Subject) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Subject) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

func (x *Subject) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Subject) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *Subject) GetRetired() bool {
	if x != nil {
		return x.Retired
	}
	return false
}

func (x *Subject) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type NameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NameReply) Reset() {
	*x = NameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameReply) ProtoMessage() {}

func (x *NameReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameReply.ProtoReflect.Descriptor instead.
func (*NameReply) Descriptor() ([]byte, []int) {
	return file_api_openidprovider_v1_openidprovider_proto_rawDescGZIP(), []int{1}
}

func (x *NameReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BasicAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Pwd string `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
}

func (x *BasicAuthReq) Reset() {
	*x = BasicAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAuthReq) ProtoMessage() {}

func (x *BasicAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicAuthReq.ProtoReflect.Descriptor instead.
func (*BasicAuthReq) Descriptor() ([]byte, []int) {
	return file_api_openidprovider_v1_openidprovider_proto_rawDescGZIP(), []int{2}
}

func (x *BasicAuthReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *BasicAuthReq) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

type BasicAuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub *Subject `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
}

func (x *BasicAuthReply) Reset() {
	*x = BasicAuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicAuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAuthReply) ProtoMessage() {}

func (x *BasicAuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicAuthReply.ProtoReflect.Descriptor instead.
func (*BasicAuthReply) Descriptor() ([]byte, []int) {
	return file_api_openidprovider_v1_openidprovider_proto_rawDescGZIP(), []int{3}
}

func (x *BasicAuthReply) GetSub() *Subject {
	if x != nil {
		return x.Sub
	}
	return nil
}

type TokenAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenAuthReq) Reset() {
	*x = TokenAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenAuthReq) ProtoMessage() {}

func (x *TokenAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenAuthReq.ProtoReflect.Descriptor instead.
func (*TokenAuthReq) Descriptor() ([]byte, []int) {
	return file_api_openidprovider_v1_openidprovider_proto_rawDescGZIP(), []int{4}
}

func (x *TokenAuthReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TokenAuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub *Subject `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	Uid string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *TokenAuthReply) Reset() {
	*x = TokenAuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenAuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenAuthReply) ProtoMessage() {}

func (x *TokenAuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenAuthReply.ProtoReflect.Descriptor instead.
func (*TokenAuthReply) Descriptor() ([]byte, []int) {
	return file_api_openidprovider_v1_openidprovider_proto_rawDescGZIP(), []int{5}
}

func (x *TokenAuthReply) GetSub() *Subject {
	if x != nil {
		return x.Sub
	}
	return nil
}

func (x *TokenAuthReply) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type SearchUidReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *SearchUidReq) Reset() {
	*x = SearchUidReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUidReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUidReq) ProtoMessage() {}

func (x *SearchUidReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUidReq.ProtoReflect.Descriptor instead.
func (*SearchUidReq) Descriptor() ([]byte, []int) {
	return file_api_openidprovider_v1_openidprovider_proto_rawDescGZIP(), []int{6}
}

func (x *SearchUidReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type SearchUidReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub *Subject `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
}

func (x *SearchUidReply) Reset() {
	*x = SearchUidReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUidReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUidReply) ProtoMessage() {}

func (x *SearchUidReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_openidprovider_v1_openidprovider_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUidReply.ProtoReflect.Descriptor instead.
func (*SearchUidReply) Descriptor() ([]byte, []int) {
	return file_api_openidprovider_v1_openidprovider_proto_rawDescGZIP(), []int{7}
}

func (x *SearchUidReply) GetSub() *Subject {
	if x != nil {
		return x.Sub
	}
	return nil
}

var File_api_openidprovider_v1_openidprovider_proto protoreflect.FileDescriptor

var file_api_openidprovider_v1_openidprovider_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6f, 0x70,
	0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x70, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x22, 0x1f, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x12, 0x1b, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x03, 0x70, 0x77, 0x64, 0x22, 0x3e, 0x0a, 0x0e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c,
	0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x03, 0x73, 0x75, 0x62, 0x22, 0x30, 0x0a, 0x0c,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07,
	0x72, 0x05, 0x10, 0x01, 0x18, 0xe8, 0x07, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x50,
	0x0a, 0x0e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2c, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x22, 0x2b, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x3e, 0x0a,
	0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2c, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x03, 0x73, 0x75, 0x62, 0x32, 0xd7, 0x03,
	0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x44, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x5d, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x76, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2d,
	0x61, 0x75, 0x74, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x22, 0x1a, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x01, 0x2a, 0x12,
	0x76, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x69, 0x64, 0x12, 0x1f, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2d, 0x75, 0x69, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x46, 0x5a, 0x44, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x61, 0x76, 0x6c, 0x79, 0x75, 0x6e, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_openidprovider_v1_openidprovider_proto_rawDescOnce sync.Once
	file_api_openidprovider_v1_openidprovider_proto_rawDescData = file_api_openidprovider_v1_openidprovider_proto_rawDesc
)

func file_api_openidprovider_v1_openidprovider_proto_rawDescGZIP() []byte {
	file_api_openidprovider_v1_openidprovider_proto_rawDescOnce.Do(func() {
		file_api_openidprovider_v1_openidprovider_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_openidprovider_v1_openidprovider_proto_rawDescData)
	})
	return file_api_openidprovider_v1_openidprovider_proto_rawDescData
}

var file_api_openidprovider_v1_openidprovider_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_openidprovider_v1_openidprovider_proto_goTypes = []interface{}{
	(*Subject)(nil),        // 0: openidprovider.v1.Subject
	(*NameReply)(nil),      // 1: openidprovider.v1.NameReply
	(*BasicAuthReq)(nil),   // 2: openidprovider.v1.BasicAuthReq
	(*BasicAuthReply)(nil), // 3: openidprovider.v1.BasicAuthReply
	(*TokenAuthReq)(nil),   // 4: openidprovider.v1.TokenAuthReq
	(*TokenAuthReply)(nil), // 5: openidprovider.v1.TokenAuthReply
	(*SearchUidReq)(nil),   // 6: openidprovider.v1.SearchUidReq
	(*SearchUidReply)(nil), // 7: openidprovider.v1.SearchUidReply
	(*emptypb.Empty)(nil),  // 8: google.protobuf.Empty
}
var file_api_openidprovider_v1_openidprovider_proto_depIdxs = []int32{
	0, // 0: openidprovider.v1.BasicAuthReply.sub:type_name -> openidprovider.v1.Subject
	0, // 1: openidprovider.v1.TokenAuthReply.sub:type_name -> openidprovider.v1.Subject
	0, // 2: openidprovider.v1.SearchUidReply.sub:type_name -> openidprovider.v1.Subject
	8, // 3: openidprovider.v1.OpenIDProvider.Name:input_type -> google.protobuf.Empty
	2, // 4: openidprovider.v1.OpenIDProvider.BasicAuth:input_type -> openidprovider.v1.BasicAuthReq
	4, // 5: openidprovider.v1.OpenIDProvider.TokenAuth:input_type -> openidprovider.v1.TokenAuthReq
	6, // 6: openidprovider.v1.OpenIDProvider.SearchUid:input_type -> openidprovider.v1.SearchUidReq
	1, // 7: openidprovider.v1.OpenIDProvider.Name:output_type -> openidprovider.v1.NameReply
	3, // 8: openidprovider.v1.OpenIDProvider.BasicAuth:output_type -> openidprovider.v1.BasicAuthReply
	5, // 9: openidprovider.v1.OpenIDProvider.TokenAuth:output_type -> openidprovider.v1.TokenAuthReply
	7, // 10: openidprovider.v1.OpenIDProvider.SearchUid:output_type -> openidprovider.v1.SearchUidReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_openidprovider_v1_openidprovider_proto_init() }
func file_api_openidprovider_v1_openidprovider_proto_init() {
	if File_api_openidprovider_v1_openidprovider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_openidprovider_v1_openidprovider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_api_openidprovider_v1_openidprovider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameReply); i {
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
		file_api_openidprovider_v1_openidprovider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicAuthReq); i {
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
		file_api_openidprovider_v1_openidprovider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicAuthReply); i {
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
		file_api_openidprovider_v1_openidprovider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenAuthReq); i {
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
		file_api_openidprovider_v1_openidprovider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenAuthReply); i {
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
		file_api_openidprovider_v1_openidprovider_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUidReq); i {
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
		file_api_openidprovider_v1_openidprovider_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUidReply); i {
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
			RawDescriptor: file_api_openidprovider_v1_openidprovider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_openidprovider_v1_openidprovider_proto_goTypes,
		DependencyIndexes: file_api_openidprovider_v1_openidprovider_proto_depIdxs,
		MessageInfos:      file_api_openidprovider_v1_openidprovider_proto_msgTypes,
	}.Build()
	File_api_openidprovider_v1_openidprovider_proto = out.File
	file_api_openidprovider_v1_openidprovider_proto_rawDesc = nil
	file_api_openidprovider_v1_openidprovider_proto_goTypes = nil
	file_api_openidprovider_v1_openidprovider_proto_depIdxs = nil
}
