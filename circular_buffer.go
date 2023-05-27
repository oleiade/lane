package lane

import "sync"

// CircularBuffer represents a circular buffer (or ring buffer) of a fixed size.
//
// A circular buffer has the property of constant-time operations for adding (putting)
// and removing (getting) items.
//
// When the buffer is full and a new item is put, the oldest item in the buffer is overwritten.
//
// The CircularBuffer struct is safe to use concurrently.
type CircularBuffer[T any] struct {
	// data holds the actual data.
	data []T

	// capacity holds the capacity of the buffer.
	capacity int

	// readPos holds the position of the next item to be retrieved.
	full bool

	// readPos holds the position of the next item to be retrieved.
	readPos int

	// writePos holds the position of the next place to write data.
	writePos int

	mu sync.Mutex
}

// NewCircularBuffer creates a new `CircularBuffer` with the specified size.
//
// It returns a pointer to the created `CircularBuffer`.
//
// Time complexity is O(n), where n is the capacity of the circular buffer.
func NewCircularBuffer[T any](capacity int) *CircularBuffer[T] {
	return &CircularBuffer[T]{
		data:     make([]T, capacity),
		capacity: capacity,
		readPos:  0,
		writePos: 0,
	}
}

// Put adds an item to the `CircularBuffer`.
//
// If the buffer is full, the oldest item is overwritten. If the
// capacity of the buffer is 0, it returns false.
//
// Time complexity is O(1). This operation is thread-safe.
func (cb *CircularBuffer[T]) Put(item T) bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.capacity == 0 {
		return false
	}

	if cb.full {
		cb.readPos = (cb.readPos + 1) % cb.capacity
	}

	// Write position points to the next place to write data.
	// Overwrite whatever is there and then move the writePos pointer.
	cb.data[cb.writePos] = item
	cb.writePos = (cb.writePos + 1) % cb.capacity

	if cb.writePos == cb.readPos && !cb.full {
		cb.full = true
	}

	return true
}

// Pop retrieves an item from the CircularBuffer and returns a boolean indicating whether
// the operation was successful.
//
// If the buffer is empty, it returns the zero value of the type and false.
//
// Time complexity is O(1). This operation is thread-safe.
func (cb *CircularBuffer[T]) Pop() (item T, ok bool) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.empty() {
		return item, false
	}

	item = cb.data[cb.readPos]

	// Move the reading position forward, ensuring to wrap around the buffer.
	cb.readPos = (cb.readPos + 1) % cb.capacity

	// If we successfully popped an item and the buffer was full, it is no longer full.
	if cb.full {
		cb.full = false
	}

	return item, true
}

// Peek returns the next item to be retrieved from the CircularBuffer without removing it.
//
// If the buffer is empty, it returns the zero value of the type and false.
//
// Time complexity is O(1). This operation is thread-safe.
func (cb *CircularBuffer[T]) Peek() (item T, ok bool) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.empty() {
		return item, false
	}

	item = cb.data[cb.readPos]

	return item, true
}

// View returns a view of the items in the CircularBuffer
// without removing them.
func (cb *CircularBuffer[T]) View() []T {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.empty() {
		return nil
	}

	if cb.full {
		return append(cb.data[cb.readPos:], cb.data[:cb.writePos]...)
	}

	if cb.writePos > cb.readPos {
		return cb.data[cb.readPos:cb.writePos]
	}

	return append(cb.data[cb.readPos:], cb.data[:cb.writePos]...)
}

// Clear removes all items from the CircularBuffer.
//
// Time complexity is O(n), where n is the capacity of the buffer.
// This operation is thread-safe.
func (cb *CircularBuffer[T]) Clear() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.readPos = 0
	cb.writePos = 0
	cb.full = false

	for i := range cb.data {
		var zero T
		cb.data[i] = zero
	}
}

// Size returns the number of items currently in the CircularBuffer.
//
// Time complexity is O(1). This operation is thread-safe.
func (cb *CircularBuffer[T]) Size() int {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	return cb.size()
}

// Capacity returns the capacity of the CircularBuffer.
func (cb *CircularBuffer[T]) Capacity() int {
	return cb.capacity
}

// Full returns whether the CircularBuffer is full.
//
// Time complexity is O(1). This operation is thread-safe.
func (cb *CircularBuffer[T]) Full() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	return cb.full
}

// Empty returns whether the CircularBuffer is empty.
//
// Time complexity is O(1). This operation is thread-safe.
func (cb *CircularBuffer[T]) Empty() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	return cb.empty()
}

func (cb *CircularBuffer[T]) size() int {
	if cb.full {
		return cb.capacity
	}

	if cb.writePos >= cb.readPos {
		return cb.writePos - cb.readPos
	}

	return cb.capacity - cb.readPos + cb.writePos
}

// empty checks if the buffer is empty.
//
// The buffer is empty if the read and write positions are
// the same and the buffer is not full. This is a helper
// method used internally.
//
// Time complexity is O(1).
func (cb *CircularBuffer[T]) empty() bool {
	return !cb.full && cb.readPos == cb.writePos
}
