package main

import (
	"go-gin/initializers"
	"go-gin/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
