package users

/*
DAO ==================>>>>> IS DATA ACCESS OBJECT
*/

import (
	"fmt"

	"github.com/egnimos/book_store_users_api/datasources/mysql/user_db"
	"github.com/egnimos/book_store_users_api/utils/errors"
	"github.com/egnimos/book_store_users_api/utils/mysql_utils"
)

const (
	errorNoRows                 = "no rows in result set"
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name , last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
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
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
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
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	println(userID)

	user.ID = userID
	return nil
}

//Update : update user in a database
func (user *User) Update() *errors.RestErr {
	// prepare the statement to update the users database
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//execute the statement
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

//Delete : delete user in a database
func (user *User) Delete() *errors.RestErr {
	//prepare the statement to delete the user from the database
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//execute the statement
	if _, err := stmt.Exec(user.ID); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

//FindByStatus : find the user by status
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := user_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to get the list of users")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInternalServerError(fmt.Sprintf("error when trying to get the user while executing the Scan : %s", err.Error()))
		}
		//append the result
		results = append(results, user)
	}

	//if the list of users is empty
	if len(results) == 0 {
		return nil, errors.NewNotFoundError("there is no user list in the database")
	}
	return results, nil
}

//FindByEmailAndPassword : this method return the user by passing the email, password, status
func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//pass the parameters in the QUERYROW
	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		mysql_utils.ParseError(err)
	}
	return nil
}
