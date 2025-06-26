package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) error {
	stream, err := c.Average(context.Background())
	if err != nil {
		return fmt.Errorf("error while calling Average: %v", err)
	}

	numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		log.Printf("Sending request: %v", number)
		stream.Send(&pb.AverageRequest{Number: number})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("error while receiving response from Average: %v", err)
	}

	log.Printf("Response from Average: %v\n", res.Result)

	return nil
}
