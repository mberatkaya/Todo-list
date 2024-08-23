package todo

import (
	"TODOproject/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestGetAllTodos(t *testing.T) {

	mockRepo := new(mocks.MockTodoRepository)

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

	mockRepo := new(mocks.MockTodoRepository)

	todo := &Todo{
		ID: primitive.NewObjectID(), Task: "New Task", Completed: false,
	}

	mockRepo.On("CreateTodo", mock.Anything, todo).Return(todo, nil)

	todoService := NewTodoService(mockRepo)

	createdTodo, err := todoService.CreateTodo(context.Background(), "New Task")

	assert.NoError(t, err)
	assert.Equal(t, todo, createdTodo)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTodoCompletion(t *testing.T) {

	mockRepo := new(mocks.MockTodoRepository)

	id := primitive.NewObjectID()
	completed := true

	mockRepo.On("UpdateTodoCompletion", mock.Anything, id, completed).Return(nil)

	todoService := NewTodoService(mockRepo)

	err := todoService.UpdateTodoCompletion(context.Background(), id, completed)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTodo(t *testing.T) {
	// Mock oluşturun
	mockRepo := new(mocks.MockTodoRepository)

	id := primitive.NewObjectID()

	mockRepo.On("DeleteTodo", mock.Anything, id).Return(nil)

	todoService := NewTodoService(mockRepo)

	err := todoService.DeleteTodo(context.Background(), id)

	// Sonuçları doğrulayın
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
