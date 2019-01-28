package user

import (
	"github.com/chelium/simple-website/todo"
)

type Service interface {
	CreateUser(user User) (string, error)
	GetUser(id string) (User, error)
	UpdateUser(id string, user User) error
	DeleteUser(id string) error
}

type service struct {
	todos todo.TodoRepository
	users UserRepository
}

func (s *service) CreateUser(user User) (string, error) {
	return s.users.Create(&user)
}

func (s *service) GetUser(id string) (User, error) {
	var result User
	user, err := s.users.ReadByID(id)
	if err != nil {
		return result, err
	}
	return *user, nil
}

func (s *service) UpdateUser(id string, user User) error {
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
func NewService(todos todo.TodoRepository, users UserRepository) Service {
	return &service{
		todos: todos,
		users: users,
	}
}
