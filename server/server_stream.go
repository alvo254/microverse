package main

import (
	"log"
	"time"

	pb "github.com/alvo254/microverse/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error{
	log.Printf("Got request with names: %v", req.Names)
	for _, name := range req.Names{
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		//Send response as a stream
		if error := stream.Send(res); error != nil{
			return error
		}
		//delay
		time.Sleep(2 * time.Second)
	}
	return nil
}