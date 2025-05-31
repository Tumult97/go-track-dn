package main

import (
    "github.com/gin-gonic/gin"
    "rest/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.Run("localhost:5100")
}