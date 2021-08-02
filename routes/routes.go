package routes

import (
	handlers "Go/GinVetEx/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	fmt.Println("Initialize Routes.")
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/events/", handlers.HandleGetEvents)
	router.POST("/json/", handlers.ShowJSONPage)
}
