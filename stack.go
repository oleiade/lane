package lane

// Stack implements a Last In First Out data structure.
//
// Built upon a Deque container, it focuses its API on the following core
// functionalities: Push, Pop, Head, Size, Empty.
//
// Every operation's has a time complexity of *O(1)*.
//
// Every operations over an instantiated Stack are goroutine-safe.
type Stack[T any] struct {
	container *Deque[T]
}

// NewStack produces a new Stack instance.
//
// When providing initialization items, those will be inserted as-is: lower index being the head of the stack.
func NewStack[T any](items ...T) (stack *Stack[T]) {
	// FIXME: unwrap here instead of depending on Deque's for clarity
	return &Stack[T]{
		container: NewDeque(items...),
	}
}

// Push adds on an item on the top of the Stack.
func (s *Stack[T]) Push(item T) {
	s.container.Prepend(item)
}

// Pop removes and returns the item on the top of the Stack.
func (s *Stack[T]) Pop() (item T, ok bool) {
	return s.container.Shift()
}

// Head returns the item on the top of the Stack.
func (s *Stack[T]) Head() (item T, ok bool) {
	return s.container.First()
}

// Size returns the size of the Stack.
func (s *Stack[T]) Size() uint {
	return s.container.Size()
}
