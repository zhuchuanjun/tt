package main

import (
	pb "ServiceA/user"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	conn, err := grpc.Dial("service_b:50052", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	res, err := client.GetUserDetails(ctx, &pb.GetUserDetailsRequest{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	return &pb.GetUserDetailsResponse{
		Id:    res.Id,
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error: %v", err)
	}
}
