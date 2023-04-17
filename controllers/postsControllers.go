package controllers

import (
	"go-gin-tuts/initializers"
	"go-gin-tuts/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"status": "error",
		})
		return
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
