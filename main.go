package main

import (
	"log"
	"time"
	"tracker/api/config"
	jwtModels "tracker/models/jwt"
	database "tracker/repository"
	"tracker/utils/jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize JWT
	jwt.Init(jwtModels.Config{
		Secret:     "your_jwt_secret_key_change_this_in_production",
		Expiration: 24 * time.Hour,
		Issuer:     "tracker-api",
	})

	// Initialize database connection
	database.Connect()
	defer database.Close()

	router := gin.Default()

	config.RegisterControllers(router)

	log.Printf("Server starting on localhost:5100")
	if err := router.Run("localhost:6900"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
