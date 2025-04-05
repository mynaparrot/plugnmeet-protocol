// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: plugnmeet_ingress.proto

package plugnmeet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IngressInput int32

const (
	IngressInput_RTMP_INPUT IngressInput = 0
	IngressInput_WHIP_INPUT IngressInput = 1
)

// Enum value maps for IngressInput.
var (
	IngressInput_name = map[int32]string{
		0: "RTMP_INPUT",
		1: "WHIP_INPUT",
	}
	IngressInput_value = map[string]int32{
		"RTMP_INPUT": 0,
		"WHIP_INPUT": 1,
	}
)

func (x IngressInput) Enum() *IngressInput {
	p := new(IngressInput)
	*p = x
	return p
}

func (x IngressInput) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IngressInput) Descriptor() protoreflect.EnumDescriptor {
	return file_plugnmeet_ingress_proto_enumTypes[0].Descriptor()
}

func (IngressInput) Type() protoreflect.EnumType {
	return &file_plugnmeet_ingress_proto_enumTypes[0]
}

func (x IngressInput) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IngressInput.Descriptor instead.
func (IngressInput) EnumDescriptor() ([]byte, []int) {
	return file_plugnmeet_ingress_proto_rawDescGZIP(), []int{0}
}

type CreateIngressReq struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	InputType       IngressInput           `protobuf:"varint,1,opt,name=input_type,json=inputType,proto3,enum=plugnmeet.IngressInput" json:"input_type,omitempty"`
	ParticipantName string                 `protobuf:"bytes,2,opt,name=participant_name,json=participantName,proto3" json:"participant_name,omitempty"`
	RoomId          string                 `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreateIngressReq) Reset() {
	*x = CreateIngressReq{}
	mi := &file_plugnmeet_ingress_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateIngressReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIngressReq) ProtoMessage() {}

func (x *CreateIngressReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_ingress_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIngressReq.ProtoReflect.Descriptor instead.
func (*CreateIngressReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_ingress_proto_rawDescGZIP(), []int{0}
}

func (x *CreateIngressReq) GetInputType() IngressInput {
	if x != nil {
		return x.InputType
	}
	return IngressInput_RTMP_INPUT
}

func (x *CreateIngressReq) GetParticipantName() string {
	if x != nil {
		return x.ParticipantName
	}
	return ""
}

func (x *CreateIngressReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type CreateIngressRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	InputType     IngressInput           `protobuf:"varint,3,opt,name=input_type,json=inputType,proto3,enum=plugnmeet.IngressInput" json:"input_type,omitempty"`
	Url           string                 `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	StreamKey     string                 `protobuf:"bytes,5,opt,name=stream_key,json=streamKey,proto3" json:"stream_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateIngressRes) Reset() {
	*x = CreateIngressRes{}
	mi := &file_plugnmeet_ingress_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateIngressRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIngressRes) ProtoMessage() {}

func (x *CreateIngressRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_ingress_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIngressRes.ProtoReflect.Descriptor instead.
func (*CreateIngressRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_ingress_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIngressRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *CreateIngressRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateIngressRes) GetInputType() IngressInput {
	if x != nil {
		return x.InputType
	}
	return IngressInput_RTMP_INPUT
}

func (x *CreateIngressRes) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateIngressRes) GetStreamKey() string {
	if x != nil {
		return x.StreamKey
	}
	return ""
}

var File_plugnmeet_ingress_proto protoreflect.FileDescriptor

const file_plugnmeet_ingress_proto_rawDesc = "" +
	"\n" +
	"\x17plugnmeet_ingress.proto\x12\tplugnmeet\"\x8e\x01\n" +
	"\x10CreateIngressReq\x126\n" +
	"\n" +
	"input_type\x18\x01 \x01(\x0e2\x17.plugnmeet.IngressInputR\tinputType\x12)\n" +
	"\x10participant_name\x18\x02 \x01(\tR\x0fparticipantName\x12\x17\n" +
	"\aroom_id\x18\x03 \x01(\tR\x06roomId\"\xa5\x01\n" +
	"\x10CreateIngressRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\x126\n" +
	"\n" +
	"input_type\x18\x03 \x01(\x0e2\x17.plugnmeet.IngressInputR\tinputType\x12\x10\n" +
	"\x03url\x18\x04 \x01(\tR\x03url\x12\x1d\n" +
	"\n" +
	"stream_key\x18\x05 \x01(\tR\tstreamKey*.\n" +
	"\fIngressInput\x12\x0e\n" +
	"\n" +
	"RTMP_INPUT\x10\x00\x12\x0e\n" +
	"\n" +
	"WHIP_INPUT\x10\x01B\x9e\x01\n" +
	"\rcom.plugnmeetB\x15PlugnmeetIngressProtoP\x01Z2github.com/mynaparrot/plugnmeet-protocol/plugnmeet\xa2\x02\x03PXX\xaa\x02\tPlugnmeet\xca\x02\tPlugnmeet\xe2\x02\x15Plugnmeet\\GPBMetadata\xea\x02\tPlugnmeetb\x06proto3"

var (
	file_plugnmeet_ingress_proto_rawDescOnce sync.Once
	file_plugnmeet_ingress_proto_rawDescData []byte
)

func file_plugnmeet_ingress_proto_rawDescGZIP() []byte {
	file_plugnmeet_ingress_proto_rawDescOnce.Do(func() {
		file_plugnmeet_ingress_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_plugnmeet_ingress_proto_rawDesc), len(file_plugnmeet_ingress_proto_rawDesc)))
	})
	return file_plugnmeet_ingress_proto_rawDescData
}

var file_plugnmeet_ingress_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_plugnmeet_ingress_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_plugnmeet_ingress_proto_goTypes = []any{
	(IngressInput)(0),        // 0: plugnmeet.IngressInput
	(*CreateIngressReq)(nil), // 1: plugnmeet.CreateIngressReq
	(*CreateIngressRes)(nil), // 2: plugnmeet.CreateIngressRes
}
var file_plugnmeet_ingress_proto_depIdxs = []int32{
	0, // 0: plugnmeet.CreateIngressReq.input_type:type_name -> plugnmeet.IngressInput
	0, // 1: plugnmeet.CreateIngressRes.input_type:type_name -> plugnmeet.IngressInput
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_plugnmeet_ingress_proto_init() }
func file_plugnmeet_ingress_proto_init() {
	if File_plugnmeet_ingress_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_plugnmeet_ingress_proto_rawDesc), len(file_plugnmeet_ingress_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_ingress_proto_goTypes,
		DependencyIndexes: file_plugnmeet_ingress_proto_depIdxs,
		EnumInfos:         file_plugnmeet_ingress_proto_enumTypes,
		MessageInfos:      file_plugnmeet_ingress_proto_msgTypes,
	}.Build()
	File_plugnmeet_ingress_proto = out.File
	file_plugnmeet_ingress_proto_goTypes = nil
	file_plugnmeet_ingress_proto_depIdxs = nil
}
