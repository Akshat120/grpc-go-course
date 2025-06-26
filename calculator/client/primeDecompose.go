package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) error {
	req := &pb.PrimeRequest{
		Number: 15485863,
	}

	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		return fmt.Errorf("error while calling Prime: %v", err)
	}

	log.Printf("Prime decomposition of %d:", req.Number)

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error while receiving response from Prime: %v", err)
		}
		log.Printf("Response from Prime: %v", res.Result)
	}

	return nil
}
