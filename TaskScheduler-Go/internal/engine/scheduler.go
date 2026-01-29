package engine

import (
    "time"
)

type Scheduler struct {
    Queue *TaskQueue
}

func NewScheduler() *Scheduler {
    return &Scheduler{
        Queue: &TaskQueue{},
    }
}

func (s *Scheduler) Start() {
    go func() {
        println("🟢 Scheduler started")

        for {
            task := s.Queue.Head
            if task == nil {
                time.Sleep(1 * time.Second)
                continue
            }

            println("⏳ Waiting for task:", task.RunAt.String())

            now := time.Now()
            if now.Before(task.RunAt) {
                time.Sleep(task.RunAt.Sub(now))
            }

            task = s.Queue.Pop()
            if task != nil && task.Action != nil {
                println("🚀 Executing task")
                task.Action()
            }
        }
    }()
}