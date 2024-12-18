package todo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	Collection *mongo.Collection
}

type TodoRepository interface {
	GetAllTodos(ctx context.Context) ([]Todo, error)
	CreateTodo(ctx context.Context, todo *Todo) (*Todo, error)
	UpdateTodoCompletion(ctx context.Context, id primitive.ObjectID, completed bool) error
	DeleteTodo(ctx context.Context, id primitive.ObjectID) error
}

func NewTodoRepository(client *mongo.Client, dbName string) TodoRepository {
	collection := client.Database(dbName).Collection("todos")
	return &todoRepository{Collection: collection}
}

func (r *todoRepository) GetAllTodos(ctx context.Context) ([]Todo, error) {
	cursor, err := r.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []Todo
	if err := cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *todoRepository) CreateTodo(ctx context.Context, todo *Todo) (*Todo, error) {
	result, err := r.Collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	return todo, nil
}

func (r *todoRepository) UpdateTodoCompletion(ctx context.Context, id primitive.ObjectID, completed bool) error {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"completed", completed}}}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *todoRepository) DeleteTodo(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := r.Collection.DeleteOne(ctx, filter)
	return err
}
