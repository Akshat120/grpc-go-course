package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
)

func doSqrt(c pb.CalculatorServiceClient) error {

	req := &pb.SqrtRequest{Number: -1}

	res, err := c.Sqrt(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error while calling Sqrt: %v", err)
	}

	log.Printf("Response from Sqrt: %v\n", res.Result)

	return nil
}
