package todo

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) GetAllTodos(ctx context.Context) ([]Todo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Todo), args.Error(1)
}

func (m *MockTodoRepository) CreateTodo(ctx context.Context, todo2 *Todo) (*Todo, error) {
	args := m.Called(ctx, todo2)
	return args.Get(0).(*Todo), args.Error(1)
}

func (m *MockTodoRepository) UpdateTodoCompletion(ctx context.Context, id primitive.ObjectID, completed bool) error {
	args := m.Called(ctx, id, completed)
	return args.Error(0)
}

func (m *MockTodoRepository) DeleteTodo(ctx context.Context, id primitive.ObjectID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
