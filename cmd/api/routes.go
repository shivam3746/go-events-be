package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.GET("/events", app.getAllEvents)
		v1.GET("/events/:id", app.getEventById)
		v1.GET("/events/:id/attendees", app.getEventAttendeesForEvent)
		v1.GET("/attendees/:id/events", app.getEventsByAttendee)

		v1.POST("/auth/register", app.registerUser)
		v1.POST("/auth/login", app.loginUser)
	}

	authGroup := v1.Group("/")
	authGroup.Use(app.AuthMiddleware())
	{
		authGroup.POST("/events", app.createEvent)
		authGroup.PUT("/events/:id", app.updateEventById)
		authGroup.DELETE("/events/:id", app.deleteEventById)
		authGroup.POST("/events/:id/attendees/:userid", app.addAttendeesToEvent)
		authGroup.DELETE("/events/:id/attendees/:userid", app.removeAttendeeFromEvent)
	}

	return g

}
