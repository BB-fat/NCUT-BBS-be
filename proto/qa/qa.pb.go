// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: qa.proto

package qaPB

import (
	//_ "account"
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

type QuestionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId   int32    `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	CreateTime int64    `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Title      string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Views      int32    `protobuf:"varint,6,opt,name=views,proto3" json:"views,omitempty"`
	Content    string   `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
	Pictures   []string `protobuf:"bytes,10,rep,name=pictures,proto3" json:"pictures,omitempty"`
}

func (x *QuestionData) Reset() {
	*x = QuestionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionData) ProtoMessage() {}

func (x *QuestionData) ProtoReflect() protoreflect.Message {
	mi := &file_qa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionData.ProtoReflect.Descriptor instead.
func (*QuestionData) Descriptor() ([]byte, []int) {
	return file_qa_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QuestionData) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *QuestionData) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *QuestionData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *QuestionData) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *QuestionData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *QuestionData) GetPictures() []string {
	if x != nil {
		return x.Pictures
	}
	return nil
}

type CreateQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Pictures string `protobuf:"bytes,3,opt,name=pictures,proto3" json:"pictures,omitempty"`
}

func (x *CreateQuestionRequest) Reset() {
	*x = CreateQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestionRequest) ProtoMessage() {}

func (x *CreateQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestionRequest.ProtoReflect.Descriptor instead.
func (*CreateQuestionRequest) Descriptor() ([]byte, []int) {
	return file_qa_proto_rawDescGZIP(), []int{1}
}

func (x *CreateQuestionRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateQuestionRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateQuestionRequest) GetPictures() string {
	if x != nil {
		return x.Pictures
	}
	return ""
}

type CreateQuestionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *QuestionData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateQuestionReply) Reset() {
	*x = CreateQuestionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qa_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestionReply) ProtoMessage() {}

func (x *CreateQuestionReply) ProtoReflect() protoreflect.Message {
	mi := &file_qa_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestionReply.ProtoReflect.Descriptor instead.
func (*CreateQuestionReply) Descriptor() ([]byte, []int) {
	return file_qa_proto_rawDescGZIP(), []int{2}
}

func (x *CreateQuestionReply) GetData() *QuestionData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_qa_proto protoreflect.FileDescriptor

var file_qa_proto_rawDesc = []byte{
	0x0a, 0x08, 0x71, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x71, 0x61, 0x1a, 0x15,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x71, 0x61, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x71, 0x61, 0x50, 0x42,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qa_proto_rawDescOnce sync.Once
	file_qa_proto_rawDescData = file_qa_proto_rawDesc
)

func file_qa_proto_rawDescGZIP() []byte {
	file_qa_proto_rawDescOnce.Do(func() {
		file_qa_proto_rawDescData = protoimpl.X.CompressGZIP(file_qa_proto_rawDescData)
	})
	return file_qa_proto_rawDescData
}

var file_qa_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_qa_proto_goTypes = []interface{}{
	(*QuestionData)(nil),          // 0: qa.QuestionData
	(*CreateQuestionRequest)(nil), // 1: qa.CreateQuestionRequest
	(*CreateQuestionReply)(nil),   // 2: qa.CreateQuestionReply
}
var file_qa_proto_depIdxs = []int32{
	0, // 0: qa.CreateQuestionReply.data:type_name -> qa.QuestionData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_qa_proto_init() }
func file_qa_proto_init() {
	if File_qa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionData); i {
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
		file_qa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuestionRequest); i {
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
		file_qa_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuestionReply); i {
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
			RawDescriptor: file_qa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qa_proto_goTypes,
		DependencyIndexes: file_qa_proto_depIdxs,
		MessageInfos:      file_qa_proto_msgTypes,
	}.Build()
	File_qa_proto = out.File
	file_qa_proto_rawDesc = nil
	file_qa_proto_goTypes = nil
	file_qa_proto_depIdxs = nil
}
