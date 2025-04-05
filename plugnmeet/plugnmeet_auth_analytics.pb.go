// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: plugnmeet_auth_analytics.proto

package plugnmeet

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type FetchAnalyticsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomIds       []string               `protobuf:"bytes,1,rep,name=room_ids,json=roomIds,proto3" json:"room_ids,omitempty"`
	From          uint32                 `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Limit         uint32                 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy       string                 `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchAnalyticsReq) Reset() {
	*x = FetchAnalyticsReq{}
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchAnalyticsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAnalyticsReq) ProtoMessage() {}

func (x *FetchAnalyticsReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAnalyticsReq.ProtoReflect.Descriptor instead.
func (*FetchAnalyticsReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *FetchAnalyticsReq) GetRoomIds() []string {
	if x != nil {
		return x.RoomIds
	}
	return nil
}

func (x *FetchAnalyticsReq) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *FetchAnalyticsReq) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchAnalyticsReq) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

type AnalyticsInfo struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RoomId           string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	FileId           string                 `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	FileName         string                 `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize         float64                `protobuf:"fixed64,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	CreationTime     int64                  `protobuf:"varint,5,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	RoomCreationTime int64                  `protobuf:"varint,6,opt,name=room_creation_time,json=roomCreationTime,proto3" json:"room_creation_time,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *AnalyticsInfo) Reset() {
	*x = AnalyticsInfo{}
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalyticsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsInfo) ProtoMessage() {}

func (x *AnalyticsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsInfo.ProtoReflect.Descriptor instead.
func (*AnalyticsInfo) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *AnalyticsInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *AnalyticsInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *AnalyticsInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *AnalyticsInfo) GetFileSize() float64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *AnalyticsInfo) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

func (x *AnalyticsInfo) GetRoomCreationTime() int64 {
	if x != nil {
		return x.RoomCreationTime
	}
	return 0
}

type FetchAnalyticsResult struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	TotalAnalytics int64                  `protobuf:"varint,1,opt,name=total_analytics,json=totalAnalytics,proto3" json:"total_analytics,omitempty"`
	From           uint32                 `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Limit          uint32                 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy        string                 `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	AnalyticsList  []*AnalyticsInfo       `protobuf:"bytes,5,rep,name=analytics_list,json=analyticsList,proto3" json:"analytics_list,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FetchAnalyticsResult) Reset() {
	*x = FetchAnalyticsResult{}
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchAnalyticsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAnalyticsResult) ProtoMessage() {}

func (x *FetchAnalyticsResult) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAnalyticsResult.ProtoReflect.Descriptor instead.
func (*FetchAnalyticsResult) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *FetchAnalyticsResult) GetTotalAnalytics() int64 {
	if x != nil {
		return x.TotalAnalytics
	}
	return 0
}

func (x *FetchAnalyticsResult) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *FetchAnalyticsResult) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchAnalyticsResult) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *FetchAnalyticsResult) GetAnalyticsList() []*AnalyticsInfo {
	if x != nil {
		return x.AnalyticsList
	}
	return nil
}

type FetchAnalyticsRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Result        *FetchAnalyticsResult  `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchAnalyticsRes) Reset() {
	*x = FetchAnalyticsRes{}
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchAnalyticsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAnalyticsRes) ProtoMessage() {}

func (x *FetchAnalyticsRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAnalyticsRes.ProtoReflect.Descriptor instead.
func (*FetchAnalyticsRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_analytics_proto_rawDescGZIP(), []int{3}
}

func (x *FetchAnalyticsRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *FetchAnalyticsRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FetchAnalyticsRes) GetResult() *FetchAnalyticsResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteAnalyticsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAnalyticsReq) Reset() {
	*x = DeleteAnalyticsReq{}
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAnalyticsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAnalyticsReq) ProtoMessage() {}

func (x *DeleteAnalyticsReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAnalyticsReq.ProtoReflect.Descriptor instead.
func (*DeleteAnalyticsReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_analytics_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAnalyticsReq) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DeleteAnalyticsRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAnalyticsRes) Reset() {
	*x = DeleteAnalyticsRes{}
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAnalyticsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAnalyticsRes) ProtoMessage() {}

