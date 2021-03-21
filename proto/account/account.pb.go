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

type GetUserInfoReply_Sex int32

const (
	GetUserInfoReply_MALE   GetUserInfoReply_Sex = 0
	GetUserInfoReply_FEMALE GetUserInfoReply_Sex = 1
)

// Enum value maps for GetUserInfoReply_Sex.
var (
	GetUserInfoReply_Sex_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	GetUserInfoReply_Sex_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x GetUserInfoReply_Sex) Enum() *GetUserInfoReply_Sex {
	p := new(GetUserInfoReply_Sex)
	*p = x
	return p
}

func (x GetUserInfoReply_Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetUserInfoReply_Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_account_proto_enumTypes[0].Descriptor()
}

func (GetUserInfoReply_Sex) Type() protoreflect.EnumType {
	return &file_account_proto_enumTypes[0]
}

func (x GetUserInfoReply_Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetUserInfoReply_Sex.Descriptor instead.
func (GetUserInfoReply_Sex) EnumDescriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2, 0}
}

type GetUserInfoReply_AccountStatus int32

const (
	GetUserInfoReply_UNNAMED        GetUserInfoReply_AccountStatus = 0
	GetUserInfoReply_PENDING_REVIEW GetUserInfoReply_AccountStatus = 1
	GetUserInfoReply_NOT_PASS       GetUserInfoReply_AccountStatus = 2
	GetUserInfoReply_ACTIVE         GetUserInfoReply_AccountStatus = 3
	GetUserInfoReply_DISABLED       GetUserInfoReply_AccountStatus = 4
)

// Enum value maps for GetUserInfoReply_AccountStatus.
var (
	GetUserInfoReply_AccountStatus_name = map[int32]string{
		0: "UNNAMED",
		1: "PENDING_REVIEW",
		2: "NOT_PASS",
		3: "ACTIVE",
		4: "DISABLED",
	}
	GetUserInfoReply_AccountStatus_value = map[string]int32{
		"UNNAMED":        0,
		"PENDING_REVIEW": 1,
		"NOT_PASS":       2,
		"ACTIVE":         3,
		"DISABLED":       4,
	}
)

func (x GetUserInfoReply_AccountStatus) Enum() *GetUserInfoReply_AccountStatus {
	p := new(GetUserInfoReply_AccountStatus)
	*p = x
	return p
}

func (x GetUserInfoReply_AccountStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetUserInfoReply_AccountStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_account_proto_enumTypes[1].Descriptor()
}

func (GetUserInfoReply_AccountStatus) Type() protoreflect.EnumType {
	return &file_account_proto_enumTypes[1]
}

func (x GetUserInfoReply_AccountStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetUserInfoReply_AccountStatus.Descriptor instead.
func (GetUserInfoReply_AccountStatus) EnumDescriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2, 1}
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

type GetUserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountName   string                         `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	RealName      string                         `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	Sex           GetUserInfoReply_Sex           `protobuf:"varint,4,opt,name=sex,proto3,enum=account.GetUserInfoReply_Sex" json:"sex,omitempty"`
	College       string                         `protobuf:"bytes,5,opt,name=college,proto3" json:"college,omitempty"`
	AccountStatus GetUserInfoReply_AccountStatus `protobuf:"varint,6,opt,name=account_status,json=accountStatus,proto3,enum=account.GetUserInfoReply_AccountStatus" json:"account_status,omitempty"`
	Grade         int32                          `protobuf:"varint,7,opt,name=grade,proto3" json:"grade,omitempty"`
	Avatar        string                         `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *GetUserInfoReply) Reset() {
	*x = GetUserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReply) ProtoMessage() {}

func (x *GetUserInfoReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetUserInfoReply.ProtoReflect.Descriptor instead.
func (*GetUserInfoReply) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserInfoReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserInfoReply) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetUserInfoReply) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *GetUserInfoReply) GetSex() GetUserInfoReply_Sex {
	if x != nil {
		return x.Sex
	}
	return GetUserInfoReply_MALE
}

func (x *GetUserInfoReply) GetCollege() string {
	if x != nil {
		return x.College
	}
	return ""
}

func (x *GetUserInfoReply) GetAccountStatus() GetUserInfoReply_AccountStatus {
	if x != nil {
		return x.AccountStatus
	}
	return GetUserInfoReply_UNNAMED
}

func (x *GetUserInfoReply) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *GetUserInfoReply) GetAvatar() string {
	if x != nil {
		return x.Avatar
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
	0xa2, 0x03, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x65, 0x78,
	0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0x12,
	0x4e, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x1b, 0x0a,
	0x03, 0x53, 0x65, 0x78, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x22, 0x58, 0x0a, 0x0d, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4e, 0x41, 0x4d, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x04, 0x42, 0x0b, 0x5a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_account_proto_goTypes = []interface{}{
	(GetUserInfoReply_Sex)(0),           // 0: account.GetUserInfoReply.Sex
	(GetUserInfoReply_AccountStatus)(0), // 1: account.GetUserInfoReply.AccountStatus
	(*LoginRequest)(nil),                // 2: account.LoginRequest
	(*LoginReply)(nil),                  // 3: account.LoginReply
	(*GetUserInfoReply)(nil),            // 4: account.GetUserInfoReply
}
var file_account_proto_depIdxs = []int32{
	0, // 0: account.GetUserInfoReply.sex:type_name -> account.GetUserInfoReply.Sex
	1, // 1: account.GetUserInfoReply.account_status:type_name -> account.GetUserInfoReply.AccountStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
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
