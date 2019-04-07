package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dreamer-nitj/grpc-go-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am calculator client!")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial connection to grpc server: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	request := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 15,
	}
	resp, err := c.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error in invoking Sum method from client: %v", err)
	}

	log.Printf("Response from greet: %v", resp.SumResult)
}
