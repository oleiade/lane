package lane

type Queue struct {
	*Deque
}

func NewQueue() *Queue {
	return &Queue{
		Deque: NewDeque(),
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Prepend(item)
}

func (q *Queue) Dequeue() interface{} {
	return q.Pop()
}

func (q *Queue) Head() interface{} {
	return q.Last()
}
