package main

import (
	"context"

	pb "github.com/pfirulo2022/go-grpc-demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello ",
	}, nil
}
