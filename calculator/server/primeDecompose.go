package main

import (
	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	number := in.Number
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{Result: divisor})
			number = number / divisor
		} else {
			divisor++
		}
	}
	return nil
}
