package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"rest/models"
)

var albums = []models.Item{
	{Id: 1, Name: "TV", Description: "", Price: 15000, IsPerItem: true},
	{Id: 2, Name: "Couch", Description: "", Price: 5000, IsPerItem: true},
	{Id: 3, Name: "Computer", Description: "", Price: 45000, IsPerItem: true},
}

// GetAlbums handles GET /albums
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
