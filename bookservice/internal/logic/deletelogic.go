package logic

import (
	"context"

	"bookservice/internal/svc"
	"bookservice/pb/book"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *book.DeleteRequest) (*book.DeleteResponse, error) {
	result := l.svcCtx.DB.Delete(in.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book.DeleteResponse{}, nil
}
