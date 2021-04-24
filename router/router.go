package router

import (
	"gin-micro-tm/controller"

	"github.com/gin-gonic/gin"
)

// NewRoutes router global
func NewRoutes() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api")

	// register router from each controller service
	controller.RoutesTask(v1)

	return router
}
