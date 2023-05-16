// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: plugnmeet_speech_services.proto

package plugnmeet

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

type SpeechToTextTranslationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId               string   `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	IsEnabled            bool     `protobuf:"varint,3,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty"`
	AllowedSpeechLangs   []string `protobuf:"bytes,4,rep,name=allowed_speech_langs,json=allowedSpeechLangs,proto3" json:"allowed_speech_langs,omitempty"`
	AllowedSpeechUsers   []string `protobuf:"bytes,5,rep,name=allowed_speech_users,json=allowedSpeechUsers,proto3" json:"allowed_speech_users,omitempty"`
	IsEnabledTranslation bool     `protobuf:"varint,6,opt,name=is_enabled_translation,json=isEnabledTranslation,proto3" json:"is_enabled_translation,omitempty"`
	AllowedTransLangs    []string `protobuf:"bytes,7,rep,name=allowed_trans_langs,json=allowedTransLangs,proto3" json:"allowed_trans_langs,omitempty"`
}

func (x *SpeechToTextTranslationReq) Reset() {
	*x = SpeechToTextTranslationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_speech_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeechToTextTranslationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeechToTextTranslationReq) ProtoMessage() {}

func (x *SpeechToTextTranslationReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_speech_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeechToTextTranslationReq.ProtoReflect.Descriptor instead.
func (*SpeechToTextTranslationReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_speech_services_proto_rawDescGZIP(), []int{0}
}

func (x *SpeechToTextTranslationReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *SpeechToTextTranslationReq) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

func (x *SpeechToTextTranslationReq) GetAllowedSpeechLangs() []string {
	if x != nil {
		return x.AllowedSpeechLangs
	}
	return nil
}

func (x *SpeechToTextTranslationReq) GetAllowedSpeechUsers() []string {
	if x != nil {
		return x.AllowedSpeechUsers
	}
	return nil
}

func (x *SpeechToTextTranslationReq) GetIsEnabledTranslation() bool {
	if x != nil {
		return x.IsEnabledTranslation
	}
	return false
}

func (x *SpeechToTextTranslationReq) GetAllowedTransLangs() []string {
	if x != nil {
		return x.AllowedTransLangs
	}
	return nil
}

type GenerateAzureTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *GenerateAzureTokenReq) Reset() {
	*x = GenerateAzureTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_speech_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAzureTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAzureTokenReq) ProtoMessage() {}

func (x *GenerateAzureTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_speech_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAzureTokenReq.ProtoReflect.Descriptor instead.
func (*GenerateAzureTokenReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_speech_services_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateAzureTokenReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type GenerateAzureTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    bool    `protobuf:"varint,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token  *string `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *GenerateAzureTokenRes) Reset() {
	*x = GenerateAzureTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_speech_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAzureTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAzureTokenRes) ProtoMessage() {}

func (x *GenerateAzureTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_speech_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAzureTokenRes.ProtoReflect.Descriptor instead.
func (*GenerateAzureTokenRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_speech_services_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateAzureTokenRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GenerateAzureTokenRes) GetMsg() bool {
	if x != nil {
		return x.Msg
	}
	return false
}

func (x *GenerateAzureTokenRes) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

var File_plugnmeet_speech_services_proto protoreflect.FileDescriptor

var file_plugnmeet_speech_services_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x22, 0x9e, 0x02, 0x0a,
	0x1a, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x54, 0x6f, 0x54, 0x65, 0x78, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68,
	0x4c, 0x61, 0x6e, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x6c,
	0x61, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x22, 0x30, 0x0a,
	0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22,
	0x66, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugnmeet_speech_services_proto_rawDescOnce sync.Once
	file_plugnmeet_speech_services_proto_rawDescData = file_plugnmeet_speech_services_proto_rawDesc
)

func file_plugnmeet_speech_services_proto_rawDescGZIP() []byte {
	file_plugnmeet_speech_services_proto_rawDescOnce.Do(func() {
		file_plugnmeet_speech_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_speech_services_proto_rawDescData)
	})
	return file_plugnmeet_speech_services_proto_rawDescData
}

var file_plugnmeet_speech_services_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_plugnmeet_speech_services_proto_goTypes = []interface{}{
	(*SpeechToTextTranslationReq)(nil), // 0: plugnmeet.SpeechToTextTranslationReq
	(*GenerateAzureTokenReq)(nil),      // 1: plugnmeet.GenerateAzureTokenReq
	(*GenerateAzureTokenRes)(nil),      // 2: plugnmeet.GenerateAzureTokenRes
}
var file_plugnmeet_speech_services_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plugnmeet_speech_services_proto_init() }
func file_plugnmeet_speech_services_proto_init() {
	if File_plugnmeet_speech_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_speech_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeechToTextTranslationReq); i {
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
		file_plugnmeet_speech_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateAzureTokenReq); i {
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
		file_plugnmeet_speech_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateAzureTokenRes); i {
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
	file_plugnmeet_speech_services_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_speech_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_speech_services_proto_goTypes,
		DependencyIndexes: file_plugnmeet_speech_services_proto_depIdxs,
		MessageInfos:      file_plugnmeet_speech_services_proto_msgTypes,
	}.Build()
	File_plugnmeet_speech_services_proto = out.File
	file_plugnmeet_speech_services_proto_rawDesc = nil
	file_plugnmeet_speech_services_proto_goTypes = nil
	file_plugnmeet_speech_services_proto_depIdxs = nil
}
