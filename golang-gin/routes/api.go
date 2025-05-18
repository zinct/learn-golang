package routes

import (
	eventController "golang-gin/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", eventController.Index)
	router.GET("/events/:id", eventController.Show)
	router.POST("/events", eventController.Store)
}
