// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: plugnmeet_nats_msg.proto

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

type NatsMsgServerToClientEvents int32

const (
	// initial data
	NatsMsgServerToClientEvents_INITIAL_DATA         NatsMsgServerToClientEvents = 0
	NatsMsgServerToClientEvents_JOINED_USERS_LIST    NatsMsgServerToClientEvents = 1
	NatsMsgServerToClientEvents_ROOM_METADATA_UPDATE NatsMsgServerToClientEvents = 2
	NatsMsgServerToClientEvents_USER_METADATA_UPDATE NatsMsgServerToClientEvents = 3
	NatsMsgServerToClientEvents_USER_JOINED          NatsMsgServerToClientEvents = 4
	NatsMsgServerToClientEvents_USER_DISCONNECTED    NatsMsgServerToClientEvents = 5
	NatsMsgServerToClientEvents_USER_OFFLINE         NatsMsgServerToClientEvents = 6
	NatsMsgServerToClientEvents_RESP_RENEW_PNM_TOKEN NatsMsgServerToClientEvents = 7
	NatsMsgServerToClientEvents_SYSTEM_NOTIFICATION  NatsMsgServerToClientEvents = 8
	NatsMsgServerToClientEvents_RESP_DATA_MSG        NatsMsgServerToClientEvents = 9
)

// Enum value maps for NatsMsgServerToClientEvents.
var (
	NatsMsgServerToClientEvents_name = map[int32]string{
		0: "INITIAL_DATA",
		1: "JOINED_USERS_LIST",
		2: "ROOM_METADATA_UPDATE",
		3: "USER_METADATA_UPDATE",
		4: "USER_JOINED",
		5: "USER_DISCONNECTED",
		6: "USER_OFFLINE",
		7: "RESP_RENEW_PNM_TOKEN",
		8: "SYSTEM_NOTIFICATION",
		9: "RESP_DATA_MSG",
	}
	NatsMsgServerToClientEvents_value = map[string]int32{
		"INITIAL_DATA":         0,
		"JOINED_USERS_LIST":    1,
		"ROOM_METADATA_UPDATE": 2,
		"USER_METADATA_UPDATE": 3,
		"USER_JOINED":          4,
		"USER_DISCONNECTED":    5,
		"USER_OFFLINE":         6,
		"RESP_RENEW_PNM_TOKEN": 7,
		"SYSTEM_NOTIFICATION":  8,
		"RESP_DATA_MSG":        9,
	}
)

func (x NatsMsgServerToClientEvents) Enum() *NatsMsgServerToClientEvents {
	p := new(NatsMsgServerToClientEvents)
	*p = x
	return p
}

func (x NatsMsgServerToClientEvents) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NatsMsgServerToClientEvents) Descriptor() protoreflect.EnumDescriptor {
	return file_plugnmeet_nats_msg_proto_enumTypes[0].Descriptor()
}

func (NatsMsgServerToClientEvents) Type() protoreflect.EnumType {
	return &file_plugnmeet_nats_msg_proto_enumTypes[0]
}

func (x NatsMsgServerToClientEvents) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NatsMsgServerToClientEvents.Descriptor instead.
func (NatsMsgServerToClientEvents) EnumDescriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{0}
}

type NatsMsgClientToServerEvents int32

const (
	NatsMsgClientToServerEvents_REQ_INITIAL_DATA    NatsMsgClientToServerEvents = 0
	NatsMsgClientToServerEvents_REQ_RENEW_PNM_TOKEN NatsMsgClientToServerEvents = 1
	NatsMsgClientToServerEvents_REQ_DATA_MSG        NatsMsgClientToServerEvents = 3
	NatsMsgClientToServerEvents_PING                NatsMsgClientToServerEvents = 4
)

// Enum value maps for NatsMsgClientToServerEvents.
var (
	NatsMsgClientToServerEvents_name = map[int32]string{
		0: "REQ_INITIAL_DATA",
		1: "REQ_RENEW_PNM_TOKEN",
		3: "REQ_DATA_MSG",
		4: "PING",
	}
	NatsMsgClientToServerEvents_value = map[string]int32{
		"REQ_INITIAL_DATA":    0,
		"REQ_RENEW_PNM_TOKEN": 1,
		"REQ_DATA_MSG":        3,
		"PING":                4,
	}
)

func (x NatsMsgClientToServerEvents) Enum() *NatsMsgClientToServerEvents {
	p := new(NatsMsgClientToServerEvents)
	*p = x
	return p
}

