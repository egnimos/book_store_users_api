package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"

	"github.com/egnimos/book_store_users_api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

//ParseError : checking and parsing the mysql error caused by the mysql database
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record found with the matching ID")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error parsing database response : %s", err.Error()))
	}

	//switch between the cases
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
