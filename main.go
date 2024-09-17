package main

import (
	"go-gin/controllers"
	"go-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {

	//gin's default router
	r := gin.Default()

	//route handler
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsFindAll)
	r.GET("/posts/:id", controllers.PostsFindOne)
	r.PATCH("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	//listen and served on 8080 default
	r.Run()
}
