package service

import (
	"context"

	pb "github.com/hmuir28/go-grpc-user-management/proto"
)

type MyUserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (userServiceClient MyUserServiceServer) CreateUser(context context.Context, createUserRequest *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	return &pb.CreateUserResponse{
		Url: "http://localhost:8089/1",
	}, nil
}
