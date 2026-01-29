package engine

import "time"

type Scheduler struct {
	queue      *TaskQueue
	addTaskCh  chan *Task
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		queue:     &TaskQueue{},
		addTaskCh: make(chan *Task),
	}
}

func (s *Scheduler) Start() {
	go func() {
		println("Scheduler started")

		for {
			// If no task, wait for new one
			if s.queue.Head == nil {
				task := <-s.addTaskCh
				s.queue.Add(task)
				continue
			}

			nextTask := s.queue.Head
			waitTime := time.Until(nextTask.RunAt)

			select {
			case task := <-s.addTaskCh:
				// new task arrived → add & re-evaluate
				s.queue.Add(task)

			case <-time.After(waitTime):
				// time to execute
				task := s.queue.Pop()
				if task != nil && task.Action != nil {
					println(" Executing task")
					task.Action()
				}
			}
		}
	}()
}

// exposed method (API uses this)
func (s *Scheduler) Schedule(task *Task) {
	s.addTaskCh <- task
}