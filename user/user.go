package user

import (
	"errors"
	"github.com/rs/xid"
)

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}
