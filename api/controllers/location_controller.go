package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tracker/models/entities"
	"tracker/utils/jwt"

	"tracker/repository/repos"

	"github.com/gin-gonic/gin"
)

var locationRepository = repos.NewLocationRepository()

func GetLocations(context *gin.Context) {
	userId, _ := jwt.ExtractUserIDFromContext(context)

	locations, err := locationRepository.GetAll(userId)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error getting locations: %s", err.Error()),
		})
		return
	}

	context.IndentedJSON(http.StatusOK, locations)
}

func GetLocationById(context *gin.Context) {
	userId, _ := jwt.ExtractUserIDFromContext(context)
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Parameter Id requried a valid number. %s", err.Error())})
	}

	location, err := locationRepository.GetById(id, userId)

	if location == nil || err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Location not found."})
		return
	}

	context.IndentedJSON(http.StatusOK, location)
}

func CreateLocation(context *gin.Context) {
	userId, _ := jwt.ExtractUserIDFromContext(context)

	var newLocation entities.Location

	err := context.BindJSON(&newLocation)

	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newLocation.UserID = userId

	id, err := locationRepository.Create(newLocation)

	if err != nil || id == 0 {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error saving item: %s", err.Error()),
		})
		return
	}

	newLocation.ID = id

	context.IndentedJSON(http.StatusCreated, newLocation)
}

func UpdateLocation(context *gin.Context) {
	userId, _ := jwt.ExtractUserIDFromContext(context)

	var updatedLocation entities.Location

	err := context.BindJSON(&updatedLocation)

	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	updatedLocation.UserID = userId

	id, err := locationRepository.Update(updatedLocation)

	if err != nil || id == 0 {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error saving item: %s", err.Error()),
		})
		return
	}

	context.IndentedJSON(http.StatusAccepted, updatedLocation)
}

func DeleteLocation(context *gin.Context) {
	userId, _ := jwt.ExtractUserIDFromContext(context)

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil || id == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Parameter Id requried a valid number. %s", err.Error())})
	}

	id, err = locationRepository.DeleteById(id, userId)

	if err != nil || id == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Location with id %d not found.", id),
		})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Item with id %d deleted successfully.", id),
	})
}
