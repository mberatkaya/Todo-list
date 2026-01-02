package todo

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockTodoService struct {
	mock.Mock
}

func (m *MockTodoService) GetAllTodosService(ctx context.Context) ([]Todo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Todo), args.Error(1)
}

func (m *MockTodoService) CreateTodoService(ctx context.Context, task string) (*Todo, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(*Todo), args.Error(1)
}

func (m *MockTodoService) UpdateTodoCompletionService(ctx context.Context, id primitive.ObjectID, completed bool) error {
	args := m.Called(ctx, id, completed)
	return args.Error(0)
}

func (m *MockTodoService) DeleteTodoService(ctx context.Context, id primitive.ObjectID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
