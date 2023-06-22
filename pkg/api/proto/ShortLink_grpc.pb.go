// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/proto/ShortLink.proto

package links

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

// ShortLinkClient is the client API for ShortLink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortLinkClient interface {
	Get(ctx context.Context, in *SlRequest, opts ...grpc.CallOption) (*SlResponse, error)
	Post(ctx context.Context, in *SlRequest, opts ...grpc.CallOption) (*SlResponse, error)
}

type shortLinkClient struct {
	cc grpc.ClientConnInterface
}

func NewShortLinkClient(cc grpc.ClientConnInterface) ShortLinkClient {
	return &shortLinkClient{cc}
}

func (c *shortLinkClient) Get(ctx context.Context, in *SlRequest, opts ...grpc.CallOption) (*SlResponse, error) {
	out := new(SlResponse)
	err := c.cc.Invoke(ctx, "/api.ShortLink/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortLinkClient) Post(ctx context.Context, in *SlRequest, opts ...grpc.CallOption) (*SlResponse, error) {
	out := new(SlResponse)
	err := c.cc.Invoke(ctx, "/api.ShortLink/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortLinkServer is the server API for ShortLink service.
// All implementations must embed UnimplementedShortLinkServer
// for forward compatibility
type ShortLinkServer interface {
	Get(context.Context, *SlRequest) (*SlResponse, error)
	Post(context.Context, *SlRequest) (*SlResponse, error)
	mustEmbedUnimplementedShortLinkServer()
}

// UnimplementedShortLinkServer must be embedded to have forward compatible implementations.
type UnimplementedShortLinkServer struct {
}

func (UnimplementedShortLinkServer) Get(context.Context, *SlRequest) (*SlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedShortLinkServer) Post(context.Context, *SlRequest) (*SlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedShortLinkServer) mustEmbedUnimplementedShortLinkServer() {}

// UnsafeShortLinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortLinkServer will
// result in compilation errors.
type UnsafeShortLinkServer interface {
	mustEmbedUnimplementedShortLinkServer()
}

func RegisterShortLinkServer(s grpc.ServiceRegistrar, srv ShortLinkServer) {
	s.RegisterService(&ShortLink_ServiceDesc, srv)
}

func _ShortLink_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortLinkServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ShortLink/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortLinkServer).Get(ctx, req.(*SlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortLink_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortLinkServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ShortLink/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortLinkServer).Post(ctx, req.(*SlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortLink_ServiceDesc is the grpc.ServiceDesc for ShortLink service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortLink_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ShortLink",
	HandlerType: (*ShortLinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ShortLink_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _ShortLink_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/ShortLink.proto",
}