func (x NatsMsgClientToServerEvents) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NatsMsgClientToServerEvents) Descriptor() protoreflect.EnumDescriptor {
	return file_plugnmeet_nats_msg_proto_enumTypes[1].Descriptor()
}

func (NatsMsgClientToServerEvents) Type() protoreflect.EnumType {
	return &file_plugnmeet_nats_msg_proto_enumTypes[1]
}

func (x NatsMsgClientToServerEvents) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NatsMsgClientToServerEvents.Descriptor instead.
func (NatsMsgClientToServerEvents) EnumDescriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{1}
}

type NatsSystemNotificationTypes int32

const (
	NatsSystemNotificationTypes_NATS_SYSTEM_NOTIFICATION_INFO    NatsSystemNotificationTypes = 0
	NatsSystemNotificationTypes_NATS_SYSTEM_NOTIFICATION_WARNING NatsSystemNotificationTypes = 1
	NatsSystemNotificationTypes_NATS_SYSTEM_NOTIFICATION_ERROR   NatsSystemNotificationTypes = 2
)

// Enum value maps for NatsSystemNotificationTypes.
var (
	NatsSystemNotificationTypes_name = map[int32]string{
		0: "NATS_SYSTEM_NOTIFICATION_INFO",
		1: "NATS_SYSTEM_NOTIFICATION_WARNING",
		2: "NATS_SYSTEM_NOTIFICATION_ERROR",
	}
	NatsSystemNotificationTypes_value = map[string]int32{
		"NATS_SYSTEM_NOTIFICATION_INFO":    0,
		"NATS_SYSTEM_NOTIFICATION_WARNING": 1,
		"NATS_SYSTEM_NOTIFICATION_ERROR":   2,
	}
)

func (x NatsSystemNotificationTypes) Enum() *NatsSystemNotificationTypes {
	p := new(NatsSystemNotificationTypes)
	*p = x
	return p
}

func (x NatsSystemNotificationTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NatsSystemNotificationTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_plugnmeet_nats_msg_proto_enumTypes[2].Descriptor()
}

func (NatsSystemNotificationTypes) Type() protoreflect.EnumType {
	return &file_plugnmeet_nats_msg_proto_enumTypes[2]
}

func (x NatsSystemNotificationTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NatsSystemNotificationTypes.Descriptor instead.
func (NatsSystemNotificationTypes) EnumDescriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{2}
}

type NatsMsgServerToClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event NatsMsgServerToClientEvents `protobuf:"varint,2,opt,name=event,proto3,enum=plugnmeet.NatsMsgServerToClientEvents" json:"event,omitempty"`
	Msg   string                      `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *NatsMsgServerToClient) Reset() {
	*x = NatsMsgServerToClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_nats_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsMsgServerToClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsMsgServerToClient) ProtoMessage() {}

func (x *NatsMsgServerToClient) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_nats_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsMsgServerToClient.ProtoReflect.Descriptor instead.
func (*NatsMsgServerToClient) Descriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{0}
}

func (x *NatsMsgServerToClient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NatsMsgServerToClient) GetEvent() NatsMsgServerToClientEvents {
	if x != nil {
		return x.Event
	}
	return NatsMsgServerToClientEvents_INITIAL_DATA
}

func (x *NatsMsgServerToClient) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type NatsMsgClientToServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event NatsMsgClientToServerEvents `protobuf:"varint,2,opt,name=event,proto3,enum=plugnmeet.NatsMsgClientToServerEvents" json:"event,omitempty"`
	Msg   string                      `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *NatsMsgClientToServer) Reset() {
	*x = NatsMsgClientToServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_nats_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsMsgClientToServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsMsgClientToServer) ProtoMessage() {}

func (x *NatsMsgClientToServer) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_nats_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsMsgClientToServer.ProtoReflect.Descriptor instead.
func (*NatsMsgClientToServer) Descriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{1}
}

func (x *NatsMsgClientToServer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NatsMsgClientToServer) GetEvent() NatsMsgClientToServerEvents {
	if x != nil {
		return x.Event
	}
	return NatsMsgClientToServerEvents_REQ_INITIAL_DATA
}

