package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrls()

	fmt.Println("server is Started....")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}