package lane

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDequeAppend(t *testing.T) {
	queue := NewDeque()

	assert.True(t, queue.Append("1"))
	assert.True(t, queue.Append("2"))
	assert.True(t, queue.Append("3"))

	assert.True(t, queue.container.Len() == 3)
	assert.Equal(t, queue.container.Front().Value, "1")
	assert.Equal(t, queue.container.Back().Value, "3")
}

func TestDequeAppendWithCapacity(t *testing.T) {
	queue := NewCappedDeque(2)

	assert.True(t, queue.Append("1"))
	assert.True(t, queue.Append("2"))
	assert.False(t, queue.Append("3"))

	assert.True(t, queue.container.Len() == 2)
	assert.Equal(t, queue.container.Front().Value, "1")
	assert.Equal(t, queue.container.Back().Value, "2")
}

func TestDequePrepend(t *testing.T) {
	queue := NewDeque()

	assert.True(t, queue.Prepend("1"))
	assert.True(t, queue.Prepend("2"))
	assert.True(t, queue.Prepend("3"))

	assert.True(t, queue.container.Len() == 3)
	assert.Equal(t, queue.container.Front().Value, "3")
	assert.Equal(t, queue.container.Back().Value, "1")
}

func TestDequePrependWithCapacity(t *testing.T) {
	queue := NewCappedDeque(2)

	assert.True(t, queue.Prepend("1"))
	assert.True(t, queue.Prepend("2"))
	assert.False(t, queue.Prepend("3"))

	assert.True(t, queue.container.Len() == 2)
	assert.Equal(t, queue.container.Front().Value, "2")
	assert.Equal(t, queue.container.Back().Value, "1")
}

func TestDequePop_fulfilled_container(t *testing.T) {
	queue := NewDeque()

	queue.Append("1")
	queue.Append("2")
	queue.Append("3")

	item := queue.Pop()
	assert.Equal(t, item, "3")
	assert.Equal(t, queue.container.Len(), 2)
	item = queue.Pop()
	assert.Equal(t, item, "2")
	assert.Equal(t, queue.container.Len(), 1)
	item = queue.Pop()
	assert.Equal(t, item, "1")
	assert.Equal(t, queue.container.Len(), 0)
}

func TestDequePop_empty_container(t *testing.T) {
	queue := NewDeque()

	item := queue.Pop()
	assert.Equal(t, item, nil)
	assert.Equal(t, queue.container.Len(), 0)
}

func TestDequeShift_fulfilled_container(t *testing.T) {
	queue := NewDeque()

	queue.Append("1")
	queue.Append("2")
	queue.Append("3")

	item := queue.Shift()
	assert.Equal(t, item, "1")
	assert.Equal(t, queue.container.Len(), 2)
	item = queue.Shift()
	assert.Equal(t, item, "2")
	assert.Equal(t, queue.container.Len(), 1)
	item = queue.Shift()
	assert.Equal(t, item, "3")
	assert.Equal(t, queue.container.Len(), 0)
}

func TestDequeShift_empty_container(t *testing.T) {
	queue := NewDeque()

	item := queue.Shift()
	assert.Equal(t, item, nil)
	assert.Equal(t, queue.container.Len(), 0)
}

func TestDequeFirst_fulfilled_container(t *testing.T) {
	queue := NewDeque()

	queue.Append("1")
	queue.Append("2")
	queue.Append("3")

	item := queue.First()
	assert.Equal(t, item, "1")
	assert.Equal(t, queue.container.Len(), 3)
}

func TestDequeFirst_empty_container(t *testing.T) {
	queue := NewDeque()

	item := queue.First()
	assert.Equal(t, item, nil)
	assert.Equal(t, queue.container.Len(), 0)
}

func TestDequeLast_fulfilled_container(t *testing.T) {
	queue := NewDeque()

	queue.Append("1")
	queue.Append("2")
	queue.Append("3")

	item := queue.Last()
	assert.Equal(t, item, "3")
	assert.Equal(t, queue.container.Len(), 3)
}

func TestDequeLast_empty_container(t *testing.T) {
	queue := NewDeque()

	item := queue.Last()
	assert.Equal(t, item, nil)
	assert.Equal(t, queue.container.Len(), 0)
}

func TestDequeEmpty_fulfilled(t *testing.T) {
	queue := NewDeque()

	queue.Append("1")
	queue.Append("2")
	queue.Append("3")

	assert.False(t, queue.Empty())
}

func TestDequeEmpty_empty_queue(t *testing.T) {
	queue := NewDeque()
	assert.True(t, queue.Empty())
}

func TestDequeFull_fulfilled(t *testing.T) {
	queue := NewCappedDeque(3)

	queue.Append("1")
	queue.Append("2")
	queue.Append("3")

	assert.True(t, queue.Full())
}

func TestDequeFull_full_queue(t *testing.T) {
	queue := NewCappedDeque(1)

	queue.Append("1")

	assert.True(t, queue.Full())
}
