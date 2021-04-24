package controller

import (
	"gin-micro-tm/model"
	"gin-micro-tm/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoutesTask ...
func RoutesTask(rg *gin.RouterGroup) {
	post := rg.Group("/post")

	post.GET("/:id", getPostByID)
	post.GET("/", getPosts)
	post.POST("/", createPost)
	post.PUT("/", updatePost)
	post.DELETE("/:id", deletePostByID)
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

func getAllTask(c *gin.Context) {

	var tasks []model.Task
	var task model.Task
	task.Title = "Task one"
	task.Body = "Task body"

	tasks = append(tasks, task)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
