package lane

import (
	"constraints"
	"sync"
)

// PQueue is a heap priority queue data structure implementation.
//
// It can be either be minimum (ascending) or maximum (descending) 
// oriented/ordered. Every operations are synchronized and goroutine-safe.
type PriorityQueue[T any] struct {
	sync.RWMutex
	items []*priorityQueueItem[T]
	itemCount uint
	comparator func(lhs, rhs int) bool
}

// NewPriorityQueue instantiates a new PriorityQueue with the provided comparison heuristic.
// The package defines the `Max` and `Min` heuristic to define a maximum-oriented or 
// minimum-oriented heuristic respectively.
func NewPriorityQueue[T any](heuristic func(lhs, rhs int)bool) *PriorityQueue[T] {
	items := make([]*priorityQueueItem[T], 1)
	items[0] = nil
	
	return &PriorityQueue[T]{
		items: items,
		itemCount: 0,
		comparator: heuristic,
	}
}

// NewMaxPriorityQueue instantiates a new maximum oriented PriorityQueue.
func NewMaxPriorityQueue[T any]() *PriorityQueue[T] {
	return NewPriorityQueue[T](Maximum[int])	
}

// NewMinPriorityQueue instantiates a new minimum oriented PriorityQueue.
func NewMinPriorityQueue[T any]() *PriorityQueue[T] {
	return NewPriorityQueue[T](Minimum[int])	
}

// Maximum returns whether `rhs` is greater than `lhs`.
//
// It can be used as a comparison heuristic during a PriorityQueue's
// instantiation.
func Maximum[T constraints.Ordered](lhs, rhs T) bool {
	return lhs < rhs
}

// Minimum returns whether `rhs` is less than `lhs`.
//
// It can be used as a comparison heuristic during a PriorityQueue's
// instantiation.
func Minimum[T constraints.Ordered](lhs, rhs T) bool {
	return lhs > rhs
}

// Push inserts the value in the PriorityQueue with the provided priority
// in at most O(log n) time complexity.
func (pq *PriorityQueue[T]) Push(value T, priority int) {
	item := newPriorityQueueItem(value, priority)

	pq.Lock()
	defer pq.Unlock()
	pq.items = append(pq.items, item)
	pq.itemCount += 1
	pq.swim(pq.size())
}

// Pop and returns the highest or lowest priority item (depending on the
// comparison heuristic of your PriorityQueue) from the PriorityQueue in
// at most O(log n) complexity.
func (pq *PriorityQueue[T]) Pop() (value T, priority int, ok bool) {
	pq.Lock()
	defer pq.Unlock()

	if pq.size() < 1 {
		ok = false
		return
	}

	var max *priorityQueueItem[T] = pq.items[1]
	
	pq.exch(1, pq.size())
	pq.items = pq.items[0:pq.size()]
	pq.itemCount -= 1
	pq.sink(1)
	
	value = max.value
	priority = max.priority
	ok = true
	
	return 
}

// Head returns the highest or lowest priority item (depending on
// the comparison heuristic of your PriorityQueue) from the PriorityQueue
// in O(1) complexity.
func (pq *PriorityQueue[T]) Head() (value T, priority int, ok bool) {
	pq.RLock()
	defer pq.RUnlock()

	if pq.size() < 1 {
		ok = false
		return
	}

	value = pq.items[1].value
	priority = pq.items[1].priority
	ok = true

	return
}

// Size returns the number of elements present in the PriorityQueue.
func (pq *PriorityQueue[T]) Size() uint {
	pq.RLock()
	defer pq.RUnlock()
	return pq.size()
}

// Empty returns whether the PriorityQueue is empty.
func (pq *PriorityQueue[T]) Empty() bool {
	pq.RLock()
	defer pq.RUnlock()
	return pq.size() == 0
}

func (pq *PriorityQueue[T]) swim(k uint) {
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k/2, k)
		k = k / 2
	}
}

func (pq *PriorityQueue[T]) sink(k uint) {
	for uint(2*k) <= pq.size() {
		var j uint = uint(2 * k)

		if j < pq.size() && pq.less(j, j+1) {
			j++
		}

		if !pq.less(k, j) {
			break
		}

		pq.exch(k, j)
		k = j
	}
}

// size is a private method that's not goroutine-safe.
// It is meant to be called by a method who has already
// acquired a lock on the PriorityQueue.
func (pq *PriorityQueue[T]) size() uint {
	return pq.itemCount
}

func (pq *PriorityQueue[T]) less(lhs, rhs uint) bool {
	return pq.comparator(pq.items[lhs].priority, pq.items[rhs].priority)
}

func (pq *PriorityQueue[T]) exch(lhs, rhs uint) {
	var tmp *priorityQueueItem[T] = pq.items[lhs]
	pq.items[lhs] = pq.items[rhs]
	pq.items[rhs] = tmp
}

// priorityQueueItem is the underlying PriorityQueue item container. 
type priorityQueueItem[T any] struct {
	value T
	priority int
}

// newPriorityQueue instantiates a new priorityQueueItem.
func newPriorityQueueItem[T any](value T, priority int) *priorityQueueItem[T] {
	return &priorityQueueItem[T]{
		value: value,
		priority: priority,
	}
} 
