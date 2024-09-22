package main

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"log"

	"myproject/user/rpc/userclient"
)

func runUserClient() error {
	// 使用 zrpc.NewClient 创建客户端
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Target: "localhost:8080",
	})
	if err != nil {
		return err
	}

	// 使用 zrpc.Client 创建 userclient
	userClient := userclient.NewUser(client)

	// 现在你可以使用 userClient 来调用 RPC 方法
	createResp, err := userClient.CreateUser(context.Background(), &userclient.CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	})
	if err != nil {
		return err
	}
	log.Printf("Created user with ID: %d", createResp.Id)

	// 获取用户
	getResp, err := userClient.GetUser(context.Background(), &userclient.GetUserRequest{
		Id: createResp.Id,
	})
	if err != nil {
		return err
	}
	log.Printf("Got user: %v", getResp)

	return nil
}
