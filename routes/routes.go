package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.GET("/event/:id", getEvent)

	authenticatedRoutes := r.Group("/")

	authenticatedRoutes.Use(middleware.Authenticate)
	authenticatedRoutes.POST("/event", createEvent)
	authenticatedRoutes.PUT("/event/:id", updateEvent)
	authenticatedRoutes.DELETE("/event/:id", deleteEvent)
	authenticatedRoutes.POST("/event/:id/register", RegisterEvent)
	authenticatedRoutes.DELETE("/event/:id/register", cancelEvent)

	r.POST("/signup", signup)
	r.POST("/login", login)

}
