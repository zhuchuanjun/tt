// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: order.proto

package server

import (
	"context"

	"tt/myproject/order/rpc/internal/logic"
	"tt/myproject/order/rpc/internal/svc"
	"tt/myproject/order/rpc/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) CreateOrder(ctx context.Context, in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

func (s *OrderServer) GetOrder(ctx context.Context, in *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	l := logic.NewGetOrderLogic(ctx, s.svcCtx)
	return l.GetOrder(in)
}