package main

import (
	"context"

	pb "github.com/Akshat120/grpc-go-course/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Result: in.A + in.B,
	}, nil
}
