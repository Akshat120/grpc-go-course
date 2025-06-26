package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) error {

	reqs := []*pb.GreetRequest{
		{Name: "Akshat"},
		{Name: "John"},
		{Name: "Jane"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		return fmt.Errorf("error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("error while receiving response from LongGreet: %v", err)
	}

	log.Printf("Response from LongGreet: %v\n", res.Result)

	return nil
}
