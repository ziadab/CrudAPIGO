package main

import (
	"api/controllers"
	"api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostGet)
	r.GET("/post/:id", controllers.PostGetById)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.PATCH(("post/:id"), controllers.UpdatePost)
	r.Run()
}
