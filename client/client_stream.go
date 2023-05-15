package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alvo254/microverse/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Client streaming started")
	stream, error := client.SayHelloClientStreaming(context.Background())
	if error != nil {
		log.Fatalf("Could not send names: %v", error)
	}

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if error := stream.Send(req); error != nil{
			log.Fatalf("Error while sending %v", error)
		}
		log.Printf("Sent the request with names: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, error := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if error != nil{
		log.Fatalf("Error while receiving %v", error)
	}
	log.Printf("%v", res.Messages)
}