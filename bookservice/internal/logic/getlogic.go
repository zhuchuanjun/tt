package logic

import (
	"bookservice/internal/model"
	"context"

	"bookservice/internal/svc"
	"bookservice/pb/book"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *book.GetRequest) (*book.GetResponse, error) {
	bo := &model.Book{}
	result := l.svcCtx.DB.First(bo, in.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book.GetResponse{
		Id:          bo.Id,
		Title:       bo.Title,
		Author:      bo.Author,
		PublishDate: bo.PublishDate,
	}, nil
}
