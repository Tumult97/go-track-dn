package item_controller

import (
	"fmt"
	"net/http"
	"strconv"
	. "tracker/models/models"

	"github.com/gin-gonic/gin"
)

var items = []Item{
	{Id: 1, Name: "TV", Description: "", Price: 15000, IsPerItem: true},
	{Id: 2, Name: "Couch", Description: "", Price: 5000, IsPerItem: true},
	{Id: 3, Name: "Computer", Description: "", Price: 45000, IsPerItem: true},
}

func GetItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func PostItems(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func GetItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Parameter Id requried a valid number. %s", err.Error())})
	}

	for _, item := range items {
		if item.Id == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found."})
}

func DeleteItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Parameter 'id' requires a valid number: %s", err.Error()),
		})
		return
	}

	// Filter items and check if any were deleted
	newItems := make([]Item, 0, len(items))
	found := false

	for _, item := range items {
		if item.Id == id {
			found = true
			continue // Skip the item to be deleted
		}
		newItems = append(newItems, item)
	}

	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Item with id %d not found.", id),
		})
		return
	}

	items = newItems
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Item with id %d deleted successfully.", id),
	})
}