func (x *DeleteAnalyticsRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAnalyticsRes.ProtoReflect.Descriptor instead.
func (*DeleteAnalyticsRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_analytics_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAnalyticsRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *DeleteAnalyticsRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetAnalyticsDownloadTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAnalyticsDownloadTokenReq) Reset() {
	*x = GetAnalyticsDownloadTokenReq{}
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAnalyticsDownloadTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnalyticsDownloadTokenReq) ProtoMessage() {}

func (x *GetAnalyticsDownloadTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnalyticsDownloadTokenReq.ProtoReflect.Descriptor instead.
func (*GetAnalyticsDownloadTokenReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_analytics_proto_rawDescGZIP(), []int{6}
}

func (x *GetAnalyticsDownloadTokenReq) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type GetAnalyticsDownloadTokenRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token         *string                `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAnalyticsDownloadTokenRes) Reset() {
	*x = GetAnalyticsDownloadTokenRes{}
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAnalyticsDownloadTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnalyticsDownloadTokenRes) ProtoMessage() {}

func (x *GetAnalyticsDownloadTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_analytics_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnalyticsDownloadTokenRes.ProtoReflect.Descriptor instead.
func (*GetAnalyticsDownloadTokenRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_analytics_proto_rawDescGZIP(), []int{7}
}

func (x *GetAnalyticsDownloadTokenRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetAnalyticsDownloadTokenRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetAnalyticsDownloadTokenRes) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

var File_plugnmeet_auth_analytics_proto protoreflect.FileDescriptor

const file_plugnmeet_auth_analytics_proto_rawDesc = "" +
	"\n" +
	"\x1eplugnmeet_auth_analytics.proto\x12\tplugnmeet\x1a\x1bbuf/validate/validate.proto\"s\n" +
	"\x11FetchAnalyticsReq\x12\x19\n" +
	"\broom_ids\x18\x01 \x03(\tR\aroomIds\x12\x12\n" +
	"\x04from\x18\x02 \x01(\rR\x04from\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\rR\x05limit\x12\x19\n" +
	"\border_by\x18\x04 \x01(\tR\aorderBy\"\xce\x01\n" +
	"\rAnalyticsInfo\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12\x17\n" +
	"\afile_id\x18\x02 \x01(\tR\x06fileId\x12\x1b\n" +
	"\tfile_name\x18\x03 \x01(\tR\bfileName\x12\x1b\n" +
	"\tfile_size\x18\x04 \x01(\x01R\bfileSize\x12#\n" +
	"\rcreation_time\x18\x05 \x01(\x03R\fcreationTime\x12,\n" +
	"\x12room_creation_time\x18\x06 \x01(\x03R\x10roomCreationTime\"\xc5\x01\n" +
	"\x14FetchAnalyticsResult\x12'\n" +
	"\x0ftotal_analytics\x18\x01 \x01(\x03R\x0etotalAnalytics\x12\x12\n" +
	"\x04from\x18\x02 \x01(\rR\x04from\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\rR\x05limit\x12\x19\n" +
	"\border_by\x18\x04 \x01(\tR\aorderBy\x12?\n" +
	"\x0eanalytics_list\x18\x05 \x03(\v2\x18.plugnmeet.AnalyticsInfoR\ranalyticsList\"v\n" +
	"\x11FetchAnalyticsRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\x127\n" +
	"\x06result\x18\x03 \x01(\v2\x1f.plugnmeet.FetchAnalyticsResultR\x06result\"5\n" +
	"\x12DeleteAnalyticsReq\x12\x1f\n" +
	"\afile_id\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\x06fileId\">\n" +
	"\x12DeleteAnalyticsRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\"?\n" +
	"\x1cGetAnalyticsDownloadTokenReq\x12\x1f\n" +
	"\afile_id\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\x06fileId\"m\n" +
	"\x1cGetAnalyticsDownloadTokenRes\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x10\n" +
	"\x03msg\x18\x02 \x01(\tR\x03msg\x12\x19\n" +
	"\x05token\x18\x03 \x01(\tH\x00R\x05token\x88\x01\x01B\b\n" +
	"\x06_tokenB\xa4\x01\n" +
	"\rcom.plugnmeetB\x1bPlugnmeetAuthAnalyticsProtoP\x01Z2github.com/mynaparrot/plugnmeet-protocol/plugnmeet\xa2\x02\x03PXX\xaa\x02\tPlugnmeet\xca\x02\tPlugnmeet\xe2\x02\x15Plugnmeet\\GPBMetadata\xea\x02\tPlugnmeetb\x06proto3"

var (
	file_plugnmeet_auth_analytics_proto_rawDescOnce sync.Once
	file_plugnmeet_auth_analytics_proto_rawDescData []byte
)

func file_plugnmeet_auth_analytics_proto_rawDescGZIP() []byte {
	file_plugnmeet_auth_analytics_proto_rawDescOnce.Do(func() {
		file_plugnmeet_auth_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_plugnmeet_auth_analytics_proto_rawDesc), len(file_plugnmeet_auth_analytics_proto_rawDesc)))
	})
	return file_plugnmeet_auth_analytics_proto_rawDescData
}

var file_plugnmeet_auth_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_plugnmeet_auth_analytics_proto_goTypes = []any{
	(*FetchAnalyticsReq)(nil),            // 0: plugnmeet.FetchAnalyticsReq
	(*AnalyticsInfo)(nil),                // 1: plugnmeet.AnalyticsInfo
	(*FetchAnalyticsResult)(nil),         // 2: plugnmeet.FetchAnalyticsResult
	(*FetchAnalyticsRes)(nil),            // 3: plugnmeet.FetchAnalyticsRes
	(*DeleteAnalyticsReq)(nil),           // 4: plugnmeet.DeleteAnalyticsReq
	(*DeleteAnalyticsRes)(nil),           // 5: plugnmeet.DeleteAnalyticsRes
	(*GetAnalyticsDownloadTokenReq)(nil), // 6: plugnmeet.GetAnalyticsDownloadTokenReq
	(*GetAnalyticsDownloadTokenRes)(nil), // 7: plugnmeet.GetAnalyticsDownloadTokenRes
}
var file_plugnmeet_auth_analytics_proto_depIdxs = []int32{
	1, // 0: plugnmeet.FetchAnalyticsResult.analytics_list:type_name -> plugnmeet.AnalyticsInfo
	2, // 1: plugnmeet.FetchAnalyticsRes.result:type_name -> plugnmeet.FetchAnalyticsResult
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_plugnmeet_auth_analytics_proto_init() }
func file_plugnmeet_auth_analytics_proto_init() {
	if File_plugnmeet_auth_analytics_proto != nil {
		return
	}
	file_plugnmeet_auth_analytics_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_plugnmeet_auth_analytics_proto_rawDesc), len(file_plugnmeet_auth_analytics_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_auth_analytics_proto_goTypes,
		DependencyIndexes: file_plugnmeet_auth_analytics_proto_depIdxs,
		MessageInfos:      file_plugnmeet_auth_analytics_proto_msgTypes,
	}.Build()
	File_plugnmeet_auth_analytics_proto = out.File
	file_plugnmeet_auth_analytics_proto_goTypes = nil
	file_plugnmeet_auth_analytics_proto_depIdxs = nil
}
