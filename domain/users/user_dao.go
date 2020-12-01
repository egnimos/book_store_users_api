package users

/*
DAO ==================>>>>> IS DATA ACCESS OBJECT
*/

import (
	"fmt"

	"github.com/egnimos/book_store_users_api/utils/errors"
)

var (
	usersDB = make(map[int]*User)
)

//Get : fetch the user from a database of a given ID
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

//Save : save user in a database
func (user *User) Save() *errors.RestErr {
	//get the user from the database to check whether the user exists or not
	currentUser := usersDB[user.ID]
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}

	// if user doesn't exists then save it to the database
	usersDB[user.ID] = user
	return nil
}
