package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) error {

	req := &pb.SumRequest{
		A: 10,
		B: 20,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error while calling Sum: %v", err)
	}

	log.Printf("Response from Sum: %v", res.Result)

	return nil
}
