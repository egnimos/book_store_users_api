package app

import (
	"github.com/egnimos/book_store_users_api/controllers/ping"
	"github.com/egnimos/book_store_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.FindUser)
	router.POST("/users", users.CreateUser)
}
