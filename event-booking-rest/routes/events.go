package routes

import (
	"example.com/event-booking/controllers"
	"example.com/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/")
	auth.Use(middlewares.Authenticate)
	{
		auth.POST("/events", controllers.CreateEvent)
		auth.PUT("/events/:id", controllers.UpdateEvent)
		auth.DELETE("/events/:id", controllers.DeleteEvent)
	}

	rg.GET("/events", controllers.GetEvents)
	rg.GET("/events/:id", controllers.GetEvent)
}
