// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: plugnmeet_datamessage.proto

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

type DataMsgType int32

const (
	DataMsgType_USER       DataMsgType = 0
	DataMsgType_SYSTEM     DataMsgType = 1
	DataMsgType_WHITEBOARD DataMsgType = 2
)

// Enum value maps for DataMsgType.
var (
	DataMsgType_name = map[int32]string{
		0: "USER",
		1: "SYSTEM",
		2: "WHITEBOARD",
	}
	DataMsgType_value = map[string]int32{
		"USER":       0,
		"SYSTEM":     1,
		"WHITEBOARD": 2,
	}
)

func (x DataMsgType) Enum() *DataMsgType {
	p := new(DataMsgType)
	*p = x
	return p
}

func (x DataMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_plugnmeet_datamessage_proto_enumTypes[0].Descriptor()
}

func (DataMsgType) Type() protoreflect.EnumType {
	return &file_plugnmeet_datamessage_proto_enumTypes[0]
}

func (x DataMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataMsgType.Descriptor instead.
func (DataMsgType) EnumDescriptor() ([]byte, []int) {
	return file_plugnmeet_datamessage_proto_rawDescGZIP(), []int{0}
}

type DataMsgBodyType int32

const (
	// SYSTEM type
	DataMsgBodyType_RAISE_HAND                   DataMsgBodyType = 0
	DataMsgBodyType_LOWER_HAND                   DataMsgBodyType = 1
	DataMsgBodyType_OTHER_USER_LOWER_HAND        DataMsgBodyType = 2
	DataMsgBodyType_FILE_UPLOAD                  DataMsgBodyType = 3
	DataMsgBodyType_INFO                         DataMsgBodyType = 4
	DataMsgBodyType_ALERT                        DataMsgBodyType = 5
	DataMsgBodyType_SEND_CHAT_MSGS               DataMsgBodyType = 6
	DataMsgBodyType_RENEW_TOKEN                  DataMsgBodyType = 7
	DataMsgBodyType_UPDATE_LOCK_SETTINGS         DataMsgBodyType = 8
	DataMsgBodyType_INIT_WHITEBOARD              DataMsgBodyType = 9
	DataMsgBodyType_USER_VISIBILITY_CHANGE       DataMsgBodyType = 10
	DataMsgBodyType_EXTERNAL_MEDIA_PLAYER_EVENTS DataMsgBodyType = 11
	DataMsgBodyType_POLL_CREATED                 DataMsgBodyType = 12
	DataMsgBodyType_NEW_POLL_RESPONSE            DataMsgBodyType = 13
	DataMsgBodyType_POLL_CLOSED                  DataMsgBodyType = 14
	DataMsgBodyType_JOIN_BREAKOUT_ROOM           DataMsgBodyType = 15
	// USER type
	DataMsgBodyType_CHAT DataMsgBodyType = 16
	// WHITEBOARD type
	DataMsgBodyType_SCENE_UPDATE                DataMsgBodyType = 17
	DataMsgBodyType_POINTER_UPDATE              DataMsgBodyType = 18
	DataMsgBodyType_ADD_WHITEBOARD_FILE         DataMsgBodyType = 19
	DataMsgBodyType_ADD_WHITEBOARD_OFFICE_FILE  DataMsgBodyType = 20
	DataMsgBodyType_PAGE_CHANGE                 DataMsgBodyType = 21
	DataMsgBodyType_WHITEBOARD_APP_STATE_CHANGE DataMsgBodyType = 22
)

// Enum value maps for DataMsgBodyType.
var (
	DataMsgBodyType_name = map[int32]string{
		0:  "RAISE_HAND",
		1:  "LOWER_HAND",
		2:  "OTHER_USER_LOWER_HAND",
		3:  "FILE_UPLOAD",
		4:  "INFO",
		5:  "ALERT",
		6:  "SEND_CHAT_MSGS",
		7:  "RENEW_TOKEN",
		8:  "UPDATE_LOCK_SETTINGS",
		9:  "INIT_WHITEBOARD",
		10: "USER_VISIBILITY_CHANGE",
		11: "EXTERNAL_MEDIA_PLAYER_EVENTS",
		12: "POLL_CREATED",
		13: "NEW_POLL_RESPONSE",
		14: "POLL_CLOSED",
		15: "JOIN_BREAKOUT_ROOM",
		16: "CHAT",
		17: "SCENE_UPDATE",
		18: "POINTER_UPDATE",
		19: "ADD_WHITEBOARD_FILE",
		20: "ADD_WHITEBOARD_OFFICE_FILE",
		21: "PAGE_CHANGE",
		22: "WHITEBOARD_APP_STATE_CHANGE",
	}
	DataMsgBodyType_value = map[string]int32{
		"RAISE_HAND":                   0,
		"LOWER_HAND":                   1,
		"OTHER_USER_LOWER_HAND":        2,
		"FILE_UPLOAD":                  3,
		"INFO":                         4,
		"ALERT":                        5,
		"SEND_CHAT_MSGS":               6,
		"RENEW_TOKEN":                  7,
		"UPDATE_LOCK_SETTINGS":         8,
		"INIT_WHITEBOARD":              9,
		"USER_VISIBILITY_CHANGE":       10,
		"EXTERNAL_MEDIA_PLAYER_EVENTS": 11,
		"POLL_CREATED":                 12,
		"NEW_POLL_RESPONSE":            13,
		"POLL_CLOSED":                  14,
		"JOIN_BREAKOUT_ROOM":           15,
		"CHAT":                         16,
		"SCENE_UPDATE":                 17,
		"POINTER_UPDATE":               18,
		"ADD_WHITEBOARD_FILE":          19,
		"ADD_WHITEBOARD_OFFICE_FILE":   20,
		"PAGE_CHANGE":                  21,
		"WHITEBOARD_APP_STATE_CHANGE":  22,
	}
)

func (x DataMsgBodyType) Enum() *DataMsgBodyType {
	p := new(DataMsgBodyType)
	*p = x
	return p
}

func (x DataMsgBodyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataMsgBodyType) Descriptor() protoreflect.EnumDescriptor {
	return file_plugnmeet_datamessage_proto_enumTypes[1].Descriptor()
}

func (DataMsgBodyType) Type() protoreflect.EnumType {
	return &file_plugnmeet_datamessage_proto_enumTypes[1]
}

func (x DataMsgBodyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataMsgBodyType.Descriptor instead.
func (DataMsgBodyType) EnumDescriptor() ([]byte, []int) {
	return file_plugnmeet_datamessage_proto_rawDescGZIP(), []int{1}
}

type DataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      DataMsgType  `protobuf:"varint,1,opt,name=type,proto3,enum=plugnmeet.DataMsgType" json:"type,omitempty"`
	MessageId *string      `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3,oneof" json:"message_id,omitempty"`
	RoomSid   string       `protobuf:"bytes,3,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	RoomId    string       `protobuf:"bytes,4,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	To        *string      `protobuf:"bytes,5,opt,name=to,proto3,oneof" json:"to,omitempty"`
	Body      *DataMsgBody `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *DataMessage) Reset() {
	*x = DataMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_datamessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataMessage) ProtoMessage() {}

func (x *DataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_datamessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataMessage.ProtoReflect.Descriptor instead.
func (*DataMessage) Descriptor() ([]byte, []int) {
	return file_plugnmeet_datamessage_proto_rawDescGZIP(), []int{0}
}

func (x *DataMessage) GetType() DataMsgType {
	if x != nil {
		return x.Type
	}
	return DataMsgType_USER
}

func (x *DataMessage) GetMessageId() string {
	if x != nil && x.MessageId != nil {
		return *x.MessageId
	}
	return ""
}

func (x *DataMessage) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *DataMessage) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *DataMessage) GetTo() string {
	if x != nil && x.To != nil {
		return *x.To
	}
	return ""
}

func (x *DataMessage) GetBody() *DataMsgBody {
	if x != nil {
		return x.Body
	}
	return nil
}

type DataMsgBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      DataMsgBodyType `protobuf:"varint,1,opt,name=type,proto3,enum=plugnmeet.DataMsgBodyType" json:"type,omitempty"`
	MessageId *string         `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3,oneof" json:"message_id,omitempty"`
	Time      *string         `protobuf:"bytes,3,opt,name=time,proto3,oneof" json:"time,omitempty"`
	From      *DataMsgReqFrom `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	Msg       string          `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	IsPrivate *uint32         `protobuf:"varint,6,opt,name=is_private,json=isPrivate,proto3,oneof" json:"is_private,omitempty"`
}

func (x *DataMsgBody) Reset() {
	*x = DataMsgBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_datamessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataMsgBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataMsgBody) ProtoMessage() {}

func (x *DataMsgBody) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_datamessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataMsgBody.ProtoReflect.Descriptor instead.
func (*DataMsgBody) Descriptor() ([]byte, []int) {
	return file_plugnmeet_datamessage_proto_rawDescGZIP(), []int{1}
}

func (x *DataMsgBody) GetType() DataMsgBodyType {
	if x != nil {
		return x.Type
	}
	return DataMsgBodyType_RAISE_HAND
}

func (x *DataMsgBody) GetMessageId() string {
	if x != nil && x.MessageId != nil {
		return *x.MessageId
	}
	return ""
}

func (x *DataMsgBody) GetTime() string {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return ""
}

func (x *DataMsgBody) GetFrom() *DataMsgReqFrom {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *DataMsgBody) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DataMsgBody) GetIsPrivate() uint32 {
	if x != nil && x.IsPrivate != nil {
		return *x.IsPrivate
	}
	return 0
}

