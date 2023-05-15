package main

import (
	"io"
	"log"

	pb "github.com/alvo254/microverse/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error{
	for {
		req, error := stream.Recv()
		if error == io.EOF{
			return nil
		}
		if error != nil {
			return error
		}
		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if error := stream.Send(res); error != nil {
			return error
		}

	}
}