package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/database"
	"example.com/rest-api/models"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	fmt.Println("Welcome to the REST-API in Go Language.")

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is the test")
	})

	r.Run(":8080")
}

func getEvents(c *gin.Context) {
	events, err := models.GetEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to fetch the data."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data."})
		return
	}
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create the data."})
		return
	}
	event.ID = 1
	event.UserID = 1
	c.JSON(http.StatusCreated, gin.H{"status": true, "data": event})

}

func getEvent(context *gin.Context) {

	params, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get Params",
		})
	}
	var event models.Event
	event, err = models.GetEventByID(params)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "No data found with Provided ID",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    event,
	})
}
