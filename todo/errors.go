package todo

import "errors"

// ErrNotFound is used when a todo could not be found.
var ErrNotFound = errors.New("todo not found")
