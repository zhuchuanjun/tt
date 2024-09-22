package main

import (
	"ServiceB/db"
	pb "ServiceB/user"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) GetUserDetails(ctx context.Context, req *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	user := db.GetById(req.UserID)
	return &pb.GetUserDetailsResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	log.Printf("server listening as %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error: %v", err)
	}
}
