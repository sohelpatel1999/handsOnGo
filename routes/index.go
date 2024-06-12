package routes

import "github.com/gin-gonic/gin"

func DocumentRoutes() *gin.Engine {

	router := gin.Default()

	routes(router)

	return router
}