package lane

import (
	"sync"
)

// Dequer is the interface that wraps the basic Deque operations.
type Dequer[T any] interface {
	Deque[T] | BoundDeque[T]
}

// Deque implements a head-tail linked list data structure.
//
// Under the hood, the Deque's implementation is built upon a doubly linked list
// container, so that every operation's time complexity is *O(1)*.
// Every operation on a Deque is goroutine-safe and ready.
//
// Note that linked-list are not CPU-cache friendly).
// for concurrent usage.
type Deque[T any] struct {
	sync.RWMutex

	// The underlying storage container.
	container *List[T]
}

// NewDeque produces a new Deque instance.
func NewDeque[T any](items ...T) *Deque[T] {
	container := New[T]()

	for _, item := range items {
		container.PushBack(item)
	}

	return &Deque[T]{
		container: container,
	}
}

// Append inserts item at the back of the Deque in an *O(1)* time complexity.
func (d *Deque[T]) Append(item T) {
	d.Lock()
	defer d.Unlock()

	d.container.PushBack(item)
}

// Prepend inserts item at the Deque's front in an *O(1)* time complexity.
func (d *Deque[T]) Prepend(item T) {
	d.Lock()
	defer d.Unlock()

	d.container.PushFront(item)
}

// Pop removes and returns the back element of the Deque in an *O(1)* time complexity.
func (d *Deque[T]) Pop() (item T, ok bool) {
	d.Lock()
	defer d.Unlock()

	lastElement := d.container.Back()
	if lastElement != nil {
		item = d.container.Remove(lastElement)
		ok = true
	}

	return
}

// Shift removes and returns the front element of the Deque in *O(1)* time complexity.
func (d *Deque[T]) Shift() (item T, ok bool) {
	d.Lock()
	defer d.Unlock()

	firstElement := d.container.Front()
	if firstElement != nil {
		item = d.container.Remove(firstElement)
		ok = true
	}

	return
}

// First returns the first value stored in the Deque in *O(1)* time complexity.
func (d *Deque[T]) First() (item T, ok bool) {
	d.RLock()
	defer d.RUnlock()

	frontItem := d.container.Front()
	if frontItem != nil {
		item = frontItem.Value
		ok = true
	}

	return
}

// Last returns the last value stored in the Deque in *O(1)* time complexity.
func (d *Deque[T]) Last() (item T, ok bool) {
	d.RLock()
	defer d.RUnlock()

	if backItem := d.container.Back(); backItem != nil {
		item = backItem.Value
		ok = true
	}

	return
}

// Size returns the Deque's size.
func (d *Deque[T]) Size() uint {
	d.RLock()
	defer d.RUnlock()

	return d.container.Len()
}

// Empty checks if the Deque is empty.
func (d *Deque[T]) Empty() bool {
	d.RLock()
	defer d.RUnlock()

	return d.container.Len() == 0
}

// Capacitor defines operations related to capacity management.
type Capacitor interface {
	// Capacity returns the current capacity of the underlying type implementation.
	Capacity() int

	// IsFull returns whether the implementing type instance is full.
	IsFull() bool
}

// BoundDeque implements a head-tail linked list data structure
// with a user-defined capacity. Once full, `Append` and `Prepend`
// operations on a BoundedDeque fail.
//
// Under the hood, BoundDeque's implementation relies upon a doubly linked list
// container. Thus, every operation on a BoundedDeque has a time complexity of *O(1)*.
// Every operation on a BoundDeque is goroutine-safe.
//
// Note that linked-list are not CPU-cache friendly).
// for concurrent usage.
type BoundDeque[T any] struct {
	Deque[T]

	// capacity defines an upper bound limit for the BoundDeque's size.
	capacity uint
}

// NewBoundDeque produces a new BoundDeque instance with the provided
// capacity.
func NewBoundDeque[T any](capacity uint, values ...T) *BoundDeque[T] {
	return &BoundDeque[T]{
		Deque:    *NewDeque(values...),
		capacity: capacity,
	}
}

// Capacity returns BoundDeque's capacity.
func (d *BoundDeque[T]) Capacity() uint {
	return d.capacity
}

// Full checks if the BoundDeque is full.
func (d *BoundDeque[T]) Full() bool {
	return d.container.Len() >= d.capacity
}

// Append inserts an item at the back of the BoundDeque in an *O(1)* time complexity.
// If BoundDeque's capacity disallows the insertion, Append returns false.
func (d *BoundDeque[T]) Append(item T) bool {
	d.Lock()
	defer d.Unlock()

	if d.Full() {
		return false
	}

	d.container.PushBack(item)

	return true
}

// Prepend inserts item at the BoundDeque's front in an *O(1)* time complexity.
// If BoundDeque's capacity disallows the insertion, Prepend returns false.
func (d *BoundDeque[T]) Prepend(item T) bool {
	d.Lock()
	defer d.Unlock()

	if d.Full() {
		return false
	}

	d.container.PushFront(item)

	return true
}
