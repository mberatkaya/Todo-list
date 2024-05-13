package todo

import "go.mongodb.org/mongo-driver/bson/primitive"

// Todo struct'ı, bir To-Do öğesini temsil eder
type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string             `json:"task"`
	Completed bool               `json:"completed"`
}

type CreateTodoRequest struct {
	Task string `json:"task"`
}

type UpdateTodoRequest struct {
	Completed bool `json:"completed"`
}
