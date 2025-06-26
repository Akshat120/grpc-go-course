package main

import (
	"context"
	"log"
	"math"

	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt function was invoked with %v", req)
	number := req.Number
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Received a negative number: %v", number)
	}
	return &pb.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
