package main

import (
	"log"
	"tracker/api/config"
	database "tracker/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	database.Connect()
	defer database.Close()

	router := gin.Default()

	config.RegisterItemController(router)

	log.Printf("Server starting on localhost:5100")
	if err := router.Run("localhost:5100"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
