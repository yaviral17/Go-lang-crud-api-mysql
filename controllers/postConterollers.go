package controllers

import (
	initializers "github.com/aviral/go-crud/Initializers"
	"github.com/aviral/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//Create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func CheckConnection(c *gin.Context) {

	c.JSON(200, gin.H{
		"response": "pong",
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {

	//Get id off url
	id := c.Param("id")

	//Get post
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {

	//Get id off url
	id := c.Param("id")

	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	//Find the post were updating
	var post models.Post

	initializers.DB.Find(&post, id)

	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// var err error
	//Get id off url
	id := c.Param("id")
	//Delete the post
	var post models.Post
	initializers.DB.Find(&post, id)
	initializers.DB.Delete(&post, id)

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}
