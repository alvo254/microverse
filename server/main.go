package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/alvo254/microverse/proto"
)

const (
	port = ":8080"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main() {
	listen, error := net.Listen("tcp", port)
	if error != nil{
		log.Fatalf("Failed to start the server %v", error)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", listen.Addr())
	if error := grpcServer.Serve(listen); error != nil {
		log.Fatalf("Failed to start %v", error)
	}
}