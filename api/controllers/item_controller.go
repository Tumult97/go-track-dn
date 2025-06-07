package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	entities "tracker/models/entities"
	"tracker/utils/jwt"

	"tracker/repository/repos"

	"github.com/gin-gonic/gin"
)

var repo = repos.NewItemRepository()

func getUserId(context *gin.Context) (int, error) {
	userId, err := jwt.ExtractUserIDFromContext(context)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func GetItems(context *gin.Context) {
	userId, _ := getUserId(context)

	items, err := repo.GetAll(userId)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error getting items: %s", err.Error()),
		})
		return
	}

	context.IndentedJSON(http.StatusOK, items)
}

func PostItems(context *gin.Context) {
	userId, _ := getUserId(context)
	var newItem entities.Item

	if err := context.BindJSON(&newItem); err != nil {
		return
	}

	var id int
	var err error

	if newItem.Id != 0 {
		id, err = repo.Update(newItem, userId)
	} else {
		id, err = repo.Create(newItem, userId)
	}

	if err != nil || id == 0 {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error saving item: %s", err.Error()),
		})
		return
	}

	context.IndentedJSON(http.StatusCreated, newItem)
}

func GetItemById(context *gin.Context) {
	userId, _ := getUserId(context)
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Parameter Id requried a valid number. %s", err.Error())})
	}

	item, err := repo.GetById(id, userId)

	if item == nil || err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found."})
		return
	}

	context.IndentedJSON(http.StatusOK, item)
}

func DeleteItemById(context *gin.Context) {
	userId, _ := getUserId(context)
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Parameter 'id' requires a valid number: %s", err.Error()),
		})
		return
	}

	id, err = repo.DeleteById(id, userId)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Item with id %d not found.", id),
		})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Item with id %d deleted successfully.", id),
	})
}
