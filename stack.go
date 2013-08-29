package lane

type Stack struct []interface{}

func NewStack(items []interface{}) *Stack {
	stack := &Stack{}
	copy(stack, items)

	return stack
}

func (s *Stack) Push(item interface{}) {
	
}

func (s *Stack) Pop() {
	
}

func (s *Stack) Item() {
	
}