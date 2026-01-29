package engine

import "time"


type Task struct {
    ID      string
    RunAt   time.Time
    Action  func()
    Next    *Task
}