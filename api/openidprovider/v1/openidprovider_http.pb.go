// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package openidproviderv1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type OpenIDProviderHTTPServer interface {
	BasicAuth(context.Context, *BasicAuthReq) (*BasicAuthReply, error)
	Name(context.Context, *emptypb.Empty) (*NameReply, error)
	SearchUid(context.Context, *SearchUidReq) (*SearchUidReply, error)
	TokenAuth(context.Context, *TokenAuthReq) (*TokenAuthReply, error)
}

func RegisterOpenIDProviderHTTPServer(s *http.Server, srv OpenIDProviderHTTPServer) {
	r := s.Route("/")
	r.POST("/openidprovider/name", _OpenIDProvider_Name0_HTTP_Handler(srv))
	r.POST("/openidprovider/basic-auth", _OpenIDProvider_BasicAuth0_HTTP_Handler(srv))
	r.POST("/openidprovider/token-auth", _OpenIDProvider_TokenAuth0_HTTP_Handler(srv))
	r.POST("/openidprovider/search-uid", _OpenIDProvider_SearchUid0_HTTP_Handler(srv))
}

func _OpenIDProvider_Name0_HTTP_Handler(srv OpenIDProviderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/openidprovider.v1.OpenIDProvider/Name")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Name(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NameReply)
		return ctx.Result(200, reply)
	}
}

func _OpenIDProvider_BasicAuth0_HTTP_Handler(srv OpenIDProviderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BasicAuthReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/openidprovider.v1.OpenIDProvider/BasicAuth")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BasicAuth(ctx, req.(*BasicAuthReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BasicAuthReply)
		return ctx.Result(200, reply)
	}
}

func _OpenIDProvider_TokenAuth0_HTTP_Handler(srv OpenIDProviderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TokenAuthReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/openidprovider.v1.OpenIDProvider/TokenAuth")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TokenAuth(ctx, req.(*TokenAuthReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TokenAuthReply)
		return ctx.Result(200, reply)
	}
}

func _OpenIDProvider_SearchUid0_HTTP_Handler(srv OpenIDProviderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SearchUidReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/openidprovider.v1.OpenIDProvider/SearchUid")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SearchUid(ctx, req.(*SearchUidReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SearchUidReply)
		return ctx.Result(200, reply)
	}
}

type OpenIDProviderHTTPClient interface {
	BasicAuth(ctx context.Context, req *BasicAuthReq, opts ...http.CallOption) (rsp *BasicAuthReply, err error)
	Name(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *NameReply, err error)
	SearchUid(ctx context.Context, req *SearchUidReq, opts ...http.CallOption) (rsp *SearchUidReply, err error)
	TokenAuth(ctx context.Context, req *TokenAuthReq, opts ...http.CallOption) (rsp *TokenAuthReply, err error)
}

type OpenIDProviderHTTPClientImpl struct {
	cc *http.Client
}

func NewOpenIDProviderHTTPClient(client *http.Client) OpenIDProviderHTTPClient {
	return &OpenIDProviderHTTPClientImpl{client}
}

func (c *OpenIDProviderHTTPClientImpl) BasicAuth(ctx context.Context, in *BasicAuthReq, opts ...http.CallOption) (*BasicAuthReply, error) {
	var out BasicAuthReply
	pattern := "/openidprovider/basic-auth"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/openidprovider.v1.OpenIDProvider/BasicAuth"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OpenIDProviderHTTPClientImpl) Name(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*NameReply, error) {
	var out NameReply
	pattern := "/openidprovider/name"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/openidprovider.v1.OpenIDProvider/Name"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OpenIDProviderHTTPClientImpl) SearchUid(ctx context.Context, in *SearchUidReq, opts ...http.CallOption) (*SearchUidReply, error) {
	var out SearchUidReply
	pattern := "/openidprovider/search-uid"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/openidprovider.v1.OpenIDProvider/SearchUid"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OpenIDProviderHTTPClientImpl) TokenAuth(ctx context.Context, in *TokenAuthReq, opts ...http.CallOption) (*TokenAuthReply, error) {
	var out TokenAuthReply
	pattern := "/openidprovider/token-auth"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/openidprovider.v1.OpenIDProvider/TokenAuth"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
