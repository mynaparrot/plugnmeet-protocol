// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: plugnmeet_recorder.proto

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

type RecordingTasks int32

const (
	RecordingTasks_START_RECORDING     RecordingTasks = 0
	RecordingTasks_STOP_RECORDING      RecordingTasks = 1
	RecordingTasks_START_RTMP          RecordingTasks = 2
	RecordingTasks_STOP_RTMP           RecordingTasks = 3
	RecordingTasks_END_RECORDING       RecordingTasks = 4
	RecordingTasks_END_RTMP            RecordingTasks = 5
	RecordingTasks_RECORDING_PROCEEDED RecordingTasks = 6
	RecordingTasks_STOP                RecordingTasks = 7
)

// Enum value maps for RecordingTasks.
var (
	RecordingTasks_name = map[int32]string{
		0: "START_RECORDING",
		1: "STOP_RECORDING",
		2: "START_RTMP",
		3: "STOP_RTMP",
		4: "END_RECORDING",
		5: "END_RTMP",
		6: "RECORDING_PROCEEDED",
		7: "STOP",
	}
	RecordingTasks_value = map[string]int32{
		"START_RECORDING":     0,
		"STOP_RECORDING":      1,
		"START_RTMP":          2,
		"STOP_RTMP":           3,
		"END_RECORDING":       4,
		"END_RTMP":            5,
		"RECORDING_PROCEEDED": 6,
		"STOP":                7,
	}
)

func (x RecordingTasks) Enum() *RecordingTasks {
	p := new(RecordingTasks)
	*p = x
	return p
}

func (x RecordingTasks) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordingTasks) Descriptor() protoreflect.EnumDescriptor {
	return file_plugnmeet_recorder_proto_enumTypes[0].Descriptor()
}

func (RecordingTasks) Type() protoreflect.EnumType {
	return &file_plugnmeet_recorder_proto_enumTypes[0]
}

func (x RecordingTasks) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordingTasks.Descriptor instead.
func (RecordingTasks) EnumDescriptor() ([]byte, []int) {
	return file_plugnmeet_recorder_proto_rawDescGZIP(), []int{0}
}

type RecorderServiceType int32

const (
	RecorderServiceType_RECORDING RecorderServiceType = 0
	RecorderServiceType_RTMP      RecorderServiceType = 1
)

// Enum value maps for RecorderServiceType.
var (
	RecorderServiceType_name = map[int32]string{
		0: "RECORDING",
		1: "RTMP",
	}
	RecorderServiceType_value = map[string]int32{
		"RECORDING": 0,
		"RTMP":      1,
	}
)

func (x RecorderServiceType) Enum() *RecorderServiceType {
	p := new(RecorderServiceType)
	*p = x
	return p
}

func (x RecorderServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecorderServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_plugnmeet_recorder_proto_enumTypes[1].Descriptor()
}

func (RecorderServiceType) Type() protoreflect.EnumType {
	return &file_plugnmeet_recorder_proto_enumTypes[1]
}

