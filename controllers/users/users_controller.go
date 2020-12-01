package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/egnimos/book_store_users_api/domain/users"
	"github.com/egnimos/book_store_users_api/services"
	"github.com/egnimos/book_store_users_api/utils/errors"
)

/*
NOTE:::: int64 can handle this number of users ==>> 1235376474764783456  BUT int can handle this number of users ==>> 1235376474764783456
CONCLUSION:::: as both int64 && int can handle the same number of request....
*/

//CreateUser : this function will create the user in data base
func CreateUser(c *gin.Context) {
	//intialize
	var user users.User
	//fetch the json request and unmarshal the json file into struct
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json request")
		c.JSON(restErr.Status, restErr)
		return
	}
	//send the user struct to the services
	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetUser : this function will get the user info of given ID
func GetUser(c *gin.Context) {
	// strconv.ParseInt(c.Param("user_id"), 10, 64)
	// strconv.Atoi(c.Param("user_id"))
	userID, userErr := strconv.Atoi(c.Param("user_id"))
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	// send the id to the services
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

//func FightUser(c *gin.Context) {
//
//}
