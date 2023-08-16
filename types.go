package main

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func NewTodo(title string, description string, completed bool) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}
