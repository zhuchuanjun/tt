package logic

import (
	"bookservice/internal/model"
	"bookservice/internal/svc"
	"bookservice/pb/book"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditLogic) Edit(in *book.EditRequest) (*book.EditResponse, error) {
	result := l.svcCtx.DB.Where("id = ?", in.Id).UpdateColumns(model.Book{
		Title:       in.Title,
		Author:      in.Author,
		PublishDate: in.PublishDate,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return &book.EditResponse{}, nil
}
