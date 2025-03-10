// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: pinax.proto

package pb

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

type AuthenticateByKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *AuthenticateByKeyRequest) Reset() {
	*x = AuthenticateByKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinax_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateByKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateByKeyRequest) ProtoMessage() {}

func (x *AuthenticateByKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pinax_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateByKeyRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateByKeyRequest) Descriptor() ([]byte, []int) {
	return file_pinax_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateByKeyRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type AuthenticateByKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ApiKeyId string `protobuf:"bytes,2,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	Meta     string `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *AuthenticateByKeyResponse) Reset() {
	*x = AuthenticateByKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinax_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateByKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateByKeyResponse) ProtoMessage() {}

func (x *AuthenticateByKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pinax_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateByKeyResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateByKeyResponse) Descriptor() ([]byte, []int) {
	return file_pinax_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateByKeyResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthenticateByKeyResponse) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

func (x *AuthenticateByKeyResponse) GetMeta() string {
	if x != nil {
		return x.Meta
	}
	return ""
}

var File_pinax_proto protoreflect.FileDescriptor

var file_pinax_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x69, 0x6e, 0x61, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65,
	0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x22,
	0x33, 0x0a, 0x18, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61,
	0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x22, 0x66, 0x0a, 0x19, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x70,
	0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x32, 0x86, 0x01, 0x0a,
	0x14, 0x50, 0x69, 0x6e, 0x61, 0x78, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x2e, 0x65, 0x6f, 0x73,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pinax_proto_rawDescOnce sync.Once
	file_pinax_proto_rawDescData = file_pinax_proto_rawDesc
)

func file_pinax_proto_rawDescGZIP() []byte {
	file_pinax_proto_rawDescOnce.Do(func() {
		file_pinax_proto_rawDescData = protoimpl.X.CompressGZIP(file_pinax_proto_rawDescData)
	})
	return file_pinax_proto_rawDescData
}

var file_pinax_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pinax_proto_goTypes = []interface{}{
	(*AuthenticateByKeyRequest)(nil),  // 0: eosn.protobuf.v1.AuthenticateByKeyRequest
	(*AuthenticateByKeyResponse)(nil), // 1: eosn.protobuf.v1.AuthenticateByKeyResponse
}
var file_pinax_proto_depIdxs = []int32{
	0, // 0: eosn.protobuf.v1.PinaxInternalService.AuthenticateByKey:input_type -> eosn.protobuf.v1.AuthenticateByKeyRequest
	1, // 1: eosn.protobuf.v1.PinaxInternalService.AuthenticateByKey:output_type -> eosn.protobuf.v1.AuthenticateByKeyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pinax_proto_init() }
func file_pinax_proto_init() {
	if File_pinax_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pinax_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateByKeyRequest); i {
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
		file_pinax_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateByKeyResponse); i {
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
			RawDescriptor: file_pinax_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pinax_proto_goTypes,
		DependencyIndexes: file_pinax_proto_depIdxs,
		MessageInfos:      file_pinax_proto_msgTypes,
	}.Build()
	File_pinax_proto = out.File
	file_pinax_proto_rawDesc = nil
	file_pinax_proto_goTypes = nil
	file_pinax_proto_depIdxs = nil
}
