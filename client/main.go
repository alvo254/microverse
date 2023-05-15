package main

import (
	"log"
	pb "github.com/alvo254/microverse/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)



func main() {
	connection, error := grpc.Dial("localhost" + port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if error != nil{
		log.Fatalf("Did not connect: %v", error)
	}
	defer connection.Close()

	// client := pb.NewGreetServiceClient(connection)

	// names := &pb.NamesList{
	// 	Names: []string{"Alvin", "Ryan", "Barry"},

	// }

	// callSayhello(client)

}