type DataMsgReqFrom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid    string  `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	UserId string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *DataMsgReqFrom) Reset() {
	*x = DataMsgReqFrom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_datamessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataMsgReqFrom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataMsgReqFrom) ProtoMessage() {}

func (x *DataMsgReqFrom) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_datamessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataMsgReqFrom.ProtoReflect.Descriptor instead.
func (*DataMsgReqFrom) Descriptor() ([]byte, []int) {
	return file_plugnmeet_datamessage_proto_rawDescGZIP(), []int{2}
}

func (x *DataMsgReqFrom) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *DataMsgReqFrom) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DataMsgReqFrom) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_plugnmeet_datamessage_proto protoreflect.FileDescriptor

var file_plugnmeet_datamessage_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70,
	0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x22, 0xe8, 0x01, 0x0a, 0x0b, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65,
	0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d,
	0x53, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x02, 0x74, 0x6f, 0x88, 0x01,
	0x01, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x74, 0x6f, 0x22, 0x86, 0x02, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x2d, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x71, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x5d, 0x0a, 0x0e,
	0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x33, 0x0a, 0x0b, 0x44,
	0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53,
	0x45, 0x52, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x48, 0x49, 0x54, 0x45, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x02,
	0x2a, 0xf1, 0x03, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x41, 0x49, 0x53, 0x45, 0x5f, 0x48, 0x41,
	0x4e, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x48, 0x41,
	0x4e, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x03,
	0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c,
	0x45, 0x52, 0x54, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x48,
	0x41, 0x54, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x4e,
	0x45, 0x57, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e,
	0x47, 0x53, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x57, 0x48, 0x49,
	0x54, 0x45, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x09, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x10, 0x0a, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4f, 0x4c, 0x4c, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x45, 0x57,
	0x5f, 0x50, 0x4f, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x0d,
	0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4f, 0x4c, 0x4c, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10,
	0x0e, 0x12, 0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x4f,
	0x55, 0x54, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x0f, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x48, 0x41,
	0x54, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x10, 0x11, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x12, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x44, 0x44,
	0x5f, 0x57, 0x48, 0x49, 0x54, 0x45, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x10, 0x13, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x44, 0x44, 0x5f, 0x57, 0x48, 0x49, 0x54, 0x45, 0x42,
	0x4f, 0x41, 0x52, 0x44, 0x5f, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x10, 0x14, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x10, 0x15, 0x12, 0x1f, 0x0a, 0x1b, 0x57, 0x48, 0x49, 0x54, 0x45, 0x42, 0x4f, 0x41, 0x52,
	0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x10, 0x16, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_plugnmeet_datamessage_proto_rawDescOnce sync.Once
	file_plugnmeet_datamessage_proto_rawDescData = file_plugnmeet_datamessage_proto_rawDesc
)

func file_plugnmeet_datamessage_proto_rawDescGZIP() []byte {
	file_plugnmeet_datamessage_proto_rawDescOnce.Do(func() {
		file_plugnmeet_datamessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_datamessage_proto_rawDescData)
	})
	return file_plugnmeet_datamessage_proto_rawDescData
}

var file_plugnmeet_datamessage_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_plugnmeet_datamessage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_plugnmeet_datamessage_proto_goTypes = []interface{}{
	(DataMsgType)(0),       // 0: plugnmeet.DataMsgType
	(DataMsgBodyType)(0),   // 1: plugnmeet.DataMsgBodyType
	(*DataMessage)(nil),    // 2: plugnmeet.DataMessage
	(*DataMsgBody)(nil),    // 3: plugnmeet.DataMsgBody
	(*DataMsgReqFrom)(nil), // 4: plugnmeet.DataMsgReqFrom
}
var file_plugnmeet_datamessage_proto_depIdxs = []int32{
	0, // 0: plugnmeet.DataMessage.type:type_name -> plugnmeet.DataMsgType
	3, // 1: plugnmeet.DataMessage.body:type_name -> plugnmeet.DataMsgBody
	1, // 2: plugnmeet.DataMsgBody.type:type_name -> plugnmeet.DataMsgBodyType
	4, // 3: plugnmeet.DataMsgBody.from:type_name -> plugnmeet.DataMsgReqFrom
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_plugnmeet_datamessage_proto_init() }
func file_plugnmeet_datamessage_proto_init() {
	if File_plugnmeet_datamessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_datamessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataMessage); i {
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
		file_plugnmeet_datamessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataMsgBody); i {
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
		file_plugnmeet_datamessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataMsgReqFrom); i {
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
	file_plugnmeet_datamessage_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_plugnmeet_datamessage_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_plugnmeet_datamessage_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_datamessage_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_datamessage_proto_goTypes,
		DependencyIndexes: file_plugnmeet_datamessage_proto_depIdxs,
		EnumInfos:         file_plugnmeet_datamessage_proto_enumTypes,
		MessageInfos:      file_plugnmeet_datamessage_proto_msgTypes,
	}.Build()
	File_plugnmeet_datamessage_proto = out.File
	file_plugnmeet_datamessage_proto_rawDesc = nil
	file_plugnmeet_datamessage_proto_goTypes = nil
	file_plugnmeet_datamessage_proto_depIdxs = nil
}
