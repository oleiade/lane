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

	assert.True(t, queue.Size() == 3)
	assert.Equal(t, queue.First(), "3")
	assert.Equal(t, queue.Last(), "1")
}

func TestQueueDequeue_fulfilled(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	item := queue.Dequeue()
	assert.Equal(t, item, "1")
	assert.Equal(t, queue.Size(), 2)
	item = queue.Dequeue()
	assert.Equal(t, item, "2")
	assert.Equal(t, queue.Size(), 1)
	item = queue.Dequeue()
	assert.Equal(t, item, "3")
	assert.Equal(t, queue.Size(), 0)
}

func TestQueueDequeue_empty(t *testing.T) {
	queue := NewQueue()

	item := queue.Dequeue()
	assert.Equal(t, item, nil)
	assert.Equal(t, queue.Size(), 0)
}

func TestQueueHead_fulfilled(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")

	item := queue.Head()
	assert.Equal(t, item, "1")
	assert.Equal(t, queue.Size(), 3)
}

func TestQueueHead_empty(t *testing.T) {
	queue := NewQueue()

	item := queue.Head()
	assert.Equal(t, item, nil)
	assert.Equal(t, queue.Size(), 0)
}
