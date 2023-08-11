// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: bounties.proto

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

type BountyState int32

const (
	BountyState_RELEASED BountyState = 0
	BountyState_DONE     BountyState = 1
	BountyState_REFUNDED BountyState = 2
)

// Enum value maps for BountyState.
var (
	BountyState_name = map[int32]string{
		0: "RELEASED",
		1: "DONE",
		2: "REFUNDED",
	}
	BountyState_value = map[string]int32{
		"RELEASED": 0,
		"DONE":     1,
		"REFUNDED": 2,
	}
)

func (x BountyState) Enum() *BountyState {
	p := new(BountyState)
	*p = x
	return p
}

func (x BountyState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BountyState) Descriptor() protoreflect.EnumDescriptor {
	return file_bounties_proto_enumTypes[0].Descriptor()
}

func (BountyState) Type() protoreflect.EnumType {
	return &file_bounties_proto_enumTypes[0]
}

func (x BountyState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BountyState.Descriptor instead.
func (BountyState) EnumDescriptor() ([]byte, []int) {
	return file_bounties_proto_rawDescGZIP(), []int{0}
}

type BountyTransferType int32

const (
	BountyTransferType_DEPOSIT    BountyTransferType = 0
	BountyTransferType_WITHDRAWAL BountyTransferType = 1
	BountyTransferType_CLAIM      BountyTransferType = 2
)

// Enum value maps for BountyTransferType.
var (
	BountyTransferType_name = map[int32]string{
		0: "DEPOSIT",
		1: "WITHDRAWAL",
		2: "CLAIM",
	}
	BountyTransferType_value = map[string]int32{
		"DEPOSIT":    0,
		"WITHDRAWAL": 1,
		"CLAIM":      2,
	}
)

func (x BountyTransferType) Enum() *BountyTransferType {
	p := new(BountyTransferType)
	*p = x
	return p
}

func (x BountyTransferType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BountyTransferType) Descriptor() protoreflect.EnumDescriptor {
	return file_bounties_proto_enumTypes[1].Descriptor()
}

func (BountyTransferType) Type() protoreflect.EnumType {
	return &file_bounties_proto_enumTypes[1]
}

func (x BountyTransferType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BountyTransferType.Descriptor instead.
func (BountyTransferType) EnumDescriptor() ([]byte, []int) {
	return file_bounties_proto_rawDescGZIP(), []int{1}
}

type BountyTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction        `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Transfer    *BountyTransferData `protobuf:"bytes,2,opt,name=transfer,proto3" json:"transfer,omitempty"`
}

func (x *BountyTransfer) Reset() {
	*x = BountyTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bounties_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BountyTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BountyTransfer) ProtoMessage() {}

func (x *BountyTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_bounties_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BountyTransfer.ProtoReflect.Descriptor instead.
func (*BountyTransfer) Descriptor() ([]byte, []int) {
	return file_bounties_proto_rawDescGZIP(), []int{0}
}

func (x *BountyTransfer) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *BountyTransfer) GetTransfer() *BountyTransferData {
	if x != nil {
		return x.Transfer
	}
	return nil
}

type BountyTransferData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        BountyTransferType `protobuf:"varint,1,opt,name=type,proto3,enum=eosn.protobuf.v1.BountyTransferType" json:"type,omitempty"`
	BountyId    string             `protobuf:"bytes,2,opt,name=bounty_id,json=bountyId,proto3" json:"bounty_id,omitempty"`
	UserId      string             `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccountName string             `protobuf:"bytes,4,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	ExtQuantity *TokenData         `protobuf:"bytes,5,opt,name=ext_quantity,json=extQuantity,proto3" json:"ext_quantity,omitempty"`
	Fee         string             `protobuf:"bytes,6,opt,name=fee,proto3" json:"fee,omitempty"`
	Value       float64            `protobuf:"fixed64,7,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BountyTransferData) Reset() {
	*x = BountyTransferData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bounties_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BountyTransferData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BountyTransferData) ProtoMessage() {}

func (x *BountyTransferData) ProtoReflect() protoreflect.Message {
	mi := &file_bounties_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BountyTransferData.ProtoReflect.Descriptor instead.
func (*BountyTransferData) Descriptor() ([]byte, []int) {
	return file_bounties_proto_rawDescGZIP(), []int{1}
}

func (x *BountyTransferData) GetType() BountyTransferType {
	if x != nil {
		return x.Type
	}
	return BountyTransferType_DEPOSIT
}

func (x *BountyTransferData) GetBountyId() string {
	if x != nil {
		return x.BountyId
	}
	return ""
}

func (x *BountyTransferData) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BountyTransferData) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *BountyTransferData) GetExtQuantity() *TokenData {
	if x != nil {
		return x.ExtQuantity
	}
	return nil
}

func (x *BountyTransferData) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *BountyTransferData) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type UpdateBountyStateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *IrreversibleTransaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	BountyId    string                   `protobuf:"bytes,2,opt,name=bounty_id,json=bountyId,proto3" json:"bounty_id,omitempty"`
	State       BountyState              `protobuf:"varint,3,opt,name=state,proto3,enum=eosn.protobuf.v1.BountyState" json:"state,omitempty"`
}

