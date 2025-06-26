package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetEveryone(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {

	log.Printf("GreetEveryone function was invoked with a streaming request")

	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error while receiving request from GreetEveryone: %v", err)
		}

		if err := stream.Send(&pb.GreetResponse{Result: fmt.Sprintf("Hello %s!", req.Name)}); err != nil {
			return fmt.Errorf("error while sending response to GreetEveryone: %v", err)
		}
	}

}
