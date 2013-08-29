package lane

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	assert.True(t, queue.container.Len() == 3)
	assert.Equal(t, queue.container.Front().Value, "3")
	assert.Equal(t, queue.container.Back().Value, "1")
}

func TestQueueDequeue_fulfilled(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	item := queue.Dequeue()
	assert.Equal(t, item, "1")
	assert.Equal(t, queue.container.Len(), 2)
	item = queue.Dequeue()
	assert.Equal(t, item, "2")
	assert.Equal(t, queue.container.Len(), 1)
	item = queue.Dequeue()
	assert.Equal(t, item, "3")
	assert.Equal(t, queue.container.Len(), 0)
}

func TestQueueDequeue_empty(t *testing.T) {
	queue := NewQueue()

	item := queue.Dequeue()
	assert.Equal(t, item, nil)
	assert.Equal(t, queue.container.Len(), 0)
}

func TestQueueHead_fulfilled(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	item := queue.Head()
	assert.Equal(t, item, "1")
	assert.Equal(t, queue.container.Len(), 3)
}

func TestQueueHead_empty(t *testing.T) {
	queue := NewQueue()

	item := queue.Head()
	assert.Equal(t, item, nil)
	assert.Equal(t, queue.container.Len(), 0)
}

func TestQueueEmpty_fulfilled(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	assert.False(t, queue.Empty())
}

func TestQueueEmpty_empty_queue(t *testing.T) {
	queue := NewQueue()
	assert.True(t, queue.Empty())
}
