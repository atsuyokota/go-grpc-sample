package main

import (
	"context"
	"fmt"
	userpb "go-grpc-sample/proto"
	config "go-grpc-sample/server/config"
	model "go-grpc-sample/server/model"
	repository "go-grpc-sample/server/repository"
	"go-grpc-sample/server/utils"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
}

func (*server) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	fmt.Println("Create user request")
	user := req.GetUser()

	data := model.UserItem{
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}

	res, _ := repository.Create(context.Background(), collection, data)
	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    res.ID.Hex(),
			Name:  res.Name,
			Email: res.Email,
		},
	}, nil
}

func (*server) ReadUser(ctx context.Context, req *userpb.ReadUserRequest) (*userpb.ReadUserResponse, error) {
	fmt.Println("Read user request")
	userID := req.GetUserId()
	data, _ := repository.Read(context.Background(), collection, userID)
	return &userpb.ReadUserResponse{
		User: dataToUserPb(data),
	}, nil
}

func dataToUserPb(data *model.UserItem) *userpb.User {
	return &userpb.User{
		Id:    data.ID.Hex(),
		Name:  data.Name,
		Email: data.Email,
	}
}

func (*server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	fmt.Println("Update user request")
	user := req.GetUser()
	data, _ := repository.Read(context.Background(), collection, user.GetId())
	data.Name = user.GetName()
	data.Email = user.GetEmail()

	res, _ := repository.Update(context.Background(), collection, data)
	return &userpb.UpdateUserResponse{
		User: dataToUserPb(res),
	}, nil
}

func (*server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	fmt.Println("Delete user request")
	userID := req.GetUserId()
	res, _ := repository.Delete(context.Background(), collection, userID)
	return &userpb.DeleteUserResponse{UserId: *res}, nil
}

// Connect to MongoDB
func connection() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

func serverRun() {
	utils.LogSettings(config.Config.LogFile)
	client, _ := connection()
	collection = client.Database(config.Config.MongoDB).Collection("user")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	userpb.RegisterUserServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Signal is received
	<-ch
	s.Stop()
	lis.Close()
	client.Disconnect(context.TODO())
	fmt.Println("Server stopped.")
}

func main() {
	serverRun()
}
