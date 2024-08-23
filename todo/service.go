package todo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct {
	Repo interface{} // Repo alanı interface{} türünde
}

func NewTodoService(repo interface{}) *TodoService {
	return &TodoService{Repo: repo}
}

func (s *TodoService) GetAllTodos(ctx context.Context) ([]Todo, error) {
	repo := s.Repo.(*TodoRepository)
	return repo.GetAllTodos(ctx)
}

func (s *TodoService) CreateTodo(ctx context.Context, task string) (*Todo, error) {
	todo := &Todo{
		Task:      task,
		Completed: false,
	}
	repo := s.Repo.(*TodoRepository)
	return repo.CreateTodo(ctx, todo)
}

func (s *TodoService) UpdateTodoCompletion(ctx context.Context, id primitive.ObjectID, completed bool) error {
	repo := s.Repo.(*TodoRepository)
	return repo.UpdateTodoCompletion(ctx, id, completed)
}

func (s *TodoService) DeleteTodo(ctx context.Context, id primitive.ObjectID) error {
	repo := s.Repo.(*TodoRepository)
	return repo.DeleteTodo(ctx, id)
}
