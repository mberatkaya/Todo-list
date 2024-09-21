package todo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService interface {
	GetAllTodosService(ctx context.Context) ([]Todo, error)
	CreateTodoService(ctx context.Context, task string) (*Todo, error)
	UpdateTodoCompletionService(ctx context.Context, id primitive.ObjectID, completed bool) error
	DeleteTodoService(ctx context.Context, id primitive.ObjectID) error
}

type todoService struct {
	Repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoService {
	return &todoService{Repo: repo}
}

func (s *todoService) GetAllTodosService(ctx context.Context) ([]Todo, error) {
	return s.Repo.GetAllTodos(ctx)
}

func (s *todoService) CreateTodoService(ctx context.Context, task string) (*Todo, error) {
	todo := &Todo{
		Task:      task,
		Completed: false,
	}
	return s.Repo.CreateTodo(ctx, todo)
}

func (s *todoService) UpdateTodoCompletionService(ctx context.Context, id primitive.ObjectID, completed bool) error {
	return s.Repo.UpdateTodoCompletion(ctx, id, completed)
}

func (s *todoService) DeleteTodoService(ctx context.Context, id primitive.ObjectID) error {
	return s.Repo.DeleteTodo(ctx, id)
}
