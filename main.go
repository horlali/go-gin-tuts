package main

import (
	"go-gin-tuts/controllers"
	"go-gin-tuts/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	// instantiate server
	router := gin.Default()

	// configure routes
	router.GET("/", controllers.PostsCreate)

	// run server
	router.Run()
}
