// client.go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math"
	"os"
	"strconv"

	"google.golang.org/grpc"
	pb "normal-distribution-grpc/go-gen"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: client <x>")
		os.Exit(1)
	}

	x, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error parsing x:", err)
		os.Exit(1)
	}

	conn, err := grpc.Dial(
		"localhost:1234",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewNormalDistributionClient(conn)

	piResult, err := client.CalculatePi(context.Background(), &pb.PiRequest{})
	if err != nil {
		log.Fatalf("Error calling CalculatePi: %v", err)
	}

	expResult, err := client.CalculateExp(context.Background(), &pb.ExpRequest{X: x})
	if err != nil {
		log.Fatalf("Error calling CalculateExp: %v", err)
	}

	result := expResult.Value / math.Sqrt(2*piResult.Value)
	fmt.Printf("e^(-x^2/2) = %f\n", expResult.Value)
	fmt.Printf("pi = %f\n", piResult.Value)
	fmt.Printf("f(x) = %f\n", result)
}
