package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to fetch the data."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	user, _ := c.Get("user")
	var event models.Event
	err := c.ShouldBindJSON(&event)
	event.UserID = user.(models.User).ID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data."})
		return
	}
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create the data."})
		return
	}
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

func updateEvent(c *gin.Context) {
	user, _ := c.Get("user")

	params, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to get params."})
	}
	event, err := models.GetEventByID(params)
	fmt.Println(event.UserID, user.(models.User).ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch the event."})
		return
	}

	if event.UserID != user.(models.User).ID {
		c.JSON(401, gin.H{"success": false, "message": "You are not authorized to update the event."})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data."})
		return
	}
	updatedEvent.ID = params
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to Update Event"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "Event Updated successfully."})

}

func deleteEvent(c *gin.Context) {
	params, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user, _ := c.Get("user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to get params."})
	}
	event, err := models.GetEventByID(params)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Event Not Found."})
		return
	}
	if event.UserID != user.(models.User).ID {
		c.JSON(401, gin.H{"success": false, "message": "You are not authorized to update the event."})
		return
	}
	err = event.DeleteEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to delete Event."})
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "Event Deleted Successfully."})

}
