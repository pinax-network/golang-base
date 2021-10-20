// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-rc.1
// 	protoc        v3.17.3
// source: pomelo.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TraidingPairId int32

const (
	TraidingPairId_EOSUSDPAIR TraidingPairId = 0
)

// Enum value maps for TraidingPairId.
var (
	TraidingPairId_name = map[int32]string{
		0: "EOSUSDPAIR",
	}
	TraidingPairId_value = map[string]int32{
		"EOSUSDPAIR": 0,
	}
)

func (x TraidingPairId) Enum() *TraidingPairId {
	p := new(TraidingPairId)
	*p = x
	return p
}

func (x TraidingPairId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TraidingPairId) Descriptor() protoreflect.EnumDescriptor {
	return file_pomelo_proto_enumTypes[0].Descriptor()
}

func (TraidingPairId) Type() protoreflect.EnumType {
	return &file_pomelo_proto_enumTypes[0]
}

func (x TraidingPairId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TraidingPairId.Descriptor instead.
func (TraidingPairId) EnumDescriptor() ([]byte, []int) {
	return file_pomelo_proto_rawDescGZIP(), []int{0}
}

type TraidingPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction           `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Id          TraidingPairId         `protobuf:"varint,2,opt,name=id,proto3,enum=eosn.protobuf.TraidingPairId" json:"id,omitempty"`
	Rate        float32                `protobuf:"fixed32,3,opt,name=rate,proto3" json:"rate,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *TraidingPair) Reset() {
	*x = TraidingPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pomelo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraidingPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraidingPair) ProtoMessage() {}

func (x *TraidingPair) ProtoReflect() protoreflect.Message {
	mi := &file_pomelo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraidingPair.ProtoReflect.Descriptor instead.
func (*TraidingPair) Descriptor() ([]byte, []int) {
	return file_pomelo_proto_rawDescGZIP(), []int{0}
}

func (x *TraidingPair) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TraidingPair) GetId() TraidingPairId {
	if x != nil {
		return x.Id
	}
	return TraidingPairId_EOSUSDPAIR
}

