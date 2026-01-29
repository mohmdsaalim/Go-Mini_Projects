package engine

type TaskQueue struct {
	Head *Task
}
// Insert task in sorted order 
func (tq *TaskQueue) Add(task *Task) {

	if tq.Head == nil || task.RunAt.Before(tq.Head.RunAt){
		println("➕ Task added:", task.RunAt.String())
		task.Next = tq.Head
		tq.Head = task
		return
	}

	// Find position
	curr := tq.Head
	for curr.Next != nil && curr.Next.RunAt.Before(task.RunAt){
		curr = curr.Next
	}
	task.Next = curr.Next
	curr.Next = task
}

// remove the first task
func (tq *TaskQueue) Pop()*Task {
	if tq.Head == nil{
		return nil
	}
	task := tq.Head
	tq.Head = tq.Head.Next
	task.Next = nil
	return task
}