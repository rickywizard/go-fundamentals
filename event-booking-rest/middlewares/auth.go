package middlewares

import (
	"net/http"
	"strings"

	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
		return
	}
	actualToken := parts[1]

	claims, err := utils.VerifyToken(actualToken)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	context.Set("claims", claims)

	context.Next()
}
