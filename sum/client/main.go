package main

import (
	"context"
	"log"

	pb "github.com/Akshat120/grpc-go-course/sum/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSumServiceClient(conn)

	resp, err := client.Sum(context.Background(), &pb.SumRequest{
		A: 343,
		B: 212,
	})
	if err != nil {
		log.Fatalf("Failed to sum: %v", err)
	}

	log.Printf("Sum: %d", resp.Result)
}
