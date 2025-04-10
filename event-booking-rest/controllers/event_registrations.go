package controllers

import (
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	claims, err := utils.GetClaims(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	userId := claims.UserID

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for the event"})
}

func CancelEventRegistration(context *gin.Context) {
	claims, err := utils.GetClaims(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	userId := claims.UserID

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err.Error()})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not cancel event registration", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled registration from event"})
}
