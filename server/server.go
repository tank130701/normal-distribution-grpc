// server.go
package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	pb "normal-distribution-grpc/go-gen"
)

type server struct {
	pb.UnimplementedNormalDistributionServer
}

func (s *server) CalculatePi(ctx context.Context, req *pb.PiRequest) (*pb.PiResponse, error) {
	log.Println("Received CalculatePi request")
	sum := 0.0
	for i := 1; i < 1000000; i++ {
		sum += 1.0 / (float64(i) * float64(i))
	}
	result := pb.PiResponse{Value: math.Sqrt(sum * 6.0)}
	return &result, nil
}

func (s *server) CalculateExp(ctx context.Context, req *pb.ExpRequest) (*pb.ExpResponse, error) {
	log.Printf("Received CalculateExp request with x=%f\n", req.X)
	result := pb.ExpResponse{Value: math.Exp(-req.X * req.X / 2.0)}
	return &result, nil
}

func main() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterNormalDistributionServer(s, &server{})

	fmt.Println("Server is listening on port 1234...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
