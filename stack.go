package lane

type Stack struct {
	*Deque
}

func NewStack() *Stack {
	return &Stack{
		Deque: NewDeque(),
	}
}

func (s *Stack) Push(item interface{}) {
	s.Prepend(item)
}

func (s *Stack) Pop() interface{} {
	return s.Shift()
}

func (s *Stack) Head() interface{} {
	return s.First()
}
