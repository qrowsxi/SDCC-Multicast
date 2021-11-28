// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: SeqClockService.proto

package api

import (
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

type MessageType int32

const (
	MessageType_APPLICATION MessageType = 0
	MessageType_ACK         MessageType = 1
	MessageType_SYSTEM      MessageType = 2
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "APPLICATION",
		1: "ACK",
		2: "SYSTEM",
	}
	MessageType_value = map[string]int32{
		"APPLICATION": 0,
		"ACK":         1,
		"SYSTEM":      2,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_SeqClockService_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_SeqClockService_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_SeqClockService_proto_rawDescGZIP(), []int{0}
}

type MessageSeq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Clock uint64      `protobuf:"varint,2,opt,name=clock,proto3" json:"clock,omitempty"`
	Src   string      `protobuf:"bytes,3,opt,name=src,proto3" json:"src,omitempty"`
	Id    string      `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Data  string      `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MessageSeq) Reset() {
	*x = MessageSeq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SeqClockService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSeq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSeq) ProtoMessage() {}

func (x *MessageSeq) ProtoReflect() protoreflect.Message {
	mi := &file_SeqClockService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSeq.ProtoReflect.Descriptor instead.
func (*MessageSeq) Descriptor() ([]byte, []int) {
	return file_SeqClockService_proto_rawDescGZIP(), []int{0}
}

func (x *MessageSeq) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_APPLICATION
}

func (x *MessageSeq) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *MessageSeq) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *MessageSeq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageSeq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type EnqueueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnqueueReply) Reset() {
	*x = EnqueueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SeqClockService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueReply) ProtoMessage() {}

func (x *EnqueueReply) ProtoReflect() protoreflect.Message {
	mi := &file_SeqClockService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueReply.ProtoReflect.Descriptor instead.
func (*EnqueueReply) Descriptor() ([]byte, []int) {
	return file_SeqClockService_proto_rawDescGZIP(), []int{1}
}

var File_SeqClockService_proto protoreflect.FileDescriptor

var file_SeqClockService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x65, 0x71, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x72, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2a, 0x33, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x02, 0x32, 0x38, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x07, 0x45,
	0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x0b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x64, 0x63, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f,
	0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SeqClockService_proto_rawDescOnce sync.Once
	file_SeqClockService_proto_rawDescData = file_SeqClockService_proto_rawDesc
)

func file_SeqClockService_proto_rawDescGZIP() []byte {
	file_SeqClockService_proto_rawDescOnce.Do(func() {
		file_SeqClockService_proto_rawDescData = protoimpl.X.CompressGZIP(file_SeqClockService_proto_rawDescData)
	})
	return file_SeqClockService_proto_rawDescData
}

var file_SeqClockService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SeqClockService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SeqClockService_proto_goTypes = []interface{}{
	(MessageType)(0),     // 0: MessageType
	(*MessageSeq)(nil),   // 1: MessageSeq
	(*EnqueueReply)(nil), // 2: EnqueueReply
}
var file_SeqClockService_proto_depIdxs = []int32{
	0, // 0: MessageSeq.type:type_name -> MessageType
	1, // 1: MessageQueueSeq.Enqueue:input_type -> MessageSeq
	2, // 2: MessageQueueSeq.Enqueue:output_type -> EnqueueReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SeqClockService_proto_init() }
func file_SeqClockService_proto_init() {
	if File_SeqClockService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SeqClockService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSeq); i {
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
		file_SeqClockService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueReply); i {
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
			RawDescriptor: file_SeqClockService_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_SeqClockService_proto_goTypes,
		DependencyIndexes: file_SeqClockService_proto_depIdxs,
		EnumInfos:         file_SeqClockService_proto_enumTypes,
		MessageInfos:      file_SeqClockService_proto_msgTypes,
	}.Build()
	File_SeqClockService_proto = out.File
	file_SeqClockService_proto_rawDesc = nil
	file_SeqClockService_proto_goTypes = nil
	file_SeqClockService_proto_depIdxs = nil
}
