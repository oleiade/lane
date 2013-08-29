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

	assert.True(t, len(queue.container) == 3)
	assert.Equal(t, queue.container[0], "1")
	assert.Equal(t, queue.container[1], "2")
	assert.Equal(t, queue.container[2], "3")
}

func TestQueueDequeue_fulfilled(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	item := queue.Dequeue()
	assert.Equal(t, item, "1")
	assert.Equal(t, len(queue.container), 2)
	item = queue.Dequeue()
	assert.Equal(t, item, "2")
	assert.Equal(t, len(queue.container), 1)
	item = queue.Dequeue()
	assert.Equal(t, item, "3")
	assert.Equal(t, len(queue.container), 0)
}

func TestQueueDequeue_empty(t *testing.T) {
	queue := NewQueue()

	item := queue.Dequeue()
	assert.Equal(t, item, nil)
	assert.Equal(t, len(queue.container), 0)
}

func TestQueueHead_fulfilled(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	item := queue.Head()
	assert.Equal(t, item, "1")
	assert.Equal(t, len(queue.container), 3)
}

func TestQueueHead_empty(t *testing.T) {
	queue := NewQueue()

	item := queue.Head()
	assert.Equal(t, item, nil)
	assert.Equal(t, len(queue.container), 0)
}
