package todo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllTodos(t *testing.T) {
	mockRepo := new(MockTodoRepository)

	expectedTodos := []Todo{
		{ID: primitive.NewObjectID(), Task: "Test Task 1", Completed: false},
		{ID: primitive.NewObjectID(), Task: "Test Task 2", Completed: true},
	}

	mockRepo.On("GetAllTodos", mock.Anything).Return(expectedTodos, nil)

	todoService := NewTodoService(mockRepo)

	todos, err := todoService.GetAllTodos(context.Background())

	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedTodos, todos)
	mockRepo.AssertExpectations(t)
}

func TestCreateTodo(t *testing.T) {
	mockRepo := new(MockTodoRepository)

	newTodo := &Todo{
		ID:        primitive.NewObjectID(),
		Task:      "New Task",
		Completed: false,
	}

	mockRepo.On("CreateTodo", mock.Anything, mock.AnythingOfType("*todo.Todo")).Return(newTodo, nil)

	todoService := NewTodoService(mockRepo)

	createdTodo, err := todoService.CreateTodo(context.Background(), "New Task")

	assert.NoError(t, err)
	assert.Equal(t, newTodo, createdTodo)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTodoCompletion(t *testing.T) {
	mockRepo := new(MockTodoRepository)

	id := primitive.NewObjectID()
	completed := true

	mockRepo.On("UpdateTodoCompletion", mock.Anything, id, completed).Return(nil)

	todoService := NewTodoService(mockRepo)

	err := todoService.UpdateTodoCompletion(context.Background(), id, completed)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTodo(t *testing.T) {
	mockRepo := new(MockTodoRepository)

	id := primitive.NewObjectID()

	mockRepo.On("DeleteTodo", mock.Anything, id).Return(nil)

	todoService := NewTodoService(mockRepo)

	err := todoService.DeleteTodo(context.Background(), id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
