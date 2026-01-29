package api

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "task-scheduler/internal/engine"
)

type AddTaskRequest struct {
    Delay int `json:"delay"` // seconds
}

func AddTaskHandler(s *engine.Scheduler) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req AddTaskRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        task := &engine.Task{
            ID:    time.Now().String(),
            RunAt: time.Now().Add(time.Duration(req.Delay) * time.Second),
            Action: func() {
                println("Task executed at:", time.Now().String())
            },
        }

        s.Queue.Add(task)

        c.JSON(http.StatusOK, gin.H{"message": "task scheduled"})
    }
}