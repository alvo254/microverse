package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/alvo254/microverse/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Bidirectional streaming started")
	stream, error := client.SayHelloBidirectionalStreaming(context.Background())

	if error != nil{
		log.Fatalf("Could not send names: %v", error)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, error := stream.Recv()
			if error == io.EOF{
				break
			}
			if error != nil {
				log.Fatalf("Error while streaming: %v", error)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if error := stream.Send(req); error != nil{
			log.Fatalf("Error while sending %V", error)
		}

		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<- waitc
	log.Printf("Bidirectional streaming finished")
}