package main

import (
	"context"
	"fmt"
	userpb "go-grpc-sample/proto"
	"log"

	"google.golang.org/grpc"
)

func CreateUser(c userpb.UserServiceClient) string {
	user := &userpb.User{
		Name:  "Tom",
		Email: "test@example.com",
	}
	createUserRes, err := c.CreateUser(context.Background(), &userpb.CreateUserRequest{User: user})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("User has been created: %v", createUserRes)
	userID := createUserRes.GetUser().GetId()
	return userID
}

func ReadUser(c userpb.UserServiceClient, userID string) {
	_, err2 := c.ReadUser(context.Background(), &userpb.ReadUserRequest{UserId: userID})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v \n", err2)
	}

	readUserReq := &userpb.ReadUserRequest{UserId: userID}
	readUserRes, readUserErr := c.ReadUser(context.Background(), readUserReq)
	if readUserErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readUserErr)
	}
	fmt.Printf("User was read: %v", readUserRes)
}

func DeleteUser(c userpb.UserServiceClient, userID string) {
	deleteUserRes, err2 := c.DeleteUser(context.Background(), &userpb.DeleteUserRequest{UserId: userID})
	if err2 != nil {
		fmt.Printf("Error happened while deleting: %v \n", err2)
	}
	fmt.Printf("User was deleted: %v", deleteUserRes)
}

func main() {
	fmt.Println("User Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := userpb.NewUserServiceClient(cc)

	// create User
	userID := CreateUser(c)
	// read User
	ReadUser(c, userID)
	// update User
	// TODO
	// delete User
	DeleteUser(c, userID)
}
