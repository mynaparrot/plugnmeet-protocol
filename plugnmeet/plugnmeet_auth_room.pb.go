// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: plugnmeet_auth_room.proto

package plugnmeet

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	livekit "github.com/livekit/protocol/livekit"
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

type GetActiveRoomInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *GetActiveRoomInfoReq) Reset() {
	*x = GetActiveRoomInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActiveRoomInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActiveRoomInfoReq) ProtoMessage() {}

func (x *GetActiveRoomInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActiveRoomInfoReq.ProtoReflect.Descriptor instead.
func (*GetActiveRoomInfoReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{0}
}

func (x *GetActiveRoomInfoReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type ActiveRoomInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status           bool                       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg              string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	RoomInfo         *ActiveRoomInfo            `protobuf:"bytes,3,opt,name=room_info,json=roomInfo,proto3,oneof" json:"room_info,omitempty"`
	ParticipantsInfo []*livekit.ParticipantInfo `protobuf:"bytes,4,rep,name=participants_info,json=participantsInfo,proto3" json:"participants_info,omitempty"`
}

func (x *ActiveRoomInfoRes) Reset() {
	*x = ActiveRoomInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRoomInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRoomInfoRes) ProtoMessage() {}

func (x *ActiveRoomInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRoomInfoRes.ProtoReflect.Descriptor instead.
func (*ActiveRoomInfoRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{1}
}

func (x *ActiveRoomInfoRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ActiveRoomInfoRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ActiveRoomInfoRes) GetRoomInfo() *ActiveRoomInfo {
	if x != nil {
		return x.RoomInfo
	}
	return nil
}

func (x *ActiveRoomInfoRes) GetParticipantsInfo() []*livekit.ParticipantInfo {
	if x != nil {
		return x.ParticipantsInfo
	}
	return nil
}

type ActiveRoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomTitle          string `protobuf:"bytes,1,opt,name=room_title,json=roomTitle,proto3" json:"room_title,omitempty"`
	RoomId             string `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Sid                string `protobuf:"bytes,3,opt,name=sid,proto3" json:"sid,omitempty"`
	JoinedParticipants int64  `protobuf:"varint,4,opt,name=joined_participants,json=joinedParticipants,proto3" json:"joined_participants,omitempty"`
	IsRunning          int32  `protobuf:"varint,5,opt,name=is_running,json=isRunning,proto3" json:"is_running,omitempty"`
	IsRecording        int32  `protobuf:"varint,6,opt,name=is_recording,json=isRecording,proto3" json:"is_recording,omitempty"`
	IsActiveRtmp       int32  `protobuf:"varint,7,opt,name=is_active_rtmp,json=isActiveRtmp,proto3" json:"is_active_rtmp,omitempty"`
	WebhookUrl         string `protobuf:"bytes,8,opt,name=webhook_url,json=webhookUrl,proto3" json:"webhook_url,omitempty"`
	IsBreakoutRoom     int32  `protobuf:"varint,9,opt,name=is_breakout_room,json=isBreakoutRoom,proto3" json:"is_breakout_room,omitempty"`
	ParentRoomId       string `protobuf:"bytes,10,opt,name=parent_room_id,json=parentRoomId,proto3" json:"parent_room_id,omitempty"`
	CreationTime       int64  `protobuf:"varint,11,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	Metadata           string `protobuf:"bytes,12,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ActiveRoomInfo) Reset() {
	*x = ActiveRoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRoomInfo) ProtoMessage() {}

func (x *ActiveRoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRoomInfo.ProtoReflect.Descriptor instead.
func (*ActiveRoomInfo) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{2}
}

func (x *ActiveRoomInfo) GetRoomTitle() string {
	if x != nil {
		return x.RoomTitle
	}
	return ""
}

func (x *ActiveRoomInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *ActiveRoomInfo) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *ActiveRoomInfo) GetJoinedParticipants() int64 {
	if x != nil {
		return x.JoinedParticipants
	}
	return 0
}

func (x *ActiveRoomInfo) GetIsRunning() int32 {
	if x != nil {
		return x.IsRunning
	}
	return 0
}

func (x *ActiveRoomInfo) GetIsRecording() int32 {
	if x != nil {
		return x.IsRecording
	}
	return 0
}

func (x *ActiveRoomInfo) GetIsActiveRtmp() int32 {
	if x != nil {
		return x.IsActiveRtmp
	}
	return 0
}

func (x *ActiveRoomInfo) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

func (x *ActiveRoomInfo) GetIsBreakoutRoom() int32 {
	if x != nil {
		return x.IsBreakoutRoom
	}
	return 0
}

func (x *ActiveRoomInfo) GetParentRoomId() string {
	if x != nil {
		return x.ParentRoomId
	}
	return ""
}

func (x *ActiveRoomInfo) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

func (x *ActiveRoomInfo) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type RoomEndReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *RoomEndReq) Reset() {
	*x = RoomEndReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomEndReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomEndReq) ProtoMessage() {}

