package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) error {
	stream, err := c.Max(context.Background())
	if err != nil {
		return fmt.Errorf("error while calling Max: %v", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
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
				log.Fatalf("error while receiving response from Max: %v", err)
			}
			log.Printf("Response from Max: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc

	return nil
}