func (x RecorderServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecorderServiceType.Descriptor instead.
func (RecorderServiceType) EnumDescriptor() ([]byte, []int) {
	return file_plugnmeet_recorder_proto_rawDescGZIP(), []int{1}
}

type PlugNmeetToRecorder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From        string         `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Task        RecordingTasks `protobuf:"varint,2,opt,name=task,proto3,enum=plugnmeet.RecordingTasks" json:"task,omitempty"`
	RoomId      string         `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomSid     string         `protobuf:"bytes,4,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	RecordingId string         `protobuf:"bytes,5,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	RecorderId  string         `protobuf:"bytes,6,opt,name=recorder_id,json=recorderId,proto3" json:"recorder_id,omitempty"`
	AccessToken string         `protobuf:"bytes,7,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RtmpUrl     *string        `protobuf:"bytes,8,opt,name=rtmp_url,json=rtmpUrl,proto3,oneof" json:"rtmp_url,omitempty"`
}

func (x *PlugNmeetToRecorder) Reset() {
	*x = PlugNmeetToRecorder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_recorder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlugNmeetToRecorder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlugNmeetToRecorder) ProtoMessage() {}

func (x *PlugNmeetToRecorder) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_recorder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlugNmeetToRecorder.ProtoReflect.Descriptor instead.
func (*PlugNmeetToRecorder) Descriptor() ([]byte, []int) {
	return file_plugnmeet_recorder_proto_rawDescGZIP(), []int{0}
}

func (x *PlugNmeetToRecorder) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PlugNmeetToRecorder) GetTask() RecordingTasks {
	if x != nil {
		return x.Task
	}
	return RecordingTasks_START_RECORDING
}

func (x *PlugNmeetToRecorder) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *PlugNmeetToRecorder) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *PlugNmeetToRecorder) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *PlugNmeetToRecorder) GetRecorderId() string {
	if x != nil {
		return x.RecorderId
	}
	return ""
}

func (x *PlugNmeetToRecorder) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *PlugNmeetToRecorder) GetRtmpUrl() string {
	if x != nil && x.RtmpUrl != nil {
		return *x.RtmpUrl
	}
	return ""
}

type RecorderToPlugNmeet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From        string         `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Task        RecordingTasks `protobuf:"varint,2,opt,name=task,proto3,enum=plugnmeet.RecordingTasks" json:"task,omitempty"`
	Status      bool           `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Msg         string         `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	RecordingId string         `protobuf:"bytes,5,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	RoomId      string         `protobuf:"bytes,6,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomSid     string         `protobuf:"bytes,7,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	RecorderId  string         `protobuf:"bytes,8,opt,name=recorder_id,json=recorderId,proto3" json:"recorder_id,omitempty"`
	FilePath    string         `protobuf:"bytes,9,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileSize    float32        `protobuf:"fixed32,10,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
}

func (x *RecorderToPlugNmeet) Reset() {
	*x = RecorderToPlugNmeet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_recorder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecorderToPlugNmeet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecorderToPlugNmeet) ProtoMessage() {}

func (x *RecorderToPlugNmeet) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_recorder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecorderToPlugNmeet.ProtoReflect.Descriptor instead.
func (*RecorderToPlugNmeet) Descriptor() ([]byte, []int) {
	return file_plugnmeet_recorder_proto_rawDescGZIP(), []int{1}
}

func (x *RecorderToPlugNmeet) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *RecorderToPlugNmeet) GetTask() RecordingTasks {
	if x != nil {
		return x.Task
	}
	return RecordingTasks_START_RECORDING
}

func (x *RecorderToPlugNmeet) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *RecorderToPlugNmeet) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RecorderToPlugNmeet) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *RecorderToPlugNmeet) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RecorderToPlugNmeet) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *RecorderToPlugNmeet) GetRecorderId() string {
	if x != nil {
		return x.RecorderId
	}
	return ""
}

func (x *RecorderToPlugNmeet) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *RecorderToPlugNmeet) GetFileSize() float32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