func (x *RoomEndReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomEndReq.ProtoReflect.Descriptor instead.
func (*RoomEndReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{3}
}

func (x *RoomEndReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type RoomEndRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RoomEndRes) Reset() {
	*x = RoomEndRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomEndRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomEndRes) ProtoMessage() {}

func (x *RoomEndRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomEndRes.ProtoReflect.Descriptor instead.
func (*RoomEndRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{4}
}

func (x *RoomEndRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *RoomEndRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type IsRoomActiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *IsRoomActiveReq) Reset() {
	*x = IsRoomActiveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsRoomActiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsRoomActiveReq) ProtoMessage() {}

func (x *IsRoomActiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsRoomActiveReq.ProtoReflect.Descriptor instead.
func (*IsRoomActiveReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{5}
}

func (x *IsRoomActiveReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type IsRoomActiveRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *IsRoomActiveRes) Reset() {
	*x = IsRoomActiveRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsRoomActiveRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsRoomActiveRes) ProtoMessage() {}

func (x *IsRoomActiveRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsRoomActiveRes.ProtoReflect.Descriptor instead.
func (*IsRoomActiveRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{6}
}

func (x *IsRoomActiveRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *IsRoomActiveRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_plugnmeet_auth_room_proto protoreflect.FileDescriptor

var file_plugnmeet_auth_room_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x75,
	0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x40, 0x76, 0x30, 0x2e, 0x36, 0x2e, 0x37, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x40, 0x76, 0x31,
	0x2e, 0x30, 0x2e, 0x30, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x12, 0x21, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x3b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d,
	0x65, 0x65, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x88, 0x01,
	0x01, 0x12, 0x45, 0x0a, 0x11, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xa5, 0x03, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x6f, 0x6f, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x12, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x72, 0x74, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x74, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a,
	0x10, 0x69, 0x73, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x69, 0x73, 0x42, 0x72, 0x65, 0x61, 0x6b,
	0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2f,
	0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x07,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22,
	0x36, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x34, 0x0a, 0x0f, 0x49, 0x73, 0x52, 0x6f, 0x6f,
	0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x3b, 0x0a,
	0x0f, 0x49, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72,
	0x72, 0x6f, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugnmeet_auth_room_proto_rawDescOnce sync.Once
	file_plugnmeet_auth_room_proto_rawDescData = file_plugnmeet_auth_room_proto_rawDesc
)

func file_plugnmeet_auth_room_proto_rawDescGZIP() []byte {
	file_plugnmeet_auth_room_proto_rawDescOnce.Do(func() {
		file_plugnmeet_auth_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_auth_room_proto_rawDescData)
	})
	return file_plugnmeet_auth_room_proto_rawDescData
}

var file_plugnmeet_auth_room_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_plugnmeet_auth_room_proto_goTypes = []interface{}{
	(*GetActiveRoomInfoReq)(nil),    // 0: plugnmeet.GetActiveRoomInfoReq
	(*ActiveRoomInfoRes)(nil),       // 1: plugnmeet.ActiveRoomInfoRes
	(*ActiveRoomInfo)(nil),          // 2: plugnmeet.ActiveRoomInfo
	(*RoomEndReq)(nil),              // 3: plugnmeet.RoomEndReq
	(*RoomEndRes)(nil),              // 4: plugnmeet.RoomEndRes
	(*IsRoomActiveReq)(nil),         // 5: plugnmeet.IsRoomActiveReq
	(*IsRoomActiveRes)(nil),         // 6: plugnmeet.IsRoomActiveRes
	(*livekit.ParticipantInfo)(nil), // 7: livekit.ParticipantInfo
}
var file_plugnmeet_auth_room_proto_depIdxs = []int32{
	2, // 0: plugnmeet.ActiveRoomInfoRes.room_info:type_name -> plugnmeet.ActiveRoomInfo
	7, // 1: plugnmeet.ActiveRoomInfoRes.participants_info:type_name -> livekit.ParticipantInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_plugnmeet_auth_room_proto_init() }
func file_plugnmeet_auth_room_proto_init() {
	if File_plugnmeet_auth_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_auth_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActiveRoomInfoReq); i {
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
		file_plugnmeet_auth_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRoomInfoRes); i {
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
		file_plugnmeet_auth_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRoomInfo); i {
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
		file_plugnmeet_auth_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomEndReq); i {
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
		file_plugnmeet_auth_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomEndRes); i {
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
		file_plugnmeet_auth_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsRoomActiveReq); i {
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
		file_plugnmeet_auth_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsRoomActiveRes); i {
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
	file_plugnmeet_auth_room_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_auth_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_auth_room_proto_goTypes,
		DependencyIndexes: file_plugnmeet_auth_room_proto_depIdxs,
		MessageInfos:      file_plugnmeet_auth_room_proto_msgTypes,
	}.Build()
	File_plugnmeet_auth_room_proto = out.File
	file_plugnmeet_auth_room_proto_rawDesc = nil
	file_plugnmeet_auth_room_proto_goTypes = nil
	file_plugnmeet_auth_room_proto_depIdxs = nil
}
