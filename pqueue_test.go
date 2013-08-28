package lane

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxPQueue_init(t *testing.T) {
	pqueue := NewPQueue(MAXPQ)

	assert.Equal(t, len(pqueue.items), 1)
	assert.Equal(t, pqueue.Size(), 0)
	assert.Nil(t, pqueue.items[0])
	assert.Equal(t, pqueue.comparator, max)
}

func TestMinPQueue_init(t *testing.T) {
	pqueue := NewPQueue(MINPQ)

	assert.Equal(t, len(pqueue.items), 1)
	assert.Equal(t, pqueue.Size(), 0)
	assert.Nil(t, pqueue.items[0])
	assert.Equal(t, pqueue.comparator, min)
}

func TestMaxPQueuePush_protects_max_order(t *testing.T) {
	pqueue := NewPQueue(MAXPQ)

	pqueue.Push("1", 1)
	pqueue.Push("2", 2)
	pqueue.Push("3", 3)

	// Heap queue index starts at one, zero index should always be nil
	firstItem := pqueue.items[1]
	firstItemValue, ok := firstItem.value.(string)
	assert.True(t, ok)
	assert.Equal(t, firstItemValue, "3")
	assert.Equal(t, firstItem.priority, 3)

	firstItem = pqueue.items[2]
	firstItemValue, ok = firstItem.value.(string)
	assert.True(t, ok)
	assert.Equal(t, firstItemValue, "1")
	assert.Equal(t, firstItem.priority, 1)

	firstItem = pqueue.items[3]
	firstItemValue, ok = firstItem.value.(string)
	assert.True(t, ok)
	assert.Equal(t, firstItemValue, "2")
	assert.Equal(t, firstItem.priority, 2)
}

func TestMinPQueuePush_protects_min_order(t *testing.T) {
	pqueue := NewPQueue(MINPQ)

	pqueue.Push("1", 1)
	pqueue.Push("2", 2)
	pqueue.Push("3", 3)

	// Heap queue index starts at one, zero index should always be nil
	firstItem := pqueue.items[1]
	firstItemValue, ok := firstItem.value.(string)
	assert.True(t, ok)
	assert.Equal(t, firstItemValue, "1")
	assert.Equal(t, firstItem.priority, 1)

	firstItem = pqueue.items[2]
	firstItemValue, ok = firstItem.value.(string)
	assert.True(t, ok)
	assert.Equal(t, firstItemValue, "2")
	assert.Equal(t, firstItem.priority, 2)

	firstItem = pqueue.items[3]
	firstItemValue, ok = firstItem.value.(string)
	assert.True(t, ok)
	assert.Equal(t, firstItemValue, "3")
	assert.Equal(t, firstItem.priority, 3)
}

func TestMaxPQueuePop_protects_max_order(t *testing.T) {
	pqueue := NewPQueue(MAXPQ)

	pqueue.Push("1", 1)
	pqueue.Push("2", 2)
	pqueue.Push("3", 3)

	value, priority := pqueue.Pop()
	castedValue, ok := value.(string)
	assert.True(t, ok)
	assert.Equal(t, castedValue, "3")
	assert.Equal(t, priority, 3)

	value, priority = pqueue.Pop()
	castedValue, ok = value.(string)
	assert.True(t, ok)
	assert.Equal(t, castedValue, "2")
	assert.Equal(t, priority, 2)

	value, priority = pqueue.Pop()
	castedValue, ok = value.(string)
	assert.True(t, ok)
	assert.Equal(t, castedValue, "1")
	assert.Equal(t, priority, 1)
}

func TestMinPQueuePop_protects_min_order(t *testing.T) {
	pqueue := NewPQueue(MINPQ)

	pqueue.Push("1", 1)
	pqueue.Push("2", 2)
	pqueue.Push("3", 3)

	value, priority := pqueue.Pop()
	castedValue, ok := value.(string)
	assert.True(t, ok)
	assert.Equal(t, castedValue, "1")
	assert.Equal(t, priority, 1)

	value, priority = pqueue.Pop()
	castedValue, ok = value.(string)
	assert.True(t, ok)
	assert.Equal(t, castedValue, "2")
	assert.Equal(t, priority, 2)

	value, priority = pqueue.Pop()
	castedValue, ok = value.(string)
	assert.True(t, ok)
	assert.Equal(t, castedValue, "3")
	assert.Equal(t, priority, 3)
}

func TestMaxPQueueHead_returns_max_element(t *testing.T) {
	pqueue := NewPQueue(MAXPQ)

	pqueue.Push("1", 1)
	pqueue.Push("2", 2)
	pqueue.Push("3", 3)

	value, priority := pqueue.Head()
	castedValue, ok := value.(string)
	assert.True(t, ok)
	assert.Equal(t, castedValue, "3")
	assert.Equal(t, priority, 3)
	assert.Equal(t, pqueue.Size(), 3)
}

func TestMinPQueueHead_returns_min_element(t *testing.T) {
	pqueue := NewPQueue(MINPQ)

	pqueue.Push("1", 1)
	pqueue.Push("2", 2)
	pqueue.Push("3", 3)

	value, priority := pqueue.Head()
	castedValue, ok := value.(string)
	assert.True(t, ok)
	assert.Equal(t, castedValue, "1")
	assert.Equal(t, priority, 1)
	assert.Equal(t, pqueue.Size(), 3)
}
