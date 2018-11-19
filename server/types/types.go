package types

type TodoItem struct {
	Id       string `json:"id"`
	Summary  string `json:"summary"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
}

type TodoItems struct {
	TodoItems []TodoItem `json:"todoItems"`
}
