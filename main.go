package main

import (
	"log"
	"os"
	"tracker/api/config"
	database "tracker/repository"
	"tracker/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	jwt.Init()

	mode := os.Getenv("API_MODE")
	gin.SetMode(mode)

	database.Connect()
	defer database.Close()

	router := gin.Default()

	config.RegisterControllers(router)

	log.Printf("Server starting on localhost:5100")
	if err := router.Run("localhost:5100"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
