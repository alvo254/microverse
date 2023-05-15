package main

import (
	"context"
	"io"
	"log"

	pb "github.com/alvo254/microverse/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Streaming started")
	stream, error := client.SayHelloServerStreaming(context.Background(), names)
	if error != nil {
		log.Fatalf("Could not send names: %v", error)
	}

	for {
		message, error := stream.Recv()
		//Check for end of file
		if error == io.EOF{
			break
		}
		if error != nil{
			log.Fatalf("Error while streaming %v", error)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}

