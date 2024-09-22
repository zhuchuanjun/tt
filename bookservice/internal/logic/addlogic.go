package logic

import (
	"bookservice/internal/model"
	"bookservice/internal/svc"
	"bookservice/pb/book"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *book.AddRequest) (*book.AddResponse, error) {
	bo := &model.Book{
		Title:       in.Title,
		Author:      in.Author,
		PublishDate: in.PublishDate,
	}
	result := l.svcCtx.DB.Create(bo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book.AddResponse{
		Id: bo.Id,
	}, nil
}
