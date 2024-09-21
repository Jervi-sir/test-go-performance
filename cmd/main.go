package main

import (
	"log"
	"octa_api_go/internal/database"
	"octa_api_go/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the PostgreSQL database
	database.ConnectDB()

	// Initialize Gin router
	router := gin.Default()

	// Serve the loader.io file for verification
	router.StaticFile("/loaderio-2bce31847666303628bdc85edc7de9df.txt", "./public/loaderio-2bce31847666303628bdc85edc7de9df.txt")

	// Define routes
	router.GET("/search", handlers.SearchItems)

	// Start the server on 0.0.0.0:8080
	log.Println("Server running on 0.0.0.0:8080")
	router.Run("0.0.0.0:8080")
}
