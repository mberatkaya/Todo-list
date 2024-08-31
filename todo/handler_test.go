package todo

import (
	"TODOproject/utility"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllTodosHandler(t *testing.T) {
	app := fiber.New()
	mockService := new(MockTodoService)
	handler := NewTodoHandler(mockService)
	handler.RegisterRoutes(app)

	todoID1 := primitive.NewObjectID()
	todo1 := &Todo{ID: todoID1, Task: "Test Task 1", Completed: false}

	todoID2 := primitive.NewObjectID()
	todo2 := &Todo{ID: todoID2, Task: "Test Task 2", Completed: true}

	mockService.On("GetAllTodos", mock.Anything).Return([]*Todo{todo1, todo2}, nil)

	req := httptest.NewRequest(fiber.MethodGet, "/api/todo", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var responseData []Todo
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, []Todo{
		{ID: todo1.ID, Task: todo1.Task, Completed: todo1.Completed},
		{ID: todo2.ID, Task: todo2.Task, Completed: todo2.Completed},
	}, responseData)

	mockService.AssertExpectations(t)
}

func TestCreateTodoHandler(t *testing.T) {
	app := fiber.New()
	mockService := new(MockTodoService)
	handler := NewTodoHandler(mockService)
	handler.RegisterRoutes(app)

	todoID := primitive.NewObjectID()
	todo := &Todo{ID: todoID, Task: "Test Task", Completed: false}
	mockService.On("CreateTodo", mock.Anything, "Test Task").Return(todo, nil)

	reqBody, _ := json.Marshal(CreateTodoDto{Task: "Test Task"})
	req := httptest.NewRequest(fiber.MethodPost, "/api/todo/", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var responseData utility.Response
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		t.Fatal(err)
	}

	var createTodoAck CreateTodoAck
	if err := json.Unmarshal(responseData.Data, &createTodoAck); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, todoID.Hex(), createTodoAck.Id)
	assert.Equal(t, todo.Task, createTodoAck.Task)
	assert.Equal(t, todo.Completed, createTodoAck.Completed)

	// Mock servisin beklentilerini kontrol et
	mockService.AssertExpectations(t)
}

func TestUpdateTodoCompletionHandler(t *testing.T) {
	app := fiber.New()
	mockService := new(MockTodoService)
	handler := NewTodoHandler(mockService)
	handler.RegisterRoutes(app)

	id := primitive.NewObjectID()
	mockService.On("UpdateTodoCompletion", mock.Anything, id, true).Return(nil)

	reqBody, _ := json.Marshal(UpdateTodoDto{Completed: true})
	req := httptest.NewRequest(fiber.MethodPut, "/api/todo/"+id.Hex(), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestDeleteTodoHandler(t *testing.T) {
	app := fiber.New()
	mockService := new(MockTodoService)
	handler := NewTodoHandler(mockService)
	handler.RegisterRoutes(app)

	id := primitive.NewObjectID()
	mockService.On("DeleteTodo", mock.Anything, id).Return(nil)

	req := httptest.NewRequest(fiber.MethodDelete, "/api/todo/"+id.Hex(), nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)

	mockService.AssertExpectations(t)
}
