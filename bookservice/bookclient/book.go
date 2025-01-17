// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: book.proto

package bookclient

import (
	"context"

	"bookservice/pb/book"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddRequest     = book.AddRequest
	AddResponse    = book.AddResponse
	DeleteRequest  = book.DeleteRequest
	DeleteResponse = book.DeleteResponse
	EditRequest    = book.EditRequest
	EditResponse   = book.EditResponse
	GetRequest     = book.GetRequest
	GetResponse    = book.GetResponse
	ListItem       = book.ListItem
	ListRequest    = book.ListRequest
	ListResponse   = book.ListResponse

	Book interface {
		Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
		List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
		Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
		Edit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error)
		Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	}

	defaultBook struct {
		cli zrpc.Client
	}
)

func NewBook(cli zrpc.Client) Book {
	return &defaultBook{
		cli: cli,
	}
}

func (m *defaultBook) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.Add(ctx, in, opts...)
}

func (m *defaultBook) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}

func (m *defaultBook) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.Get(ctx, in, opts...)
}

func (m *defaultBook) Edit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.Edit(ctx, in, opts...)
}

func (m *defaultBook) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	client := book.NewBookClient(m.cli.Conn())
	return client.Delete(ctx, in, opts...)
}
