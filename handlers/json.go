package handlers

import (
	"Go/GinVetEx/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {

	events := GetAllEvents

	c.HTML(
		http.StatusOK,
		"index.html",

		gin.H{
			"title":   "Home Page",
			"payload": events,
		},
	)

}

func ShowJSONPage(c *gin.Context) {

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	if !ValidateJson(jsonData) {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON")
	}

	var f interface{}
	error := json.Unmarshal(jsonData, &f)
	if error != nil {
		fmt.Println(error)
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON")
	}
	c.JSON(http.StatusOK, f)
}

func ValidateJson(input []byte) bool {
	event := models.Event{
		Name:    "Log",
		Type:    "Info",
		Message: "Json Beautified",
		Date:    time.Now(),
	}

	val := json.Valid(input)
	if val {
		fmt.Println(val)
	} else {
		fmt.Println("Not valid JSON.")
		event.Message = "Json parse failed."
	}

	insertResult, err := Create(&event)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult)

	return val
}
