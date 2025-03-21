package dfuse

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Query                string          `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Variables            *_struct.Struct `protobuf:"bytes,2,opt,name=variables,proto3" json:"variables,omitempty"`
	OperationName        string          `protobuf:"bytes,3,opt,name=operationName,proto3" json:"operationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1746f99f7140fb04, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *Request) GetVariables() *_struct.Struct {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *Request) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

type Response struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Errors               []*Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1746f99f7140fb04, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

// GraphQL Error
type Error struct {
	// Description of the error intended for the developer.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// The source location for the error.
	Locations []*SourceLocation `protobuf:"bytes,2,rep,name=locations,proto3" json:"locations,omitempty"`
	// Path to the `null` value justified by this error.
	Path *_struct.ListValue `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// Free-form extensions (starts with a map)
	Extensions           *_struct.Struct `protobuf:"bytes,4,opt,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1746f99f7140fb04, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetLocations() []*SourceLocation {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *Error) GetPath() *_struct.ListValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Error) GetExtensions() *_struct.Struct {
	if m != nil {
		return m.Extensions
	}
	return nil
}

// The source location of an error.
type SourceLocation struct {
	// The line the error occurred at.
	Line int32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	// The column the error occurred at.
	Column               int32    `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceLocation) Reset()         { *m = SourceLocation{} }
func (m *SourceLocation) String() string { return proto.CompactTextString(m) }
func (*SourceLocation) ProtoMessage()    {}
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1746f99f7140fb04, []int{3}
}

func (m *SourceLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceLocation.Unmarshal(m, b)
}
func (m *SourceLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceLocation.Marshal(b, m, deterministic)
}
func (m *SourceLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceLocation.Merge(m, src)
}
func (m *SourceLocation) XXX_Size() int {
	return xxx_messageInfo_SourceLocation.Size(m)
}
func (m *SourceLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceLocation.DiscardUnknown(m)
}

var xxx_messageInfo_SourceLocation proto.InternalMessageInfo

func (m *SourceLocation) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *SourceLocation) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "dfuse.eosio.v1.Request")
	proto.RegisterType((*Response)(nil), "dfuse.eosio.v1.Response")
	proto.RegisterType((*Error)(nil), "dfuse.eosio.v1.Error")
	proto.RegisterType((*SourceLocation)(nil), "dfuse.eosio.v1.SourceLocation")
}

func init() { proto.RegisterFile("lib/lib.proto", fileDescriptor_1746f99f7140fb04) }

var fileDescriptor_1746f99f7140fb04 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0xcb, 0xd3, 0x40,
	0x10, 0x36, 0xb6, 0x69, 0xcc, 0x04, 0x7b, 0x58, 0xac, 0x0d, 0x45, 0xa4, 0x04, 0x0f, 0xbd, 0xb8,
	0xd5, 0x88, 0x78, 0xe9, 0x41, 0x84, 0xe2, 0xc1, 0x2a, 0xb8, 0x05, 0x0f, 0xde, 0x36, 0xe9, 0x34,
	0x0d, 0xa4, 0xd9, 0x74, 0x3f, 0x4a, 0x05, 0x7f, 0xa0, 0x3f, 0x4b, 0xb2, 0xd9, 0xd2, 0xb7, 0x2d,
	0xbc, 0xa7, 0x9d, 0x8f, 0x67, 0x9f, 0x99, 0x67, 0x66, 0x60, 0x54, 0x48, 0xde, 0xec, 0x0e, 0xd5,
	0xdc, 0xbd, 0xb4, 0x91, 0x42, 0x0b, 0x32, 0xdc, 0x6c, 0x8d, 0x42, 0x8a, 0x42, 0x95, 0x82, 0x1e,
	0xdf, 0x4f, 0x5e, 0x15, 0x42, 0x14, 0x15, 0xce, 0x6d, 0x36, 0x33, 0xdb, 0xb9, 0xd2, 0xd2, 0xe4,
	0xba, 0x43, 0x27, 0x7f, 0x21, 0x60, 0x78, 0x30, 0xa8, 0x34, 0x79, 0x01, 0xfe, 0xc1, 0xa0, 0xfc,
	0x13, 0x7b, 0x53, 0x6f, 0x16, 0xb2, 0xce, 0x21, 0x1f, 0x21, 0x3c, 0x72, 0x59, 0xf2, 0xac, 0x42,
	0x15, 0x3f, 0x9d, 0x7a, 0xb3, 0x28, 0x1d, 0xd3, 0x8e, 0x92, 0x9e, 0x29, 0xe9, 0xda, 0x52, 0xb2,
	0x0b, 0x92, 0xbc, 0x81, 0xe7, 0xa2, 0x41, 0xc9, 0x75, 0x29, 0xea, 0x1f, 0x7c, 0x8f, 0x71, 0xcf,
	0x92, 0x5e, 0x07, 0x93, 0xef, 0xf0, 0x8c, 0xa1, 0x6a, 0x44, 0xad, 0x90, 0x10, 0xe8, 0x6f, 0xb8,
	0xe6, 0xae, 0xba, 0xb5, 0xc9, 0x5b, 0x18, 0xa0, 0x94, 0x42, 0xb6, 0x95, 0x7b, 0xb3, 0x28, 0x1d,
	0xd1, 0x6b, 0x71, 0x74, 0xd9, 0x66, 0x99, 0x03, 0x25, 0xff, 0x3c, 0xf0, 0x6d, 0x84, 0xc4, 0x10,
	0xec, 0x51, 0x29, 0x5e, 0xa0, 0xe3, 0x3b, 0xbb, 0x64, 0x01, 0x61, 0x25, 0x72, 0xdb, 0xc2, 0x99,
	0xf5, 0xf5, 0x2d, 0xeb, 0x5a, 0x18, 0x99, 0xe3, 0xca, 0xc1, 0xd8, 0xe5, 0x03, 0xa1, 0xd0, 0x6f,
	0xb8, 0xde, 0x59, 0x35, 0x51, 0x3a, 0xb9, 0x1b, 0xc4, 0xaa, 0x54, 0xfa, 0x17, 0xaf, 0x0c, 0x32,
	0x8b, 0x23, 0x9f, 0x00, 0xf0, 0xa4, 0xb1, 0x56, 0xb6, 0x5c, 0xff, 0xf1, 0xf1, 0x3d, 0x80, 0x26,
	0x0b, 0x18, 0x5e, 0x77, 0xd1, 0xce, 0xa7, 0x2a, 0xeb, 0x4e, 0x8f, 0xcf, 0xac, 0x4d, 0x5e, 0xc2,
	0x20, 0x17, 0x95, 0xd9, 0xd7, 0x76, 0x33, 0x3e, 0x73, 0x5e, 0xfa, 0x0d, 0x82, 0xaf, 0xed, 0x51,
	0xfc, 0x5c, 0x91, 0xcf, 0x10, 0x2c, 0x4f, 0x98, 0x1b, 0x8d, 0x64, 0x7c, 0xab, 0xd3, 0x6d, 0x7e,
	0x12, 0xdf, 0x27, 0xba, 0xa5, 0x24, 0x4f, 0xde, 0x79, 0x5f, 0xa2, 0xdf, 0x61, 0x93, 0xb9, 0x1b,
	0xcb, 0x06, 0xb6, 0xe9, 0x0f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x38, 0x11, 0xe1, 0x33, 0x7d,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GraphQLClient is the client API for GraphQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GraphQLClient interface {
	Execute(ctx context.Context, in *Request, opts ...grpc.CallOption) (GraphQL_ExecuteClient, error)
}

type graphQLClient struct {
	cc *grpc.ClientConn
}

func NewGraphQLClient(cc *grpc.ClientConn) GraphQLClient {
	return &graphQLClient{cc}
}

func (c *graphQLClient) Execute(ctx context.Context, in *Request, opts ...grpc.CallOption) (GraphQL_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GraphQL_serviceDesc.Streams[0], "/dfuse.graphql.v1.GraphQL/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &graphQLExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GraphQL_ExecuteClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type graphQLExecuteClient struct {
	grpc.ClientStream
}

func (x *graphQLExecuteClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GraphQLServer is the server API for GraphQL service.
type GraphQLServer interface {
	Execute(*Request, GraphQL_ExecuteServer) error
}

func RegisterGraphQLServer(s *grpc.Server, srv GraphQLServer) {
	s.RegisterService(&_GraphQL_serviceDesc, srv)
}

func _GraphQL_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GraphQLServer).Execute(m, &graphQLExecuteServer{stream})
}

type GraphQL_ExecuteServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type graphQLExecuteServer struct {
	grpc.ServerStream
}

func (x *graphQLExecuteServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _GraphQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.graphql.v1.GraphQL",
	HandlerType: (*GraphQLServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _GraphQL_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lib/lib.proto",
}
