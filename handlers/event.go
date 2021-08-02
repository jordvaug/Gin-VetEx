package handlers

import (
	"Go/GinVetEx/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetEvents(c *gin.Context) {
	var loadedEvents, err = GetAllEvents()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": loadedEvents})
}

func HandleCreateTask(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := Create(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
