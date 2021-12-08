// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-rc.1
// 	protoc        v3.17.3
// source: accounts.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthId int32

const (
	AuthId_POMELO     AuthId = 0
	AuthId_EOSN_LOGIN AuthId = 1
)

// Enum value maps for AuthId.
var (
	AuthId_name = map[int32]string{
		0: "POMELO",
		1: "EOSN_LOGIN",
	}
	AuthId_value = map[string]int32{
		"POMELO":     0,
		"EOSN_LOGIN": 1,
	}
)

func (x AuthId) Enum() *AuthId {
	p := new(AuthId)
	*p = x
	return p
}

func (x AuthId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthId) Descriptor() protoreflect.EnumDescriptor {
	return file_accounts_proto_enumTypes[0].Descriptor()
}

func (AuthId) Type() protoreflect.EnumType {
	return &file_accounts_proto_enumTypes[0]
}

func (x AuthId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthId.Descriptor instead.
func (AuthId) EnumDescriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{0}
}

type UserTransactionSignatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EosnId        string   `protobuf:"bytes,1,opt,name=eosn_id,json=eosnId,proto3" json:"eosn_id,omitempty"`
	ChainId       []byte   `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	AuthIds       []AuthId `protobuf:"varint,3,rep,packed,name=auth_ids,json=authIds,proto3,enum=eosn.protobuf.v1.AuthId" json:"auth_ids,omitempty"`
	SerializedTrx []byte   `protobuf:"bytes,4,opt,name=serialized_trx,json=serializedTrx,proto3" json:"serialized_trx,omitempty"`
}

func (x *UserTransactionSignatureRequest) Reset() {
	*x = UserTransactionSignatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTransactionSignatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTransactionSignatureRequest) ProtoMessage() {}

func (x *UserTransactionSignatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTransactionSignatureRequest.ProtoReflect.Descriptor instead.
func (*UserTransactionSignatureRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *UserTransactionSignatureRequest) GetEosnId() string {
	if x != nil {
		return x.EosnId
	}
	return ""
}

func (x *UserTransactionSignatureRequest) GetChainId() []byte {
	if x != nil {
		return x.ChainId
	}
	return nil
}

func (x *UserTransactionSignatureRequest) GetAuthIds() []AuthId {
	if x != nil {
		return x.AuthIds
	}
	return nil
}

func (x *UserTransactionSignatureRequest) GetSerializedTrx() []byte {
	if x != nil {
		return x.SerializedTrx
	}
	return nil
}

type RequireUserKycRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EosnId string `protobuf:"bytes,1,opt,name=eosn_id,json=eosnId,proto3" json:"eosn_id,omitempty"`
}

func (x *RequireUserKycRequest) Reset() {
	*x = RequireUserKycRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequireUserKycRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequireUserKycRequest) ProtoMessage() {}

func (x *RequireUserKycRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequireUserKycRequest.ProtoReflect.Descriptor instead.
func (*RequireUserKycRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *RequireUserKycRequest) GetEosnId() string {
	if x != nil {
		return x.EosnId
	}
	return ""
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction                                  `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	ActionData  *CreateAccountRequest_CreateAccountActionData `protobuf:"bytes,2,opt,name=action_data,json=actionData,proto3" json:"action_data,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *CreateAccountRequest) GetActionData() *CreateAccountRequest_CreateAccountActionData {
	if x != nil {
		return x.ActionData
	}
	return nil
}

type LinkAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction                              `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	ActionData  *LinkAccountRequest_LinkAccountActionData `protobuf:"bytes,2,opt,name=action_data,json=actionData,proto3" json:"action_data,omitempty"`
}

func (x *LinkAccountRequest) Reset() {
	*x = LinkAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkAccountRequest) ProtoMessage() {}

func (x *LinkAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkAccountRequest.ProtoReflect.Descriptor instead.
func (*LinkAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *LinkAccountRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *LinkAccountRequest) GetActionData() *LinkAccountRequest_LinkAccountActionData {
	if x != nil {
		return x.ActionData
	}
	return nil
}

type UnlinkAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction                                  `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	ActionData  *UnlinkAccountRequest_UnlinkAccountActionData `protobuf:"bytes,2,opt,name=action_data,json=actionData,proto3" json:"action_data,omitempty"`
}

func (x *UnlinkAccountRequest) Reset() {
	*x = UnlinkAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlinkAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlinkAccountRequest) ProtoMessage() {}

func (x *UnlinkAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlinkAccountRequest.ProtoReflect.Descriptor instead.
func (*UnlinkAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{4}
}

func (x *UnlinkAccountRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *UnlinkAccountRequest) GetActionData() *UnlinkAccountRequest_UnlinkAccountActionData {
	if x != nil {
		return x.ActionData
	}
	return nil
}

type SignedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedTrx []byte `protobuf:"bytes,2,opt,name=signed_trx,json=signedTrx,proto3" json:"signed_trx,omitempty"`
}

func (x *SignedTransaction) Reset() {
	*x = SignedTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedTransaction) ProtoMessage() {}

func (x *SignedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedTransaction.ProtoReflect.Descriptor instead.
func (*SignedTransaction) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{5}
}

func (x *SignedTransaction) GetSignedTrx() []byte {
	if x != nil {
		return x.SignedTrx
	}
	return nil
}

type CreateAccountRequest_CreateAccountActionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EosnId     string   `protobuf:"bytes,1,opt,name=eosn_id,json=eosnId,proto3" json:"eosn_id,omitempty"`
	PublicKeys []string `protobuf:"bytes,2,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
}

