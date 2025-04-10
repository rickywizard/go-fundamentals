package routes

import (
	"example.com/event-booking/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/signup", controllers.Signup)
	rg.POST("/signin", controllers.Signin)
}
