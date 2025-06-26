package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	log.Printf("Max function was invoked with a streaming request")

	max := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error while receiving request from Max: %v", err)
		}

		if req.Number > max {
			max = req.Number
			log.Printf("New max: %v", max)
		}

		if err := stream.Send(&pb.MaxResponse{Result: max}); err != nil {
			return fmt.Errorf("error while sending response to Max: %v", err)
		}
	}
}
