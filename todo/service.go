package todo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService interface {
	GetAllTodos(ctx context.Context) ([]Todo, error)
	CreateTodo(ctx context.Context, task string) (*Todo, error)
	UpdateTodoCompletion(ctx context.Context, id primitive.ObjectID, completed bool) error
	DeleteTodo(ctx context.Context, id primitive.ObjectID) error
}

type todoService struct {
	Repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoService {
	return &todoService{Repo: repo}
}

func (s *todoService) GetAllTodos(ctx context.Context) ([]Todo, error) {
	return s.Repo.GetAllTodos(ctx)
}

func (s *todoService) CreateTodo(ctx context.Context, task string) (*Todo, error) {
	todo := &Todo{
		Task:      task,
		Completed: false,
	}
	return s.Repo.CreateTodo(ctx, todo)
}

func (s *todoService) UpdateTodoCompletion(ctx context.Context, id primitive.ObjectID, completed bool) error {
	return s.Repo.UpdateTodoCompletion(ctx, id, completed)
}

func (s *todoService) DeleteTodo(ctx context.Context, id primitive.ObjectID) error {
	return s.Repo.DeleteTodo(ctx, id)
}
