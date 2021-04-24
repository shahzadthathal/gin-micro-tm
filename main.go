package main

//gin-micro-tm (gin microservice task management)
//support for versioned modules to our application by running this command:
//go mod init gin-micro-tm
//It will create go.mod file
//Lets get gin by running this command:
//go get -u github.com/gin-gonic/gin
//This will download the Gin dependency and add it to the go.mod file
//new file appear go.sum. That file contains checksums for direct and indirect dependencies of our application

import (
	"log"
	"net/http"

	"gin-micro-tm/config"
	"gin-micro-tm/model"

	"github.com/gin-gonic/gin"
)

// type Task struct {
// 	Title string
// 	Body  string
// }

func handleGetTasks(c *gin.Context) {
	var tasks []model.Task
	var task model.Task
	task.Title = "Task one"
	task.Body = "Task body"

	tasks = append(tasks, task)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func main() {
	//log.Println("Hello from go")

	var err error
	// Setup database
	config.DB, err = config.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	//In Golang we can defer the execution of
	//code to when the function is exited
	//This is a really nice way to free resources
	defer config.DB.Close()

	r := gin.Default()
	r.GET("/tasks/", handleGetTasks)
	r.Run()
}
