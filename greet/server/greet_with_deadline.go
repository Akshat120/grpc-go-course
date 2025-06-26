package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline function was invoked with %v", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("The client cancelled the request!")
			return nil, status.Error(codes.Canceled, "The client cancelled the request!")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{Result: "Hello " + req.Name}, nil
}
