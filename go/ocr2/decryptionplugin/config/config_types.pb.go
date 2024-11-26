// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: config/config_types.proto

package config

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

type OracleIDtoKeyShareIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OracleId      uint32 `protobuf:"varint,1,opt,name=oracle_id,json=oracleId,proto3" json:"oracle_id,omitempty"`
	KeyShareIndex uint32 `protobuf:"varint,2,opt,name=key_share_index,json=keyShareIndex,proto3" json:"key_share_index,omitempty"`
}

func (x *OracleIDtoKeyShareIndex) Reset() {
	*x = OracleIDtoKeyShareIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleIDtoKeyShareIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleIDtoKeyShareIndex) ProtoMessage() {}

func (x *OracleIDtoKeyShareIndex) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleIDtoKeyShareIndex.ProtoReflect.Descriptor instead.
func (*OracleIDtoKeyShareIndex) Descriptor() ([]byte, []int) {
	return file_config_config_types_proto_rawDescGZIP(), []int{0}
}

func (x *OracleIDtoKeyShareIndex) GetOracleId() uint32 {
	if x != nil {
		return x.OracleId
	}
	return 0
}

func (x *OracleIDtoKeyShareIndex) GetKeyShareIndex() uint32 {
	if x != nil {
		return x.KeyShareIndex
	}
	return 0
}

type ReportingPluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxQueryLengthBytes       uint32                     `protobuf:"varint,1,opt,name=max_query_length_bytes,json=maxQueryLengthBytes,proto3" json:"max_query_length_bytes,omitempty"`
	MaxObservationLengthBytes uint32                     `protobuf:"varint,2,opt,name=max_observation_length_bytes,json=maxObservationLengthBytes,proto3" json:"max_observation_length_bytes,omitempty"`
	MaxReportLengthBytes      uint32                     `protobuf:"varint,3,opt,name=max_report_length_bytes,json=maxReportLengthBytes,proto3" json:"max_report_length_bytes,omitempty"`
	RequestCountLimit         uint32                     `protobuf:"varint,4,opt,name=request_count_limit,json=requestCountLimit,proto3" json:"request_count_limit,omitempty"`
	RequestTotalBytesLimit    uint32                     `protobuf:"varint,5,opt,name=request_total_bytes_limit,json=requestTotalBytesLimit,proto3" json:"request_total_bytes_limit,omitempty"`
	RequireLocalRequestCheck  bool                       `protobuf:"varint,6,opt,name=require_local_request_check,json=requireLocalRequestCheck,proto3" json:"require_local_request_check,omitempty"`
	PublicKey                 []byte                     `protobuf:"bytes,7,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PrivKeyShare              []byte                     `protobuf:"bytes,8,opt,name=priv_key_share,json=privKeyShare,proto3" json:"priv_key_share,omitempty"`
	OracleIdToKeyIndex        []*OracleIDtoKeyShareIndex `protobuf:"bytes,9,rep,name=oracle_id_to_key_index,json=oracleIdToKeyIndex,proto3" json:"oracle_id_to_key_index,omitempty"`
}

func (x *ReportingPluginConfig) Reset() {
	*x = ReportingPluginConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_config_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportingPluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportingPluginConfig) ProtoMessage() {}

func (x *ReportingPluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_config_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportingPluginConfig.ProtoReflect.Descriptor instead.
func (*ReportingPluginConfig) Descriptor() ([]byte, []int) {
	return file_config_config_types_proto_rawDescGZIP(), []int{1}
}

func (x *ReportingPluginConfig) GetMaxQueryLengthBytes() uint32 {
	if x != nil {
		return x.MaxQueryLengthBytes
	}
	return 0
}

func (x *ReportingPluginConfig) GetMaxObservationLengthBytes() uint32 {
	if x != nil {
		return x.MaxObservationLengthBytes
	}
	return 0
}

func (x *ReportingPluginConfig) GetMaxReportLengthBytes() uint32 {
	if x != nil {
		return x.MaxReportLengthBytes
	}
	return 0
}

func (x *ReportingPluginConfig) GetRequestCountLimit() uint32 {
	if x != nil {
		return x.RequestCountLimit
	}
	return 0
}

func (x *ReportingPluginConfig) GetRequestTotalBytesLimit() uint32 {
	if x != nil {
		return x.RequestTotalBytesLimit
	}
	return 0
}

func (x *ReportingPluginConfig) GetRequireLocalRequestCheck() bool {
	if x != nil {
		return x.RequireLocalRequestCheck
	}
	return false
}

func (x *ReportingPluginConfig) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *ReportingPluginConfig) GetPrivKeyShare() []byte {
	if x != nil {
		return x.PrivKeyShare
	}
	return nil
}

func (x *ReportingPluginConfig) GetOracleIdToKeyIndex() []*OracleIDtoKeyShareIndex {
	if x != nil {
		return x.OracleIdToKeyIndex
	}
	return nil
}

var File_config_config_types_proto protoreflect.FileDescriptor

var file_config_config_types_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x17, 0x4f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x49, 0x44, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x8e, 0x04, 0x0a, 0x15, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x33, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x1c, 0x6d, 0x61, 0x78, 0x5f,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19,
	0x6d, 0x61, 0x78, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x61, 0x78,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x39, 0x0a, 0x19, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x69,
	0x76, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x4b, 0x65, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12,
	0x59, 0x0a, 0x16, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x12, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x64,
	0x54, 0x6f, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f,
	0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_config_types_proto_rawDescOnce sync.Once
	file_config_config_types_proto_rawDescData = file_config_config_types_proto_rawDesc
)

func file_config_config_types_proto_rawDescGZIP() []byte {
	file_config_config_types_proto_rawDescOnce.Do(func() {
		file_config_config_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_config_types_proto_rawDescData)
	})
	return file_config_config_types_proto_rawDescData
}

var file_config_config_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_config_types_proto_goTypes = []interface{}{
	(*OracleIDtoKeyShareIndex)(nil), // 0: config_types.OracleIDtoKeyShareIndex
	(*ReportingPluginConfig)(nil),   // 1: config_types.ReportingPluginConfig
}
var file_config_config_types_proto_depIdxs = []int32{
	0, // 0: config_types.ReportingPluginConfig.oracle_id_to_key_index:type_name -> config_types.OracleIDtoKeyShareIndex
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_config_types_proto_init() }
func file_config_config_types_proto_init() {
	if File_config_config_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_config_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleIDtoKeyShareIndex); i {
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
		file_config_config_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportingPluginConfig); i {
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
			RawDescriptor: file_config_config_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_config_types_proto_goTypes,
		DependencyIndexes: file_config_config_types_proto_depIdxs,
		MessageInfos:      file_config_config_types_proto_msgTypes,
	}.Build()
	File_config_config_types_proto = out.File
	file_config_config_types_proto_rawDesc = nil
	file_config_config_types_proto_goTypes = nil
	file_config_config_types_proto_depIdxs = nil
}