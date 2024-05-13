package routes

import (
	"net/http"
	"strconv"

	"elbolaky.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// registerForEvent registers a user for an event based on the provided context.
//
// Parameters:
// - context: The gin.Context object representing the current HTTP request.
//
// Returns:
// - None.
func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

// cancelRegistration cancels a user's registration for an event.
//
// Parameters:
// - context: The gin.Context object representing the current HTTP request.
//
// Returns:
// - None.
func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event := models.Event{
		ID: eventID,
	}

	err = event.CancelRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancelregistration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Canceled!"})
}
