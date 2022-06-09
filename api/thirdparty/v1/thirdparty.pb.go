// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.1
// source: api/thirdparty/v1/thirdparty.proto

package thirdpartyv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CallbackUrl       string `protobuf:"bytes,3,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	KeyValidityPeriod uint64 `protobuf:"varint,4,opt,name=key_validity_period,json=keyValidityPeriod,proto3" json:"key_validity_period,omitempty"`
	AutoLogin         bool   `protobuf:"varint,5,opt,name=auto_login,json=autoLogin,proto3" json:"auto_login,omitempty"`
	Secret            string `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty"`
	DevMode           bool   `protobuf:"varint,7,opt,name=dev_mode,json=devMode,proto3" json:"dev_mode,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Info) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *Info) GetKeyValidityPeriod() uint64 {
	if x != nil {
		return x.KeyValidityPeriod
	}
	return 0
}

func (x *Info) GetAutoLogin() bool {
	if x != nil {
		return x.AutoLogin
	}
	return false
}

func (x *Info) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Info) GetDevMode() bool {
	if x != nil {
		return x.DevMode
	}
	return false
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner       string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	CallbackUrl string `protobuf:"bytes,4,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	AutoLogin   bool   `protobuf:"varint,5,opt,name=auto_login,json=autoLogin,proto3" json:"auto_login,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{1}
}

func (x *AddRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AddRequest) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *AddRequest) GetAutoLogin() bool {
	if x != nil {
		return x.AutoLogin
	}
	return false
}

type AddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AddReply) Reset() {
	*x = AddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReply) ProtoMessage() {}

func (x *AddReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReply.ProtoReflect.Descriptor instead.
func (*AddReply) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{2}
}

func (x *AddReply) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type BindAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *BindAdminRequest) Reset() {
	*x = BindAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindAdminRequest) ProtoMessage() {}

func (x *BindAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindAdminRequest.ProtoReflect.Descriptor instead.
func (*BindAdminRequest) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{3}
}

func (x *BindAdminRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BindAdminRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type AllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AllRequest) Reset() {
	*x = AllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRequest) ProtoMessage() {}

func (x *AllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRequest.ProtoReflect.Descriptor instead.
func (*AllRequest) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{4}
}

type AllReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Info `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AllReply) Reset() {
	*x = AllReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllReply) ProtoMessage() {}

func (x *AllReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllReply.ProtoReflect.Descriptor instead.
func (*AllReply) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{5}
}

func (x *AllReply) GetData() []*Info {
	if x != nil {
		return x.Data
	}
	return nil
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{6}
}

func (x *InfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResultReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultReply) Reset() {
	*x = ResultReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultReply) ProtoMessage() {}

func (x *ResultReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultReply.ProtoReflect.Descriptor instead.
func (*ResultReply) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{7}
}

func (x *ResultReply) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type ServiceTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ServiceTokenReply) Reset() {
	*x = ServiceTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceTokenReply) ProtoMessage() {}

func (x *ServiceTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceTokenReply.ProtoReflect.Descriptor instead.
func (*ServiceTokenReply) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{8}
}

func (x *ServiceTokenReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GrantTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PeriodOfValidity int32  `protobuf:"varint,2,opt,name=period_of_validity,json=periodOfValidity,proto3" json:"period_of_validity,omitempty"`
}

func (x *GrantTokenReq) Reset() {
	*x = GrantTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantTokenReq) ProtoMessage() {}

func (x *GrantTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantTokenReq.ProtoReflect.Descriptor instead.
func (*GrantTokenReq) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{9}
}

func (x *GrantTokenReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GrantTokenReq) GetPeriodOfValidity() int32 {
	if x != nil {
		return x.PeriodOfValidity
	}
	return 0
}

type GrantTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiredAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
}

func (x *GrantTokenReply) Reset() {
	*x = GrantTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantTokenReply) ProtoMessage() {}

func (x *GrantTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_thirdparty_v1_thirdparty_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantTokenReply.ProtoReflect.Descriptor instead.
func (*GrantTokenReply) Descriptor() ([]byte, []int) {
	return file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP(), []int{10}
}

func (x *GrantTokenReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GrantTokenReply) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

var File_api_thirdparty_v1_thirdparty_proto protoreflect.FileDescriptor

var file_api_thirdparty_v1_thirdparty_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x6b,
	0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x76, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0xb5, 0x01,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10,
	0x01, 0x18, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x75, 0x74, 0x6f,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x33, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x4a, 0x0a, 0x10, 0x42, 0x69,
	0x6e, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18,
	0x32, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x0b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x29, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x66, 0x0a, 0x0d, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01,
	0x18, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x12, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x5f, 0x6f, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x1a, 0x07, 0x10, 0x80, 0xe7, 0x84, 0x0f, 0x20, 0x1e,
	0x52, 0x10, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x4f, 0x66, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x22, 0x62, 0x0a, 0x0f, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x32, 0xf3, 0x03, 0x0a, 0x0a, 0x54, 0x68, 0x69, 0x72, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x55, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x19, 0x2e, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x61, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x07,
	0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x22, 0x13, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x69, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6b, 0x0a, 0x09, 0x42, 0x69, 0x6e, 0x64,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x2d, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x61, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x0a,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x22, 0x17, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x3e, 0x5a, 0x3c,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x76, 0x6c, 0x79, 0x75, 0x6e, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x73, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_thirdparty_v1_thirdparty_proto_rawDescOnce sync.Once
	file_api_thirdparty_v1_thirdparty_proto_rawDescData = file_api_thirdparty_v1_thirdparty_proto_rawDesc
)

func file_api_thirdparty_v1_thirdparty_proto_rawDescGZIP() []byte {
	file_api_thirdparty_v1_thirdparty_proto_rawDescOnce.Do(func() {
		file_api_thirdparty_v1_thirdparty_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_thirdparty_v1_thirdparty_proto_rawDescData)
	})
	return file_api_thirdparty_v1_thirdparty_proto_rawDescData
}

var file_api_thirdparty_v1_thirdparty_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_thirdparty_v1_thirdparty_proto_goTypes = []interface{}{
	(*Info)(nil),                  // 0: thirdparty.v1.Info
	(*AddRequest)(nil),            // 1: thirdparty.v1.AddRequest
	(*AddReply)(nil),              // 2: thirdparty.v1.AddReply
	(*BindAdminRequest)(nil),      // 3: thirdparty.v1.BindAdminRequest
	(*AllRequest)(nil),            // 4: thirdparty.v1.AllRequest
	(*AllReply)(nil),              // 5: thirdparty.v1.AllReply
	(*InfoRequest)(nil),           // 6: thirdparty.v1.InfoRequest
	(*ResultReply)(nil),           // 7: thirdparty.v1.ResultReply
	(*ServiceTokenReply)(nil),     // 8: thirdparty.v1.ServiceTokenReply
	(*GrantTokenReq)(nil),         // 9: thirdparty.v1.GrantTokenReq
	(*GrantTokenReply)(nil),       // 10: thirdparty.v1.GrantTokenReply
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_api_thirdparty_v1_thirdparty_proto_depIdxs = []int32{
	0,  // 0: thirdparty.v1.AddReply.info:type_name -> thirdparty.v1.Info
	0,  // 1: thirdparty.v1.AllReply.data:type_name -> thirdparty.v1.Info
	11, // 2: thirdparty.v1.GrantTokenReply.expired_at:type_name -> google.protobuf.Timestamp
	1,  // 3: thirdparty.v1.ThirdParty.Add:input_type -> thirdparty.v1.AddRequest
	6,  // 4: thirdparty.v1.ThirdParty.Inspect:input_type -> thirdparty.v1.InfoRequest
	3,  // 5: thirdparty.v1.ThirdParty.BindAdmin:input_type -> thirdparty.v1.BindAdminRequest
	4,  // 6: thirdparty.v1.ThirdParty.All:input_type -> thirdparty.v1.AllRequest
	9,  // 7: thirdparty.v1.ThirdParty.GrantToken:input_type -> thirdparty.v1.GrantTokenReq
	2,  // 8: thirdparty.v1.ThirdParty.Add:output_type -> thirdparty.v1.AddReply
	0,  // 9: thirdparty.v1.ThirdParty.Inspect:output_type -> thirdparty.v1.Info
	7,  // 10: thirdparty.v1.ThirdParty.BindAdmin:output_type -> thirdparty.v1.ResultReply
	5,  // 11: thirdparty.v1.ThirdParty.All:output_type -> thirdparty.v1.AllReply
	10, // 12: thirdparty.v1.ThirdParty.GrantToken:output_type -> thirdparty.v1.GrantTokenReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_thirdparty_v1_thirdparty_proto_init() }
func file_api_thirdparty_v1_thirdparty_proto_init() {
	if File_api_thirdparty_v1_thirdparty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReply); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindAdminRequest); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllRequest); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllReply); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultReply); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceTokenReply); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantTokenReq); i {
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
		file_api_thirdparty_v1_thirdparty_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantTokenReply); i {
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
			RawDescriptor: file_api_thirdparty_v1_thirdparty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_thirdparty_v1_thirdparty_proto_goTypes,
		DependencyIndexes: file_api_thirdparty_v1_thirdparty_proto_depIdxs,
		MessageInfos:      file_api_thirdparty_v1_thirdparty_proto_msgTypes,
	}.Build()
	File_api_thirdparty_v1_thirdparty_proto = out.File
	file_api_thirdparty_v1_thirdparty_proto_rawDesc = nil
	file_api_thirdparty_v1_thirdparty_proto_goTypes = nil
	file_api_thirdparty_v1_thirdparty_proto_depIdxs = nil
}
