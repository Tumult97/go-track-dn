package helpers

import (
	"fmt"
	"net/http"
	"tracker/models/entities"
	"tracker/utils/jwt"

	"github.com/gin-gonic/gin"
)

func HandleEntitySave[T entities.Entity](
	context *gin.Context,
	repoFunc func(item T) (T, error)) {

	userId, _ := jwt.ExtractUserIDFromContext(context)

	var updatedEntity T

	err := context.BindJSON(&updatedEntity)

	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	updatedEntity.SetUserId(userId)

	result, err := repoFunc(updatedEntity)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error saving: %s", err.Error()),
		})
		return
	}

	context.IndentedJSON(http.StatusOK, result)
}

//func HandleGetAll[T any]
