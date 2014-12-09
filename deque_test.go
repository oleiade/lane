package lane

import (
	"strconv"
	"testing"
)

func TestDequeAppend(t *testing.T) {
	deque := NewDeque()
	sampleSize := 100

	// Append elements in the Deque and assert it does not fail
	for i := 0; i < sampleSize; i++ {
		var value string = strconv.Itoa(i)
		var ok bool = deque.Append(value)

		assert(
			t,
			ok == true,
			"deque.Append(%d) = %t; want %t", i, ok, true,
		)
	}

	assert(
		t,
		deque.container.Len() == sampleSize,
		"deque.container.Len() = %d; want %d", deque.container.Len(), sampleSize,
	)

	assert(
		t,
		deque.container.Front().Value == "0",
		"deque.container.Front().Value = %s; want %s", deque.container.Front().Value, "0",
	)

	assert(
		t,
		deque.container.Back().Value == "99",
		"deque.container.Back().Value = %s; want %s", deque.container.Back().Value, "99",
	)
}

func TestDequeAppendWithCapacity(t *testing.T) {
	dequeSize := 20
	deque := NewCappedDeque(dequeSize)

	// Append the maximum number of elements in the Deque
	// and assert it does not fail
	for i := 0; i < dequeSize; i++ {
		var value string = strconv.Itoa(i)
		var ok bool = deque.Append(value)

		assert(
			t,
			ok == true,
			"deque.Append(%d) = %t; want %t", i, ok, true,
		)
	}

	// Try to overflow the Deque size limit, and make
	// sure appending fails
	var ok bool = deque.Append("should not be ok")
	assert(
		t,
		ok == false,
		"deque.Append(%s) = %t; want %t", "should not be ok", ok, false,
	)

	assert(
		t,
		deque.container.Len() == dequeSize,
		"deque.container.Len() = %d; want %d", deque.container.Len(), dequeSize,
	)

	assert(
		t,
		deque.container.Front().Value == "0",
		"deque.container.Front().Value = %s; want %s", deque.container.Front().Value, "0",
	)

	assert(
		t,
		deque.container.Back().Value == "19",
		"deque.container.Back().Value = %s; want %s", deque.container.Back().Value, "19",
	)
}

func TestDequePrepend(t *testing.T) {
	deque := NewDeque()
	sampleSize := 100

	// Prepend elements in the Deque and assert it does not fail
	for i := 0; i < sampleSize; i++ {
		var value string = strconv.Itoa(i)
		var ok bool = deque.Prepend(value)

		assert(
			t,
			ok == true,
			"deque.Prepend(%d) = %t; want %t", i, ok, true,
		)
	}

	assert(
		t,
		deque.container.Len() == sampleSize,
		"deque.container.Len() = %d; want %d", deque.container.Len(), sampleSize,
	)

	assert(
		t,
		deque.container.Front().Value == "99",
		"deque.container.Front().Value = %s; want %s", deque.container.Front().Value, "99",
	)

	assert(
		t,
		deque.container.Back().Value == "0",
		"deque.container.Back().Value = %s; want %s", deque.container.Back().Value, "0",
	)
}

func TestDequePrependWithCapacity(t *testing.T) {
	dequeSize := 20
	deque := NewCappedDeque(dequeSize)

	// Prepend elements in the Deque and assert it does not fail
	for i := 0; i < dequeSize; i++ {
		var value string = strconv.Itoa(i)
		var ok bool = deque.Prepend(value)

		assert(
			t,
			ok == true,
			"deque.Prepend(%d) = %t; want %t", i, ok, true,
		)
	}

	// Try to overflow the Deque size limit, and make
	// sure appending fails
	var ok bool = deque.Prepend("should not be ok")
	assert(
		t,
		ok == false,
		"deque.Prepend(%s) = %t; want %t", "should not be ok", ok, false,
	)

	assert(
		t,
		deque.container.Len() == dequeSize,
		"deque.container.Len() = %d; want %d", deque.container.Len(), dequeSize,
	)

	assert(
		t,
		deque.container.Front().Value == "19",
		"deque.container.Front().Value = %s; want %s", deque.container.Front().Value, "19",
	)

	assert(
		t,
		deque.container.Back().Value == "0",
		"deque.container.Back().Value = %s; want %s", deque.container.Back().Value, "0",
	)
}

