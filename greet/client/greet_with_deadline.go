package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
)

func doGreetWithDeadline(c pb.GreetServiceClient) error {
	req := &pb.GreetRequest{Name: "Akshat"}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		return fmt.Errorf("error while calling GreetWithDeadline: %v", err)
	}

	log.Printf("Response from GreetWithDeadline: %v\n", res.Result)

	return nil
}
