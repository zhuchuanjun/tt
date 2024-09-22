// client.go
package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"grpc_example/pb"
	"log"
	"time"
)

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

	// 从etcd获取ServiceA的地址
	resp, err := etcdClient.Get(context.Background(), "/services/serviceA")
	if err != nil {
		log.Fatalf("Failed to get ServiceA address: %v", err)
	}
	if len(resp.Kvs) == 0 {
		log.Fatal("ServiceA address not found in etcd")
	}
	serviceAAddr := string(resp.Kvs[0].Value)

	// 连接到ServiceA
	conn, err := grpc.Dial(serviceAAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewServiceAClient(conn)

	// 调用GetUser
	r, err := client.GetUser(context.Background(), &pb.GetUserRequest{UserId: 1})
	if err != nil {
		log.Fatalf("Could not get user: %v", err)
	}
	fmt.Printf("User: %s (%s)\n", r.Name, r.Email)
}
