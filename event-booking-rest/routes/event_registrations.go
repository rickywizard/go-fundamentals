package routes

import (
	"example.com/event-booking/controllers"
	"example.com/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterEventRegistrationsRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/")
	auth.Use(middlewares.Authenticate)
	{
		auth.POST("/events/:id/register", controllers.RegisterForEvent)
		auth.DELETE("/events/:id/register", controllers.CancelEventRegistration)
	}
}
