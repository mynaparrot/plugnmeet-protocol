// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: plugnmeet_auth_room.proto

package plugnmeet

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	livekit "github.com/livekit/protocol/livekit"
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

type GetActiveRoomInfoReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetActiveRoomInfoReq) Reset() {
	*x = GetActiveRoomInfoReq{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetActiveRoomInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActiveRoomInfoReq) ProtoMessage() {}

func (x *GetActiveRoomInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[0]
	if x != nil {
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
	state            protoimpl.MessageState     `protogen:"open.v1"`
	Status           bool                       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg              string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	RoomInfo         *ActiveRoomInfo            `protobuf:"bytes,3,opt,name=room_info,json=roomInfo,proto3,oneof" json:"room_info,omitempty"`
	ParticipantsInfo []*livekit.ParticipantInfo `protobuf:"bytes,4,rep,name=participants_info,json=participantsInfo,proto3" json:"participants_info,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ActiveRoomInfoRes) Reset() {
	*x = ActiveRoomInfoRes{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActiveRoomInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRoomInfoRes) ProtoMessage() {}

func (x *ActiveRoomInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[1]
	if x != nil {
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
	state              protoimpl.MessageState `protogen:"open.v1"`
	RoomTitle          string                 `protobuf:"bytes,1,opt,name=room_title,json=roomTitle,proto3" json:"room_title,omitempty"`
	RoomId             string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Sid                string                 `protobuf:"bytes,3,opt,name=sid,proto3" json:"sid,omitempty"`
	JoinedParticipants int64                  `protobuf:"varint,4,opt,name=joined_participants,json=joinedParticipants,proto3" json:"joined_participants,omitempty"`
	IsRunning          int32                  `protobuf:"varint,5,opt,name=is_running,json=isRunning,proto3" json:"is_running,omitempty"`
	IsRecording        int32                  `protobuf:"varint,6,opt,name=is_recording,json=isRecording,proto3" json:"is_recording,omitempty"`
	IsActiveRtmp       int32                  `protobuf:"varint,7,opt,name=is_active_rtmp,json=isActiveRtmp,proto3" json:"is_active_rtmp,omitempty"`
	WebhookUrl         string                 `protobuf:"bytes,8,opt,name=webhook_url,json=webhookUrl,proto3" json:"webhook_url,omitempty"`
	IsBreakoutRoom     int32                  `protobuf:"varint,9,opt,name=is_breakout_room,json=isBreakoutRoom,proto3" json:"is_breakout_room,omitempty"`
	ParentRoomId       string                 `protobuf:"bytes,10,opt,name=parent_room_id,json=parentRoomId,proto3" json:"parent_room_id,omitempty"`
	CreationTime       int64                  `protobuf:"varint,11,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	Metadata           string                 `protobuf:"bytes,12,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ActiveRoomInfo) Reset() {
	*x = ActiveRoomInfo{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActiveRoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRoomInfo) ProtoMessage() {}

func (x *ActiveRoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[2]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomEndReq) Reset() {
	*x = RoomEndReq{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomEndReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomEndReq) ProtoMessage() {}

func (x *RoomEndReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[3]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomEndRes) Reset() {
	*x = RoomEndRes{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomEndRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomEndRes) ProtoMessage() {}

func (x *RoomEndRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[4]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsRoomActiveReq) Reset() {
	*x = IsRoomActiveReq{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsRoomActiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsRoomActiveReq) ProtoMessage() {}

func (x *IsRoomActiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[5]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	IsActive      bool                   `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Msg           string                 `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsRoomActiveRes) Reset() {
	*x = IsRoomActiveRes{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsRoomActiveRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsRoomActiveRes) ProtoMessage() {}

func (x *IsRoomActiveRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[6]
	if x != nil {
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

func (x *IsRoomActiveRes) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *IsRoomActiveRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ActiveRoomWithParticipant struct {
	state            protoimpl.MessageState     `protogen:"open.v1"`
	RoomInfo         *ActiveRoomInfo            `protobuf:"bytes,3,opt,name=room_info,json=roomInfo,proto3,oneof" json:"room_info,omitempty"`
	ParticipantsInfo []*livekit.ParticipantInfo `protobuf:"bytes,4,rep,name=participants_info,json=participantsInfo,proto3" json:"participants_info,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ActiveRoomWithParticipant) Reset() {
	*x = ActiveRoomWithParticipant{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActiveRoomWithParticipant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRoomWithParticipant) ProtoMessage() {}

func (x *ActiveRoomWithParticipant) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRoomWithParticipant.ProtoReflect.Descriptor instead.
func (*ActiveRoomWithParticipant) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{7}
}

func (x *ActiveRoomWithParticipant) GetRoomInfo() *ActiveRoomInfo {
	if x != nil {
		return x.RoomInfo
	}
	return nil
}

func (x *ActiveRoomWithParticipant) GetParticipantsInfo() []*livekit.ParticipantInfo {
	if x != nil {
		return x.ParticipantsInfo
	}
	return nil
}

type GetActiveRoomInfoRes struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Status        bool                       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Room          *ActiveRoomWithParticipant `protobuf:"bytes,3,opt,name=room,proto3" json:"room,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetActiveRoomInfoRes) Reset() {
	*x = GetActiveRoomInfoRes{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetActiveRoomInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActiveRoomInfoRes) ProtoMessage() {}

func (x *GetActiveRoomInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActiveRoomInfoRes.ProtoReflect.Descriptor instead.
func (*GetActiveRoomInfoRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{8}
}

func (x *GetActiveRoomInfoRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetActiveRoomInfoRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetActiveRoomInfoRes) GetRoom() *ActiveRoomWithParticipant {
	if x != nil {
		return x.Room
	}
	return nil
}

type GetActiveRoomsInfoRes struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Status        bool                         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Rooms         []*ActiveRoomWithParticipant `protobuf:"bytes,3,rep,name=rooms,proto3" json:"rooms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetActiveRoomsInfoRes) Reset() {
	*x = GetActiveRoomsInfoRes{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetActiveRoomsInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActiveRoomsInfoRes) ProtoMessage() {}

func (x *GetActiveRoomsInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActiveRoomsInfoRes.ProtoReflect.Descriptor instead.
func (*GetActiveRoomsInfoRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{9}
}

func (x *GetActiveRoomsInfoRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetActiveRoomsInfoRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetActiveRoomsInfoRes) GetRooms() []*ActiveRoomWithParticipant {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type PastRoomInfo struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	RoomTitle          string                 `protobuf:"bytes,1,opt,name=room_title,json=roomTitle,proto3" json:"room_title,omitempty"`
	RoomId             string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomSid            string                 `protobuf:"bytes,3,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	JoinedParticipants int64                  `protobuf:"varint,4,opt,name=joined_participants,json=joinedParticipants,proto3" json:"joined_participants,omitempty"`
	WebhookUrl         string                 `protobuf:"bytes,5,opt,name=webhook_url,json=webhookUrl,proto3" json:"webhook_url,omitempty"`
	Created            string                 `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Ended              string                 `protobuf:"bytes,7,opt,name=ended,proto3" json:"ended,omitempty"`
	AnalyticsFileId    string                 `protobuf:"bytes,8,opt,name=analytics_file_id,json=analyticsFileId,proto3" json:"analytics_file_id,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *PastRoomInfo) Reset() {
	*x = PastRoomInfo{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PastRoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PastRoomInfo) ProtoMessage() {}

func (x *PastRoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PastRoomInfo.ProtoReflect.Descriptor instead.
func (*PastRoomInfo) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{10}
}

func (x *PastRoomInfo) GetRoomTitle() string {
	if x != nil {
		return x.RoomTitle
	}
	return ""
}

func (x *PastRoomInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *PastRoomInfo) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *PastRoomInfo) GetJoinedParticipants() int64 {
	if x != nil {
		return x.JoinedParticipants
	}
	return 0
}

func (x *PastRoomInfo) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

func (x *PastRoomInfo) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *PastRoomInfo) GetEnded() string {
	if x != nil {
		return x.Ended
	}
	return ""
}

func (x *PastRoomInfo) GetAnalyticsFileId() string {
	if x != nil {
		return x.AnalyticsFileId
	}
	return ""
}

type FetchPastRoomsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomIds       []string               `protobuf:"bytes,1,rep,name=room_ids,json=roomIds,proto3" json:"room_ids,omitempty"`
	From          uint32                 `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Limit         uint32                 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy       string                 `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchPastRoomsReq) Reset() {
	*x = FetchPastRoomsReq{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchPastRoomsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPastRoomsReq) ProtoMessage() {}

func (x *FetchPastRoomsReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPastRoomsReq.ProtoReflect.Descriptor instead.
func (*FetchPastRoomsReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{11}
}

func (x *FetchPastRoomsReq) GetRoomIds() []string {
	if x != nil {
		return x.RoomIds
	}
	return nil
}

func (x *FetchPastRoomsReq) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *FetchPastRoomsReq) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchPastRoomsReq) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

type FetchPastRoomsResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TotalRooms    int64                  `protobuf:"varint,1,opt,name=total_rooms,json=totalRooms,proto3" json:"total_rooms,omitempty"`
	From          uint32                 `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Limit         uint32                 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy       string                 `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	RoomsList     []*PastRoomInfo        `protobuf:"bytes,5,rep,name=rooms_list,json=roomsList,proto3" json:"rooms_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchPastRoomsResult) Reset() {
	*x = FetchPastRoomsResult{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchPastRoomsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPastRoomsResult) ProtoMessage() {}

func (x *FetchPastRoomsResult) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPastRoomsResult.ProtoReflect.Descriptor instead.
func (*FetchPastRoomsResult) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{12}
}

func (x *FetchPastRoomsResult) GetTotalRooms() int64 {
	if x != nil {
		return x.TotalRooms
	}
	return 0
}

func (x *FetchPastRoomsResult) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *FetchPastRoomsResult) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchPastRoomsResult) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *FetchPastRoomsResult) GetRoomsList() []*PastRoomInfo {
	if x != nil {
		return x.RoomsList
	}
	return nil
}

type FetchPastRoomsRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Result        *FetchPastRoomsResult  `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchPastRoomsRes) Reset() {
	*x = FetchPastRoomsRes{}
	mi := &file_plugnmeet_auth_room_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchPastRoomsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPastRoomsRes) ProtoMessage() {}

func (x *FetchPastRoomsRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_room_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPastRoomsRes.ProtoReflect.Descriptor instead.
func (*FetchPastRoomsRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_room_proto_rawDescGZIP(), []int{13}
}

func (x *FetchPastRoomsRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *FetchPastRoomsRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FetchPastRoomsRes) GetResult() *FetchPastRoomsResult {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_plugnmeet_auth_room_proto protoreflect.FileDescriptor

const file_plugnmeet_auth_room_proto_rawDesc = "" +
	"\n" +
	"\x19plugnmeet_auth_room.proto\x12\tplugnmeet\x1a\x14livekit_models.proto\x1a\x1bbuf/validate/validate.proto\"7\n" +
	"\x14GetActiveRoomInfoReq\x12\x1f\n" +
	"\aroom_id\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\x06roomId\"\xcf\x01\n" +
	"\x11ActiveRoomInfoRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\x12;\n" +
	"\troom_info\x18\x03 \x01(\v2\x19.plugnmeet.ActiveRoomInfoH\x00R\broomInfo\x88\x01\x01\x12E\n" +
	"\x11participants_info\x18\x04 \x03(\v2\x18.livekit.ParticipantInfoR\x10participantsInfoB\f\n" +
	"\n" +
	"_room_info\"\xa5\x03\n" +
	"\x0eActiveRoomInfo\x12\x1d\n" +
	"\n" +
	"room_title\x18\x01 \x01(\tR\troomTitle\x12\x17\n" +
	"\aroom_id\x18\x02 \x01(\tR\x06roomId\x12\x10\n" +
	"\x03sid\x18\x03 \x01(\tR\x03sid\x12/\n" +
	"\x13joined_participants\x18\x04 \x01(\x03R\x12joinedParticipants\x12\x1d\n" +
	"\n" +
	"is_running\x18\x05 \x01(\x05R\tisRunning\x12!\n" +
	"\fis_recording\x18\x06 \x01(\x05R\visRecording\x12$\n" +
	"\x0eis_active_rtmp\x18\a \x01(\x05R\fisActiveRtmp\x12\x1f\n" +
	"\vwebhook_url\x18\b \x01(\tR\n" +
	"webhookUrl\x12(\n" +
	"\x10is_breakout_room\x18\t \x01(\x05R\x0eisBreakoutRoom\x12$\n" +
	"\x0eparent_room_id\x18\n" +
	" \x01(\tR\fparentRoomId\x12#\n" +
	"\rcreation_time\x18\v \x01(\x03R\fcreationTime\x12\x1a\n" +
	"\bmetadata\x18\f \x01(\tR\bmetadata\"-\n" +
	"\n" +
	"RoomEndReq\x12\x1f\n" +
	"\aroom_id\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\x06roomId\"6\n" +
	"\n" +
	"RoomEndRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\"2\n" +
	"\x0fIsRoomActiveReq\x12\x1f\n" +
	"\aroom_id\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\x06roomId\"X\n" +
	"\x0fIsRoomActiveRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x1b\n" +
	"\tis_active\x18\x02 \x01(\bR\bisActive\x12\x10\n" +
	"\x03msg\x18\x03 \x01(\tR\x03msg\"\xad\x01\n" +
	"\x19ActiveRoomWithParticipant\x12;\n" +
	"\troom_info\x18\x03 \x01(\v2\x19.plugnmeet.ActiveRoomInfoH\x00R\broomInfo\x88\x01\x01\x12E\n" +
	"\x11participants_info\x18\x04 \x03(\v2\x18.livekit.ParticipantInfoR\x10participantsInfoB\f\n" +
	"\n" +
	"_room_info\"z\n" +
	"\x14GetActiveRoomInfoRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\x128\n" +
	"\x04room\x18\x03 \x01(\v2$.plugnmeet.ActiveRoomWithParticipantR\x04room\"}\n" +
	"\x15GetActiveRoomsInfoRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\x12:\n" +
	"\x05rooms\x18\x03 \x03(\v2$.plugnmeet.ActiveRoomWithParticipantR\x05rooms\"\x8f\x02\n" +
	"\fPastRoomInfo\x12\x1d\n" +
	"\n" +
	"room_title\x18\x01 \x01(\tR\troomTitle\x12\x17\n" +
	"\aroom_id\x18\x02 \x01(\tR\x06roomId\x12\x19\n" +
	"\broom_sid\x18\x03 \x01(\tR\aroomSid\x12/\n" +
	"\x13joined_participants\x18\x04 \x01(\x03R\x12joinedParticipants\x12\x1f\n" +
	"\vwebhook_url\x18\x05 \x01(\tR\n" +
	"webhookUrl\x12\x18\n" +
	"\acreated\x18\x06 \x01(\tR\acreated\x12\x14\n" +
	"\x05ended\x18\a \x01(\tR\x05ended\x12*\n" +
	"\x11analytics_file_id\x18\b \x01(\tR\x0fanalyticsFileId\"s\n" +
	"\x11FetchPastRoomsReq\x12\x19\n" +
	"\broom_ids\x18\x01 \x03(\tR\aroomIds\x12\x12\n" +
	"\x04from\x18\x02 \x01(\rR\x04from\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\rR\x05limit\x12\x19\n" +
	"\border_by\x18\x04 \x01(\tR\aorderBy\"\xb4\x01\n" +
	"\x14FetchPastRoomsResult\x12\x1f\n" +
	"\vtotal_rooms\x18\x01 \x01(\x03R\n" +
	"totalRooms\x12\x12\n" +
	"\x04from\x18\x02 \x01(\rR\x04from\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\rR\x05limit\x12\x19\n" +
	"\border_by\x18\x04 \x01(\tR\aorderBy\x126\n" +
	"\n" +
	"rooms_list\x18\x05 \x03(\v2\x17.plugnmeet.PastRoomInfoR\troomsList\"v\n" +
	"\x11FetchPastRoomsRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\x127\n" +
	"\x06result\x18\x03 \x01(\v2\x1f.plugnmeet.FetchPastRoomsResultR\x06resultB\x9f\x01\n" +
	"\rcom.plugnmeetB\x16PlugnmeetAuthRoomProtoP\x01Z2github.com/mynaparrot/plugnmeet-protocol/plugnmeet\xa2\x02\x03PXX\xaa\x02\tPlugnmeet\xca\x02\tPlugnmeet\xe2\x02\x15Plugnmeet\\GPBMetadata\xea\x02\tPlugnmeetb\x06proto3"

var (
	file_plugnmeet_auth_room_proto_rawDescOnce sync.Once
	file_plugnmeet_auth_room_proto_rawDescData []byte
)

func file_plugnmeet_auth_room_proto_rawDescGZIP() []byte {
	file_plugnmeet_auth_room_proto_rawDescOnce.Do(func() {
		file_plugnmeet_auth_room_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_plugnmeet_auth_room_proto_rawDesc), len(file_plugnmeet_auth_room_proto_rawDesc)))
	})
	return file_plugnmeet_auth_room_proto_rawDescData
}

var file_plugnmeet_auth_room_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_plugnmeet_auth_room_proto_goTypes = []any{
	(*GetActiveRoomInfoReq)(nil),      // 0: plugnmeet.GetActiveRoomInfoReq
	(*ActiveRoomInfoRes)(nil),         // 1: plugnmeet.ActiveRoomInfoRes
	(*ActiveRoomInfo)(nil),            // 2: plugnmeet.ActiveRoomInfo
	(*RoomEndReq)(nil),                // 3: plugnmeet.RoomEndReq
	(*RoomEndRes)(nil),                // 4: plugnmeet.RoomEndRes
	(*IsRoomActiveReq)(nil),           // 5: plugnmeet.IsRoomActiveReq
	(*IsRoomActiveRes)(nil),           // 6: plugnmeet.IsRoomActiveRes
	(*ActiveRoomWithParticipant)(nil), // 7: plugnmeet.ActiveRoomWithParticipant
	(*GetActiveRoomInfoRes)(nil),      // 8: plugnmeet.GetActiveRoomInfoRes
	(*GetActiveRoomsInfoRes)(nil),     // 9: plugnmeet.GetActiveRoomsInfoRes
	(*PastRoomInfo)(nil),              // 10: plugnmeet.PastRoomInfo
	(*FetchPastRoomsReq)(nil),         // 11: plugnmeet.FetchPastRoomsReq
	(*FetchPastRoomsResult)(nil),      // 12: plugnmeet.FetchPastRoomsResult
	(*FetchPastRoomsRes)(nil),         // 13: plugnmeet.FetchPastRoomsRes
	(*livekit.ParticipantInfo)(nil),   // 14: livekit.ParticipantInfo
}
var file_plugnmeet_auth_room_proto_depIdxs = []int32{
	2,  // 0: plugnmeet.ActiveRoomInfoRes.room_info:type_name -> plugnmeet.ActiveRoomInfo
	14, // 1: plugnmeet.ActiveRoomInfoRes.participants_info:type_name -> livekit.ParticipantInfo
	2,  // 2: plugnmeet.ActiveRoomWithParticipant.room_info:type_name -> plugnmeet.ActiveRoomInfo
	14, // 3: plugnmeet.ActiveRoomWithParticipant.participants_info:type_name -> livekit.ParticipantInfo
	7,  // 4: plugnmeet.GetActiveRoomInfoRes.room:type_name -> plugnmeet.ActiveRoomWithParticipant
	7,  // 5: plugnmeet.GetActiveRoomsInfoRes.rooms:type_name -> plugnmeet.ActiveRoomWithParticipant
	10, // 6: plugnmeet.FetchPastRoomsResult.rooms_list:type_name -> plugnmeet.PastRoomInfo
	12, // 7: plugnmeet.FetchPastRoomsRes.result:type_name -> plugnmeet.FetchPastRoomsResult
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_plugnmeet_auth_room_proto_init() }
func file_plugnmeet_auth_room_proto_init() {
	if File_plugnmeet_auth_room_proto != nil {
		return
	}
	file_plugnmeet_auth_room_proto_msgTypes[1].OneofWrappers = []any{}
	file_plugnmeet_auth_room_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_plugnmeet_auth_room_proto_rawDesc), len(file_plugnmeet_auth_room_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_auth_room_proto_goTypes,
		DependencyIndexes: file_plugnmeet_auth_room_proto_depIdxs,
		MessageInfos:      file_plugnmeet_auth_room_proto_msgTypes,
	}.Build()
	File_plugnmeet_auth_room_proto = out.File
	file_plugnmeet_auth_room_proto_goTypes = nil
	file_plugnmeet_auth_room_proto_depIdxs = nil
}
