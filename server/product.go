package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"net"
	"time"
	"tt/product"
)

type Server struct {
	product.UnimplementedProductServiceServer
}

func (s *Server) GetProduct(ctx context.Context, in *product.GetProductRequest) (*product.GetProductResponse, error) {
	return &product.GetProductResponse{
		Id:    in.Id,
		Name:  "aa",
		Price: 1.2,
	}, nil
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	serviceKey := "/product"
	serviceValue := "localhost:1212"

	_, err = cli.Put(context.Background(), serviceKey, serviceValue)
	if err != nil {
		panic(err)
	}

	g := grpc.NewServer()
	product.RegisterProductServiceServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}
	g.Serve(lis)
}
