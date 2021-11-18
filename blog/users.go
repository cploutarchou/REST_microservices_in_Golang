package blog

import (
	"blog/domain/users"
	"blog/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.APIError) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
func CreateUser(user users.User) (*users.User, *errors.APIError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
