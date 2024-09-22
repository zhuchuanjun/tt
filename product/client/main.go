package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc"
	"time"
	"tt/product/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetProduct(ctx, &proto.GetProductRequest{Id: 1})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	g.Dump(r.String())
}
