package main

import (
	pb "HelloService" // Import generated code
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type greeterService struct {
}

func (s *greeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("Fourth Here")
	return &pb.HelloResponse{Message: "Hello, MINHCV5666222 " + req.Name + "!"}, nil
}

func (s *greeterService) SayGoodbye(ctx context.Context, req *pb.GoodbyeRequest) (*pb.GoodbyeResponse, error) {
	return &pb.GoodbyeResponse{Message: "Goodbye, MINHCV5 " + req.Name + "!"}, nil
}

func makeSayHelloEndpoint(svc pb.HelloServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.HelloRequest)
		return svc.SayHello(ctx, req)
	}
}

func makeSayGoodbyeEndpoint(svc pb.HelloServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("There Here")
		req := request.(*pb.GoodbyeRequest)
		return svc.SayGoodbye(ctx, req)
	}
}

func decodeSayHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	fmt.Println("Second Here")
	req := grpcReq.(*pb.HelloRequest)
	return req, nil
}

func encodeSayHelloResponse(_ context.Context, response interface{}) (interface{}, error) {
	fmt.Println("Fifth Here")
	resp := response.(*pb.HelloResponse)
	return resp, nil
}

func decodeSayGoodbyeRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GoodbyeRequest)
	return req, nil
}

func encodeSayGoodbyeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GoodbyeResponse)
	return resp, nil
}

type greeterServiceServer struct {
	sayHello   kitgrpc.Handler
	sayGoodbye kitgrpc.Handler
}

func (s *greeterServiceServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("First Here")
	_, resp, err := s.sayHello.ServeGRPC(ctx, req)
	fmt.Println("Switch sit")
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HelloResponse), nil
}

func (s *greeterServiceServer) SayGoodbye(ctx context.Context, req *pb.GoodbyeRequest) (*pb.GoodbyeResponse, error) {
	_, resp, err := s.sayGoodbye.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GoodbyeResponse), nil
}
func main() {
	svc := &greeterService{}

	sayHelloHandler := kitgrpc.NewServer(
		makeSayHelloEndpoint(svc),
		decodeSayHelloRequest,
		encodeSayHelloResponse,
	)

	sayGoodbyeHandler := kitgrpc.NewServer(
		makeSayGoodbyeEndpoint(svc),
		decodeSayGoodbyeRequest,
		encodeSayGoodbyeResponse,
	)

	grpcServer := &greeterServiceServer{
		sayHello:   sayHelloHandler,
		sayGoodbye: sayGoodbyeHandler,
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	baseServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(baseServer, grpcServer)

	log.Println("Server listening on port 50051")
	if err := baseServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
