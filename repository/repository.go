package repository

import (
	model "TODOproject/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	Collection *mongo.Collection
}

func NewTodoRepository(client *mongo.Client, dbName, collectionName string) *TodoRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &TodoRepository{Collection: collection}
}

func (r *TodoRepository) GetAllTodos(ctx context.Context) ([]model.Todo, error) {
	cursor, err := r.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []model.Todo
	if err := cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	result, err := r.Collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	return todo, nil
}

func (r *TodoRepository) UpdateTodoCompletion(ctx context.Context, id primitive.ObjectID, completed bool) error {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"completed", completed}}}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := r.Collection.DeleteOne(ctx, filter)
	return err
}
