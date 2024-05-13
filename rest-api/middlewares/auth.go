package middlewares

import (
	"net/http"

	"elbolaky.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

// Authenticate authenticates the user by validating the token in the request header.
//
// Parameters:
// - context: The gin.Context object representing the current HTTP request.
// Return type: None.
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userID, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
