package config

import (
	controllers "tracker/api/controllers"
	"tracker/middleware"

	"github.com/gin-gonic/gin"
)

func registerAuthController(router *gin.Engine) {
	const authEndpoint = "/auth"

	router.POST(authEndpoint+"/register", controllers.Register)
	router.POST(authEndpoint+"/login", controllers.Login)
	router.POST(authEndpoint+"/refresh", controllers.RefreshToken)
}

func registerItemController(router *gin.RouterGroup) {
	const controllerEndpoint = "/items"

	router.GET(controllerEndpoint, controllers.GetItems)
	router.GET(controllerEndpoint+"/:id", controllers.GetItemById)
	router.POST(controllerEndpoint, controllers.CreateItem)
	router.PUT(controllerEndpoint, controllers.UpdateItem)
	router.DELETE(controllerEndpoint+"/:id", controllers.DeleteItemById)
}

func registerProfileController(router *gin.RouterGroup) {
	router.GET("/profile", controllers.GetProfile)
}

func registerLocationController(router *gin.RouterGroup) {
	controllerEndpoint := "/location"
	router.GET(controllerEndpoint, controllers.GetLocations)
	router.GET(controllerEndpoint+"/:id", controllers.GetLocationById)
	router.POST(controllerEndpoint, controllers.CreateLocation)
	router.PUT(controllerEndpoint, controllers.UpdateLocation)
	router.DELETE(controllerEndpoint+"/:id", controllers.DeleteLocation)
}

func RegisterControllers(router *gin.Engine) {
	registerAuthController(router)

	protected := router.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		registerProfileController(protected)
		registerItemController(protected)
		registerLocationController(protected)
	}
}
