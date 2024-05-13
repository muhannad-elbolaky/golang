package routes

import (
	"net/http"

	"elbolaky.com/rest-api/models"
	"elbolaky.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

// getUsers retrieves all users from the database and returns them in the response.
//
// Parameters:
// - context: The gin.Context object representing the current HTTP request.
// Return type: None.
func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users."})
		return
	}

	context.JSON(http.StatusOK, users)
}

// signup handles the signup functionality.
//
// It expects a JSON payload containing user information in the request body.
// The function validates the request data and saves the user to the database.
// If the request data is invalid, it returns a JSON response with a "Could not parse request data." message and a 400 status code.
// If there is an error saving the user to the database, it returns a JSON response with a "Could not create user." message and a 500 status code.
// If the user is successfully saved, it returns a JSON response with a "User created successfully." message, the created user object, and a 201 status code.
//
// Parameters:
// - context: The gin.Context object representing the current HTTP request.
//
// Returns:
// - None.
func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "user": user})
}

// login handles the login functionality.
//
// It expects a JSON payload containing user information in the request body.
// The function validates the request data and generates a token for the user.
// If the request data is invalid, it returns a JSON response with a "Could not parse request data." message and a 400 status code.
// If the user credentials are invalid, it returns a JSON response with a "Could not authenticate user." message and a 401 status code.
// If there is an error generating the token, it returns a JSON response with a "Could not authenticate user." message and a 500 status code.
// If the login is successful, it returns a JSON response with a "Login successful." message and the generated token.
//
// Parameters:
// - context: The gin.Context object representing the current HTTP request.
//
// Returns:
// - None.
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(&user.Email, &user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}
