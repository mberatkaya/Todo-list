package todo

type CreateTodoDto struct {
	Task string `json:"task"`
}

type CreateTodoAck struct {
	Id        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type UpdateTodoDto struct {
	Completed bool `json:"completed"`
}
