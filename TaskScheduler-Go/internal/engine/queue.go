package engine

type TaskQueue struct {
	Head *Task
}

func (q *TaskQueue) Add(task *Task) {
	if q.Head == nil || task.RunAt.Before(q.Head.RunAt) {
		task.Next = q.Head
		q.Head = task
		return
	}

	curr := q.Head
	for curr.Next != nil && curr.Next.RunAt.Before(task.RunAt) {
		curr = curr.Next
	}

	task.Next = curr.Next
	curr.Next = task
}

func (q *TaskQueue) Pop() *Task {
	if q.Head == nil {
		return nil
	}
	t := q.Head
	q.Head = q.Head.Next
	t.Next = nil
	return t
}