package lane

import (
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func TestMaxPQueue_init(t *testing.T) {
	pqueue := NewPQueue(MAXPQ)

	assert(
		t,
		len(pqueue.items) == 1,
		"len(pqueue.items) == %d; want %d", len(pqueue.items), 1,
	)

	assert(
		t,
		pqueue.Size() == 0,
		"pqueue.Size() = %d; want %d", pqueue.Size(), 0,
	)

	assert(
		t,
		pqueue.items[0] == nil,
		"pqueue.items[0] = %v; want %v", pqueue.items[0], nil,
	)

	assert(
		t,
		reflect.ValueOf(pqueue.comparator).Pointer() == reflect.ValueOf(max).Pointer(),
		"pqueue.comparator != max",
	)
}

func TestMinPQueue_init(t *testing.T) {
	pqueue := NewPQueue(MINPQ)

	assert(
		t,
		len(pqueue.items) == 1,
		"len(pqueue.items) = %d; want %d", len(pqueue.items), 1,
	)

	assert(
		t,
		pqueue.Size() == 0,
		"pqueue.Size() = %d; want %d", pqueue.Size(), 0,
	)

	assert(
		t,
		pqueue.items[0] == nil,
		"pqueue.items[0] = %v; want %v", pqueue.items[0], nil,
	)

	assert(
		t,
		reflect.ValueOf(pqueue.comparator).Pointer() == reflect.ValueOf(min).Pointer(),
		"pqueue.comparator != min",
	)
}

func TestMaxPQueuePushAndPop_protects_max_order(t *testing.T) {
	pqueue := NewPQueue(MAXPQ)
	pqueueSize := 100

	// Populate the test priority queue with dummy elements
	// in asc ordered.
	for i := 0; i < pqueueSize; i++ {
		var value string = strconv.Itoa(i)
		var priority int = i

		pqueue.Push(value, priority)
	}

	containerIndex := 1 // binary heap are 1 indexed
	for i := 99; i >= 0; i-- {
		var expectedValue string = strconv.Itoa(i)
		var expectedPriority int = i

		// Avoiding testing arithmetics headaches by using the pop function directly
		value, priority := pqueue.Pop()
		assert(
			t,
			value == expectedValue,
			"value = %v; want %v", containerIndex, value, expectedValue,
		)
		assert(
			t,
			priority == expectedPriority,
			"priority = %v; want %v", containerIndex, priority, expectedValue,
		)

		containerIndex++
	}
}

func TestMaxPQueuePushAndPop_concurrently_protects_max_order(t *testing.T) {
	var wg sync.WaitGroup

	pqueue := NewPQueue(MAXPQ)
	pqueueSize := 100

	// Populate the test priority queue with dummy elements
	// in asc ordered.
	for i := 0; i < pqueueSize; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			var value string = strconv.Itoa(i)
			var priority int = i
			pqueue.Push(value, priority)
		}(i)
	}

	wg.Wait()

	containerIndex := 1 // binary heap are 1 indexed
	for i := 99; i >= 0; i-- {
		var expectedValue string = strconv.Itoa(i)
		var expectedPriority int = i

		// Avoiding testing arithmetics headaches by using the pop function directly
		value, priority := pqueue.Pop()
		assert(
			t,
			value == expectedValue,
			"value = %v; want %v", containerIndex, value, expectedValue,
		)
		assert(
			t,
			priority == expectedPriority,
			"priority = %v; want %v", containerIndex, priority, expectedValue,
		)

		containerIndex++
	}
}

func TestMinPQueuePushAndPop_protects_min_order(t *testing.T) {
	pqueue := NewPQueue(MINPQ)
	pqueueSize := 100

	// Populate the test priority queue with dummy elements
	// in asc ordered.
	for i := 0; i < pqueueSize; i++ {
		var value string = strconv.Itoa(i)
		var priority int = i

		pqueue.Push(value, priority)
	}

	for i := 0; i < pqueueSize; i++ {
		var expectedValue string = strconv.Itoa(i)
		var expectedPriority int = i

		// Avoiding testing arithmetics headaches by using the pop function directly
		value, priority := pqueue.Pop()
		assert(
			t,
			value == expectedValue,
			"value = %v; want %v", value, expectedValue,
		)
		assert(
			t,
			priority == expectedPriority,
			"priority = %v; want %v", priority, expectedValue,
		)
	}
}

func TestMinPQueuePushAndPop_concurrently_protects_min_order(t *testing.T) {
	pqueue := NewPQueue(MINPQ)
	pqueueSize := 100

	var wg sync.WaitGroup

	// Populate the test priority queue with dummy elements
	// in asc ordered.
	for i := 0; i < pqueueSize; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			var value string = strconv.Itoa(i)
			var priority int = i

			pqueue.Push(value, priority)
		}(i)
	}

	wg.Wait()

	for i := 0; i < pqueueSize; i++ {
		var expectedValue string = strconv.Itoa(i)
		var expectedPriority int = i

		// Avoiding testing arithmetics headaches by using the pop function directly
		value, priority := pqueue.Pop()
		assert(
			t,
			value == expectedValue,
			"value = %v; want %v", value, expectedValue,
		)
		assert(
			t,
			priority == expectedPriority,
			"priority = %v; want %v", priority, expectedValue,
		)
	}
}

func TestMaxPQueueHead_returns_max_element(t *testing.T) {
	pqueue := NewPQueue(MAXPQ)

	pqueue.Push("1", 1)
	pqueue.Push("2", 2)

	value, priority := pqueue.Head()

	// First element of the binary heap is always left empty, so container
	// size is the number of elements actually stored + 1
	assert(t, len(pqueue.items) == 3, "len(pqueue.items) = %d; want %d", len(pqueue.items), 3)

	assert(t, value == "2", "pqueue.Head().value = %v; want %v", value, "2")
	assert(t, priority == 2, "pqueue.Head().priority = %d; want %d", priority, 2)
}

func TestMinPQueueHead_returns_min_element(t *testing.T) {
	pqueue := NewPQueue(MINPQ)

	pqueue.Push("1", 1)
	pqueue.Push("2", 2)

	value, priority := pqueue.Head()

	// First element of the binary heap is always left empty, so container
	// size is the number of elements actually stored + 1
	assert(t, len(pqueue.items) == 3, "len(pqueue.items) = %d; want %d", len(pqueue.items), 3)

	assert(t, value == "1", "pqueue.Head().value = %v; want %v", value, "1")
	assert(t, priority == 1, "pqueue.Head().priority = %d; want %d", priority, 1)
}
