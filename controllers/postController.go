package controllers

import (
	"go-gin/initializers"
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	//create schema for valiadation of request

	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}

	ctx.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": post})
}

func PostsFindAll(ctx *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	ctx.JSON(http.StatusOK, gin.H{"message": "All posts", "posts": posts})
}

func PostsFindOne(ctx *gin.Context) {
	var post models.Post
	id := ctx.Param("id")

	initializers.DB.First(&post, id)
	ctx.JSON(http.StatusOK, gin.H{"message": "Post found", "post": post})
}

func PostUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var input struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	ctx.Bind(&input)

	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := initializers.DB.Model(&post).Updates(&models.Post{Title: input.Title, Body: input.Body}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post updated", "post": post})
}

func PostDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	if (initializers.DB.First(&models.Post{}, id).RowsAffected == 0) {
		ctx.JSON(404, gin.H{"message": "Post not found"})
		return
	}

	initializers.DB.Delete(&models.Post{}, id)

	ctx.JSON(200, gin.H{"message": "Post deleted successfully"})
}
