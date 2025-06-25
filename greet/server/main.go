package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.UnimplementedGreetServiceServer
	GreetRequest pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("server listening at %s", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
