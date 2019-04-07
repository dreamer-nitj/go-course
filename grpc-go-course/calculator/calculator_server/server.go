package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dreamer-nitj/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse,
	error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	sum := firstNumber + secondNumber
	result := &calculatorpb.SumResponse{
		SumResult: sum,
	}

	return result, nil
}

func main() {
	fmt.Println("Starting calculator gRPC server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to run grpc server: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
