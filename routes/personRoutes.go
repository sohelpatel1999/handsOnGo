package routes

import (
	"handsOnGO/controller"

	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {
	api := router.Group("/person")
	{
		api.GET("/", controller.GetAlllPersonDetails)
		api.POST("/", controller.CreatePerson)
		api.GET("/details", controller.GetPersonDetailsById)
		api.PUT("/details", controller.UpdatePersonDetailsById)
		// api.POST("/", controller.CreatePerson)
	}
}
