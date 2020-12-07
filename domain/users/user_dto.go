package users

/*
DTO ==================>>>>> IS DATA TRANSFER OBJECT
*/

import (
	"strings"

	"github.com/egnimos/book_store_users_api/utils/errors"
)

const (
	StatusActive = "active"
)

//User : User struct for performing some operations
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

//Validate : it filter the email
func (user *User) Validate() *errors.RestErr {
	//validate the firstname
	user.FirstName = strings.TrimSpace(user.FirstName)
	//validate the lastname
	user.LastName = strings.TrimSpace(user.LastName)
	//validate the email
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	//validate the password
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
