package controllers

import (
	"api/initializers"
	"api/models"
	"api/requests"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostCreate(c *gin.Context) {
	// Get all posts
	err := c.BindJSON(&requests.PostCreate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()}) // Send 400 with error message
		return
	}

	post := models.Post{Title: requests.PostCreate.Title, Body: requests.PostCreate.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		log.Fatal(result.Error.Error())
		c.AbortWithStatus(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})

}

func PostGet(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func PostGetById(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	result := initializers.DB.First(&post, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// log.Fatal(result.Error.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": "Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	result := initializers.DB.Delete(&post, id)
	// log.Print(result.Error.Error())
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": "Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete"})
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	err := c.BindJSON(&requests.PostCreate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()}) // Send 400 with error message
		return
	}
	result := initializers.DB.First(&post, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": "Not found"})
		return
	}
	post.Title = requests.PostCreate.Title
	post.Body = requests.PostCreate.Body
	initializers.DB.Save(&post)
	c.JSON(http.StatusOK, gin.H{"post": post})
}
