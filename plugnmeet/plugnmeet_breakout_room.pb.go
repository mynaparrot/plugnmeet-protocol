// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: plugnmeet_breakout_room.proto

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

type CreateBreakoutRoomsReq struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RoomId          string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RequestedUserId string                 `protobuf:"bytes,2,opt,name=requested_user_id,json=requestedUserId,proto3" json:"requested_user_id,omitempty"`
	Duration        uint64                 `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	WelcomeMsg      *string                `protobuf:"bytes,4,opt,name=welcome_msg,json=welcomeMsg,proto3,oneof" json:"welcome_msg,omitempty"`
	Rooms           []*BreakoutRoom        `protobuf:"bytes,5,rep,name=rooms,proto3" json:"rooms,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreateBreakoutRoomsReq) Reset() {
	*x = CreateBreakoutRoomsReq{}
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBreakoutRoomsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBreakoutRoomsReq) ProtoMessage() {}

func (x *CreateBreakoutRoomsReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBreakoutRoomsReq.ProtoReflect.Descriptor instead.
func (*CreateBreakoutRoomsReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBreakoutRoomsReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *CreateBreakoutRoomsReq) GetRequestedUserId() string {
	if x != nil {
		return x.RequestedUserId
	}
	return ""
}

func (x *CreateBreakoutRoomsReq) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *CreateBreakoutRoomsReq) GetWelcomeMsg() string {
	if x != nil && x.WelcomeMsg != nil {
		return *x.WelcomeMsg
	}
	return ""
}

func (x *CreateBreakoutRoomsReq) GetRooms() []*BreakoutRoom {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type BreakoutRoom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Duration      uint64                 `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Started       bool                   `protobuf:"varint,4,opt,name=started,proto3" json:"started,omitempty"`
	Created       uint64                 `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Users         []*BreakoutRoomUser    `protobuf:"bytes,6,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BreakoutRoom) Reset() {
	*x = BreakoutRoom{}
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BreakoutRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutRoom) ProtoMessage() {}

func (x *BreakoutRoom) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakoutRoom.ProtoReflect.Descriptor instead.
func (*BreakoutRoom) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{1}
}

func (x *BreakoutRoom) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BreakoutRoom) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BreakoutRoom) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *BreakoutRoom) GetStarted() bool {
	if x != nil {
		return x.Started
	}
	return false
}

func (x *BreakoutRoom) GetCreated() uint64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *BreakoutRoom) GetUsers() []*BreakoutRoomUser {
	if x != nil {
		return x.Users
	}
	return nil
}

type BreakoutRoomUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Joined        bool                   `protobuf:"varint,3,opt,name=joined,proto3" json:"joined,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BreakoutRoomUser) Reset() {
	*x = BreakoutRoomUser{}
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BreakoutRoomUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutRoomUser) ProtoMessage() {}

func (x *BreakoutRoomUser) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakoutRoomUser.ProtoReflect.Descriptor instead.
func (*BreakoutRoomUser) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{2}
}

func (x *BreakoutRoomUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BreakoutRoomUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BreakoutRoomUser) GetJoined() bool {
	if x != nil {
		return x.Joined
	}
	return false
}

type IncreaseBreakoutRoomDurationReq struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	BreakoutRoomId string                 `protobuf:"bytes,1,opt,name=breakout_room_id,json=breakoutRoomId,proto3" json:"breakout_room_id,omitempty"`
	Duration       uint64                 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	RoomId         string                 `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *IncreaseBreakoutRoomDurationReq) Reset() {
	*x = IncreaseBreakoutRoomDurationReq{}
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncreaseBreakoutRoomDurationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseBreakoutRoomDurationReq) ProtoMessage() {}

func (x *IncreaseBreakoutRoomDurationReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseBreakoutRoomDurationReq.ProtoReflect.Descriptor instead.
func (*IncreaseBreakoutRoomDurationReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{3}
}

func (x *IncreaseBreakoutRoomDurationReq) GetBreakoutRoomId() string {
	if x != nil {
		return x.BreakoutRoomId
	}
	return ""
}

func (x *IncreaseBreakoutRoomDurationReq) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *IncreaseBreakoutRoomDurationReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type BroadcastBreakoutRoomMsgReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	RoomId        string                 `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BroadcastBreakoutRoomMsgReq) Reset() {
	*x = BroadcastBreakoutRoomMsgReq{}
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastBreakoutRoomMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastBreakoutRoomMsgReq) ProtoMessage() {}

func (x *BroadcastBreakoutRoomMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastBreakoutRoomMsgReq.ProtoReflect.Descriptor instead.
func (*BroadcastBreakoutRoomMsgReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{4}
}

func (x *BroadcastBreakoutRoomMsgReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BroadcastBreakoutRoomMsgReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type JoinBreakoutRoomReq struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	BreakoutRoomId string                 `protobuf:"bytes,1,opt,name=breakout_room_id,json=breakoutRoomId,proto3" json:"breakout_room_id,omitempty"`
	UserId         string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoomId         string                 `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	IsAdmin        bool                   `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *JoinBreakoutRoomReq) Reset() {
	*x = JoinBreakoutRoomReq{}
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinBreakoutRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinBreakoutRoomReq) ProtoMessage() {}

func (x *JoinBreakoutRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinBreakoutRoomReq.ProtoReflect.Descriptor instead.
func (*JoinBreakoutRoomReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{5}
}

func (x *JoinBreakoutRoomReq) GetBreakoutRoomId() string {
	if x != nil {
		return x.BreakoutRoomId
	}
	return ""
}

func (x *JoinBreakoutRoomReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *JoinBreakoutRoomReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *JoinBreakoutRoomReq) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type EndBreakoutRoomReq struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	BreakoutRoomId string                 `protobuf:"bytes,1,opt,name=breakout_room_id,json=breakoutRoomId,proto3" json:"breakout_room_id,omitempty"`
	RoomId         string                 `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EndBreakoutRoomReq) Reset() {
	*x = EndBreakoutRoomReq{}
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndBreakoutRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndBreakoutRoomReq) ProtoMessage() {}

func (x *EndBreakoutRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndBreakoutRoomReq.ProtoReflect.Descriptor instead.
func (*EndBreakoutRoomReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{6}
}

func (x *EndBreakoutRoomReq) GetBreakoutRoomId() string {
	if x != nil {
		return x.BreakoutRoomId
	}
	return ""
}

func (x *EndBreakoutRoomReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type BreakoutRoomRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token         *string                `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token,omitempty"` // join token
	Room          *BreakoutRoom          `protobuf:"bytes,4,opt,name=room,proto3,oneof" json:"room,omitempty"`   // for my breakout room
	Rooms         []*BreakoutRoom        `protobuf:"bytes,5,rep,name=rooms,proto3" json:"rooms,omitempty"`       // rooms list
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BreakoutRoomRes) Reset() {
	*x = BreakoutRoomRes{}
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BreakoutRoomRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutRoomRes) ProtoMessage() {}

func (x *BreakoutRoomRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakoutRoomRes.ProtoReflect.Descriptor instead.
func (*BreakoutRoomRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{7}
}

func (x *BreakoutRoomRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *BreakoutRoomRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BreakoutRoomRes) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *BreakoutRoomRes) GetRoom() *BreakoutRoom {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *BreakoutRoomRes) GetRooms() []*BreakoutRoom {
	if x != nil {
		return x.Rooms
	}
	return nil
}

var File_plugnmeet_breakout_room_proto protoreflect.FileDescriptor

const file_plugnmeet_breakout_room_proto_rawDesc = "" +
	"\n" +
	"\x1dplugnmeet_breakout_room.proto\x12\tplugnmeet\"\xde\x01\n" +
	"\x16CreateBreakoutRoomsReq\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12*\n" +
	"\x11requested_user_id\x18\x02 \x01(\tR\x0frequestedUserId\x12\x1a\n" +
	"\bduration\x18\x03 \x01(\x04R\bduration\x12$\n" +
	"\vwelcome_msg\x18\x04 \x01(\tH\x00R\n" +
	"welcomeMsg\x88\x01\x01\x12-\n" +
	"\x05rooms\x18\x05 \x03(\v2\x17.plugnmeet.BreakoutRoomR\x05roomsB\x0e\n" +
	"\f_welcome_msg\"\xb7\x01\n" +
	"\fBreakoutRoom\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x1a\n" +
	"\bduration\x18\x03 \x01(\x04R\bduration\x12\x18\n" +
	"\astarted\x18\x04 \x01(\bR\astarted\x12\x18\n" +
	"\acreated\x18\x05 \x01(\x04R\acreated\x121\n" +
	"\x05users\x18\x06 \x03(\v2\x1b.plugnmeet.BreakoutRoomUserR\x05users\"N\n" +
	"\x10BreakoutRoomUser\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x16\n" +
	"\x06joined\x18\x03 \x01(\bR\x06joined\"\x80\x01\n" +
	"\x1fIncreaseBreakoutRoomDurationReq\x12(\n" +
	"\x10breakout_room_id\x18\x01 \x01(\tR\x0ebreakoutRoomId\x12\x1a\n" +
	"\bduration\x18\x02 \x01(\x04R\bduration\x12\x17\n" +
	"\aroom_id\x18\x03 \x01(\tR\x06roomId\"H\n" +
	"\x1bBroadcastBreakoutRoomMsgReq\x12\x10\n" +
	"\x03msg\x18\x01 \x01(\tR\x03msg\x12\x17\n" +
	"\aroom_id\x18\x03 \x01(\tR\x06roomId\"\x8c\x01\n" +
	"\x13JoinBreakoutRoomReq\x12(\n" +
	"\x10breakout_room_id\x18\x01 \x01(\tR\x0ebreakoutRoomId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x17\n" +
	"\aroom_id\x18\x03 \x01(\tR\x06roomId\x12\x19\n" +
	"\bis_admin\x18\x04 \x01(\bR\aisAdmin\"W\n" +
	"\x12EndBreakoutRoomReq\x12(\n" +
	"\x10breakout_room_id\x18\x01 \x01(\tR\x0ebreakoutRoomId\x12\x17\n" +
	"\aroom_id\x18\x03 \x01(\tR\x06roomId\"\xca\x01\n" +
	"\x0fBreakoutRoomRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\x12\x19\n" +
	"\x05token\x18\x03 \x01(\tH\x00R\x05token\x88\x01\x01\x120\n" +
	"\x04room\x18\x04 \x01(\v2\x17.plugnmeet.BreakoutRoomH\x01R\x04room\x88\x01\x01\x12-\n" +
	"\x05rooms\x18\x05 \x03(\v2\x17.plugnmeet.BreakoutRoomR\x05roomsB\b\n" +
	"\x06_tokenB\a\n" +
	"\x05_roomB\xa3\x01\n" +
	"\rcom.plugnmeetB\x1aPlugnmeetBreakoutRoomProtoP\x01Z2github.com/mynaparrot/plugnmeet-protocol/plugnmeet\xa2\x02\x03PXX\xaa\x02\tPlugnmeet\xca\x02\tPlugnmeet\xe2\x02\x15Plugnmeet\\GPBMetadata\xea\x02\tPlugnmeetb\x06proto3"

var (
	file_plugnmeet_breakout_room_proto_rawDescOnce sync.Once
	file_plugnmeet_breakout_room_proto_rawDescData []byte
)

func file_plugnmeet_breakout_room_proto_rawDescGZIP() []byte {
	file_plugnmeet_breakout_room_proto_rawDescOnce.Do(func() {
		file_plugnmeet_breakout_room_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_plugnmeet_breakout_room_proto_rawDesc), len(file_plugnmeet_breakout_room_proto_rawDesc)))
	})
	return file_plugnmeet_breakout_room_proto_rawDescData
}

