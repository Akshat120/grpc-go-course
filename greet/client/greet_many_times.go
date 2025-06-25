package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	req := &pb.GreetRequest{
		Name: "Akshat",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("GreetManyTimes: %s", msg.Result)
	}
}
