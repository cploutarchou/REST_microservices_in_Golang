package blog

import (
	user "blog/domain/users"
	"blog/utils/errors"
	"net/http"
)

func CreateUser(user user.User) (*user.User, *errors.APIError) {
	return &user, &errors.APIError{
		Status: http.StatusOK,
	}
}
