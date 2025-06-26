package main

import (
	"fmt"
	"io"

	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {

	var result, count float64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if count == 0 {
				return fmt.Errorf("no numbers received")
			}
			return stream.SendAndClose(&pb.AverageResponse{Result: result / count})
		}
		if err != nil {
			return fmt.Errorf("error while receiving request from Average: %v", err)
		}
		result += float64(req.Number)
		count++
	}
}
