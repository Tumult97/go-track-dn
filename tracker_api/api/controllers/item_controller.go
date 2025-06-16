package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tracker/utils/helpers"
	"tracker/utils/jwt"

	"tracker/repository/repos"

	"github.com/gin-gonic/gin"
)

var itemReposiotory = repos.NewItemRepository()

func GetItems(context *gin.Context) {
	userId, _ := jwt.ExtractUserIDFromContext(context)

	items, err := itemReposiotory.GetAll(userId)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error getting items: %s", err.Error()),
		})
		return
	}

	context.IndentedJSON(http.StatusOK, items)
}

func GetItemById(context *gin.Context) {
	userId, _ := jwt.ExtractUserIDFromContext(context)
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Parameter Id requried a valid number. %s", err.Error())})
	}

	item, err := itemReposiotory.GetById(id, userId)

	if item == nil || err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found."})
		return
	}

	context.IndentedJSON(http.StatusOK, item)
}

func CreateItem(context *gin.Context) {
	helpers.HandleEntitySave(
		context,
		itemReposiotory.Create)
}

func UpdateItem(context *gin.Context) {
	helpers.HandleEntitySave(
		context,
		itemReposiotory.Update)
}

func DeleteItemById(context *gin.Context) {
	userId, _ := jwt.ExtractUserIDFromContext(context)
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Parameter 'id' requires a valid number: %s", err.Error()),
		})
		return
	}

	id, err = itemReposiotory.DeleteById(id, userId)

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
