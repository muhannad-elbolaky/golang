package main

import (
	"os"

	"elbolaky.com/rest-api/db"
	"elbolaky.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main is the entry point of the Go program.
//
// It loads the environment variables from the .env file using the godotenv package.
// If there is an error loading the .env file, it panics with the error message.
// It sets the GIN_MODE environment variable using the os package.
// It initializes the database connection using the db package.
// It creates a new Gin server with the default middleware.
// It registers the routes using the routes package.
// It starts the server on port 8080.
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	gin.SetMode(os.Getenv("GIN_MODE"))

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
