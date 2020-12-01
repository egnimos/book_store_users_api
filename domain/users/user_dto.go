package users

/*
DTO ==================>>>>> IS DATA TRANSFER OBJECT
*/

import (
	"strings"

	"github.com/egnimos/book_store_users_api/utils/errors"
)

//User : User struct for performing some operations
type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
