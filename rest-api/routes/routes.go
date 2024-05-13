package routes

import (
	"elbolaky.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all the routes for the server.
//
// Parameters:
// - server: The gin.Engine server to register the routes on.
// Return type: None.
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.GET("/users", getUsers)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
