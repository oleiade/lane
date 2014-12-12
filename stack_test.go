package lane

import (
	"strconv"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := NewStack()
	stackSize := 100

	// Fulfill the test Stack
	for i := 0; i < stackSize; i++ {
		var value string = strconv.Itoa(i)
		stack.Push(value)
	}

	assert(
		t,
		stack.container.Len() == stackSize,
		"stack.container.Len() = %d; want %d", stack.container.Len(), stackSize,
	)

	assert(
		t,
		stack.container.Front().Value == "99",
		"stack.container.Front().Value = %s; want %s", stack.container.Front().Value, "99",
	)

	assert(
		t,
		stack.container.Back().Value == "0",
		"stack.container.Back().Value = %s; want %s", stack.container.Back().Value, "0",
	)
}

func TestStackPop_fulfilled(t *testing.T) {
	stack := NewStack()
	stackSize := 100

	// Add elements to the test Stack
	for i := 0; i < stackSize; i++ {
		var value string = strconv.Itoa(i)
		stack.Push(value)
	}

	// Pop elements and assert they come out in LIFO order
	for i := 99; i >= 0; i-- {
		var expectedValue string = strconv.Itoa(i)
		item := stack.Pop()

		assert(
			t,
			item == expectedValue,
			"stack.Pop() = %v; want %v", item, expectedValue,
		)

		assert(
			t,
			stack.container.Len() == i,
			"stack.container.Len() = %d; want %d", stack.container.Len(), i,
		)
	}
}

func TestStackPop_empty(t *testing.T) {
	stack := NewStack()

	item := stack.Pop()
	assert(
		t,
		item == nil,
		"stack.Pop() = %v; want %v", item, nil,
	)

	assert(
		t,
		stack.container.Len() == 0,
		"stack.container.Len() = %d; want %d", stack.container.Len(), 0,
	)
}

func TestStackHead_fulfilled(t *testing.T) {
	stack := NewStack()

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	item := stack.Head()
	assert(
		t,
		item == "3",
		"stack.Head() = %v; want %v", item, "3",
	)

	assert(
		t,
		stack.container.Len() == 3,
		"stack.container.Len() = %d; want %d", stack.container.Len(), 3,
	)
}

func TestStackHead_empty(t *testing.T) {
	stack := NewStack()

	item := stack.Head()
	assert(
		t,
		item == nil,
		"stack.Head() = %v; want %v", item, nil,
	)

	assert(
		t,
		stack.container.Len() == 0,
		"stack.container.Len() = %d; want %d", stack.container.Len(), 0,
	)
}

func TestStackEmpty_fulfilled(t *testing.T) {
	stack := NewStack()
	stack.Push("1")
	assert(
		t,
		stack.Empty() == false,
		"stack.Empty() = %b; want %b", stack.Empty(), false,
	)
}

func TestStackEmpty_empty_queue(t *testing.T) {
	stack := NewStack()
	assert(
		t,
		stack.Empty() == true,
		"stack.Empty() = %b; want %b", stack.Empty(), true,
	)
}
