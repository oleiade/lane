package lane

// Queue is a First In First Out data structure implementation.
//
// Built upon a Deque container, its API focuses on the following core
// functionalities: Enqueue, Dequeue, Head, Size, Empty.
//
// Every operation has a time complexity of *O(1)*.
//
// Every operation over an instantiated Queue are goroutine-safe.
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

// Enqueue adds an item at the back of the Queue in *O(1)* time complexity.
func (q *Queue[T]) Enqueue(item T) {
	q.container.Prepend(item)
}

// Dequeue removes and returns the Queue's front item in *O(1)* time complexity.
func (q *Queue[T]) Dequeue() (item T, ok bool) {
	return q.container.Pop()
}

// Head returns the Queue's front queue item in *O(1)* time complexity.
func (q *Queue[T]) Head() (item T, ok bool) {
	return q.container.Last()
}

// Size returns the size of the Queue.
func (q *Queue[T]) Size() uint {
	return q.container.Size()
}
