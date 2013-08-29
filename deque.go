package lane

import (
	"container/list"
	"sync"
)

type Deque struct {
	sync.RWMutex
	container *list.List
}

func NewDeque() *Deque {
	return &Deque{
		container: list.New(),
	}
}

// Append inserts element at the back of the Deque in a O(1) time complexity
func (s *Deque) Append(item interface{}) {
	s.Lock()
	defer s.Unlock()

	s.container.PushBack(item)

	return
}

// Prepend inserts element at the Deques front in a O(1) time complexity
func (s *Deque) Prepend(item interface{}) {
	s.Lock()
	defer s.Unlock()

	s.container.PushFront(item)

	return
}

// Pop removes the last element of the deque in a O(1) time complexity
func (s *Deque) Pop() interface{} {
	s.Lock()
	defer s.Unlock()

	var item interface{} = nil
	var lastContainerItem *list.Element = nil

	lastContainerItem = s.container.Back()
	if lastContainerItem != nil {
		item = s.container.Remove(lastContainerItem)
	}

	return item
}

// Shift removes the last element of the deque in a O(1) time complexity
func (s *Deque) Shift() interface{} {
	s.Lock()
	defer s.Unlock()

	var item interface{} = nil
	var firstContainerItem *list.Element = nil

	firstContainerItem = s.container.Front()
	if firstContainerItem != nil {
		item = s.container.Remove(firstContainerItem)
	}

	return item
}

// First returns the first value stored in the deque in a O(1) time complexity
func (s *Deque) First() interface{} {
	s.RLock()
	defer s.RUnlock()

	item := s.container.Front()
	if item != nil {
		return item.Value
	} else {
		return nil
	}
}

// Last returns the last value stored in the deque in a O(1) time complexity
func (s *Deque) Last() interface{} {
	s.RLock()
	defer s.RUnlock()

	item := s.container.Back()
	if item != nil {
		return item.Value
	} else {
		return nil
	}
}

// Size returns the actual deque size
func (s *Deque) Size() int {
	s.RLock()
	defer s.RUnlock()

	return s.container.Len()
}

// Empty checks if the deque is empty
func (s *Deque) Empty() bool {
	s.RLock()
	defer s.RUnlock()

	return s.container.Len() == 0
}
