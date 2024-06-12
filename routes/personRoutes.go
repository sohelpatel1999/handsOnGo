package routes

import (
	"handsOnGO/controller"

	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {
	api := router.Group("/person")
	{
		api.GET("/", controller.GerPersonDetails)
		api.POST("/", controller.CreatePerson)
	}
}
