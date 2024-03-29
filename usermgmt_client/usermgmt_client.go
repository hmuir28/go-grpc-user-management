package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hmuir/go-user-management/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var new_users = make(map[string]int32)

	new_users["Alice"] = 32
	new_users["Bob"] = 29
	new_users["Harry"] = 12

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{
			Name: name,
			Age:  age,
		})

		if err != nil {
			log.Fatalf("could not create user: %v", err)
			return
		}

		log.Printf(`
			User Details:
				NAME: %s
				AGE: %d
				ID: %d
		`, r.GetName(), r.GetAge(), r.GetId())
	}

	params := &pb.GetUsersParams{}

	r, err := c.GetUsers(ctx, params)

	if err != nil {
		log.Fatalf("could not retrieve users: %v", err)
	}

	log.Print("\nUSER LIST: \n")
	log.Printf("r.GetUsers(): %v\n", r.GetUsers())
}
