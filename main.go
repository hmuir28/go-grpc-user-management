package main

import (
	"log"

	pb "github.com/hmuir28/go-grpc-user-management/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	createUserRequest := &pb.CreateUserRequest{
		FirstName: "Harry",
		LastName:  "Muir",
		Username:  "hmuir@gmail.com",
		Password:  "12345",
		Email:     "hmuir@gmail.com",
	}

	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	// callSayHelloClientStreaming(client, names)
	us.callCreateUser(client, createUserRequest)
}
