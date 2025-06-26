package main

import (
	"fmt"
	"io"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	var result string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: result})
		}
		if err != nil {
			return fmt.Errorf("error while receiving request from LongGreet: %v", err)
		}
		result += fmt.Sprintf("Hello %s!\n", req.Name)
	}
}
