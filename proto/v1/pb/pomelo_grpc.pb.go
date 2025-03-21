// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: pomelo.proto

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

// PomeloInternalServiceClient is the client API for PomeloInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PomeloInternalServiceClient interface {
	GetLastTraidingPair(ctx context.Context, in *GetLastTradingPairRequest, opts ...grpc.CallOption) (*TraidingPair, error)
	AddTraidingPair(ctx context.Context, in *TraidingPair, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddContribution(ctx context.Context, in *Contribution, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateMatching(ctx context.Context, in *MatchingData, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pomeloInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPomeloInternalServiceClient(cc grpc.ClientConnInterface) PomeloInternalServiceClient {
	return &pomeloInternalServiceClient{cc}
}

func (c *pomeloInternalServiceClient) GetLastTraidingPair(ctx context.Context, in *GetLastTradingPairRequest, opts ...grpc.CallOption) (*TraidingPair, error) {
	out := new(TraidingPair)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.v1.PomeloInternalService/GetLastTraidingPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pomeloInternalServiceClient) AddTraidingPair(ctx context.Context, in *TraidingPair, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.v1.PomeloInternalService/AddTraidingPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pomeloInternalServiceClient) AddContribution(ctx context.Context, in *Contribution, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.v1.PomeloInternalService/AddContribution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pomeloInternalServiceClient) UpdateMatching(ctx context.Context, in *MatchingData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/eosn.protobuf.v1.PomeloInternalService/UpdateMatching", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PomeloInternalServiceServer is the server API for PomeloInternalService service.
// All implementations must embed UnimplementedPomeloInternalServiceServer
// for forward compatibility
type PomeloInternalServiceServer interface {
	GetLastTraidingPair(context.Context, *GetLastTradingPairRequest) (*TraidingPair, error)
	AddTraidingPair(context.Context, *TraidingPair) (*emptypb.Empty, error)
	AddContribution(context.Context, *Contribution) (*emptypb.Empty, error)
	UpdateMatching(context.Context, *MatchingData) (*emptypb.Empty, error)
	mustEmbedUnimplementedPomeloInternalServiceServer()
}

// UnimplementedPomeloInternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPomeloInternalServiceServer struct {
}

func (UnimplementedPomeloInternalServiceServer) GetLastTraidingPair(context.Context, *GetLastTradingPairRequest) (*TraidingPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastTraidingPair not implemented")
}
func (UnimplementedPomeloInternalServiceServer) AddTraidingPair(context.Context, *TraidingPair) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTraidingPair not implemented")
}
func (UnimplementedPomeloInternalServiceServer) AddContribution(context.Context, *Contribution) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContribution not implemented")
}
func (UnimplementedPomeloInternalServiceServer) UpdateMatching(context.Context, *MatchingData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMatching not implemented")
}
func (UnimplementedPomeloInternalServiceServer) mustEmbedUnimplementedPomeloInternalServiceServer() {}

// UnsafePomeloInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PomeloInternalServiceServer will
// result in compilation errors.
type UnsafePomeloInternalServiceServer interface {
	mustEmbedUnimplementedPomeloInternalServiceServer()
}

func RegisterPomeloInternalServiceServer(s grpc.ServiceRegistrar, srv PomeloInternalServiceServer) {
	s.RegisterService(&PomeloInternalService_ServiceDesc, srv)
}

func _PomeloInternalService_GetLastTraidingPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastTradingPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PomeloInternalServiceServer).GetLastTraidingPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.v1.PomeloInternalService/GetLastTraidingPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PomeloInternalServiceServer).GetLastTraidingPair(ctx, req.(*GetLastTradingPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PomeloInternalService_AddTraidingPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraidingPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PomeloInternalServiceServer).AddTraidingPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.v1.PomeloInternalService/AddTraidingPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PomeloInternalServiceServer).AddTraidingPair(ctx, req.(*TraidingPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _PomeloInternalService_AddContribution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contribution)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PomeloInternalServiceServer).AddContribution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.v1.PomeloInternalService/AddContribution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PomeloInternalServiceServer).AddContribution(ctx, req.(*Contribution))
	}
	return interceptor(ctx, in, info, handler)
}

func _PomeloInternalService_UpdateMatching_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchingData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PomeloInternalServiceServer).UpdateMatching(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eosn.protobuf.v1.PomeloInternalService/UpdateMatching",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PomeloInternalServiceServer).UpdateMatching(ctx, req.(*MatchingData))
	}
	return interceptor(ctx, in, info, handler)
}

// PomeloInternalService_ServiceDesc is the grpc.ServiceDesc for PomeloInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PomeloInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eosn.protobuf.v1.PomeloInternalService",
	HandlerType: (*PomeloInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastTraidingPair",
			Handler:    _PomeloInternalService_GetLastTraidingPair_Handler,
		},
		{
			MethodName: "AddTraidingPair",
			Handler:    _PomeloInternalService_AddTraidingPair_Handler,
		},
		{
			MethodName: "AddContribution",
			Handler:    _PomeloInternalService_AddContribution_Handler,
		},
		{
			MethodName: "UpdateMatching",
			Handler:    _PomeloInternalService_UpdateMatching_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pomelo.proto",
}
