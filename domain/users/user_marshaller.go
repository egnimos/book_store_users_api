package users

import (
	"encoding/json"
	"log"
)

type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

//Marshall : this marshalls the list of structs into the json format baised on the headers request
func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

//Marshall : this marshalls the given struct into json format baised on the headers request
func (user *User) Marshall(isPublic bool) interface{} {
	//if the request is form the PUBLIC network then it returns the PUBLICUSER struct as a response
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	//if the request is form the private network then it returns the PRIVATEUSER struct as a response
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
