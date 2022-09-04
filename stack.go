package lane

// Stack is a LIFO (Last in first out) data structure implementation.
// It is based on a Deque container and focuses its API on core
// functionalities: Push, Pop, Head, Size, Empty.
//
// Every operation's time complexity is O(1).
//
// As it is implemented using a Deque container, every operations
// over an instiated Stack are synchronized and goroutine-safe.
type Stack[T any] struct {
	container *Deque[T]
}

// NewStack produces a new Stack instance.
//
// If any initialization variadic items are provided, they
// will be inserted as is: lower index being the head of stack.
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