func (x *UpdateBountyStateData) Reset() {
	*x = UpdateBountyStateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bounties_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBountyStateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBountyStateData) ProtoMessage() {}

func (x *UpdateBountyStateData) ProtoReflect() protoreflect.Message {
	mi := &file_bounties_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBountyStateData.ProtoReflect.Descriptor instead.
func (*UpdateBountyStateData) Descriptor() ([]byte, []int) {
	return file_bounties_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBountyStateData) GetTransaction() *IrreversibleTransaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *UpdateBountyStateData) GetBountyId() string {
	if x != nil {
		return x.BountyId
	}
	return ""
}

func (x *UpdateBountyStateData) GetState() BountyState {
	if x != nil {
		return x.State
	}
	return BountyState_RELEASED
}

var File_bounties_proto protoreflect.FileDescriptor

var file_bounties_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x31, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01,
	0x0a, 0x0e, 0x42, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x40, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x22, 0x8f, 0x02, 0x0a, 0x12, 0x42, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e,
	0x74, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a,
	0x0c, 0x65, 0x78, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0b, 0x65, 0x78, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x4b, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x72, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e,
	0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x33,
	0x0a, 0x0b, 0x42, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a,
	0x08, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44, 0x45,
	0x44, 0x10, 0x02, 0x2a, 0x3c, 0x0a, 0x12, 0x42, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x50,
	0x4f, 0x53, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52,
	0x41, 0x57, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x10,
	0x02, 0x32, 0xc8, 0x01, 0x0a, 0x1d, 0x50, 0x6f, 0x6d, 0x65, 0x6c, 0x6f, 0x42, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x75, 0x6e, 0x74, 0x79,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e,
	0x74, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x75, 0x6e, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x65, 0x6f, 0x73, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bounties_proto_rawDescOnce sync.Once
	file_bounties_proto_rawDescData = file_bounties_proto_rawDesc
)

func file_bounties_proto_rawDescGZIP() []byte {
	file_bounties_proto_rawDescOnce.Do(func() {
		file_bounties_proto_rawDescData = protoimpl.X.CompressGZIP(file_bounties_proto_rawDescData)
	})
	return file_bounties_proto_rawDescData
}

var file_bounties_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bounties_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bounties_proto_goTypes = []interface{}{
	(BountyState)(0),                // 0: eosn.protobuf.v1.BountyState
	(BountyTransferType)(0),         // 1: eosn.protobuf.v1.BountyTransferType
	(*BountyTransfer)(nil),          // 2: eosn.protobuf.v1.BountyTransfer
	(*BountyTransferData)(nil),      // 3: eosn.protobuf.v1.BountyTransferData
	(*UpdateBountyStateData)(nil),   // 4: eosn.protobuf.v1.UpdateBountyStateData
	(*Transaction)(nil),             // 5: eosn.protobuf.v1.Transaction
	(*TokenData)(nil),               // 6: eosn.protobuf.v1.TokenData
	(*IrreversibleTransaction)(nil), // 7: eosn.protobuf.v1.IrreversibleTransaction
	(*emptypb.Empty)(nil),           // 8: google.protobuf.Empty
}
var file_bounties_proto_depIdxs = []int32{
	5, // 0: eosn.protobuf.v1.BountyTransfer.transaction:type_name -> eosn.protobuf.v1.Transaction
	3, // 1: eosn.protobuf.v1.BountyTransfer.transfer:type_name -> eosn.protobuf.v1.BountyTransferData
	1, // 2: eosn.protobuf.v1.BountyTransferData.type:type_name -> eosn.protobuf.v1.BountyTransferType
	6, // 3: eosn.protobuf.v1.BountyTransferData.ext_quantity:type_name -> eosn.protobuf.v1.TokenData
	7, // 4: eosn.protobuf.v1.UpdateBountyStateData.transaction:type_name -> eosn.protobuf.v1.IrreversibleTransaction
	0, // 5: eosn.protobuf.v1.UpdateBountyStateData.state:type_name -> eosn.protobuf.v1.BountyState
	2, // 6: eosn.protobuf.v1.PomeloBountiesInternalService.AddBountyTransfer:input_type -> eosn.protobuf.v1.BountyTransfer
	4, // 7: eosn.protobuf.v1.PomeloBountiesInternalService.UpdateBountyState:input_type -> eosn.protobuf.v1.UpdateBountyStateData
	8, // 8: eosn.protobuf.v1.PomeloBountiesInternalService.AddBountyTransfer:output_type -> google.protobuf.Empty
	8, // 9: eosn.protobuf.v1.PomeloBountiesInternalService.UpdateBountyState:output_type -> google.protobuf.Empty
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_bounties_proto_init() }
func file_bounties_proto_init() {
	if File_bounties_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bounties_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BountyTransfer); i {
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
		file_bounties_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BountyTransferData); i {
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
		file_bounties_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBountyStateData); i {
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
			RawDescriptor: file_bounties_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bounties_proto_goTypes,
		DependencyIndexes: file_bounties_proto_depIdxs,
		EnumInfos:         file_bounties_proto_enumTypes,
		MessageInfos:      file_bounties_proto_msgTypes,
	}.Build()
	File_bounties_proto = out.File
	file_bounties_proto_rawDesc = nil
	file_bounties_proto_goTypes = nil
	file_bounties_proto_depIdxs = nil
}
