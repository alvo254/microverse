// Client stream to server and server sends a response
package main

import (
	"log"
	"io"
	pb "github.com/alvo254/microverse/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error{
	var messages []string
	for {
		req, error := stream.Recv()
		if error == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if error != nil {
			return error
		}
		log.Printf("Got requests with name: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}