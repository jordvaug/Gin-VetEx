package routes

import (
	handlers "Go/GinVetEx/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/events/", handlers.HandleGetEvents)
	router.POST("/json/", handlers.ShowJSONPage)
}
