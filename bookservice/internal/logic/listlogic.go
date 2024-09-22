package logic

import (
	"bookservice/internal/model"
	"context"

	"bookservice/internal/svc"
	"bookservice/pb/book"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *book.ListRequest) (*book.ListResponse, error) {
	var bos []*model.Book
	result := l.svcCtx.DB.Find(&bos)
	if result.Error != nil {
		return nil, result.Error
	}
	resp := &book.ListResponse{
		Items: make([]*book.ListItem, 0),
	}
	for _, bo := range bos {
		resp.Items = append(resp.Items, &book.ListItem{
			Id:          bo.Id,
			Title:       bo.Title,
			Author:      bo.Author,
			PublishDate: bo.PublishDate,
		})
	}
	return resp, nil
}
