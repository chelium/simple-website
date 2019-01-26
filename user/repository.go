package user

// UserRepository provides access to user store
type UserRepository interface {
	Create(user *User) (string, error)
	ReadByID(id string) (*User, error)
	ReadByName(username string) (*User, error)
	Update(id string, user *User) error
	Delete(id string) error
}
