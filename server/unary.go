package main

import (
	"context"
	pb "github.com/alvo254/microverse/proto"

)


//Unary api
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "Hello ",
	}, nil
}