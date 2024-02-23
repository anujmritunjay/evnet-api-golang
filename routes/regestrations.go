package routes

import (
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterEvent(context *gin.Context) {
	user, _ := context.Get("user")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(404, gin.H{"success": false, "message": "Please Provide a valid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(404, gin.H{"success": true, "message": err.Error()})
		return
	}

	err = event.RegisterEvent(user.(models.User).ID)

	if err != nil {
		context.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}
	context.JSON(201, gin.H{"success": true, "message": "You are successfully Registered for the event."})
	return
}

func cancelEvent(context *gin.Context) {
	user, _ := context.Get("user")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(404, gin.H{"success": false, "message": "Please Provide a valid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(404, gin.H{"success": true, "message": err.Error()})
		return
	}

	err = event.CancelEvent(user.(models.User).ID)
	if err != nil {
		context.JSON(500, gin.H{"success": false, "message": "Failed to cancel event."})
		return
	}

	context.JSON(200, gin.H{"success": true, "message": "Your event has been cancelled."})
}
