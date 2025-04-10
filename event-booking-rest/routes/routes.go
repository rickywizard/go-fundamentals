package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	api := server.Group("/api")

	RegisterEventRoutes(api)
	RegisterAuthRoutes(api)
	RegisterEventRegistrationsRoute(api)
}