func (x *NatsMsgClientToServer) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type NatsKvRoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId      string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomSid     string `protobuf:"bytes,2,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	EnabledE2Ee bool   `protobuf:"varint,3,opt,name=enabled_e2ee,json=enabledE2ee,proto3" json:"enabled_e2ee,omitempty"`
	Metadata    string `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt   uint64 `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *NatsKvRoomInfo) Reset() {
	*x = NatsKvRoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_nats_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsKvRoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsKvRoomInfo) ProtoMessage() {}

func (x *NatsKvRoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_nats_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsKvRoomInfo.ProtoReflect.Descriptor instead.
func (*NatsKvRoomInfo) Descriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{2}
}

func (x *NatsKvRoomInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *NatsKvRoomInfo) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *NatsKvRoomInfo) GetEnabledE2Ee() bool {
	if x != nil {
		return x.EnabledE2Ee
	}
	return false
}

func (x *NatsKvRoomInfo) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *NatsKvRoomInfo) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type NatsKvUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserSid        string `protobuf:"bytes,2,opt,name=user_sid,json=userSid,proto3" json:"user_sid,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	RoomId         string `protobuf:"bytes,4,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	IsAdmin        bool   `protobuf:"varint,5,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	IsPresenter    bool   `protobuf:"varint,6,opt,name=is_presenter,json=isPresenter,proto3" json:"is_presenter,omitempty"`
	Metadata       string `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	JoinedAt       uint64 `protobuf:"varint,8,opt,name=joined_at,json=joinedAt,proto3" json:"joined_at,omitempty"`
	ReconnectedAt  uint64 `protobuf:"varint,9,opt,name=reconnected_at,json=reconnectedAt,proto3" json:"reconnected_at,omitempty"`
	DisconnectedAt uint64 `protobuf:"varint,10,opt,name=disconnected_at,json=disconnectedAt,proto3" json:"disconnected_at,omitempty"`
}

func (x *NatsKvUserInfo) Reset() {
	*x = NatsKvUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_nats_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsKvUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsKvUserInfo) ProtoMessage() {}

func (x *NatsKvUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_nats_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsKvUserInfo.ProtoReflect.Descriptor instead.
func (*NatsKvUserInfo) Descriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{3}
}

func (x *NatsKvUserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *NatsKvUserInfo) GetUserSid() string {
	if x != nil {
		return x.UserSid
	}
	return ""
}

func (x *NatsKvUserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NatsKvUserInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *NatsKvUserInfo) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *NatsKvUserInfo) GetIsPresenter() bool {
	if x != nil {
		return x.IsPresenter
	}
	return false
}

func (x *NatsKvUserInfo) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *NatsKvUserInfo) GetJoinedAt() uint64 {
	if x != nil {
		return x.JoinedAt
	}
	return 0
}

func (x *NatsKvUserInfo) GetReconnectedAt() uint64 {
	if x != nil {
		return x.ReconnectedAt
	}
	return 0
}

func (x *NatsKvUserInfo) GetDisconnectedAt() uint64 {
	if x != nil {
		return x.DisconnectedAt
	}
	return 0
}

type MediaServerConnInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Token       string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	EnabledE2Ee bool   `protobuf:"varint,3,opt,name=enabled_e2ee,json=enabledE2ee,proto3" json:"enabled_e2ee,omitempty"`
}

func (x *MediaServerConnInfo) Reset() {
	*x = MediaServerConnInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_nats_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaServerConnInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaServerConnInfo) ProtoMessage() {}

func (x *MediaServerConnInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_nats_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaServerConnInfo.ProtoReflect.Descriptor instead.
func (*MediaServerConnInfo) Descriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{4}
}

func (x *MediaServerConnInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MediaServerConnInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MediaServerConnInfo) GetEnabledE2Ee() bool {
	if x != nil {
		return x.EnabledE2Ee
	}
	return false
}

