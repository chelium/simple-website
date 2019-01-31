package todo

import (
	//"errors"
	"github.com/rs/xid"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	Owner       string `json:"owner"`
	CreatedBy   string `json:"createdBy"`
	AssignedTo  string `json:"assignedTo"`
}

// NewTodo creates a new todo.
func NewTodo(title, description string) *Todo {
	return &Todo{
		ID:          xid.New().String(),
		Title:       title,
		Description: description,
		Complete:    false,
	}
}
