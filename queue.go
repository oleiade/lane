package lane

import (
	"container/list"
	"sync"
)

type Queue struct {
	sync.RWMutex
	container *list.List
}

func NewQueue() *Queue {
	return &Queue{
		container: list.New(),
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Lock()
	defer q.Unlock()

	q.container.PushFront(item)

	return
}

func (q *Queue) Dequeue() interface{} {
	q.Lock()
	defer q.Unlock()

	var item interface{} = nil
	var lastListItem *list.Element = nil

	lastListItem = q.container.Back()

	if lastListItem != nil {
		item = q.container.Remove(lastListItem)
	}

	return item
}

func (q *Queue) Size() int {
	q.RLock()
	defer q.RUnlock()

	return q.container.Len()
}

func (q *Queue) Empty() bool {
	q.RLock()
	defer q.RUnlock()

	return q.container.Len() == 0
}

func (q *Queue) Head() interface{} {
	q.RLock()
	defer q.RUnlock()

	item := q.container.Back()
	if item != nil {
		return item.Value
	} else {
		return nil
	}
}