func TestDequePop_fulfilled_container(t *testing.T) {
	deque := NewDeque()
	dequeSize := 100

	// Populate the test deque
	for i := 0; i < dequeSize; i++ {
		var value string = strconv.Itoa(i)
		deque.Append(value)
	}

	// Pop elements of the deque and assert elements come out
	// in order and container size is updated accordingly
	for i := dequeSize - 1; i >= 0; i-- {
		item := deque.Pop()

		var itemValue string = item.(string)
		var expectedValue string = strconv.Itoa(i)

		assert(
			t,
			itemValue == expectedValue,
			"deque.Pop() = %s; want %s", itemValue, expectedValue,
		)

		assert(
			t,
			deque.container.Len() == i,
			"deque.container.Len() = %d; want %d", deque.container.Len(), i,
		)

	}
}

func TestDequePop_empty_container(t *testing.T) {
	deque := NewDeque()
	item := deque.Pop()

	assert(
		t,
		item == nil,
		"item = %v; want %v", item, nil,
	)

	assert(
		t,
		deque.container.Len() == 0,
		"deque.container.Len() = %d; want %d", deque.container.Len(), 0,
	)
}

func TestDequeShift_fulfilled_container(t *testing.T) {
	deque := NewDeque()
	dequeSize := 100

	// Populate the test deque
	for i := 0; i < dequeSize; i++ {
		var value string = strconv.Itoa(i)
		deque.Append(value)
	}

	// Pop elements of the deque and assert elements come out
	// in order and container size is updated accordingly
	for i := 0; i < dequeSize; i++ {
		item := deque.Shift()

		var itemValue string = item.(string)
		var expectedValue string = strconv.Itoa(i)

		assert(
			t,
			itemValue == expectedValue,
			"deque.Shift() = %s; want %s", itemValue, expectedValue,
		)

		assert(
			t,
			// Len should be equal to dequeSize - (i + 1) as i is zero indexed
			deque.container.Len() == (dequeSize-(i+1)),
			"deque.container.Len() = %d; want %d", deque.container.Len(), dequeSize-i,
		)
	}
}

func TestDequeShift_empty_container(t *testing.T) {
	deque := NewDeque()

	item := deque.Shift()
	assert(
		t,
		item == nil,
		"deque.Shift() = %v; want %v", item, nil,
	)

	assert(
		t,
		deque.container.Len() == 0,
		"deque.container.Len() = %d; want %d", deque.container.Len(), 0,
	)
}

func TestDequeFirst_fulfilled_container(t *testing.T) {
	deque := NewDeque()
	deque.Append("1")
	item := deque.First()

	assert(
		t,
		item == "1",
		"deque.First() = %s; want %s", item, "1",
	)

	assert(
		t,
		deque.container.Len() == 1,
		"deque.container.Len() = %d; want %d", deque.container.Len(), 1,
	)
}

func TestDequeFirst_empty_container(t *testing.T) {
	deque := NewDeque()
	item := deque.First()

	assert(
		t,
		item == nil,
		"deque.First() = %v; want %v", item, nil,
	)

	assert(
		t,
		deque.container.Len() == 0,
		"deque.container.Len() = %d; want %d", deque.container.Len(), 0,
	)
}

func TestDequeLast_fulfilled_container(t *testing.T) {
	deque := NewDeque()

	deque.Append("1")
	deque.Append("2")
	deque.Append("3")

	item := deque.Last()

	assert(
		t,
		item == "3",
		"deque.Last() = %s; want %s", item, "3",
	)

	assert(
		t,
		deque.container.Len() == 3,
		"deque.container.Len() = %d; want %d", deque.container.Len(), 3,
	)
}

func TestDequeLast_empty_container(t *testing.T) {
	deque := NewDeque()
	item := deque.Last()

	assert(
		t,
		item == nil,
		"deque.Last() = %v; want %v", item, nil,
	)

	assert(
		t,
		deque.container.Len() == 0,
		"deque.container.Len() = %d; want %d", deque.container.Len(), 0,
	)
}

func TestDequeEmpty_fulfilled(t *testing.T) {
	deque := NewDeque()
	deque.Append("1")

	assert(
		t,
		deque.Empty() == false,
		"deque.Empty() = %t; want %t", deque.Empty(), false)
}

func TestDequeEmpty_empty_deque(t *testing.T) {
	deque := NewDeque()
	assert(
		t,
		deque.Empty() == true,
		"deque.Empty() = %t; want %t", deque.Empty(), true,
	)
}

func TestDequeFull_fulfilled(t *testing.T) {
	deque := NewCappedDeque(3)

	deque.Append("1")
	deque.Append("2")
	deque.Append("3")

	assert(
		t,
		deque.Full() == true,
		"deque.Full() = %t; want %t", deque.Full(), true,
	)
}

func TestDequeFull_non_full_deque(t *testing.T) {
	deque := NewCappedDeque(3)
	deque.Append("1")

	assert(
		t,
		deque.Full() == false,
		"deque.Full() = %t; want %t", deque.Full(), false,
	)
}
