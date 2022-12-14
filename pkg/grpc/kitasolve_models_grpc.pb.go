// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/kitasolve_models.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KitaSolveModelsClient is the client API for KitaSolveModels service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KitaSolveModelsClient interface {
	AddSolve(ctx context.Context, in *Solve, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	GetSolve(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Solve, error)
	DeleteSolve(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	SearchSolve(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (KitaSolveModels_SearchSolveClient, error)
	UpdateSolve(ctx context.Context, opts ...grpc.CallOption) (KitaSolveModels_UpdateSolveClient, error)
}

type kitaSolveModelsClient struct {
	cc grpc.ClientConnInterface
}

func NewKitaSolveModelsClient(cc grpc.ClientConnInterface) KitaSolveModelsClient {
	return &kitaSolveModelsClient{cc}
}

func (c *kitaSolveModelsClient) AddSolve(ctx context.Context, in *Solve, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/kitasolve_models.kitaSolveModels/addSolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitaSolveModelsClient) GetSolve(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Solve, error) {
	out := new(Solve)
	err := c.cc.Invoke(ctx, "/kitasolve_models.kitaSolveModels/getSolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitaSolveModelsClient) DeleteSolve(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/kitasolve_models.kitaSolveModels/deleteSolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitaSolveModelsClient) SearchSolve(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (KitaSolveModels_SearchSolveClient, error) {
	stream, err := c.cc.NewStream(ctx, &KitaSolveModels_ServiceDesc.Streams[0], "/kitasolve_models.kitaSolveModels/searchSolve", opts...)
	if err != nil {
		return nil, err
	}
	x := &kitaSolveModelsSearchSolveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KitaSolveModels_SearchSolveClient interface {
	Recv() (*Solve, error)
	grpc.ClientStream
}

type kitaSolveModelsSearchSolveClient struct {
	grpc.ClientStream
}

func (x *kitaSolveModelsSearchSolveClient) Recv() (*Solve, error) {
	m := new(Solve)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kitaSolveModelsClient) UpdateSolve(ctx context.Context, opts ...grpc.CallOption) (KitaSolveModels_UpdateSolveClient, error) {
	stream, err := c.cc.NewStream(ctx, &KitaSolveModels_ServiceDesc.Streams[1], "/kitasolve_models.kitaSolveModels/updateSolve", opts...)
	if err != nil {
		return nil, err
	}
	x := &kitaSolveModelsUpdateSolveClient{stream}
	return x, nil
}

type KitaSolveModels_UpdateSolveClient interface {
	Send(*Solve) error
	CloseAndRecv() (*wrapperspb.StringValue, error)
	grpc.ClientStream
}

type kitaSolveModelsUpdateSolveClient struct {
	grpc.ClientStream
}

func (x *kitaSolveModelsUpdateSolveClient) Send(m *Solve) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kitaSolveModelsUpdateSolveClient) CloseAndRecv() (*wrapperspb.StringValue, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(wrapperspb.StringValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KitaSolveModelsServer is the server API for KitaSolveModels service.
// All implementations must embed UnimplementedKitaSolveModelsServer
// for forward compatibility
type KitaSolveModelsServer interface {
	AddSolve(context.Context, *Solve) (*wrapperspb.StringValue, error)
	GetSolve(context.Context, *wrapperspb.StringValue) (*Solve, error)
	DeleteSolve(context.Context, *wrapperspb.StringValue) (*wrapperspb.BoolValue, error)
	SearchSolve(*wrapperspb.StringValue, KitaSolveModels_SearchSolveServer) error
	UpdateSolve(KitaSolveModels_UpdateSolveServer) error
	mustEmbedUnimplementedKitaSolveModelsServer()
}

// UnimplementedKitaSolveModelsServer must be embedded to have forward compatible implementations.
type UnimplementedKitaSolveModelsServer struct {
}

func (UnimplementedKitaSolveModelsServer) AddSolve(context.Context, *Solve) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSolve not implemented")
}
func (UnimplementedKitaSolveModelsServer) GetSolve(context.Context, *wrapperspb.StringValue) (*Solve, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSolve not implemented")
}
func (UnimplementedKitaSolveModelsServer) DeleteSolve(context.Context, *wrapperspb.StringValue) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSolve not implemented")
}
func (UnimplementedKitaSolveModelsServer) SearchSolve(*wrapperspb.StringValue, KitaSolveModels_SearchSolveServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchSolve not implemented")
}
func (UnimplementedKitaSolveModelsServer) UpdateSolve(KitaSolveModels_UpdateSolveServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateSolve not implemented")
}
func (UnimplementedKitaSolveModelsServer) mustEmbedUnimplementedKitaSolveModelsServer() {}

// UnsafeKitaSolveModelsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KitaSolveModelsServer will
// result in compilation errors.
type UnsafeKitaSolveModelsServer interface {
	mustEmbedUnimplementedKitaSolveModelsServer()
}

func RegisterKitaSolveModelsServer(s grpc.ServiceRegistrar, srv KitaSolveModelsServer) {
	s.RegisterService(&KitaSolveModels_ServiceDesc, srv)
}

func _KitaSolveModels_AddSolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Solve)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitaSolveModelsServer).AddSolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitasolve_models.kitaSolveModels/addSolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitaSolveModelsServer).AddSolve(ctx, req.(*Solve))
	}
	return interceptor(ctx, in, info, handler)
}

func _KitaSolveModels_GetSolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitaSolveModelsServer).GetSolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitasolve_models.kitaSolveModels/getSolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitaSolveModelsServer).GetSolve(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _KitaSolveModels_DeleteSolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitaSolveModelsServer).DeleteSolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitasolve_models.kitaSolveModels/deleteSolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitaSolveModelsServer).DeleteSolve(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _KitaSolveModels_SearchSolve_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KitaSolveModelsServer).SearchSolve(m, &kitaSolveModelsSearchSolveServer{stream})
}

type KitaSolveModels_SearchSolveServer interface {
	Send(*Solve) error
	grpc.ServerStream
}

type kitaSolveModelsSearchSolveServer struct {
	grpc.ServerStream
}

func (x *kitaSolveModelsSearchSolveServer) Send(m *Solve) error {
	return x.ServerStream.SendMsg(m)
}

func _KitaSolveModels_UpdateSolve_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KitaSolveModelsServer).UpdateSolve(&kitaSolveModelsUpdateSolveServer{stream})
}

type KitaSolveModels_UpdateSolveServer interface {
	SendAndClose(*wrapperspb.StringValue) error
	Recv() (*Solve, error)
	grpc.ServerStream
}

type kitaSolveModelsUpdateSolveServer struct {
	grpc.ServerStream
}

func (x *kitaSolveModelsUpdateSolveServer) SendAndClose(m *wrapperspb.StringValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kitaSolveModelsUpdateSolveServer) Recv() (*Solve, error) {
	m := new(Solve)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KitaSolveModels_ServiceDesc is the grpc.ServiceDesc for KitaSolveModels service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KitaSolveModels_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kitasolve_models.kitaSolveModels",
	HandlerType: (*KitaSolveModelsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addSolve",
			Handler:    _KitaSolveModels_AddSolve_Handler,
		},
		{
			MethodName: "getSolve",
			Handler:    _KitaSolveModels_GetSolve_Handler,
		},
		{
			MethodName: "deleteSolve",
			Handler:    _KitaSolveModels_DeleteSolve_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchSolve",
			Handler:       _KitaSolveModels_SearchSolve_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "updateSolve",
			Handler:       _KitaSolveModels_UpdateSolve_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/kitasolve_models.proto",
}
