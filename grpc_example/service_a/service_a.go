// service_a.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"grpc_example/pb"
)

type serverA struct {
	pb.UnimplementedServiceAServer
	serviceBClient pb.ServiceBClient
}

func (s *serverA) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.serviceBClient.GetUserDetails(ctx, req)
}

func main() {
	// 连接到etcd
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	// 从etcd获取ServiceB的地址
	resp, err := etcdClient.Get(context.Background(), "/services/serviceB")
	if err != nil {
		log.Fatalf("Failed to get ServiceB address: %v", err)
	}
	if len(resp.Kvs) == 0 {
		log.Fatal("ServiceB address not found in etcd")
	}
	serviceBAddr := string(resp.Kvs[0].Value)

	// 连接到ServiceB
	conn, err := grpc.Dial(serviceBAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to ServiceB: %v", err)
	}
	defer conn.Close()
	serviceBClient := pb.NewServiceBClient(conn)

	// 启动ServiceA的gRPC服务器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServiceAServer(s, &serverA{serviceBClient: serviceBClient})

	// 注册ServiceA到etcd
	_, err = etcdClient.Put(context.Background(), "/services/serviceA", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to register ServiceA: %v", err)
	}

	fmt.Println("service_a register")

	// 定义要写入文件的内容
	content := "Hello, World!\nThis is a text file written in Go."

	// 打开一个文件，如果文件不存在则创建，如果存在则清空内容
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// 将内容写入文件
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("ServiceA is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
