package main

import (
	"log"
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

	gin.SetMode(gin.ReleaseMode)

	database.Connect()
	defer database.Close()

	router := gin.Default()

	config.RegisterControllers(router)

	log.Printf("Server starting on localhost:8080")
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
