package controller

import (
	"gin-micro-tm/model"
	"gin-micro-tm/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoutesTask ...
func RoutesTask(rg *gin.RouterGroup) {
	route := rg.Group("/task")

	//route.GET("/:id", getTaskByID)
	route.GET("/", getTasks)
	route.POST("/", createTask)
	//route.PUT("/", updateTask)
	//route.DELETE("/:id", deleteTaskByID)
}

func createTask(c *gin.Context) {

	var item model.Task

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json", "error": err.Error()})
		return
	}

	item, err := repository.CreateItem(item)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)

}

func getTasks(c *gin.Context) {

	items, err := repository.GetItemAll()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}
