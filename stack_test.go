package lane

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := NewStack()

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	assert.True(t, stack.container.Len() == 3)
	assert.Equal(t, stack.container.Front().Value, "3")
	assert.Equal(t, stack.container.Back().Value, "1")
}

func TestStackPop_fulfilled(t *testing.T) {
	stack := NewStack()

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	item := stack.Pop()
	assert.Equal(t, item, "3")
	assert.Equal(t, stack.container.Len(), 2)
	item = stack.Pop()
	assert.Equal(t, item, "2")
	assert.Equal(t, stack.container.Len(), 1)
	item = stack.Pop()
	assert.Equal(t, item, "1")
	assert.Equal(t, stack.container.Len(), 0)
}

func TestStackPop_empty(t *testing.T) {
	stack := NewStack()

	item := stack.Pop()
	assert.Equal(t, item, nil)
	assert.Equal(t, stack.container.Len(), 0)
}

func TestStackHead_fulfilled(t *testing.T) {
	stack := NewStack()

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	item := stack.Head()
	assert.Equal(t, item, "3")
	assert.Equal(t, stack.container.Len(), 3)
}

func TestStackHead_empty(t *testing.T) {
	stack := NewStack()

	item := stack.Head()
	assert.Equal(t, item, nil)
	assert.Equal(t, stack.container.Len(), 0)
}

func TestStackEmpty_fulfilled(t *testing.T) {
	stack := NewStack()

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	assert.False(t, stack.Empty())
}

func TestStackEmpty_empty_queue(t *testing.T) {
	stack := NewStack()
	assert.True(t, stack.Empty())
}
