// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.1
// source: api/openidprovider/v1/openidprovider.proto

package openidproviderv1

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

// OpenIDProviderClient is the client API for OpenIDProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenIDProviderClient interface {
	Name(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NameReply, error)
	BasicAuth(ctx context.Context, in *BasicAuthReq, opts ...grpc.CallOption) (*BasicAuthReply, error)
	TokenAuth(ctx context.Context, in *TokenAuthReq, opts ...grpc.CallOption) (*TokenAuthReply, error)
	SearchUid(ctx context.Context, in *SearchUidReq, opts ...grpc.CallOption) (*SearchUidReply, error)
}

type openIDProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenIDProviderClient(cc grpc.ClientConnInterface) OpenIDProviderClient {
	return &openIDProviderClient{cc}
}

func (c *openIDProviderClient) Name(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NameReply, error) {
	out := new(NameReply)
	err := c.cc.Invoke(ctx, "/openidprovider.v1.OpenIDProvider/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openIDProviderClient) BasicAuth(ctx context.Context, in *BasicAuthReq, opts ...grpc.CallOption) (*BasicAuthReply, error) {
	out := new(BasicAuthReply)
	err := c.cc.Invoke(ctx, "/openidprovider.v1.OpenIDProvider/BasicAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openIDProviderClient) TokenAuth(ctx context.Context, in *TokenAuthReq, opts ...grpc.CallOption) (*TokenAuthReply, error) {
	out := new(TokenAuthReply)
	err := c.cc.Invoke(ctx, "/openidprovider.v1.OpenIDProvider/TokenAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openIDProviderClient) SearchUid(ctx context.Context, in *SearchUidReq, opts ...grpc.CallOption) (*SearchUidReply, error) {
	out := new(SearchUidReply)
	err := c.cc.Invoke(ctx, "/openidprovider.v1.OpenIDProvider/SearchUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenIDProviderServer is the server API for OpenIDProvider service.
// All implementations must embed UnimplementedOpenIDProviderServer
// for forward compatibility
type OpenIDProviderServer interface {
	Name(context.Context, *emptypb.Empty) (*NameReply, error)
	BasicAuth(context.Context, *BasicAuthReq) (*BasicAuthReply, error)
	TokenAuth(context.Context, *TokenAuthReq) (*TokenAuthReply, error)
	SearchUid(context.Context, *SearchUidReq) (*SearchUidReply, error)
	mustEmbedUnimplementedOpenIDProviderServer()
}

// UnimplementedOpenIDProviderServer must be embedded to have forward compatible implementations.
type UnimplementedOpenIDProviderServer struct {
}

func (UnimplementedOpenIDProviderServer) Name(context.Context, *emptypb.Empty) (*NameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (UnimplementedOpenIDProviderServer) BasicAuth(context.Context, *BasicAuthReq) (*BasicAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasicAuth not implemented")
}
func (UnimplementedOpenIDProviderServer) TokenAuth(context.Context, *TokenAuthReq) (*TokenAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenAuth not implemented")
}
func (UnimplementedOpenIDProviderServer) SearchUid(context.Context, *SearchUidReq) (*SearchUidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUid not implemented")
}
func (UnimplementedOpenIDProviderServer) mustEmbedUnimplementedOpenIDProviderServer() {}

// UnsafeOpenIDProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenIDProviderServer will
// result in compilation errors.
type UnsafeOpenIDProviderServer interface {
	mustEmbedUnimplementedOpenIDProviderServer()
}

func RegisterOpenIDProviderServer(s grpc.ServiceRegistrar, srv OpenIDProviderServer) {
	s.RegisterService(&OpenIDProvider_ServiceDesc, srv)
}

func _OpenIDProvider_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenIDProviderServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openidprovider.v1.OpenIDProvider/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenIDProviderServer).Name(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenIDProvider_BasicAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenIDProviderServer).BasicAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openidprovider.v1.OpenIDProvider/BasicAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenIDProviderServer).BasicAuth(ctx, req.(*BasicAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenIDProvider_TokenAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenIDProviderServer).TokenAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openidprovider.v1.OpenIDProvider/TokenAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenIDProviderServer).TokenAuth(ctx, req.(*TokenAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenIDProvider_SearchUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenIDProviderServer).SearchUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openidprovider.v1.OpenIDProvider/SearchUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenIDProviderServer).SearchUid(ctx, req.(*SearchUidReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OpenIDProvider_ServiceDesc is the grpc.ServiceDesc for OpenIDProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenIDProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openidprovider.v1.OpenIDProvider",
	HandlerType: (*OpenIDProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Name",
			Handler:    _OpenIDProvider_Name_Handler,
		},
		{
			MethodName: "BasicAuth",
			Handler:    _OpenIDProvider_BasicAuth_Handler,
		},
		{
			MethodName: "TokenAuth",
			Handler:    _OpenIDProvider_TokenAuth_Handler,
		},
		{
			MethodName: "SearchUid",
			Handler:    _OpenIDProvider_SearchUid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/openidprovider/v1/openidprovider.proto",
}
