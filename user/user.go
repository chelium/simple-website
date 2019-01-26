package user

import (
	"github.com/rs/xid"
)

type User struct {
	ID           string   `json:"id"`
	Username     string   `json:"username"`
	PasswordHash string   `json:"password"`
	Todos        []string `json:"todos"`
}

// NewUser creates a new user.
func NewUser(username, passwordHash string) *User {
	todos := make([]string, 0)
	return &User{
		ID:           xid.New().String(),
		Username:     username,
		PasswordHash: passwordHash,
		Todos:        todos,
	}
}
