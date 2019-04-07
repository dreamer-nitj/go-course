package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dreamer-nitj/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am a client!")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial connection to grpc server: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ankit",
			LastName:  "Kumar",
		},
	}
	resp, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Error in invoking Greet method from client: %v", err)
	}

	log.Printf("Response from greet: %v", resp.Result)
}
