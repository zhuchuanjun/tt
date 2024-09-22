// service_b.go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/lib/pq"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"grpc_example/pb"
)

type serverB struct {
	pb.UnimplementedServiceBServer
	db *sql.DB
}

func (s *serverB) GetUserDetails(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var name, email string
	err := s.db.QueryRow("SELECT name, email FROM users WHERE id = $1", req.UserId).Scan(&name, &email)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{UserId: req.UserId, Name: name, Email: email}, nil
}

func main() {
	// 连接到PostgreSQL
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=userdb sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 连接到etcd
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	// 启动ServiceB的gRPC服务器
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServiceBServer(s, &serverB{db: db})

	// 注册ServiceB到etcd
	_, err = etcdClient.Put(context.Background(), "/services/serviceB", "localhost:50052")
	if err != nil {
		log.Fatalf("Failed to register ServiceB: %v", err)
	}

	fmt.Println("service_b register")

	fmt.Println("ServiceB is running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
