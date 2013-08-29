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
	defer q.Unlock()

	q.container = append(q.container, item)

	return
}

func (q *Queue) Dequeue() interface{} {
	q.Lock()
	defer q.Unlock()

	item, err := q.getFirstItem()
	if err != nil {
		return nil
	}

	if len(q.container) > 1 {
		q.container = q.container[1:]
	} else {
		q.container = *new([]interface{})
	}

	return item
}

func (q *Queue) Size() int {
	q.RLock()
	defer q.RUnlock()

	return len(q.container)
}

func (q *Queue) Empty() bool {
	q.RLock()
	defer q.RUnlock()

	return len(q.container) == 0
}

func (q *Queue) Head() interface{} {
	q.RLock()
	defer q.RUnlock()

	item, err := q.getFirstItem()
	if err != nil {
		return nil
	}

	return item
}

func (q *Queue) getFirstItem() (interface{}, error) {
	if len(q.container) >= 1 {
		return q.container[0], nil
	} else {
		return nil, errors.New("The queue is empty")
	}
}
