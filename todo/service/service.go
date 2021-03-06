package todo

import (
	"github.com/chelium/simple-website/todo"
	"github.com/chelium/simple-website/user"
)

type Service interface {
	CreateUserTodo(userID string, todo todo.Todo) (string, error)
	GetUserTodo(userID, todoID string) (todo.Todo, error)
	GetUserTodos(userID string) ([]todo.Todo, error)
	UpdateUserTodo(userID, todoID string, todo todo.Todo) error
	DeleteUserTodo(userID, todoID string) error
}

type service struct {
	todos todo.TodoRepository
	users user.UserRepository
}

func (s *service) CreateUserTodo(userID string, todo todo.Todo) (string, error) {
	user, err := s.users.ReadByID(userID)
	if err != nil {
		return "", err
	}
	todoID, err := s.todos.Create(userID, &todo)
	if err != nil {
		return "", err
	}
	user.Todos = append(user.Todos, todoID)
	err = s.users.Update(userID, user)
	if err != nil {
		return "", err
	}
	return todoID, nil
}

func (s *service) GetUserTodo(userID, todoID string) (todo.Todo, error) {
	var result todo.Todo
	todo, err := s.todos.Read(userID, todoID)
	if err != nil {
		return result, err
	}
	return *todo, nil
}

func (s *service) GetUserTodos(userID string) ([]todo.Todo, error) {
	var result []todo.Todo
	todos, err := s.todos.ReadAll(userID)
	if err != nil {
		return result, err
	}
	for _, v := range todos {
		result = append(result, *v)
	}
	return result, nil
}

func (s *service) UpdateUserTodo(userID, todoID string, todo todo.Todo) error {
	return s.todos.Update(userID, todoID, &todo)
}

func (s *service) DeleteUserTodo(userID, todoID string) error {
	user, err := s.users.ReadByID(userID)
	if err != nil {
		return err
	}
	removeIdx := 0
	for ; removeIdx < len(user.Todos); removeIdx++ {
		if user.Todos[removeIdx] == todoID {
			break
		}
	}
	if removeIdx < len(user.Todos) {
		user.Todos = append(user.Todos[:removeIdx], user.Todos[removeIdx+1:]...)
	}
	if err := s.users.Update(userID, user); err != nil {
		return err
	}
	return s.todos.Delete(userID, todoID)
}

// NewService creates a todo service with necessary dependencies.
func NewService(todos todo.TodoRepository, users user.UserRepository) Service {
	return &service{
		todos: todos,
		users: users,
	}
}
