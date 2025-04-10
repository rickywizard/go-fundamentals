package controllers

import (
	"net/http"

	"example.com/event-booking/models"
	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not sign up", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully sign up"})
}

func Signin(context *gin.Context) {
	var user models.UserSignin

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login success", "token": token})
}
