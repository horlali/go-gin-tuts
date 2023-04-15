package controllers

import (
	"go-gin-tuts/initializers"
	"go-gin-tuts/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	post := models.Post{Title: "Hello", Body: "Body"}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
