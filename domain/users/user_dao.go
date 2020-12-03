package users

/*
DAO ==================>>>>> IS DATA ACCESS OBJECT
*/

import (
	"github.com/egnimos/book_store_users_api/datasources/mysql/user_db"
	"github.com/egnimos/book_store_users_api/utils/date_utils"
	"github.com/egnimos/book_store_users_api/utils/errors"
	"github.com/egnimos/book_store_users_api/utils/mysql_utils"
)

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

//Get : fetch the user from a database of a given ID
func (user *User) Get() *errors.RestErr {
	//Prepare creates a prepare statement for later queries
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//QueryRow executes the prepared statement by taking the argument and returned *ROW which is always not-nil...and the error will be return after the SCAN function will call
	result := stmt.QueryRow(user.ID)

	//call the scan function to executes the final statement
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}

	//after all the function and statement executes this will return nil
	return nil
}

//Save : save user in a database
func (user *User) Save() *errors.RestErr {
	//get the user from the database to check whether the user exists or not
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//execute the statement
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		mysql_utils.ParseError(saveErr)
	}

	// userID, err := insertResult.LastInsertId()
	// if err != nil {
	// 	return mysql_utils.ParseError(err)
	// }
	println(insertResult)

	// user.ID = userID
	return nil
}
