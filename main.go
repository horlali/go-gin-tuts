package main

import (
	"go-gin-tuts/controllers"
	"go-gin-tuts/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	router.POST("/", controllers.PostsCreate)
	router.Run() // listen and serve on 0.0.0.0:8080
}
