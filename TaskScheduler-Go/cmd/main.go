package main

import (
	"github.com/gin-gonic/gin"
	"task-scheduler/internal/api"
	"task-scheduler/internal/engine"
)

func main() {
	scheduler := engine.NewScheduler()
	scheduler.Start()

	r := gin.Default()
	r.POST("/task", api.AddTaskHandler(scheduler))
	r.Run(":8080")
}