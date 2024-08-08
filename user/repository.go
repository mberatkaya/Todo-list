package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName string) *UserRepository {
	collection := client.Database(dbName).Collection("users")
	return &UserRepository{Collection: collection}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
	var user User
	err := r.Collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByNickname(ctx context.Context, nickname string) (*User, error) {
	var user User
	err := r.Collection.FindOne(ctx, bson.D{{"nickname", nickname}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id primitive.ObjectID, updateFields bson.D) error {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", updateFields}}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := r.Collection.DeleteOne(ctx, filter)
	return err
}
