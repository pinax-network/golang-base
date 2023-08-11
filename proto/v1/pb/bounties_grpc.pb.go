// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: bounties.proto

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

// PomeloBountiesInternalServiceClient is the client API for PomeloBountiesInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PomeloBountiesInternalServiceClient interface {
	AddBountyTransfer(ctx context.Context, in *BountyTransfer, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// deprecated
	UpdateBountyState(ctx context.Context, in *UpdateBountyStateData, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pomeloBountiesInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPomeloBountiesInternalServiceClient(cc grpc.ClientConnInterface) PomeloBountiesInternalServiceClient {
	return &pomeloBountiesInternalServiceClient{cc}
}

func (c *pomeloBountiesInternalServiceClient) AddBountyTransfer(ctx context.Context, in *BountyTransfer, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.v1.PomeloBountiesInternalService/AddBountyTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pomeloBountiesInternalServiceClient) UpdateBountyState(ctx context.Context, in *UpdateBountyStateData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.v1.PomeloBountiesInternalService/UpdateBountyState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PomeloBountiesInternalServiceServer is the server API for PomeloBountiesInternalService service.
// All implementations must embed UnimplementedPomeloBountiesInternalServiceServer
// for forward compatibility
type PomeloBountiesInternalServiceServer interface {
	AddBountyTransfer(context.Context, *BountyTransfer) (*emptypb.Empty, error)
	// deprecated
	UpdateBountyState(context.Context, *UpdateBountyStateData) (*emptypb.Empty, error)
	mustEmbedUnimplementedPomeloBountiesInternalServiceServer()
}

// UnimplementedPomeloBountiesInternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPomeloBountiesInternalServiceServer struct {
}

func (UnimplementedPomeloBountiesInternalServiceServer) AddBountyTransfer(context.Context, *BountyTransfer) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBountyTransfer not implemented")
}
func (UnimplementedPomeloBountiesInternalServiceServer) UpdateBountyState(context.Context, *UpdateBountyStateData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBountyState not implemented")
}
func (UnimplementedPomeloBountiesInternalServiceServer) mustEmbedUnimplementedPomeloBountiesInternalServiceServer() {
}

// UnsafePomeloBountiesInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PomeloBountiesInternalServiceServer will
// result in compilation errors.
type UnsafePomeloBountiesInternalServiceServer interface {
	mustEmbedUnimplementedPomeloBountiesInternalServiceServer()
}

func RegisterPomeloBountiesInternalServiceServer(s grpc.ServiceRegistrar, srv PomeloBountiesInternalServiceServer) {
	s.RegisterService(&PomeloBountiesInternalService_ServiceDesc, srv)
}

func _PomeloBountiesInternalService_AddBountyTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BountyTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PomeloBountiesInternalServiceServer).AddBountyTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.v1.PomeloBountiesInternalService/AddBountyTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PomeloBountiesInternalServiceServer).AddBountyTransfer(ctx, req.(*BountyTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _PomeloBountiesInternalService_UpdateBountyState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBountyStateData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PomeloBountiesInternalServiceServer).UpdateBountyState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.v1.PomeloBountiesInternalService/UpdateBountyState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PomeloBountiesInternalServiceServer).UpdateBountyState(ctx, req.(*UpdateBountyStateData))
	}
	return interceptor(ctx, in, info, handler)
}

// PomeloBountiesInternalService_ServiceDesc is the grpc.ServiceDesc for PomeloBountiesInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PomeloBountiesInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eosn.protobuf.v1.PomeloBountiesInternalService",
	HandlerType: (*PomeloBountiesInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBountyTransfer",
			Handler:    _PomeloBountiesInternalService_AddBountyTransfer_Handler,
		},
		{
			MethodName: "UpdateBountyState",
			Handler:    _PomeloBountiesInternalService_UpdateBountyState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bounties.proto",
}
