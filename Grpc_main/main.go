package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	pb "HelloService" // Import generated code
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("Khong co luon")
	return &pb.HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func (s *server) SayGoodbye(ctx context.Context, req *pb.GoodbyeRequest) (*pb.GoodbyeResponse, error) {
	return &pb.GoodbyeResponse{Message: "Goodbye, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	test := server{}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &test)

	log.Println("Server is listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
