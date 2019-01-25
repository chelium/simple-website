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
}

// Get a todo
func Get() (*Todo, error) {
	t := new(Todo)
	t.Title = "Testing"
	t.Description = "Test description"
	t.Complete = false
	return t, nil
}

// Add a todo
func Add(title, description string) string {
	t := NewTodo(title, description)
	return t.ID
}

// Update a todo
func Update(id string, todo *Todo) error {
	return nil
}

// Delete a todo
func Delete(id string) error {
	return nil
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