type FromParentToChild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task        RecordingTasks `protobuf:"varint,1,opt,name=task,proto3,enum=plugnmeet.RecordingTasks" json:"task,omitempty"`
	RecordingId string         `protobuf:"bytes,2,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	RoomId      string         `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomSid     string         `protobuf:"bytes,4,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
}

func (x *FromParentToChild) Reset() {
	*x = FromParentToChild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_recorder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromParentToChild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromParentToChild) ProtoMessage() {}

func (x *FromParentToChild) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_recorder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromParentToChild.ProtoReflect.Descriptor instead.
func (*FromParentToChild) Descriptor() ([]byte, []int) {
	return file_plugnmeet_recorder_proto_rawDescGZIP(), []int{2}
}

func (x *FromParentToChild) GetTask() RecordingTasks {
	if x != nil {
		return x.Task
	}
	return RecordingTasks_START_RECORDING
}

func (x *FromParentToChild) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *FromParentToChild) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *FromParentToChild) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

type FromChildToParent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task        RecordingTasks `protobuf:"varint,1,opt,name=task,proto3,enum=plugnmeet.RecordingTasks" json:"task,omitempty"`
	Status      bool           `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Msg         string         `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	RecordingId string         `protobuf:"bytes,4,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	RoomId      string         `protobuf:"bytes,5,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomSid     string         `protobuf:"bytes,6,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
}

func (x *FromChildToParent) Reset() {
	*x = FromChildToParent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_recorder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromChildToParent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromChildToParent) ProtoMessage() {}

func (x *FromChildToParent) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_recorder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromChildToParent.ProtoReflect.Descriptor instead.
func (*FromChildToParent) Descriptor() ([]byte, []int) {
	return file_plugnmeet_recorder_proto_rawDescGZIP(), []int{3}
}

func (x *FromChildToParent) GetTask() RecordingTasks {
	if x != nil {
		return x.Task
	}
	return RecordingTasks_START_RECORDING
}

func (x *FromChildToParent) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *FromChildToParent) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FromChildToParent) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *FromChildToParent) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *FromChildToParent) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

type StartRecorderChildArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId           string              `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RecordingId      string              `protobuf:"bytes,2,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	RoomSid          string              `protobuf:"bytes,3,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	AccessToken      string              `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	PlugNMeetInfo    *PlugNmeetInfo      `protobuf:"bytes,5,opt,name=plug_n_meet_info,json=plugNMeetInfo,proto3" json:"plug_n_meet_info,omitempty"`
	PostMp4Convert   bool                `protobuf:"varint,6,opt,name=post_mp4_convert,json=postMp4Convert,proto3" json:"post_mp4_convert,omitempty"`
	CopyToPath       *CopyToPath         `protobuf:"bytes,7,opt,name=copy_to_path,json=copyToPath,proto3" json:"copy_to_path,omitempty"`
	ServiceType      RecorderServiceType `protobuf:"varint,8,opt,name=serviceType,proto3,enum=plugnmeet.RecorderServiceType" json:"serviceType,omitempty"`
	RecorderId       *string             `protobuf:"bytes,9,opt,name=recorder_id,json=recorderId,proto3,oneof" json:"recorder_id,omitempty"`
	RtmpUrl          *string             `protobuf:"bytes,10,opt,name=rtmp_url,json=rtmpUrl,proto3,oneof" json:"rtmp_url,omitempty"`
	WebsocketUrl     string              `protobuf:"bytes,11,opt,name=websocket_url,json=websocketUrl,proto3" json:"websocket_url,omitempty"`
	CustomChromePath *string             `protobuf:"bytes,12,opt,name=custom_chrome_path,json=customChromePath,proto3,oneof" json:"custom_chrome_path,omitempty"`
}

func (x *StartRecorderChildArgs) Reset() {
	*x = StartRecorderChildArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_recorder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRecorderChildArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRecorderChildArgs) ProtoMessage() {}

func (x *StartRecorderChildArgs) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_recorder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRecorderChildArgs.ProtoReflect.Descriptor instead.
func (*StartRecorderChildArgs) Descriptor() ([]byte, []int) {
	return file_plugnmeet_recorder_proto_rawDescGZIP(), []int{4}
}

func (x *StartRecorderChildArgs) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *StartRecorderChildArgs) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *StartRecorderChildArgs) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *StartRecorderChildArgs) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *StartRecorderChildArgs) GetPlugNMeetInfo() *PlugNmeetInfo {
	if x != nil {
		return x.PlugNMeetInfo
	}
	return nil
}

func (x *StartRecorderChildArgs) GetPostMp4Convert() bool {
	if x != nil {
		return x.PostMp4Convert
	}
	return false
}

func (x *StartRecorderChildArgs) GetCopyToPath() *CopyToPath {
	if x != nil {
		return x.CopyToPath
	}
	return nil
}

func (x *StartRecorderChildArgs) GetServiceType() RecorderServiceType {
	if x != nil {
		return x.ServiceType
	}
	return RecorderServiceType_RECORDING
}

func (x *StartRecorderChildArgs) GetRecorderId() string {
	if x != nil && x.RecorderId != nil {
		return *x.RecorderId
	}
	return ""
}

func (x *StartRecorderChildArgs) GetRtmpUrl() string {
	if x != nil && x.RtmpUrl != nil {
		return *x.RtmpUrl
	}
	return ""
}

func (x *StartRecorderChildArgs) GetWebsocketUrl() string {
	if x != nil {
		return x.WebsocketUrl
	}
	return ""
}

