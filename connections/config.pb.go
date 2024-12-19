// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.0
// source: config.proto

// This is just enough protobuf to be able to talk with the Server's
// API Query() endpoint.

package connections

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

type ApiClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaCertificate       string `protobuf:"bytes,1,opt,name=ca_certificate,json=caCertificate,proto3" json:"ca_certificate,omitempty"`
	ClientCert          string `protobuf:"bytes,2,opt,name=client_cert,json=clientCert,proto3" json:"client_cert,omitempty"`
	ClientPrivateKey    string `protobuf:"bytes,3,opt,name=client_private_key,json=clientPrivateKey,proto3" json:"client_private_key,omitempty"`
	ApiConnectionString string `protobuf:"bytes,4,opt,name=api_connection_string,json=apiConnectionString,proto3" json:"api_connection_string,omitempty"`
	Name                string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	PinnedServerName    string `protobuf:"bytes,6,opt,name=pinned_server_name,json=pinnedServerName,proto3" json:"pinned_server_name,omitempty"`
}

func (x *ApiClientConfig) Reset() {
	*x = ApiClientConfig{}
	mi := &file_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiClientConfig) ProtoMessage() {}

func (x *ApiClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiClientConfig.ProtoReflect.Descriptor instead.
func (*ApiClientConfig) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *ApiClientConfig) GetCaCertificate() string {
	if x != nil {
		return x.CaCertificate
	}
	return ""
}

func (x *ApiClientConfig) GetClientCert() string {
	if x != nil {
		return x.ClientCert
	}
	return ""
}

func (x *ApiClientConfig) GetClientPrivateKey() string {
	if x != nil {
		return x.ClientPrivateKey
	}
	return ""
}

func (x *ApiClientConfig) GetApiConnectionString() string {
	if x != nil {
		return x.ApiConnectionString
	}
	return ""
}

func (x *ApiClientConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiClientConfig) GetPinnedServerName() string {
	if x != nil {
		return x.PinnedServerName
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x04, 0x0a, 0x0f, 0x41, 0x70, 0x69, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x61, 0x0a, 0x0e, 0x63, 0x61, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x3a, 0xe2, 0xfc, 0xe3, 0xc4, 0x01, 0x34, 0x12, 0x32, 0x54, 0x68, 0x65, 0x20, 0x43,
	0x41, 0x20, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x41, 0x50, 0x49,
	0x20, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x0d, 0x63,
	0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x97, 0x01, 0x0a,
	0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x76, 0xe2, 0xfc, 0xe3, 0xc4, 0x01, 0x70, 0x12, 0x6e, 0x41, 0x20, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x20, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x73, 0x20, 0x74, 0x6f,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x20, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x27, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x20, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x20, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x27, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x5d, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x2f, 0xe2, 0xfc, 0xe3, 0xc4, 0x01, 0x29, 0x12, 0x27, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x20, 0x6b, 0x65, 0x79, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x6f, 0x20, 0x77, 0x69,
	0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x2e, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x6e, 0x0a, 0x15, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xe2, 0xfc, 0xe3, 0xc4, 0x01, 0x34, 0x12, 0x32, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x77, 0x68, 0x65, 0x72, 0x65, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x61, 0x70, 0x69, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x77, 0x69,
	0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x52, 0x13, 0x61, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2b, 0xe2, 0xfc, 0xe3, 0xc4, 0x01, 0x25, 0x12, 0x23, 0x54, 0x68, 0x65,
	0x20, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_proto_goTypes = []any{
	(*ApiClientConfig)(nil), // 0: proto.ApiClientConfig
}
var file_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	file_semantic_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}