// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: messages.proto

package messages

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts       int64   `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Serial   string  `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Tenant   string  `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Network  string  `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	Model    string  `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	Tags     string  `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	Kind     string  `protobuf:"bytes,7,opt,name=kind,proto3" json:"kind,omitempty"`
	Value    float32 `protobuf:"fixed32,8,opt,name=value,proto3" json:"value,omitempty"`
	MetaJSON string  `protobuf:"bytes,9,opt,name=metaJSON,proto3" json:"metaJSON,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *Event) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *Event) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *Event) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Event) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Event) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Event) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Event) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Event) GetMetaJSON() string {
	if x != nil {
		return x.MetaJSON
	}
	return ""
}

type Init struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WASMUrl            string `protobuf:"bytes,1,opt,name=WASMUrl,proto3" json:"WASMUrl,omitempty"`
	OutputsJSON        string `protobuf:"bytes,2,opt,name=outputsJSON,proto3" json:"outputsJSON,omitempty"`
	EnvVars            string `protobuf:"bytes,3,opt,name=envVars,proto3" json:"envVars,omitempty"`
	AcceptsEmptyOutput bool   `protobuf:"varint,4,opt,name=acceptsEmptyOutput,proto3" json:"acceptsEmptyOutput,omitempty"`
}

func (x *Init) Reset() {
	*x = Init{}
	mi := &file_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Init) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Init) ProtoMessage() {}

func (x *Init) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Init.ProtoReflect.Descriptor instead.
func (*Init) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Init) GetWASMUrl() string {
	if x != nil {
		return x.WASMUrl
	}
	return ""
}

func (x *Init) GetOutputsJSON() string {
	if x != nil {
		return x.OutputsJSON
	}
	return ""
}

func (x *Init) GetEnvVars() string {
	if x != nil {
		return x.EnvVars
	}
	return ""
}

func (x *Init) GetAcceptsEmptyOutput() bool {
	if x != nil {
		return x.AcceptsEmptyOutput
	}
	return false
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x8c,
	0x01, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x41, 0x53, 0x4d, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x57, 0x41, 0x53, 0x4d, 0x55, 0x72,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x4a,
	0x53, 0x4f, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x12, 0x2e, 0x0a,
	0x12, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x61, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x31, 0x5a,
	0x2f, 0x62, 0x69, 0x7a, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x74, 0x2f, 0x57, 0x41, 0x53, 0x4d,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messages_proto_goTypes = []any{
	(*Event)(nil), // 0: messages.Event
	(*Init)(nil),  // 1: messages.Init
}
var file_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}