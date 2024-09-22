package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"tt/product/proto"
)

type server struct {
	proto.UnsafeProductServiceServer
}

func (s *server) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	return &proto.GetProductResponse{
		Id:    req.GetId(),
		Name:  "lalala",
		Price: 10.01,
	}, nil
}

func main() {
	// 启动 gRPC 服务器
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		proto.RegisterProductServiceServer(grpcServer, &server{})
		reflection.Register(grpcServer)

		log.Printf("gRPC server listening on %v", listener.Addr())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// 启动 gRPC-Gateway 反向代理
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterProductServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}

	log.Printf("HTTP server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
