package main

import (
	pb "client/user"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

type server struct {
	pb.UnimplementedUserServer
}

func main() {
	// 连接到 service_a
	conn, err := grpc.Dial("service_a:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	// 设置一个超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用 GetUser 方法
	req := &pb.GetUserDetailsRequest{UserID: 1}
	res, err := client.GetUserDetails(ctx, req)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	fmt.Printf("User: %v\n", res)
}
