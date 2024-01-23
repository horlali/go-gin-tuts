package controllers

import (
	"go-gin-tuts/entity"
	"go-gin-tuts/initializers"
	"go-gin-tuts/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body entity.Post

	// handle error
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Could not create post"})
		return
	}

	// Return it
	c.JSON(200, post)
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
