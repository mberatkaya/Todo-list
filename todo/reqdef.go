package todo

type CreateTodoDto struct {
	Task string `json:"task"`
}

type UpdateTodoDto struct {
	Completed bool `json:"completed"`
}