type NatsInitialData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room            *NatsKvRoomInfo      `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	LocalUser       *NatsKvUserInfo      `protobuf:"bytes,2,opt,name=local_user,json=localUser,proto3" json:"local_user,omitempty"`
	MediaServerInfo *MediaServerConnInfo `protobuf:"bytes,3,opt,name=media_server_info,json=mediaServerInfo,proto3" json:"media_server_info,omitempty"`
}

func (x *NatsInitialData) Reset() {
	*x = NatsInitialData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_nats_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsInitialData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsInitialData) ProtoMessage() {}

func (x *NatsInitialData) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_nats_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsInitialData.ProtoReflect.Descriptor instead.
func (*NatsInitialData) Descriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{5}
}

func (x *NatsInitialData) GetRoom() *NatsKvRoomInfo {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *NatsInitialData) GetLocalUser() *NatsKvUserInfo {
	if x != nil {
		return x.LocalUser
	}
	return nil
}

func (x *NatsInitialData) GetMediaServerInfo() *MediaServerConnInfo {
	if x != nil {
		return x.MediaServerInfo
	}
	return nil
}

type NatsSystemNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type   NatsSystemNotificationTypes `protobuf:"varint,2,opt,name=type,proto3,enum=plugnmeet.NatsSystemNotificationTypes" json:"type,omitempty"`
	Msg    string                      `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	SentAt int64                       `protobuf:"varint,4,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *NatsSystemNotification) Reset() {
	*x = NatsSystemNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_nats_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsSystemNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsSystemNotification) ProtoMessage() {}

func (x *NatsSystemNotification) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_nats_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsSystemNotification.ProtoReflect.Descriptor instead.
func (*NatsSystemNotification) Descriptor() ([]byte, []int) {
	return file_plugnmeet_nats_msg_proto_rawDescGZIP(), []int{6}
}

func (x *NatsSystemNotification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NatsSystemNotification) GetType() NatsSystemNotificationTypes {
	if x != nil {
		return x.Type
	}
	return NatsSystemNotificationTypes_NATS_SYSTEM_NOTIFICATION_INFO
}

func (x *NatsSystemNotification) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *NatsSystemNotification) GetSentAt() int64 {
	if x != nil {
		return x.SentAt
	}
	return 0
}

var File_plugnmeet_nats_msg_proto protoreflect.FileDescriptor

var file_plugnmeet_nats_msg_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x74, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x75, 0x67,
	0x6e, 0x6d, 0x65, 0x65, 0x74, 0x22, 0x77, 0x0a, 0x15, 0x4e, 0x61, 0x74, 0x73, 0x4d, 0x73, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x4e, 0x61, 0x74, 0x73, 0x4d, 0x73,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x77,
	0x0a, 0x15, 0x4e, 0x61, 0x74, 0x73, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65,
	0x65, 0x74, 0x2e, 0x4e, 0x61, 0x74, 0x73, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x4e, 0x61, 0x74, 0x73,
	0x4b, 0x76, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x65, 0x32, 0x65, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x45, 0x32, 0x65,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb8, 0x02, 0x0a,
	0x0e, 0x4e, 0x61, 0x74, 0x73, 0x4b, 0x76, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x53, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f,
	0x69, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6a,
	0x6f, 0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x60, 0x0a, 0x13, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x5f, 0x65, 0x32, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x45, 0x32, 0x65, 0x65, 0x22, 0xc6, 0x01, 0x0a, 0x0f, 0x4e, 0x61,
	0x74, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a,
	0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x4e, 0x61, 0x74, 0x73, 0x4b, 0x76, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x38, 0x0a, 0x0a,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x4e, 0x61, 0x74,
	0x73, 0x4b, 0x76, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x11, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x16, 0x4e, 0x61, 0x74, 0x73, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x4e, 0x61, 0x74, 0x73, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x2a, 0xfa, 0x01, 0x0a, 0x1b, 0x4e, 0x61, 0x74, 0x73, 0x4d, 0x73, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4a, 0x4f, 0x49, 0x4e, 0x45, 0x44,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x53, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x01, 0x12, 0x18, 0x0a,
	0x14, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x52,
	0x45, 0x53, 0x50, 0x5f, 0x52, 0x45, 0x4e, 0x45, 0x57, 0x5f, 0x50, 0x4e, 0x4d, 0x5f, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x11,
	0x0a, 0x0d, 0x52, 0x45, 0x53, 0x50, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4d, 0x53, 0x47, 0x10,
	0x09, 0x2a, 0x68, 0x0a, 0x1b, 0x4e, 0x61, 0x74, 0x73, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x51, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x51, 0x5f, 0x52, 0x45,
	0x4e, 0x45, 0x57, 0x5f, 0x50, 0x4e, 0x4d, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x52, 0x45, 0x51, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4d, 0x53, 0x47, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x2a, 0x8a, 0x01, 0x0a, 0x1b,
	0x4e, 0x61, 0x74, 0x73, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x4e,
	0x41, 0x54, 0x53, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x24,
	0x0a, 0x20, 0x4e, 0x41, 0x54, 0x53, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x4e, 0x4f,
	0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x4e, 0x41, 0x54, 0x53, 0x5f, 0x53, 0x59, 0x53,
	0x54, 0x45, 0x4d, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x42, 0x9e, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x42, 0x15, 0x50, 0x6c, 0x75, 0x67,
	0x6e, 0x6d, 0x65, 0x65, 0x74, 0x4e, 0x61, 0x74, 0x73, 0x4d, 0x73, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e,
	0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09,
	0x50, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0xca, 0x02, 0x09, 0x50, 0x6c, 0x75, 0x67,
	0x6e, 0x6d, 0x65, 0x65, 0x74, 0xe2, 0x02, 0x15, 0x50, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65,
	0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09,
	0x50, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_plugnmeet_nats_msg_proto_rawDescOnce sync.Once
	file_plugnmeet_nats_msg_proto_rawDescData = file_plugnmeet_nats_msg_proto_rawDesc
)

func file_plugnmeet_nats_msg_proto_rawDescGZIP() []byte {
	file_plugnmeet_nats_msg_proto_rawDescOnce.Do(func() {
		file_plugnmeet_nats_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_nats_msg_proto_rawDescData)
	})
	return file_plugnmeet_nats_msg_proto_rawDescData
}

var file_plugnmeet_nats_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_plugnmeet_nats_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_plugnmeet_nats_msg_proto_goTypes = []any{
	(NatsMsgServerToClientEvents)(0), // 0: plugnmeet.NatsMsgServerToClientEvents
	(NatsMsgClientToServerEvents)(0), // 1: plugnmeet.NatsMsgClientToServerEvents
	(NatsSystemNotificationTypes)(0), // 2: plugnmeet.NatsSystemNotificationTypes
	(*NatsMsgServerToClient)(nil),    // 3: plugnmeet.NatsMsgServerToClient
	(*NatsMsgClientToServer)(nil),    // 4: plugnmeet.NatsMsgClientToServer
	(*NatsKvRoomInfo)(nil),           // 5: plugnmeet.NatsKvRoomInfo
	(*NatsKvUserInfo)(nil),           // 6: plugnmeet.NatsKvUserInfo
	(*MediaServerConnInfo)(nil),      // 7: plugnmeet.MediaServerConnInfo
	(*NatsInitialData)(nil),          // 8: plugnmeet.NatsInitialData
	(*NatsSystemNotification)(nil),   // 9: plugnmeet.NatsSystemNotification
}
var file_plugnmeet_nats_msg_proto_depIdxs = []int32{
	0, // 0: plugnmeet.NatsMsgServerToClient.event:type_name -> plugnmeet.NatsMsgServerToClientEvents
	1, // 1: plugnmeet.NatsMsgClientToServer.event:type_name -> plugnmeet.NatsMsgClientToServerEvents
	5, // 2: plugnmeet.NatsInitialData.room:type_name -> plugnmeet.NatsKvRoomInfo
	6, // 3: plugnmeet.NatsInitialData.local_user:type_name -> plugnmeet.NatsKvUserInfo
	7, // 4: plugnmeet.NatsInitialData.media_server_info:type_name -> plugnmeet.MediaServerConnInfo
	2, // 5: plugnmeet.NatsSystemNotification.type:type_name -> plugnmeet.NatsSystemNotificationTypes
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_plugnmeet_nats_msg_proto_init() }
func file_plugnmeet_nats_msg_proto_init() {
	if File_plugnmeet_nats_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_nats_msg_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NatsMsgServerToClient); i {
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
		file_plugnmeet_nats_msg_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NatsMsgClientToServer); i {
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
		file_plugnmeet_nats_msg_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NatsKvRoomInfo); i {
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
		file_plugnmeet_nats_msg_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*NatsKvUserInfo); i {
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
		file_plugnmeet_nats_msg_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MediaServerConnInfo); i {
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
		file_plugnmeet_nats_msg_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*NatsInitialData); i {
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
		file_plugnmeet_nats_msg_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*NatsSystemNotification); i {
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
			RawDescriptor: file_plugnmeet_nats_msg_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_nats_msg_proto_goTypes,
		DependencyIndexes: file_plugnmeet_nats_msg_proto_depIdxs,
		EnumInfos:         file_plugnmeet_nats_msg_proto_enumTypes,
		MessageInfos:      file_plugnmeet_nats_msg_proto_msgTypes,
	}.Build()
	File_plugnmeet_nats_msg_proto = out.File
	file_plugnmeet_nats_msg_proto_rawDesc = nil
	file_plugnmeet_nats_msg_proto_goTypes = nil
	file_plugnmeet_nats_msg_proto_depIdxs = nil
}
