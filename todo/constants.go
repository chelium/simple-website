package todo

type TodoStatus int

const (
	NotStarted TodoStatus = iota
	NotAssigned
	InProgress
	Removed
	Completed
)
