package todo

// TodoRepository provides access to todo store.
type TodoRepository interface {
	Create(userID string, todo *Todo) (string, error)
	Read(userID, todoID string) (*Todo, error)
	ReadAll(userID string) ([]*Todo, error)
	Update(userID, todoID string, todo *Todo) error
	Delete(userID, todoID string) error
}
