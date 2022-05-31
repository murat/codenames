// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protos/codenames.proto

package protos

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CodenamesClient is the client API for Codenames service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodenamesClient interface {
	CreateGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameResponse, error)
	JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type codenamesClient struct {
	cc grpc.ClientConnInterface
}

func NewCodenamesClient(cc grpc.ClientConnInterface) CodenamesClient {
	return &codenamesClient{cc}
}

func (c *codenamesClient) CreateGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameResponse, error) {
	out := new(GameResponse)
	err := c.cc.Invoke(ctx, "/protos.Codenames/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codenamesClient) JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protos.Codenames/JoinGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodenamesServer is the server API for Codenames service.
// All implementations must embed UnimplementedCodenamesServer
// for forward compatibility
type CodenamesServer interface {
	CreateGame(context.Context, *GameRequest) (*GameResponse, error)
	JoinGame(context.Context, *JoinGameRequest) (*empty.Empty, error)
	mustEmbedUnimplementedCodenamesServer()
}

// UnimplementedCodenamesServer must be embedded to have forward compatible implementations.
type UnimplementedCodenamesServer struct {
}

func (UnimplementedCodenamesServer) CreateGame(context.Context, *GameRequest) (*GameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedCodenamesServer) JoinGame(context.Context, *JoinGameRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedCodenamesServer) mustEmbedUnimplementedCodenamesServer() {}

// UnsafeCodenamesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodenamesServer will
// result in compilation errors.
type UnsafeCodenamesServer interface {
	mustEmbedUnimplementedCodenamesServer()
}

func RegisterCodenamesServer(s grpc.ServiceRegistrar, srv CodenamesServer) {
	s.RegisterService(&Codenames_ServiceDesc, srv)
}

func _Codenames_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodenamesServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Codenames/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodenamesServer).CreateGame(ctx, req.(*GameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Codenames_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodenamesServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Codenames/JoinGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodenamesServer).JoinGame(ctx, req.(*JoinGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Codenames_ServiceDesc is the grpc.ServiceDesc for Codenames service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Codenames_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Codenames",
	HandlerType: (*CodenamesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _Codenames_CreateGame_Handler,
		},
		{
			MethodName: "JoinGame",
			Handler:    _Codenames_JoinGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/codenames.proto",
}
