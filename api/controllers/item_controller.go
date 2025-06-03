package item_controller

import (
	"fmt"
	"net/http"
	"strconv"
	entities "tracker/models/entities"

	"tracker/repository/repos"

	"github.com/gin-gonic/gin"
)

var items = []entities.Item{
	{Id: 1, Name: "TV", Description: "", Price: 15000, IsPerItem: true},
	{Id: 2, Name: "Couch", Description: "", Price: 5000, IsPerItem: true},
	{Id: 3, Name: "Computer", Description: "", Price: 45000, IsPerItem: true},
}

var repo = repos.NewItemRepository()

func GetItems(c *gin.Context) {
	items, err := repo.GetAll()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error getting items: %s", err.Error()),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, items)
}

func PostItems(c *gin.Context) {
	var newItem entities.Item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	var id int
	var err error

	if newItem.Id != 0 {
		id, err = repo.Update(newItem)
	} else {
		id, err = repo.Create(newItem)
	}

	if err != nil || id == 0 {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error saving item: %s", err.Error()),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, newItem)
}

func GetItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Parameter Id requried a valid number. %s", err.Error())})
	}

	item, err := repo.GetById(id)

	if item == nil || err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}

func DeleteItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Parameter 'id' requires a valid number: %s", err.Error()),
		})
		return
	}

	id, err = repo.DeleteById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Item with id %d not found.", id),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Item with id %d deleted successfully.", id),
	})
}