func (x *StartRecorderChildArgs) GetCustomChromePath() string {
	if x != nil && x.CustomChromePath != nil {
		return *x.CustomChromePath
	}
	return ""
}

type PlugNmeetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host      string  `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	ApiKey    string  `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret string  `protobuf:"bytes,3,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	JoinHost  *string `protobuf:"bytes,4,opt,name=join_host,json=joinHost,proto3,oneof" json:"join_host,omitempty"`
}

func (x *PlugNmeetInfo) Reset() {
	*x = PlugNmeetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_recorder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlugNmeetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlugNmeetInfo) ProtoMessage() {}

func (x *PlugNmeetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_recorder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlugNmeetInfo.ProtoReflect.Descriptor instead.
func (*PlugNmeetInfo) Descriptor() ([]byte, []int) {
	return file_plugnmeet_recorder_proto_rawDescGZIP(), []int{5}
}

func (x *PlugNmeetInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PlugNmeetInfo) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *PlugNmeetInfo) GetApiSecret() string {
	if x != nil {
		return x.ApiSecret
	}
	return ""
}

func (x *PlugNmeetInfo) GetJoinHost() string {
	if x != nil && x.JoinHost != nil {
		return *x.JoinHost
	}
	return ""
}

type CopyToPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainPath string  `protobuf:"bytes,1,opt,name=main_path,json=mainPath,proto3" json:"main_path,omitempty"`
	SubPath  *string `protobuf:"bytes,2,opt,name=sub_path,json=subPath,proto3,oneof" json:"sub_path,omitempty"`
}

func (x *CopyToPath) Reset() {
	*x = CopyToPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_recorder_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyToPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyToPath) ProtoMessage() {}

func (x *CopyToPath) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_recorder_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyToPath.ProtoReflect.Descriptor instead.
func (*CopyToPath) Descriptor() ([]byte, []int) {
	return file_plugnmeet_recorder_proto_rawDescGZIP(), []int{6}
}

func (x *CopyToPath) GetMainPath() string {
	if x != nil {
		return x.MainPath
	}
	return ""
}

func (x *CopyToPath) GetSubPath() string {
	if x != nil && x.SubPath != nil {
		return *x.SubPath
	}
	return ""
}

var File_plugnmeet_recorder_proto protoreflect.FileDescriptor

var file_plugnmeet_recorder_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x75, 0x67,
	0x6e, 0x6d, 0x65, 0x65, 0x74, 0x22, 0xa0, 0x02, 0x0a, 0x13, 0x50, 0x6c, 0x75, 0x67, 0x4e, 0x6d,
	0x65, 0x65, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x53, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x08, 0x72,
	0x74, 0x6d, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x72, 0x74, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x72, 0x74, 0x6d, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0xb4, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x6c, 0x75, 0x67, 0x4e, 0x6d, 0x65, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x53, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x99, 0x01, 0x0a, 0x11, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x6f,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x69, 0x64, 0x22, 0xc3, 0x01, 0x0a, 0x11,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x54, 0x6f, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x69,
	0x64, 0x22, 0xcc, 0x04, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x41, 0x72, 0x67, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d,
	0x53, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x41, 0x0a, 0x10, 0x70, 0x6c, 0x75, 0x67, 0x5f, 0x6e,
	0x5f, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x4e, 0x6d, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x70, 0x6c, 0x75, 0x67,
	0x4e, 0x4d, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x6d, 0x70, 0x34, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x4d, 0x70, 0x34, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x70, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x54, 0x6f, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x0a, 0x63, 0x6f, 0x70, 0x79, 0x54, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x12, 0x40, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x72, 0x74, 0x6d, 0x70, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x72, 0x74, 0x6d, 0x70, 0x55, 0x72,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x65, 0x62,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x31, 0x0a, 0x12, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43,
	0x68, 0x72, 0x6f, 0x6d, 0x65, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x72, 0x74, 0x6d, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x8b, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x4e, 0x6d, 0x65, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x20,
	0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x56,
	0x0a, 0x0a, 0x43, 0x6f, 0x70, 0x79, 0x54, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x75,
	0x62, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x2a, 0x9c, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x52, 0x54, 0x4d, 0x50,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x52, 0x54, 0x4d, 0x50, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x54, 0x4d, 0x50,
	0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x5f,
	0x50, 0x52, 0x4f, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x53,
	0x54, 0x4f, 0x50, 0x10, 0x07, 0x2a, 0x2e, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52,
	0x54, 0x4d, 0x50, 0x10, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_plugnmeet_recorder_proto_rawDescOnce sync.Once
	file_plugnmeet_recorder_proto_rawDescData = file_plugnmeet_recorder_proto_rawDesc
)

func file_plugnmeet_recorder_proto_rawDescGZIP() []byte {
	file_plugnmeet_recorder_proto_rawDescOnce.Do(func() {
		file_plugnmeet_recorder_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_recorder_proto_rawDescData)
	})
	return file_plugnmeet_recorder_proto_rawDescData
}

var file_plugnmeet_recorder_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_plugnmeet_recorder_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_plugnmeet_recorder_proto_goTypes = []interface{}{
	(RecordingTasks)(0),            // 0: plugnmeet.RecordingTasks
	(RecorderServiceType)(0),       // 1: plugnmeet.RecorderServiceType
	(*PlugNmeetToRecorder)(nil),    // 2: plugnmeet.PlugNmeetToRecorder
	(*RecorderToPlugNmeet)(nil),    // 3: plugnmeet.RecorderToPlugNmeet
	(*FromParentToChild)(nil),      // 4: plugnmeet.FromParentToChild
	(*FromChildToParent)(nil),      // 5: plugnmeet.FromChildToParent
	(*StartRecorderChildArgs)(nil), // 6: plugnmeet.StartRecorderChildArgs
	(*PlugNmeetInfo)(nil),          // 7: plugnmeet.PlugNmeetInfo
	(*CopyToPath)(nil),             // 8: plugnmeet.CopyToPath
}
var file_plugnmeet_recorder_proto_depIdxs = []int32{
	0, // 0: plugnmeet.PlugNmeetToRecorder.task:type_name -> plugnmeet.RecordingTasks
	0, // 1: plugnmeet.RecorderToPlugNmeet.task:type_name -> plugnmeet.RecordingTasks
	0, // 2: plugnmeet.FromParentToChild.task:type_name -> plugnmeet.RecordingTasks
	0, // 3: plugnmeet.FromChildToParent.task:type_name -> plugnmeet.RecordingTasks
	7, // 4: plugnmeet.StartRecorderChildArgs.plug_n_meet_info:type_name -> plugnmeet.PlugNmeetInfo
	8, // 5: plugnmeet.StartRecorderChildArgs.copy_to_path:type_name -> plugnmeet.CopyToPath
	1, // 6: plugnmeet.StartRecorderChildArgs.serviceType:type_name -> plugnmeet.RecorderServiceType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_plugnmeet_recorder_proto_init() }
func file_plugnmeet_recorder_proto_init() {
	if File_plugnmeet_recorder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_recorder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlugNmeetToRecorder); i {
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
		file_plugnmeet_recorder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecorderToPlugNmeet); i {
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
		file_plugnmeet_recorder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromParentToChild); i {
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
		file_plugnmeet_recorder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromChildToParent); i {
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
		file_plugnmeet_recorder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRecorderChildArgs); i {
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
		file_plugnmeet_recorder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlugNmeetInfo); i {
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
		file_plugnmeet_recorder_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyToPath); i {
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
	file_plugnmeet_recorder_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_plugnmeet_recorder_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_plugnmeet_recorder_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_plugnmeet_recorder_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_recorder_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_recorder_proto_goTypes,
		DependencyIndexes: file_plugnmeet_recorder_proto_depIdxs,
		EnumInfos:         file_plugnmeet_recorder_proto_enumTypes,
		MessageInfos:      file_plugnmeet_recorder_proto_msgTypes,
	}.Build()
	File_plugnmeet_recorder_proto = out.File
	file_plugnmeet_recorder_proto_rawDesc = nil
	file_plugnmeet_recorder_proto_goTypes = nil
	file_plugnmeet_recorder_proto_depIdxs = nil
}
