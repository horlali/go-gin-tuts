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
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", controllers.PostsCreate)
	}

	router.Run()
}