func (x *CreateAccountRequest_CreateAccountActionData) Reset() {
	*x = CreateAccountRequest_CreateAccountActionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest_CreateAccountActionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest_CreateAccountActionData) ProtoMessage() {}

func (x *CreateAccountRequest_CreateAccountActionData) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest_CreateAccountActionData.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest_CreateAccountActionData) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CreateAccountRequest_CreateAccountActionData) GetEosnId() string {
	if x != nil {
		return x.EosnId
	}
	return ""
}

func (x *CreateAccountRequest_CreateAccountActionData) GetPublicKeys() []string {
	if x != nil {
		return x.PublicKeys
	}
	return nil
}

type LinkAccountRequest_LinkAccountActionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EosnId      string `protobuf:"bytes,1,opt,name=eosn_id,json=eosnId,proto3" json:"eosn_id,omitempty"`
	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	Signature   string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *LinkAccountRequest_LinkAccountActionData) Reset() {
	*x = LinkAccountRequest_LinkAccountActionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkAccountRequest_LinkAccountActionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkAccountRequest_LinkAccountActionData) ProtoMessage() {}

func (x *LinkAccountRequest_LinkAccountActionData) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkAccountRequest_LinkAccountActionData.ProtoReflect.Descriptor instead.
func (*LinkAccountRequest_LinkAccountActionData) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{3, 0}
}

func (x *LinkAccountRequest_LinkAccountActionData) GetEosnId() string {
	if x != nil {
		return x.EosnId
	}
	return ""
}

func (x *LinkAccountRequest_LinkAccountActionData) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *LinkAccountRequest_LinkAccountActionData) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type UnlinkAccountRequest_UnlinkAccountActionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EosnId      string `protobuf:"bytes,1,opt,name=eosn_id,json=eosnId,proto3" json:"eosn_id,omitempty"`
	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
}

func (x *UnlinkAccountRequest_UnlinkAccountActionData) Reset() {
	*x = UnlinkAccountRequest_UnlinkAccountActionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlinkAccountRequest_UnlinkAccountActionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlinkAccountRequest_UnlinkAccountActionData) ProtoMessage() {}

func (x *UnlinkAccountRequest_UnlinkAccountActionData) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlinkAccountRequest_UnlinkAccountActionData.ProtoReflect.Descriptor instead.
func (*UnlinkAccountRequest_UnlinkAccountActionData) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{4, 0}
}

func (x *UnlinkAccountRequest_UnlinkAccountActionData) GetEosnId() string {
	if x != nil {
		return x.EosnId
	}
	return ""
}

func (x *UnlinkAccountRequest_UnlinkAccountActionData) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

var File_accounts_proto protoreflect.FileDescriptor

var file_accounts_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x31, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01,
	0x0a, 0x1f, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6f, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x6f, 0x73, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49,
	0x64, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x72,
	0x78, 0x22, 0x30, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6f,
	0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6f, 0x73,
	0x6e, 0x49, 0x64, 0x22, 0x8d, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x53,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6f, 0x73,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6f, 0x73, 0x6e,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x73, 0x22, 0xa5, 0x02, 0x0a, 0x12, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x71, 0x0a, 0x15, 0x4c, 0x69, 0x6e, 0x6b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6f, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x6f, 0x73, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x8f, 0x02, 0x0a, 0x14,
	0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x73, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6f, 0x73,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e,
	0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x55, 0x0a, 0x17, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6f, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x6f, 0x73, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a,
	0x11, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72,
	0x78, 0x2a, 0x24, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x0a, 0x0a, 0x06, 0x50,
	0x4f, 0x4d, 0x45, 0x4c, 0x4f, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4f, 0x53, 0x4e, 0x5f,
	0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x32, 0xd4, 0x03, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x65, 0x6f, 0x73,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x4b, 0x79, 0x63, 0x12, 0x27, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x65, 0x6f, 0x73,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b,
	0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x65, 0x6f,
	0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x55,
	0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x65,
	0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounts_proto_rawDescOnce sync.Once
	file_accounts_proto_rawDescData = file_accounts_proto_rawDesc
)

func file_accounts_proto_rawDescGZIP() []byte {
	file_accounts_proto_rawDescOnce.Do(func() {
		file_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounts_proto_rawDescData)
	})
	return file_accounts_proto_rawDescData
}

