package lane

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                   string
		items                  []int
		wantContainerLastOk    bool
		wantContainerLastValue int
		wantContainerSize      uint
	}{
		{
			desc:                "NewQueue initializes a Queue",
			wantContainerLastOk: false,
			wantContainerSize:   0,
		},
		{
			desc:                   "NewQueue with initializer produces FIFO ordering",
			items:                  []int{1, 2, 3},
			wantContainerLastOk:    true,
			wantContainerLastValue: 1,
			wantContainerSize:      3,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotQueue := NewQueue(tC.items...)
			gotContainerFirstValue, gotContainerFirstOk := gotQueue.container.Last()
			gotContainerSize := gotQueue.container.Size()

			assert.Equal(t, tC.wantContainerLastOk, gotContainerFirstOk)
			assert.Equal(t, tC.wantContainerLastValue, gotContainerFirstValue)
			assert.Equal(t, tC.wantContainerSize, gotContainerSize)
		})
	}
}

func BenchmarkNewQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewQueue[int]()
	}
}

func TestQueueEnqueue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                   string
		queue                  *Queue[int]
		enqueueValue           int
		wantContainerSize      uint
		wantContainerFirst     bool
		wantContainerLastValue int
	}{
		{
			desc:                   "Enqueue to an empty Queue inserts value",
			queue:                  NewQueue[int](),
			enqueueValue:           42,
			wantContainerSize:      1,
			wantContainerFirst:     true,
			wantContainerLastValue: 42,
		},
		{
			desc:                   "Enqueue inserts value at the head",
			queue:                  NewQueue([]int{41, 40}...),
			enqueueValue:           42,
			wantContainerSize:      3,
			wantContainerFirst:     true,
			wantContainerLastValue: 41,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			tC.queue.Enqueue(tC.enqueueValue)
			gotContainerLen := tC.queue.container.Size()
			gotContainerFirstValue, gotContainerFirstOk := tC.queue.container.Last()

			assert.Equal(t, tC.wantContainerSize, gotContainerLen)
			assert.Equal(t, tC.wantContainerFirst, gotContainerFirstOk)

			if tC.wantContainerFirst {
				assert.Equal(t, tC.wantContainerLastValue, gotContainerFirstValue)
			}
		})
	}
}

func BenchmarkQueueEnqueue(b *testing.B) {
	b.ReportAllocs()

	queue := NewQueue[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func TestQueueDequeue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                   string
		queue                  *Queue[int]
		wantDequeueOk          bool
		wantDequeueValue       int
		wantContainerSize      uint
		wantContainerFirst     bool
		wantContainerLastValue int
	}{
		{
			desc:               "Dequeue on an empty Queue",
			queue:              NewQueue[int](),
			wantDequeueOk:      false,
			wantContainerSize:  0,
			wantContainerFirst: false,
		},
		{
			desc:                   "Dequeue removes and returns the head value",
			queue:                  NewQueue([]int{42, 41, 40}...),
			wantDequeueOk:          true,
			wantDequeueValue:       42,
			wantContainerSize:      2,
			wantContainerFirst:     true,
			wantContainerLastValue: 41,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotDequeueValue, gotDequeueOk := tC.queue.Dequeue()
			gotContainerLen := tC.queue.container.Size()
			gotContainerFirstValue, gotContainerFirstOk := tC.queue.container.Last()

			assert.Equal(t, tC.wantDequeueOk, gotDequeueOk)
			assert.Equal(t, tC.wantDequeueValue, gotDequeueValue)
			assert.Equal(t, tC.wantContainerSize, gotContainerLen)
			assert.Equal(t, tC.wantContainerFirst, gotContainerFirstOk)

			if tC.wantContainerFirst {
				assert.Equal(t, tC.wantContainerLastValue, gotContainerFirstValue)
			}
		})
	}
}

func BenchmarkQueueDequeue(b *testing.B) {
	b.ReportAllocs()

	queue := NewQueue[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}

func TestQueueHead(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                   string
		queue                  *Queue[int]
		wantHeadOk             bool
		wantHeadValue          int
		wantContainerSize      uint
		wantContainerFirst     bool
		wantContainerLastValue int
	}{
		{
			desc:               "Head on an empty Queue",
			queue:              NewQueue[int](),
			wantHeadOk:         false,
			wantContainerSize:  0,
			wantContainerFirst: false,
		},
		{
			desc:                   "Head returns the head value and leaves the Queue untouched",
			queue:                  NewQueue([]int{42, 41, 40}...),
			wantHeadOk:             true,
			wantHeadValue:          42,
			wantContainerSize:      3,
			wantContainerFirst:     true,
			wantContainerLastValue: 42,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotHeadValue, gotHeadOk := tC.queue.Head()
			gotContainerLen := tC.queue.container.Size()
			gotContainerFirstValue, gotContainerFirstOk := tC.queue.container.Last()

			assert.Equal(t, tC.wantHeadOk, gotHeadOk)
			assert.Equal(t, tC.wantHeadValue, gotHeadValue)
			assert.Equal(t, tC.wantContainerSize, gotContainerLen)
			assert.Equal(t, tC.wantContainerFirst, gotContainerFirstOk)

			if tC.wantContainerFirst {
				assert.Equal(t, tC.wantContainerLastValue, gotContainerFirstValue)
			}
		})
	}
}

func BenchmarkQueueHead(b *testing.B) {
	b.ReportAllocs()

	queue := NewQueue[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Head()
	}
}
