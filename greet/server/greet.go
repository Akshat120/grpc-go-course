package main

import (
	"context"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Result: "Hello " + req.Name,
	}, nil
}
