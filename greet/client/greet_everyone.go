package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) error {
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		return fmt.Errorf("error while calling GreetEveryone: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{Name: "Akshat"},
		{Name: "John"},
		{Name: "Jane"},
		{Name: "Jim"},
		{Name: "Jill"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving response from GreetEveryone: %v", err)
			}
			log.Printf("Response from GreetEveryone: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc

	return nil
}
