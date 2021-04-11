// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: account.proto

package accountPB

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type AccountStatus int32

const (
	AccountStatus_UNNAMED        AccountStatus = 0
	AccountStatus_PENDING_REVIEW AccountStatus = 1
	AccountStatus_NOT_PASS       AccountStatus = 2
	AccountStatus_ACTIVE         AccountStatus = 3
	AccountStatus_DISABLED       AccountStatus = 4
)

// Enum value maps for AccountStatus.
var (
	AccountStatus_name = map[int32]string{
		0: "UNNAMED",
		1: "PENDING_REVIEW",
		2: "NOT_PASS",
		3: "ACTIVE",
		4: "DISABLED",
	}
	AccountStatus_value = map[string]int32{
		"UNNAMED":        0,
		"PENDING_REVIEW": 1,
		"NOT_PASS":       2,
		"ACTIVE":         3,
		"DISABLED":       4,
	}
)

func (x AccountStatus) Enum() *AccountStatus {
	p := new(AccountStatus)
	*p = x
	return p
}

func (x AccountStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_account_proto_enumTypes[0].Descriptor()
}

func (AccountStatus) Type() protoreflect.EnumType {
	return &file_account_proto_enumTypes[0]
}

func (x AccountStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountStatus.Descriptor instead.
func (AccountStatus) EnumDescriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

type Sex int32

const (
	Sex_MALE   Sex = 0
	Sex_FEMALE Sex = 1
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	Sex_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_account_proto_enumTypes[1].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_account_proto_enumTypes[1]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName string `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoginReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountName   string        `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	RealName      string        `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	Sex           Sex           `protobuf:"varint,4,opt,name=sex,proto3,enum=account.Sex" json:"sex,omitempty"`
	College       string        `protobuf:"bytes,5,opt,name=college,proto3" json:"college,omitempty"`
	AccountStatus AccountStatus `protobuf:"varint,6,opt,name=account_status,json=accountStatus,proto3,enum=account.AccountStatus" json:"account_status,omitempty"`
	Grade         int32         `protobuf:"varint,7,opt,name=grade,proto3" json:"grade,omitempty"`
	Avatar        string        `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *UserInfo) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *UserInfo) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_MALE
}

func (x *UserInfo) GetCollege() string {
	if x != nil {
		return x.College
	}
	return ""
}

func (x *UserInfo) GetAccountStatus() AccountStatus {
	if x != nil {
		return x.AccountStatus
	}
	return AccountStatus_UNNAMED
}

func (x *UserInfo) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *UserInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type GetUserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	Message  string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetUserInfoReply) Reset() {
	*x = GetUserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReply) ProtoMessage() {}

func (x *GetUserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReply.ProtoReflect.Descriptor instead.
func (*GetUserInfoReply) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserInfoReply) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *GetUserInfoReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	Password    string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	RealName    string `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	Sex         Sex    `protobuf:"varint,4,opt,name=sex,proto3,enum=account.Sex" json:"sex,omitempty"`
	College     string `protobuf:"bytes,5,opt,name=college,proto3" json:"college,omitempty"`
	Grade       int32  `protobuf:"varint,7,opt,name=grade,proto3" json:"grade,omitempty"`
	Avatar      string `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAccountRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *CreateAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateAccountRequest) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *CreateAccountRequest) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_MALE
}

func (x *CreateAccountRequest) GetCollege() string {
	if x != nil {
		return x.College
	}
	return ""
}

func (x *CreateAccountRequest) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *CreateAccountRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type CreateAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateAccountReply) Reset() {
	*x = CreateAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountReply) ProtoMessage() {}

func (x *CreateAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountReply.ProtoReflect.Descriptor instead.
func (*CreateAccountReply) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAccountReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateAccountReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateAccountReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SubmitVerifyInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image  string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Remark string `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *SubmitVerifyInfoRequest) Reset() {
	*x = SubmitVerifyInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitVerifyInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitVerifyInfoRequest) ProtoMessage() {}

func (x *SubmitVerifyInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitVerifyInfoRequest.ProtoReflect.Descriptor instead.
func (*SubmitVerifyInfoRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{6}
}

func (x *SubmitVerifyInfoRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *SubmitVerifyInfoRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type SubmitVerifyInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SubmitVerifyInfoReply) Reset() {
	*x = SubmitVerifyInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitVerifyInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitVerifyInfoReply) ProtoMessage() {}

func (x *SubmitVerifyInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitVerifyInfoReply.ProtoReflect.Descriptor instead.
func (*SubmitVerifyInfoReply) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{7}
}

func (x *SubmitVerifyInfoReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SubmitVerifyInfoReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x56, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x81, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x03,
	0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x22, 0x5b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xda, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x78,
	0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x5e, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x47, 0x0a,
	0x17, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x4b, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0x58, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4e, 0x41, 0x4d, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x56,
	0x49, 0x45, 0x57, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x41, 0x53,
	0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x1b, 0x0a,
	0x03, 0x53, 0x65, 0x78, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_account_proto_goTypes = []interface{}{
	(AccountStatus)(0),              // 0: account.AccountStatus
	(Sex)(0),                        // 1: account.Sex
	(*LoginRequest)(nil),            // 2: account.LoginRequest
	(*LoginReply)(nil),              // 3: account.LoginReply
	(*UserInfo)(nil),                // 4: account.UserInfo
	(*GetUserInfoReply)(nil),        // 5: account.GetUserInfoReply
	(*CreateAccountRequest)(nil),    // 6: account.CreateAccountRequest
	(*CreateAccountReply)(nil),      // 7: account.CreateAccountReply
	(*SubmitVerifyInfoRequest)(nil), // 8: account.SubmitVerifyInfoRequest
	(*SubmitVerifyInfoReply)(nil),   // 9: account.SubmitVerifyInfoReply
}
var file_account_proto_depIdxs = []int32{
	1, // 0: account.UserInfo.sex:type_name -> account.Sex
	0, // 1: account.UserInfo.account_status:type_name -> account.AccountStatus
	4, // 2: account.GetUserInfoReply.userInfo:type_name -> account.UserInfo
	1, // 3: account.CreateAccountRequest.sex:type_name -> account.Sex
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReply); i {
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
		file_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountReply); i {
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
		file_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitVerifyInfoRequest); i {
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
		file_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitVerifyInfoReply); i {
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
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		EnumInfos:         file_account_proto_enumTypes,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
