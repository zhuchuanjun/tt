package logic

import (
	"context"

	"tt/hello-grpc/greet/internal/svc"
	"tt/hello-grpc/greet/types/greet"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *greet.HelloRequest) (*greet.HelloResponse, error) {
	return &greet.HelloResponse{
		Message: "Hello, " + in.Name,
	}, nil
}
