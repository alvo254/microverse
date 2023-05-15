package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alvo254/microverse/proto"
)

func callSayHello(client pb.GreetServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res,error := client.SayHello(ctx, &pb.NoParam{})
	if error != nil{
		log.Fatalf("Could not greet: %v", error)
	}
	log.Printf("%s", res.Message)
}

