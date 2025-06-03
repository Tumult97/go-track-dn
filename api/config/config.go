package config

import (
	controllers "tracker/api/controllers"
	"tracker/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterItemController(router *gin.RouterGroup) {
	const controllerEndpoint = "/items"

	router.GET(controllerEndpoint, controllers.GetItems)
	router.GET(controllerEndpoint+"/:id", controllers.GetItemById)
	router.POST(controllerEndpoint, controllers.PostItems)
	router.DELETE(controllerEndpoint+"/:id", controllers.DeleteItemById)
}

func RegisterAuthController(router *gin.Engine) {
	const authEndpoint = "/auth"

	router.POST(authEndpoint+"/register", controllers.Register)
	router.POST(authEndpoint+"/login", controllers.Login)
}

func RegisterProfileController(router *gin.RouterGroup) {
	router.GET("/profile", controllers.GetProfile)
}

func RegisterControllers(router *gin.Engine) {
	RegisterAuthController(router)

	protected := router.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		RegisterProfileController(protected)
		RegisterItemController(protected)
	}
}
