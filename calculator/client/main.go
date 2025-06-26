package main

import (
	"log"

	pb "github.com/Akshat120/grpc-go-course/calculator/proto"
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

	client := pb.NewCalculatorServiceClient(conn)

	// doSum(client)
	// doPrime(client)
	// doAverage(client)
	// doMax(client)
	if err := doSqrt(client); err != nil {
		log.Fatalf("Error while calling Sqrt: %v", err)
	}
}
