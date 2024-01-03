package main

import (
	"log"
	"net"

	pb "github.com/hmuir28/go-grpc-user-management/proto"
	"github.com/hmuir28/go-grpc-user-management/service"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8089")

	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	grpcServer := grpc.NewServer()

	newServer := &service.MyUserServiceServer{}

	pb.RegisterUserServiceUnaryServer(grpcServer, newServer)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Impossible to serve: %s", err)
	}
}
