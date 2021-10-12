// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccountsInternalServiceClient is the client API for AccountsInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsInternalServiceClient interface {
	SignUserTransaction(ctx context.Context, in *UserTransactionSignatureRequest, opts ...grpc.CallOption) (*SignedTransaction, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LinkAccount(ctx context.Context, in *LinkAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UnlinkAccount(ctx context.Context, in *UnlinkAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StreamUpdatedUserSocials(ctx context.Context, in *GetUpdatedSocialsRequest, opts ...grpc.CallOption) (AccountsInternalService_StreamUpdatedUserSocialsClient, error)
}

type accountsInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsInternalServiceClient(cc grpc.ClientConnInterface) AccountsInternalServiceClient {
	return &accountsInternalServiceClient{cc}
}

func (c *accountsInternalServiceClient) SignUserTransaction(ctx context.Context, in *UserTransactionSignatureRequest, opts ...grpc.CallOption) (*SignedTransaction, error) {
	out := new(SignedTransaction)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.AccountsInternalService/SignUserTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsInternalServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.AccountsInternalService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsInternalServiceClient) LinkAccount(ctx context.Context, in *LinkAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.AccountsInternalService/LinkAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsInternalServiceClient) UnlinkAccount(ctx context.Context, in *UnlinkAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.AccountsInternalService/UnlinkAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsInternalServiceClient) StreamUpdatedUserSocials(ctx context.Context, in *GetUpdatedSocialsRequest, opts ...grpc.CallOption) (AccountsInternalService_StreamUpdatedUserSocialsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AccountsInternalService_ServiceDesc.Streams[0], "/eosn.protobuf.AccountsInternalService/StreamUpdatedUserSocials", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountsInternalServiceStreamUpdatedUserSocialsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccountsInternalService_StreamUpdatedUserSocialsClient interface {
	Recv() (*LinkedSocial, error)
	grpc.ClientStream
}

type accountsInternalServiceStreamUpdatedUserSocialsClient struct {
	grpc.ClientStream
}

func (x *accountsInternalServiceStreamUpdatedUserSocialsClient) Recv() (*LinkedSocial, error) {
	m := new(LinkedSocial)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountsInternalServiceServer is the server API for AccountsInternalService service.
// All implementations must embed UnimplementedAccountsInternalServiceServer
// for forward compatibility
type AccountsInternalServiceServer interface {
	SignUserTransaction(context.Context, *UserTransactionSignatureRequest) (*SignedTransaction, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*emptypb.Empty, error)
	LinkAccount(context.Context, *LinkAccountRequest) (*emptypb.Empty, error)
	UnlinkAccount(context.Context, *UnlinkAccountRequest) (*emptypb.Empty, error)
	StreamUpdatedUserSocials(*GetUpdatedSocialsRequest, AccountsInternalService_StreamUpdatedUserSocialsServer) error
	mustEmbedUnimplementedAccountsInternalServiceServer()
}

// UnimplementedAccountsInternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsInternalServiceServer struct {
}

func (UnimplementedAccountsInternalServiceServer) SignUserTransaction(context.Context, *UserTransactionSignatureRequest) (*SignedTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUserTransaction not implemented")
}
func (UnimplementedAccountsInternalServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountsInternalServiceServer) LinkAccount(context.Context, *LinkAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkAccount not implemented")
}
func (UnimplementedAccountsInternalServiceServer) UnlinkAccount(context.Context, *UnlinkAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkAccount not implemented")
}
func (UnimplementedAccountsInternalServiceServer) StreamUpdatedUserSocials(*GetUpdatedSocialsRequest, AccountsInternalService_StreamUpdatedUserSocialsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamUpdatedUserSocials not implemented")
}
func (UnimplementedAccountsInternalServiceServer) mustEmbedUnimplementedAccountsInternalServiceServer() {
}

// UnsafeAccountsInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsInternalServiceServer will
// result in compilation errors.
type UnsafeAccountsInternalServiceServer interface {
	mustEmbedUnimplementedAccountsInternalServiceServer()
}

func RegisterAccountsInternalServiceServer(s grpc.ServiceRegistrar, srv AccountsInternalServiceServer) {
	s.RegisterService(&AccountsInternalService_ServiceDesc, srv)
}

func _AccountsInternalService_SignUserTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTransactionSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsInternalServiceServer).SignUserTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.AccountsInternalService/SignUserTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsInternalServiceServer).SignUserTransaction(ctx, req.(*UserTransactionSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsInternalService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsInternalServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.AccountsInternalService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsInternalServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsInternalService_LinkAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsInternalServiceServer).LinkAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.AccountsInternalService/LinkAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsInternalServiceServer).LinkAccount(ctx, req.(*LinkAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsInternalService_UnlinkAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlinkAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsInternalServiceServer).UnlinkAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.AccountsInternalService/UnlinkAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsInternalServiceServer).UnlinkAccount(ctx, req.(*UnlinkAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsInternalService_StreamUpdatedUserSocials_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUpdatedSocialsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountsInternalServiceServer).StreamUpdatedUserSocials(m, &accountsInternalServiceStreamUpdatedUserSocialsServer{stream})
}

type AccountsInternalService_StreamUpdatedUserSocialsServer interface {
	Send(*LinkedSocial) error
	grpc.ServerStream
}

type accountsInternalServiceStreamUpdatedUserSocialsServer struct {
	grpc.ServerStream
}

func (x *accountsInternalServiceStreamUpdatedUserSocialsServer) Send(m *LinkedSocial) error {
	return x.ServerStream.SendMsg(m)
}

// AccountsInternalService_ServiceDesc is the grpc.ServiceDesc for AccountsInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountsInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eosn.protobuf.AccountsInternalService",
	HandlerType: (*AccountsInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUserTransaction",
			Handler:    _AccountsInternalService_SignUserTransaction_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _AccountsInternalService_CreateAccount_Handler,
		},
		{
			MethodName: "LinkAccount",
			Handler:    _AccountsInternalService_LinkAccount_Handler,
		},
		{
			MethodName: "UnlinkAccount",
			Handler:    _AccountsInternalService_UnlinkAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamUpdatedUserSocials",
			Handler:       _AccountsInternalService_StreamUpdatedUserSocials_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "accounts.proto",
}
