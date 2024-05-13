package todo

type CreateTodoRequest struct {
	Task string `json:"task"`
}

type UpdateTodoRequest struct {
	Completed bool `json:"completed"`
}
