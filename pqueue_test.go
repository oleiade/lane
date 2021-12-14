package lane

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestPriorityQueuePush(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc             string
		heuristic 			 func(lhs, rhs int) bool
		pushItems 			 []*priorityQueueItem[string]
		wantItemCount    uint
		wantItems 			 []*priorityQueueItem[string]
	}{
		{
			desc:             "Push on empty PriorityQueue",
			heuristic: 				Max[int],
			pushItems: 				[]*priorityQueueItem[string]{
				newPriorityQueueItem("a", 1),
			},
			wantItemCount:         1,
			wantItems:				[]*priorityQueueItem[string]{
				nil,
				newPriorityQueueItem("a", 1),
			},
		},
		{
			desc:             "Push on multiple values on max oriented PriorityQueue",
			heuristic: 				Max[int],
			pushItems: 				[]*priorityQueueItem[string]{
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
			wantItemCount:         3,
			wantItems:				[]*priorityQueueItem[string]{
				nil,
				newPriorityQueueItem("c", 3),
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
			},
		},
		{
			desc:             "Push on multiple values on min oriented PriorityQueue",
			heuristic: 				Min[int],
			pushItems: 				[]*priorityQueueItem[string]{
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
			wantItemCount:         3,
			wantItems:				[]*priorityQueueItem[string]{
				nil,
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			pqueue := NewPriorityQueue[string](testCase.heuristic)
			for _, item := range testCase.pushItems {
				pqueue.Push(item.value, item.priority)
			}

			gotItemCount := pqueue.itemCount
			gotItems := pqueue.items

			assert.Equal(t, testCase.wantItemCount, gotItemCount)
			assert.Equal(t, testCase.wantItems, gotItems)
		})
	}
}

func TestPriorityQueuePop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc             string
		heuristic 			 func(lhs, rhs int) bool
		pushItems 			 []*priorityQueueItem[string]
		wantOk 					 bool
		wantValue 			 string
		wantPriority     int
		wantItemCount    uint
		wantItems 			 []*priorityQueueItem[string]
	}{
		{
			desc:             "Pop from an empty PriorityQueue",
			heuristic: 				Max[int],
			pushItems: 				[]*priorityQueueItem[string]{},
			wantOk: 					false,
			wantValue: 				"",
			wantPriority: 		0,
			wantItemCount:    0,
			wantItems:				[]*priorityQueueItem[string]{
				nil,
			},
		},
		{
			desc:             "Pop from a filled max oriented PriorityQueue",
			heuristic: 				Max[int],
			pushItems: 				[]*priorityQueueItem[string]{
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
			wantOk: true,
			wantValue: "c",
			wantPriority: 3,
			wantItemCount:         2,
			wantItems:				[]*priorityQueueItem[string]{
				nil,
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("a", 1),
			},
		},
		{
			desc:             "Pop from a filled min oriented PriorityQueue",
			heuristic: 				Min[int],
			pushItems: 				[]*priorityQueueItem[string]{
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
			wantOk: true,
			wantValue: "a",
			wantPriority: 1,
			wantItemCount:         2,
			wantItems:				[]*priorityQueueItem[string]{
				nil,
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			pqueue := NewPriorityQueue[string](testCase.heuristic)
			for _, item := range testCase.pushItems {
				pqueue.Push(item.value, item.priority)
			}

			gotValue, gotPriority, gotOk := pqueue.Pop()
			gotItemCount := pqueue.itemCount
			gotItems := pqueue.items

			assert.Equal(t, testCase.wantOk, gotOk)
			assert.Equal(t, testCase.wantValue, gotValue)
			assert.Equal(t, testCase.wantPriority, gotPriority)
			assert.Equal(t, testCase.wantItemCount, gotItemCount)
			assert.Equal(t, testCase.wantItems, gotItems)
		})
	}
}

func TestPriorityQueueHead(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc             string
		heuristic 			 func(lhs, rhs int) bool
		pushItems 			 []*priorityQueueItem[string]
		wantOk 					 bool
		wantValue 			 string
		wantPriority     int
		wantItemCount    uint
		wantItems 			 []*priorityQueueItem[string]
	}{
		{
			desc:             "Head of an empty PriorityQueue",
			heuristic: 				Max[int],
			pushItems: 				[]*priorityQueueItem[string]{},
			wantOk: 					false,
			wantValue: 				"",
			wantPriority: 		0,
			wantItemCount:    0,
			wantItems:				[]*priorityQueueItem[string]{
				nil,
			},
		},
		{
			desc:             "Head of a filled max oriented PriorityQueue",
			heuristic: 				Max[int],
			pushItems: 				[]*priorityQueueItem[string]{
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
			wantOk: true,
			wantValue: "c",
			wantPriority: 3,
			wantItemCount: 3,
			wantItems:				[]*priorityQueueItem[string]{
				nil,
				newPriorityQueueItem("c", 3),
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
			},
		},
		{
			desc:             "Head of a filled min oriented PriorityQueue",
			heuristic: 				Min[int],
			pushItems: 				[]*priorityQueueItem[string]{
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
			wantOk: true,
			wantValue: "a",
			wantPriority: 1,
			wantItemCount:         3,
			wantItems:				[]*priorityQueueItem[string]{
				nil,
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			pqueue := NewPriorityQueue[string](testCase.heuristic)
			for _, item := range testCase.pushItems {
				pqueue.Push(item.value, item.priority)
			}

			gotValue, gotPriority, gotOk := pqueue.Head()
			gotItemCount := pqueue.itemCount
			gotItems := pqueue.items

			assert.Equal(t, testCase.wantOk, gotOk)
			assert.Equal(t, testCase.wantValue, gotValue)
			assert.Equal(t, testCase.wantPriority, gotPriority)
			assert.Equal(t, testCase.wantItemCount, gotItemCount)
			assert.Equal(t, testCase.wantItems, gotItems)
		})
	}
}

func TestPriorityQueueSize(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc             string
		pushItems 			 []*priorityQueueItem[string]
		wantValue 			 uint
	}{
		{
			desc:             "Head of an empty PriorityQueue",
			pushItems: 				[]*priorityQueueItem[string]{},
			wantValue: 				0,
		},
		{
			desc:             "Head of a filled max oriented PriorityQueue",
			pushItems: 				[]*priorityQueueItem[string]{
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
			wantValue: 3,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			pqueue := NewMaxPriorityQueue[string]()
			for _, item := range testCase.pushItems {
				pqueue.Push(item.value, item.priority)
			}

			gotValue := pqueue.Size()

			assert.Equal(t, testCase.wantValue, gotValue)
		})
	}
}

func TestPriorityQueueEmpty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc             string
		pushItems 			 []*priorityQueueItem[string]
		wantValue 			 bool
	}{
		{
			desc:             "Head of an empty PriorityQueue",
			pushItems: 				[]*priorityQueueItem[string]{},
			wantValue: 				true,
		},
		{
			desc:             "Head of a filled max oriented PriorityQueue",
			pushItems: 				[]*priorityQueueItem[string]{
				newPriorityQueueItem("a", 1),
				newPriorityQueueItem("b", 2),
				newPriorityQueueItem("c", 3),
			},
			wantValue: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()

			pqueue := NewMaxPriorityQueue[string]()
			for _, item := range testCase.pushItems {
				pqueue.Push(item.value, item.priority)
			}

			gotValue := pqueue.Empty()

			assert.Equal(t, testCase.wantValue, gotValue)
		})
	}
}