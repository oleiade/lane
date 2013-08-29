package lane

import (
	"container/list"
	"sync"
)

type Stack struct {
	sync.RWMutex
	container *list.List
}

func NewStack() *Stack {
	return &Stack{
		container: list.New(),
	}
}

func (s *Stack) Push(item interface{}) {
	s.Lock()
	defer s.Unlock()

	s.container.PushFront(item)

	return
}

func (s *Stack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()

	var item interface{} = nil
	var firstListItem *list.Element = nil

	firstListItem = s.container.Front()
	if firstListItem != nil {
		item = s.container.Remove(firstListItem)
	}

	return item
}

func (s *Stack) Size() int {
	s.RLock()
	defer s.RUnlock()

	return s.container.Len()
}

func (s *Stack) Empty() bool {
	s.RLock()
	defer s.RUnlock()

	return s.container.Len() == 0
}

func (s *Stack) Head() interface{} {
	s.RLock()
	defer s.RUnlock()

	item := s.container.Front()
	if item != nil {
		return item.Value
	} else {
		return nil
	}
}
