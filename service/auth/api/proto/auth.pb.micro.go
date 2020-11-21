// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: micro/micro/auth/api/proto/auth.proto

package go_micro_api_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/yadisnel/go-ms/v2/client"
	server "github.com/yadisnel/go-ms/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	Verify(ctx context.Context, in *VerifyRequest, opts ...client.CallOption) (*VerifyResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Verify(ctx context.Context, in *VerifyRequest, opts ...client.CallOption) (*VerifyResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Verify", in)
	out := new(VerifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Verify(context.Context, *VerifyRequest, *VerifyResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Verify(ctx context.Context, in *VerifyRequest, out *VerifyResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Verify(ctx context.Context, in *VerifyRequest, out *VerifyResponse) error {
	return h.AuthHandler.Verify(ctx, in, out)
}
