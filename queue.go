package lane

import (
	"errors"
	"sync"
)

type Queue struct {
	sync.RWMutex
	container []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		container: *new([]interface{}),
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Lock()

	q.container = append(q.container, item)

	q.Unlock()

	return
}

func (q *Queue) Dequeue() interface{} {
	q.Lock()

	item, err := q.getFirstItem()
	if err != nil {
		return nil
	}

	if len(q.container) > 1 {
		q.container = q.container[1:]
	} else {
		q.container = *new([]interface{})
	}

	q.Unlock()

	return item
}

func (q *Queue) Size() int {
	q.Lock()

	queueSize := len(q.container)

	q.Unlock()

	return queueSize
}

func (q *Queue) getFirstItem() (interface{}, error) {
	if len(q.container) >= 1 {
		return q.container[0], nil
	} else {
		return nil, errors.New("The queue is empty")
	}
}
