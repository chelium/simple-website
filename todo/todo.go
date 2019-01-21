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

// Gets a todo
func Get() (Todo, error) {
	t := new(Todo)
	t.Title = "Testing"
	t.Description = "Test description"
	t.Complete = false
	return *t, nil
}

func newTodo(title, description string) Todo {
	return Todo{
		ID:          xid.New().String(),
		Title:       title,
		Description: description,
		Complete:    false,
	}
}
