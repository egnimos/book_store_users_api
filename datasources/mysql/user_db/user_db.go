package user_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "MYSQL_USERS_USERNAME"
	mysqlUsersPassword = "MYSQL_USERS_PASSWORD"
	mysqlUsersHost     = "MYSQL_USERS_HOST"
	mysqlUsersSchema   = "MYSQL_USERS_SCHEMA"
)

//Client
var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)

	var err error
	//connecting to the database server
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	//checking the connection
	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database suucessfully configured")
}
