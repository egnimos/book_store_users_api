package services

import (
	"github.com/egnimos/book_store_users_api/domain/users"
	"github.com/egnimos/book_store_users_api/utils/crypto_utils"
	"github.com/egnimos/book_store_users_api/utils/date_utils"
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
	user.Password = crypto_utils.GetMd5(user.Password)
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser : update the user services
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser := &users.User{ID: user.ID}
	if err := currentUser.Get(); err != nil {
		return nil, err
	}

	//if request method is PATCH
	if isPartial {
		//if user FIRSTNAME is not empty then assign the value to the CURRENTUSER
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		//if user LASTNAME is not empty then assign the value to the CURRENTUSER
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		//if user EMAIL is not empty then assign the value to the CURRENTUSER
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	//pass it to the domain section
	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}

//DeleteUser : delete the user from the given ID
func DeleteUser(userID int64) (map[string]string, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Delete(); err != nil {
		return nil, err
	}

	return map[string]string{"message": "user is successfully deleted"}, nil
}

//SearchUser : search the user from the given status
func SearchUser(status string) (users.Users, *errors.RestErr) {
	result := &users.User{}
	return result.FindByStatus(status)
}
