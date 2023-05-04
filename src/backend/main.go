package main

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()

	rGroup := routes.Group("/chat")

	rGroup.GET("", controller.FirstPage)
	rGroup.GET("/:id", controller.GetChatsController)
	rGroup.PUT("/add_chat", controller.AddChats)
	rGroup.DELETE("/:id/delete_chat", controller.DeleteChats)
	rGroup.POST("/:id/post_chat", controller.PostChatsController)

	routes.Run()
}
