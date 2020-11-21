// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: crufter/micro/debug/stats/proto/debug.proto

package go_micro_debug_stats

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

// Client API for Stats service

type StatsService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error)
	Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Stats_StreamService, error)
}

type statsService struct {
	c    client.Client
	name string
}

func NewStatsService(name string, c client.Client) StatsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.debug.stats"
	}
	return &statsService{
		c:    c,
		name: name,
	}
}

func (c *statsService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Stats.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsService) Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error) {
	req := c.c.NewRequest(c.name, "Stats.Write", in)
	out := new(WriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsService) Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Stats_StreamService, error) {
	req := c.c.NewRequest(c.name, "Stats.Stream", &StreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &statsServiceStream{stream}, nil
}

type Stats_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamResponse, error)
}

type statsServiceStream struct {
	stream client.Stream
}

func (x *statsServiceStream) Close() error {
	return x.stream.Close()
}

func (x *statsServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *statsServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *statsServiceStream) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Stats service

type StatsHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Write(context.Context, *WriteRequest, *WriteResponse) error
	Stream(context.Context, *StreamRequest, Stats_StreamStream) error
}

func RegisterStatsHandler(s server.Server, hdlr StatsHandler, opts ...server.HandlerOption) error {
	type stats interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error
		Stream(ctx context.Context, stream server.Stream) error
	}
	type Stats struct {
		stats
	}
	h := &statsHandler{hdlr}
	return s.Handle(s.NewHandler(&Stats{h}, opts...))
}

type statsHandler struct {
	StatsHandler
}

func (h *statsHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.StatsHandler.Read(ctx, in, out)
}

func (h *statsHandler) Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error {
	return h.StatsHandler.Write(ctx, in, out)
}

func (h *statsHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.StatsHandler.Stream(ctx, m, &statsStreamStream{stream})
}

type Stats_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamResponse) error
}

type statsStreamStream struct {
	stream server.Stream
}

func (x *statsStreamStream) Close() error {
	return x.stream.Close()
}

func (x *statsStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *statsStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *statsStreamStream) Send(m *StreamResponse) error {
	return x.stream.Send(m)
}
