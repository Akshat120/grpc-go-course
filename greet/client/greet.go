package main

import (
	"context"
	"log"

	pb "github.com/Akshat120/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {

	req := &pb.GreetRequest{
		Name: "Akshat",
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)

}