var file_plugnmeet_breakout_room_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_plugnmeet_breakout_room_proto_goTypes = []any{
	(*CreateBreakoutRoomsReq)(nil),          // 0: plugnmeet.CreateBreakoutRoomsReq
	(*BreakoutRoom)(nil),                    // 1: plugnmeet.BreakoutRoom
	(*BreakoutRoomUser)(nil),                // 2: plugnmeet.BreakoutRoomUser
	(*IncreaseBreakoutRoomDurationReq)(nil), // 3: plugnmeet.IncreaseBreakoutRoomDurationReq
	(*BroadcastBreakoutRoomMsgReq)(nil),     // 4: plugnmeet.BroadcastBreakoutRoomMsgReq
	(*JoinBreakoutRoomReq)(nil),             // 5: plugnmeet.JoinBreakoutRoomReq
	(*EndBreakoutRoomReq)(nil),              // 6: plugnmeet.EndBreakoutRoomReq
	(*BreakoutRoomRes)(nil),                 // 7: plugnmeet.BreakoutRoomRes
}
var file_plugnmeet_breakout_room_proto_depIdxs = []int32{
	1, // 0: plugnmeet.CreateBreakoutRoomsReq.rooms:type_name -> plugnmeet.BreakoutRoom
	2, // 1: plugnmeet.BreakoutRoom.users:type_name -> plugnmeet.BreakoutRoomUser
	1, // 2: plugnmeet.BreakoutRoomRes.room:type_name -> plugnmeet.BreakoutRoom
	1, // 3: plugnmeet.BreakoutRoomRes.rooms:type_name -> plugnmeet.BreakoutRoom
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_plugnmeet_breakout_room_proto_init() }
func file_plugnmeet_breakout_room_proto_init() {
	if File_plugnmeet_breakout_room_proto != nil {
		return
	}
	file_plugnmeet_breakout_room_proto_msgTypes[0].OneofWrappers = []any{}
	file_plugnmeet_breakout_room_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_plugnmeet_breakout_room_proto_rawDesc), len(file_plugnmeet_breakout_room_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_breakout_room_proto_goTypes,
		DependencyIndexes: file_plugnmeet_breakout_room_proto_depIdxs,
		MessageInfos:      file_plugnmeet_breakout_room_proto_msgTypes,
	}.Build()
	File_plugnmeet_breakout_room_proto = out.File
	file_plugnmeet_breakout_room_proto_goTypes = nil
	file_plugnmeet_breakout_room_proto_depIdxs = nil
}
