package config

import (
	item_controller "tracker/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterItemController(router *gin.Engine) {
	const controllerEndpoint = "/items"

	router.GET(controllerEndpoint, item_controller.GetItems)
	router.GET(controllerEndpoint+"/:id", item_controller.GetItemById)
	router.POST(controllerEndpoint, item_controller.PostItems)
	router.DELETE(controllerEndpoint+"/:id", item_controller.DeleteItemById)
}