func (x *TraidingPair) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *TraidingPair) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type Contribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	From        string       `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To          string       `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Memo        string       `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
	Quantity    string       `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Value       float32      `protobuf:"fixed32,6,opt,name=value,proto3" json:"value,omitempty"`
	UserId      string       `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoundId     int32        `protobuf:"varint,8,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
	SeasonId    int32        `protobuf:"varint,9,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
}

func (x *Contribution) Reset() {
	*x = Contribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pomelo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contribution) ProtoMessage() {}

func (x *Contribution) ProtoReflect() protoreflect.Message {
	mi := &file_pomelo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contribution.ProtoReflect.Descriptor instead.
func (*Contribution) Descriptor() ([]byte, []int) {
	return file_pomelo_proto_rawDescGZIP(), []int{1}
}

func (x *Contribution) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *Contribution) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Contribution) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Contribution) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Contribution) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *Contribution) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Contribution) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Contribution) GetRoundId() int32 {
	if x != nil {
		return x.RoundId
	}
	return 0
}

func (x *Contribution) GetSeasonId() int32 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

type GetLastTradingPairRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id TraidingPairId `protobuf:"varint,1,opt,name=id,proto3,enum=eosn.protobuf.TraidingPairId" json:"id,omitempty"`
}

func (x *GetLastTradingPairRequest) Reset() {
	*x = GetLastTradingPairRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pomelo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastTradingPairRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastTradingPairRequest) ProtoMessage() {}

func (x *GetLastTradingPairRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pomelo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastTradingPairRequest.ProtoReflect.Descriptor instead.
func (*GetLastTradingPairRequest) Descriptor() ([]byte, []int) {
	return file_pomelo_proto_rawDescGZIP(), []int{2}
}

func (x *GetLastTradingPairRequest) GetId() TraidingPairId {
	if x != nil {
		return x.Id
	}
	return TraidingPairId_EOSUSDPAIR
}

var File_pomelo_proto protoreflect.FileDescriptor

var file_pomelo_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x6d, 0x65, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x0c, 0x54, 0x72,
	0x61, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69,
	0x72, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x87, 0x02, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x72, 0x61, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x49, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x2a, 0x20, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69,
	0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4f, 0x53, 0x55, 0x53, 0x44, 0x50, 0x41, 0x49,
	0x52, 0x10, 0x00, 0x32, 0x8b, 0x02, 0x0a, 0x15, 0x50, 0x6f, 0x6d, 0x65, 0x6c, 0x6f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x69, 0x64, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x28, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x72, 0x61, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0f, 0x41, 0x64, 0x64, 0x54, 0x72, 0x61, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72,
	0x12, 0x1b, 0x2e, 0x65, 0x6f, 0x73, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x72, 0x61, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x65, 0x6f, 0x73,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pomelo_proto_rawDescOnce sync.Once
	file_pomelo_proto_rawDescData = file_pomelo_proto_rawDesc
)

func file_pomelo_proto_rawDescGZIP() []byte {
	file_pomelo_proto_rawDescOnce.Do(func() {
		file_pomelo_proto_rawDescData = protoimpl.X.CompressGZIP(file_pomelo_proto_rawDescData)
	})
	return file_pomelo_proto_rawDescData
}

var file_pomelo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pomelo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pomelo_proto_goTypes = []interface{}{
	(TraidingPairId)(0),               // 0: eosn.protobuf.TraidingPairId
	(*TraidingPair)(nil),              // 1: eosn.protobuf.TraidingPair
	(*Contribution)(nil),              // 2: eosn.protobuf.Contribution
	(*GetLastTradingPairRequest)(nil), // 3: eosn.protobuf.GetLastTradingPairRequest
	(*Transaction)(nil),               // 4: eosn.protobuf.Transaction
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_pomelo_proto_depIdxs = []int32{
	4, // 0: eosn.protobuf.TraidingPair.transaction:type_name -> eosn.protobuf.Transaction
	0, // 1: eosn.protobuf.TraidingPair.id:type_name -> eosn.protobuf.TraidingPairId
	5, // 2: eosn.protobuf.TraidingPair.time:type_name -> google.protobuf.Timestamp
	4, // 3: eosn.protobuf.Contribution.transaction:type_name -> eosn.protobuf.Transaction
	0, // 4: eosn.protobuf.GetLastTradingPairRequest.id:type_name -> eosn.protobuf.TraidingPairId
	3, // 5: eosn.protobuf.PomeloInternalService.GetLastTraidingPair:input_type -> eosn.protobuf.GetLastTradingPairRequest
	1, // 6: eosn.protobuf.PomeloInternalService.AddTraidingPair:input_type -> eosn.protobuf.TraidingPair
	2, // 7: eosn.protobuf.PomeloInternalService.AddContribution:input_type -> eosn.protobuf.Contribution
	1, // 8: eosn.protobuf.PomeloInternalService.GetLastTraidingPair:output_type -> eosn.protobuf.TraidingPair
	6, // 9: eosn.protobuf.PomeloInternalService.AddTraidingPair:output_type -> google.protobuf.Empty
	6, // 10: eosn.protobuf.PomeloInternalService.AddContribution:output_type -> google.protobuf.Empty
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pomelo_proto_init() }
func file_pomelo_proto_init() {
	if File_pomelo_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pomelo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraidingPair); i {
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
		file_pomelo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contribution); i {
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
		file_pomelo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastTradingPairRequest); i {
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
			RawDescriptor: file_pomelo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pomelo_proto_goTypes,
		DependencyIndexes: file_pomelo_proto_depIdxs,
		EnumInfos:         file_pomelo_proto_enumTypes,
		MessageInfos:      file_pomelo_proto_msgTypes,
	}.Build()
	File_pomelo_proto = out.File
	file_pomelo_proto_rawDesc = nil
	file_pomelo_proto_goTypes = nil
	file_pomelo_proto_depIdxs = nil
}
