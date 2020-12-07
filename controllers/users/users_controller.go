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
		restErr := errors.NewBadRequestError("invalid json request while creating a user")
		c.JSON(restErr.Status, restErr)
		return
	}
	//send the user struct to the services
	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//GetUser : this function will get the user info of given ID
func GetUser(c *gin.Context) {
	// strconv.ParseInt(c.Param("user_id"), 10, 64)
	// strconv.Atoi(c.Param("user_id"))
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
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
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

//UpdateUser : this user usually update the data from the database...
func UpdateUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	//intialize
	var user users.User
	//check whether the given json body is valid or not
	if err := c.ShouldBindJSON(&user); err != nil {
		invalidErr := errors.NewInternalServerError("invalid json body")
		c.JSON(invalidErr.Status, invalidErr)
		return
	}

	//send the user struct to the services
	user.ID = userID
	//check whether the request method is PATCH and PUT
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	//final implementation
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//DeleteUser : delete the data from the users database of given ID
func DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		paramErr := errors.NewBadRequestError("user id should be a number")
		c.JSON(paramErr.Status, paramErr)
		return
	}

	//send the userID to the services
	result, deleteErr := services.DeleteUser(userID)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

//SearchUser : search the user on the basis of ID or name or status or email
func SearchUser(c *gin.Context) {
	status := c.Query("status")

	usersList, err := services.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, usersList.Marshall(c.GetHeader("X-Public") == "true"))
}
