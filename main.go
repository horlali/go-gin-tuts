package main

import (
	"go-gin-tuts/controllers"
	"go-gin-tuts/initializers"
	"go-gin-tuts/migration"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.CreateDB()   // create db
	migration.MigrateTables() // migrate tables
}

func main() {
	// instantiate server
	router := gin.Default()

	// configure routes
	router.GET("/posts", controllers.PostsIndex)
	router.POST("/posts", controllers.PostsCreate)

	// run server
	router.Run("localhost:8080")
}
