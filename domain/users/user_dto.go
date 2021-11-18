package users

import (
	"blog/utils/errors"
	"strings"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (user User) Validate() *errors.APIError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequest("Invalid email address.")
	}
	return nil
}
