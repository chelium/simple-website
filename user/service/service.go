package service

import (
	"github.com/chelium/simple-website/todo"
	"github.com/chelium/simple-website/user"
)

type Service interface {
	CreateUser(user user.User) (string, error)
	GetUser(id string) (user.User, error)
	UpdateUser(id string, user user.User) error
	DeleteUser(id string) error
}

type service struct {
	todos todo.TodoRepository
	users user.UserRepository
}

func (s *service) CreateUser(user user.User) (string, error) {
	return s.users.Create(&user)
}

func (s *service) GetUser(id string) (user.User, error) {
	var result user.User
	user, err := s.users.ReadByID(id)
	if err != nil {
		return result, err
	}
	return *user, nil
}

func (s *service) UpdateUser(id string, user user.User) error {
	return s.users.Update(id, &user)
}

func (s *service) DeleteUser(id string) error {
	user, err := s.users.ReadByID(id)
	if err != nil {
		return err
	}
	for _, todoID := range user.Todos {
		if err := s.todos.Delete(id, todoID); err != nil {
			return err
		}
	}
	return s.users.Delete(id)
}

// NewService creates a user service with necessary dependencies.
func NewService(todos todo.TodoRepository, users user.UserRepository) Service {
	return &service{
		todos: todos,
		users: users,
	}
}
