package lane

// Queue is a FIFO (First in first out) data structure implementation.
// It is based on a Deque container and focuses its API on core
// functionalities: Enqueue, Dequeue, Head, Size, Empty.
//
// Every operation's time complexity is O(1).
//
// As it is implemented using a Deque container, every operations
// over an instiated Queue are synchronized and goroutine-safe.
type Queue[T any] struct {
	container *Deque[T]
}

// NewQueue produces a new Queue instance.
func NewQueue[T any](items ...T) *Queue[T] {
	deque := NewDeque[T]()

	for _, item := range items {
		deque.container.PushFront(item)
	}

	return &Queue[T]{
		container: deque,
	}
}

// Enqueue adds an item at the back of the Queue in O(1) time complexity.
func (q *Queue[T]) Enqueue(item T) {
	q.container.Prepend(item)
}

// Dequeue removes and returns the Queue's front item in O(1) time complexity.
func (q *Queue[T]) Dequeue() (item T, ok bool) {
	return q.container.Pop()
}

// Head returns the Queue's front queue item in O(1) time complexity.
func (q *Queue[T]) Head() (item T, ok bool) {
	return q.container.Last()
}

// Size returns the size of the Queue.
func (q *Queue[T]) Size() uint {
	return q.container.Size()
}
