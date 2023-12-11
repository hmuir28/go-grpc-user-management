package restuser

import (
	"context"
	"io"
	"log"

	pb "github.com/hmuir28/go-grpc-user-management/proto"
)

func callCreateUser(client pb.UserServiceClient, createUserRequest *pb.CreateUserRequest) {
	log.Printf("User creation flow has started...")

	stream, err := client.CreateUser(context.Background())

	if err != nil {
		log.Fatalf("could not user info : %v", createUserRequest)
	}

	waitChannel := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}

			log.Println(message)
		}

		close(waitChannel)
	}()

	log.Println(createUserRequest)

	stream.CloseSend()

	<-waitChannel

	log.Printf("streaming finished...")
}
