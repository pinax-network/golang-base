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

type SignedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedTrx []byte `protobuf:"bytes,2,opt,name=signed_trx,json=signedTrx,proto3" json:"signed_trx,omitempty"`
}

func (x *SignedTransaction) Reset() {
	*x = SignedTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedTransaction) ProtoMessage() {}

func (x *SignedTransaction) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SignedTransaction.ProtoReflect.Descriptor instead.
func (*SignedTransaction) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *SignedTransaction) GetSignedTrx() []byte {
	if x != nil {
		return x.SignedTrx
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Undo  bool   `protobuf:"varint,1,opt,name=undo,proto3" json:"undo,omitempty"`
	TrxId []byte `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetUndo() bool {
	if x != nil {
		return x.Undo
	}
	return false
}

func (x *Transaction) GetTrxId() []byte {
	if x != nil {
		return x.TrxId
	}
	return nil
}

type UserTransactionSignatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EosnId        string   `protobuf:"bytes,1,opt,name=eosn_id,json=eosnId,proto3" json:"eosn_id,omitempty"`
	ChainId       []byte   `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	AuthIds       []AuthId `protobuf:"varint,3,rep,packed,name=auth_ids,json=authIds,proto3,enum=accounts.AuthId" json:"auth_ids,omitempty"`
	SerializedTrx []byte   `protobuf:"bytes,4,opt,name=serialized_trx,json=serializedTrx,proto3" json:"serialized_trx,omitempty"`
}

func (x *UserTransactionSignatureRequest) Reset() {
	*x = UserTransactionSignatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTransactionSignatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTransactionSignatureRequest) ProtoMessage() {}

func (x *UserTransactionSignatureRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserTransactionSignatureRequest.ProtoReflect.Descriptor instead.
func (*UserTransactionSignatureRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{2}
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
		mi := &file_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{3}
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
		mi := &file_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkAccountRequest) ProtoMessage() {}

func (x *LinkAccountRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LinkAccountRequest.ProtoReflect.Descriptor instead.
func (*LinkAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{4}
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
		mi := &file_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlinkAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlinkAccountRequest) ProtoMessage() {}

func (x *UnlinkAccountRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UnlinkAccountRequest.ProtoReflect.Descriptor instead.
func (*UnlinkAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{5}
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
	return file_accounts_proto_rawDescGZIP(), []int{3, 0}
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
	return file_accounts_proto_rawDescGZIP(), []int{4, 0}
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
	return file_accounts_proto_rawDescGZIP(), []int{5, 0}
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
	0x12, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x78, 0x22, 0x38, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e,
	0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x64, 0x6f, 0x12, 0x15,
	0x0a, 0x06, 0x74, 0x72, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x74, 0x72, 0x78, 0x49, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x1f, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6f, 0x73,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6f, 0x73, 0x6e,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49,
	0x64, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x72,
	0x78, 0x22, 0xfd, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x53, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6f, 0x73, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6f, 0x73, 0x6e, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x73, 0x22, 0x95, 0x02, 0x0a, 0x12, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x53, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x71, 0x0a, 0x15, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x17, 0x0a, 0x07, 0x65, 0x6f, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x6f, 0x73, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x14, 0x55, 0x6e,
	0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x6e, 0x6c, 0x69,
	0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x55, 0x0a, 0x17, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x17, 0x0a, 0x07, 0x65, 0x6f, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x6f, 0x73, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x24, 0x0a, 0x06, 0x41,
	0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x4f, 0x4d, 0x45, 0x4c, 0x4f, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4f, 0x53, 0x4e, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10,
	0x01, 0x32, 0xd7, 0x02, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a,
	0x13, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b, 0x4c, 0x69, 0x6e,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x0d, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x6e, 0x6c,
	0x69, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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
	(AuthId)(0),                                          // 0: accounts.AuthId
	(*SignedTransaction)(nil),                            // 1: accounts.SignedTransaction
	(*Transaction)(nil),                                  // 2: accounts.Transaction
	(*UserTransactionSignatureRequest)(nil),              // 3: accounts.UserTransactionSignatureRequest
	(*CreateAccountRequest)(nil),                         // 4: accounts.CreateAccountRequest
	(*LinkAccountRequest)(nil),                           // 5: accounts.LinkAccountRequest
	(*UnlinkAccountRequest)(nil),                         // 6: accounts.UnlinkAccountRequest
	(*CreateAccountRequest_CreateAccountActionData)(nil), // 7: accounts.CreateAccountRequest.CreateAccountActionData
	(*LinkAccountRequest_LinkAccountActionData)(nil),     // 8: accounts.LinkAccountRequest.LinkAccountActionData
	(*UnlinkAccountRequest_UnlinkAccountActionData)(nil), // 9: accounts.UnlinkAccountRequest.UnlinkAccountActionData
	(*emptypb.Empty)(nil),                                // 10: google.protobuf.Empty
}
var file_accounts_proto_depIdxs = []int32{
	0,  // 0: accounts.UserTransactionSignatureRequest.auth_ids:type_name -> accounts.AuthId
	2,  // 1: accounts.CreateAccountRequest.transaction:type_name -> accounts.Transaction
	7,  // 2: accounts.CreateAccountRequest.action_data:type_name -> accounts.CreateAccountRequest.CreateAccountActionData
	2,  // 3: accounts.LinkAccountRequest.transaction:type_name -> accounts.Transaction
	8,  // 4: accounts.LinkAccountRequest.action_data:type_name -> accounts.LinkAccountRequest.LinkAccountActionData
	2,  // 5: accounts.UnlinkAccountRequest.transaction:type_name -> accounts.Transaction
	9,  // 6: accounts.UnlinkAccountRequest.action_data:type_name -> accounts.UnlinkAccountRequest.UnlinkAccountActionData
	3,  // 7: accounts.AccountsInternalService.SignUserTransaction:input_type -> accounts.UserTransactionSignatureRequest
	4,  // 8: accounts.AccountsInternalService.CreateAccount:input_type -> accounts.CreateAccountRequest
	5,  // 9: accounts.AccountsInternalService.LinkAccount:input_type -> accounts.LinkAccountRequest
	6,  // 10: accounts.AccountsInternalService.UnlinkAccount:input_type -> accounts.UnlinkAccountRequest
	1,  // 11: accounts.AccountsInternalService.SignUserTransaction:output_type -> accounts.SignedTransaction
	10, // 12: accounts.AccountsInternalService.CreateAccount:output_type -> google.protobuf.Empty
	10, // 13: accounts.AccountsInternalService.LinkAccount:output_type -> google.protobuf.Empty
	10, // 14: accounts.AccountsInternalService.UnlinkAccount:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_accounts_proto_init() }
func file_accounts_proto_init() {
	if File_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_accounts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_accounts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
