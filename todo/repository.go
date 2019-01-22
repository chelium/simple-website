package todo

// TodoRepository provides access to todo store.
type TodoRepository interface {
	Create(todo *Todo)
	Read(id string) (*Todo, error)
	ReadAll() ([]*Todo, error)
	Update(id string, todo *Todo) error
	Delete(id string) error
}
