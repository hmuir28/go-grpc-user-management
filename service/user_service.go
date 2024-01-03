package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "github.com/hmuir28/go-grpc-user-management/proto"
)

type messageUnit struct {
	ClientName        string
	MessageBody       pb.CreateUserRequest
	MessageUniqueCode int
	ClientUniqueCode  int
}

type messageHandle struct {
	MQue []messageUnit
	mu   sync.Mutex
}

var messageHandleObject = messageHandle{}

type MyUserServiceServer struct {
	pb.UnimplementedUserServiceUnaryServer
}

func (userServiceClient MyUserServiceServer) UnaryCreateUser(context context.Context, createUserRequest *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	return &pb.CreateUserResponse{
		Url: "http://localhost:8089/1",
	}, nil
}

func (userServiceClient MyUserServiceServer) BidirectionalCreateUser(userServiceServer pb.UserServiceBidrectional_BidirectionalCreateUserServer) error {

	fmt.Println("__________===wpdionwdnw")

	clientUnique := rand.Intn(1e8)
	errorCh := make(chan error)

	go receiveFromStream(userServiceServer, clientUnique, errorCh)

	go sendToStream(userServiceServer, clientUnique, errorCh)

	return <-errorCh
}

func receiveFromStream(userServiceServer_ pb.UserServiceBidrectional_BidirectionalCreateUserServer, clientUniqueCode_ int, errorCh_ chan error) {

	for {

		msg, err := userServiceServer_.Recv()

		if err != nil {
			log.Printf("Error in receiving message from client :: %v", err)
			errorCh_ <- err
		} else {

			messageHandleObject.mu.Lock()

			user := pb.CreateUserRequest{
				Username:  msg.Username,
				FirstName: msg.FirstName,
				LastName:  msg.LastName,
				Email:     msg.Email,
				Password:  msg.Password,
			}

			messageHandleObject.MQue = append(messageHandleObject.MQue, messageUnit{
				ClientName:        msg.FirstName + " " + msg.LastName,
				MessageBody:       user,
				ClientUniqueCode:  clientUniqueCode_,
				MessageUniqueCode: rand.Intn(1e8),
			})

			messageHandleObject.mu.Unlock()

			log.Printf("%v", messageHandleObject.MQue[len(messageHandleObject.MQue)-1])
		}
	}

}

func sendToStream(userServiceServer_ pb.UserServiceBidrectional_BidirectionalCreateUserServer, clientUniqueCode_ int, errorCh_ chan error) {

	for {

		for {

			time.Sleep(500 * time.Millisecond)

			messageHandleObject.mu.Lock()

			if len(messageHandleObject.MQue) == 0 {

				messageHandleObject.mu.Unlock()
				break
			}

			senderClientName := messageHandleObject.MQue[0].ClientName
			senderUniqueCode := messageHandleObject.MQue[0].ClientUniqueCode
			senderBody := messageHandleObject.MQue[0].MessageBody

			log.Println(senderBody)

			messageHandleObject.mu.Unlock()

			if senderUniqueCode != clientUniqueCode_ {

				url := string(senderUniqueCode) + " - " + senderClientName

				err := userServiceServer_.Send(&pb.CreateUserResponse{
					Url: url,
				})

				if err != nil {
					errorCh_ <- err
				}

				messageHandleObject.mu.Lock()

				if len(messageHandleObject.MQue) > 1 {
					messageHandleObject.MQue = messageHandleObject.MQue[1:]
				} else {
					messageHandleObject.MQue = []messageUnit{}
				}

				messageHandleObject.mu.Unlock()
			}
		}

		time.Sleep(100 * time.Millisecond)
	}

}
