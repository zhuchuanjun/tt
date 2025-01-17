// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: greet.proto

package greetclient

import (
	"context"

	"tt/hello-grpc/greet/types/greet"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	HelloRequest  = greet.HelloRequest
	HelloResponse = greet.HelloResponse

	Greet interface {
		SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	}

	defaultGreet struct {
		cli zrpc.Client
	}
)

func NewGreet(cli zrpc.Client) Greet {
	return &defaultGreet{
		cli: cli,
	}
}

func (m *defaultGreet) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	client := greet.NewGreetClient(m.cli.Conn())
	return client.SayHello(ctx, in, opts...)
}