var file_accounts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_accounts_proto_goTypes = []interface{}{
	(AuthId)(0),                                          // 0: eosn.protobuf.v1.AuthId
	(*UserTransactionSignatureRequest)(nil),              // 1: eosn.protobuf.v1.UserTransactionSignatureRequest
	(*RequireUserKycRequest)(nil),                        // 2: eosn.protobuf.v1.RequireUserKycRequest
	(*CreateAccountRequest)(nil),                         // 3: eosn.protobuf.v1.CreateAccountRequest
	(*LinkAccountRequest)(nil),                           // 4: eosn.protobuf.v1.LinkAccountRequest
	(*UnlinkAccountRequest)(nil),                         // 5: eosn.protobuf.v1.UnlinkAccountRequest
	(*SignedTransaction)(nil),                            // 6: eosn.protobuf.v1.SignedTransaction
	(*CreateAccountRequest_CreateAccountActionData)(nil), // 7: eosn.protobuf.v1.CreateAccountRequest.CreateAccountActionData
	(*LinkAccountRequest_LinkAccountActionData)(nil),     // 8: eosn.protobuf.v1.LinkAccountRequest.LinkAccountActionData
	(*UnlinkAccountRequest_UnlinkAccountActionData)(nil), // 9: eosn.protobuf.v1.UnlinkAccountRequest.UnlinkAccountActionData
	(*Transaction)(nil),                                  // 10: eosn.protobuf.v1.Transaction
	(*emptypb.Empty)(nil),                                // 11: google.protobuf.Empty
}
var file_accounts_proto_depIdxs = []int32{
	0,  // 0: eosn.protobuf.v1.UserTransactionSignatureRequest.auth_ids:type_name -> eosn.protobuf.v1.AuthId
	10, // 1: eosn.protobuf.v1.CreateAccountRequest.transaction:type_name -> eosn.protobuf.v1.Transaction
	7,  // 2: eosn.protobuf.v1.CreateAccountRequest.action_data:type_name -> eosn.protobuf.v1.CreateAccountRequest.CreateAccountActionData
	10, // 3: eosn.protobuf.v1.LinkAccountRequest.transaction:type_name -> eosn.protobuf.v1.Transaction
	8,  // 4: eosn.protobuf.v1.LinkAccountRequest.action_data:type_name -> eosn.protobuf.v1.LinkAccountRequest.LinkAccountActionData
	10, // 5: eosn.protobuf.v1.UnlinkAccountRequest.transaction:type_name -> eosn.protobuf.v1.Transaction
	9,  // 6: eosn.protobuf.v1.UnlinkAccountRequest.action_data:type_name -> eosn.protobuf.v1.UnlinkAccountRequest.UnlinkAccountActionData
	1,  // 7: eosn.protobuf.v1.AccountsInternalService.SignUserTransaction:input_type -> eosn.protobuf.v1.UserTransactionSignatureRequest
	2,  // 8: eosn.protobuf.v1.AccountsInternalService.RequireUserKyc:input_type -> eosn.protobuf.v1.RequireUserKycRequest
	3,  // 9: eosn.protobuf.v1.AccountsInternalService.CreateAccount:input_type -> eosn.protobuf.v1.CreateAccountRequest
	4,  // 10: eosn.protobuf.v1.AccountsInternalService.LinkAccount:input_type -> eosn.protobuf.v1.LinkAccountRequest
	5,  // 11: eosn.protobuf.v1.AccountsInternalService.UnlinkAccount:input_type -> eosn.protobuf.v1.UnlinkAccountRequest
	6,  // 12: eosn.protobuf.v1.AccountsInternalService.SignUserTransaction:output_type -> eosn.protobuf.v1.SignedTransaction
	11, // 13: eosn.protobuf.v1.AccountsInternalService.RequireUserKyc:output_type -> google.protobuf.Empty
	11, // 14: eosn.protobuf.v1.AccountsInternalService.CreateAccount:output_type -> google.protobuf.Empty
	11, // 15: eosn.protobuf.v1.AccountsInternalService.LinkAccount:output_type -> google.protobuf.Empty
	11, // 16: eosn.protobuf.v1.AccountsInternalService.UnlinkAccount:output_type -> google.protobuf.Empty
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_accounts_proto_init() }
func file_accounts_proto_init() {
	if File_accounts_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTransactionSignatureRequest); i {
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
		file_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequireUserKycRequest); i {
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
		file_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkAccountRequest); i {
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
		file_accounts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlinkAccountRequest); i {
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
		file_accounts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedTransaction); i {
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
		file_accounts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest_CreateAccountActionData); i {
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
		file_accounts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkAccountRequest_LinkAccountActionData); i {
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
		file_accounts_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlinkAccountRequest_UnlinkAccountActionData); i {
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
			RawDescriptor: file_accounts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounts_proto_goTypes,
		DependencyIndexes: file_accounts_proto_depIdxs,
		EnumInfos:         file_accounts_proto_enumTypes,
		MessageInfos:      file_accounts_proto_msgTypes,
	}.Build()
	File_accounts_proto = out.File
	file_accounts_proto_rawDesc = nil
	file_accounts_proto_goTypes = nil
	file_accounts_proto_depIdxs = nil
}
