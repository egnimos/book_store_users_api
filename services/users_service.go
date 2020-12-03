package services

import (
	"github.com/egnimos/book_store_users_api/domain/users"
	"github.com/egnimos/book_store_users_api/utils/errors"
)

//GetUser : fetch the user from database
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//CreateUser : create user service
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//filter and check the email
	if err := user.Validate(); err != nil {
		return nil, err
	}

	//save the user in database
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
