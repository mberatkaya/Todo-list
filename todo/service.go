package todo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct {
	Repo *TodoRepository
}

func NewTodoService(repo *TodoRepository) *TodoService {
	return &TodoService{Repo: repo}
}

func (s *TodoService) GetAllTodos(ctx context.Context) ([]Todo, error) {
	return s.Repo.GetAllTodos(ctx)
}

func (s *TodoService) CreateTodo(ctx context.Context, task string) (*Todo, error) {
	todo := &Todo{
		Task:      task,
		Completed: false,
	}
	return s.Repo.CreateTodo(ctx, todo)
}

func (s *TodoService) UpdateTodoCompletion(ctx context.Context, id primitive.ObjectID, completed bool) error {
	return s.Repo.UpdateTodoCompletion(ctx, id, completed)
}

func (s *TodoService) DeleteTodo(ctx context.Context, id primitive.ObjectID) error {
	return s.Repo.DeleteTodo(ctx, id)
}
