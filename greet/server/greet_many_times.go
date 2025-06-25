package main

import (
	"fmt"
	"time"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	for i := range 10 {
		stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s, number %d", in.Name, i),
		})
		time.Sleep(1 * time.Second)
	}
	return nil
}
