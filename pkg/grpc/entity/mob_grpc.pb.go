// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: grpc/entity/mob.proto

package entity_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MobServiceClient is the client API for MobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MobServiceClient interface {
	CreateMob(ctx context.Context, in *CreateMobDto, opts ...grpc.CallOption) (*Mob, error)
	GetMobs(ctx context.Context, in *GetMobsRequestDto, opts ...grpc.CallOption) (*ReceiveMobsDto, error)
}

type mobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMobServiceClient(cc grpc.ClientConnInterface) MobServiceClient {
	return &mobServiceClient{cc}
}

func (c *mobServiceClient) CreateMob(ctx context.Context, in *CreateMobDto, opts ...grpc.CallOption) (*Mob, error) {
	out := new(Mob)
	err := c.cc.Invoke(ctx, "/grpc.entity.v1.MobService/CreateMob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobServiceClient) GetMobs(ctx context.Context, in *GetMobsRequestDto, opts ...grpc.CallOption) (*ReceiveMobsDto, error) {
	out := new(ReceiveMobsDto)
	err := c.cc.Invoke(ctx, "/grpc.entity.v1.MobService/GetMobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobServiceServer is the server API for MobService service.
// All implementations must embed UnimplementedMobServiceServer
// for forward compatibility
type MobServiceServer interface {
	CreateMob(context.Context, *CreateMobDto) (*Mob, error)
	GetMobs(context.Context, *GetMobsRequestDto) (*ReceiveMobsDto, error)
	mustEmbedUnimplementedMobServiceServer()
}

// UnimplementedMobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMobServiceServer struct {
}

func (UnimplementedMobServiceServer) CreateMob(context.Context, *CreateMobDto) (*Mob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMob not implemented")
}
func (UnimplementedMobServiceServer) GetMobs(context.Context, *GetMobsRequestDto) (*ReceiveMobsDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMobs not implemented")
}
func (UnimplementedMobServiceServer) mustEmbedUnimplementedMobServiceServer() {}

// UnsafeMobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobServiceServer will
// result in compilation errors.
type UnsafeMobServiceServer interface {
	mustEmbedUnimplementedMobServiceServer()
}

func RegisterMobServiceServer(s grpc.ServiceRegistrar, srv MobServiceServer) {
	s.RegisterService(&MobService_ServiceDesc, srv)
}

func _MobService_CreateMob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMobDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobServiceServer).CreateMob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.entity.v1.MobService/CreateMob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobServiceServer).CreateMob(ctx, req.(*CreateMobDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobService_GetMobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMobsRequestDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobServiceServer).GetMobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.entity.v1.MobService/GetMobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobServiceServer).GetMobs(ctx, req.(*GetMobsRequestDto))
	}
	return interceptor(ctx, in, info, handler)
}

// MobService_ServiceDesc is the grpc.ServiceDesc for MobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.entity.v1.MobService",
	HandlerType: (*MobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMob",
			Handler:    _MobService_CreateMob_Handler,
		},
		{
			MethodName: "GetMobs",
			Handler:    _MobService_GetMobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/entity/mob.proto",
}
