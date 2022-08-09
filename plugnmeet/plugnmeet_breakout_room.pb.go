// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: plugnmeet_breakout_room.proto

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

type CreateBreakoutRoomsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId          string          `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RequestedUserId string          `protobuf:"bytes,2,opt,name=requested_user_id,json=requestedUserId,proto3" json:"requested_user_id,omitempty"`
	Duration        uint64          `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	WelcomeMsg      *string         `protobuf:"bytes,4,opt,name=welcome_msg,json=welcomeMsg,proto3,oneof" json:"welcome_msg,omitempty"`
	Rooms           []*BreakoutRoom `protobuf:"bytes,5,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *CreateBreakoutRoomsReq) Reset() {
	*x = CreateBreakoutRoomsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBreakoutRoomsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBreakoutRoomsReq) ProtoMessage() {}

func (x *CreateBreakoutRoomsReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Duration uint64              `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Started  bool                `protobuf:"varint,4,opt,name=started,proto3" json:"started,omitempty"`
	Created  uint64              `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Users    []*BreakoutRoomUser `protobuf:"bytes,6,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *BreakoutRoom) Reset() {
	*x = BreakoutRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakoutRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutRoom) ProtoMessage() {}

func (x *BreakoutRoom) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Joined bool   `protobuf:"varint,3,opt,name=joined,proto3" json:"joined,omitempty"`
}

func (x *BreakoutRoomUser) Reset() {
	*x = BreakoutRoomUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakoutRoomUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutRoomUser) ProtoMessage() {}

func (x *BreakoutRoomUser) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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

var File_plugnmeet_breakout_room_proto protoreflect.FileDescriptor

var file_plugnmeet_breakout_room_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x62, 0x72, 0x65, 0x61,
	0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d,
	0x65, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x77,
	0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x05,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0xb7, 0x01, 0x0a, 0x0c,
	0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x4e, 0x0a, 0x10, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6a,
	0x6f, 0x69, 0x6e, 0x65, 0x64, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_plugnmeet_breakout_room_proto_rawDescOnce sync.Once
	file_plugnmeet_breakout_room_proto_rawDescData = file_plugnmeet_breakout_room_proto_rawDesc
)

func file_plugnmeet_breakout_room_proto_rawDescGZIP() []byte {
	file_plugnmeet_breakout_room_proto_rawDescOnce.Do(func() {
		file_plugnmeet_breakout_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_breakout_room_proto_rawDescData)
	})
	return file_plugnmeet_breakout_room_proto_rawDescData
}

var file_plugnmeet_breakout_room_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_plugnmeet_breakout_room_proto_goTypes = []interface{}{
	(*CreateBreakoutRoomsReq)(nil), // 0: plugnmeet.CreateBreakoutRoomsReq
	(*BreakoutRoom)(nil),           // 1: plugnmeet.BreakoutRoom
	(*BreakoutRoomUser)(nil),       // 2: plugnmeet.BreakoutRoomUser
}
var file_plugnmeet_breakout_room_proto_depIdxs = []int32{
	1, // 0: plugnmeet.CreateBreakoutRoomsReq.rooms:type_name -> plugnmeet.BreakoutRoom
	2, // 1: plugnmeet.BreakoutRoom.users:type_name -> plugnmeet.BreakoutRoomUser
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_plugnmeet_breakout_room_proto_init() }
func file_plugnmeet_breakout_room_proto_init() {
	if File_plugnmeet_breakout_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_breakout_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBreakoutRoomsReq); i {
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
		file_plugnmeet_breakout_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakoutRoom); i {
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
		file_plugnmeet_breakout_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakoutRoomUser); i {
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
	file_plugnmeet_breakout_room_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_breakout_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_breakout_room_proto_goTypes,
		DependencyIndexes: file_plugnmeet_breakout_room_proto_depIdxs,
		MessageInfos:      file_plugnmeet_breakout_room_proto_msgTypes,
	}.Build()
	File_plugnmeet_breakout_room_proto = out.File
	file_plugnmeet_breakout_room_proto_rawDesc = nil
	file_plugnmeet_breakout_room_proto_goTypes = nil
	file_plugnmeet_breakout_room_proto_depIdxs = nil
}
