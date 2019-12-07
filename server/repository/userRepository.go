package repository

import (
	"context"
	"fmt"
	model "go-grpc-sample/server/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Create(ctx context.Context, collection *mongo.Collection, user model.UserItem) (*model.UserItem, error) {
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Failed to insert data: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Failed to convert ObjectID"),
		)
	}
	user.ID = oid
	return &user, nil
}

func Read(ctx context.Context, collection *mongo.Collection, userId string) (*model.UserItem, error) {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("Failed to parse ObjectID"),
		)
	}

	// create an empty struct
	data := &model.UserItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound, fmt.Sprintf("Cannot find User ID: %v", err),
		)
	}
	return data, nil
}

func Update(ctx context.Context, collection *mongo.Collection, user *model.UserItem) (*model.UserItem, error) {

	filter := bson.M{"_id": user.ID}
	_, err := collection.ReplaceOne(ctx, filter, user)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Cannot update object %v", err),
		)
	}
	return user, nil
}

func Delete(ctx context.Context, collection *mongo.Collection, userId string) (*string, error) {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("Failed to parse ObjectID"),
		)
	}

	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot delete object: %v", err),
		)
	}
	return &userId, nil
}
