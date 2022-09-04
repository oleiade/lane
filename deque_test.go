package lane

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeAppend(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                   string
		deque                  *Deque[int]
		appendValue            int
		wantContainerLen       uint
		wantContainerBack      bool
		wantContainerBackValue int
	}{
		{
			desc:                   "append to empty deque inserts value",
			deque:                  NewDeque[int](),
			appendValue:            42,
			wantContainerLen:       1,
			wantContainerBack:      true,
			wantContainerBackValue: 42,
		},
		{
			desc:                   "append inserts value at the back",
			deque:                  NewDeque([]int{40, 41}...),
			appendValue:            42,
			wantContainerLen:       3,
			wantContainerBack:      true,
			wantContainerBackValue: 42,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			tC.deque.Append(tC.appendValue)
			gotContainerLen := tC.deque.container.Len()
			gotContainerBack := tC.deque.container.Back()

			assert.Equal(t, tC.wantContainerLen, gotContainerLen)
			assert.Equal(t, tC.wantContainerBack, gotContainerBack != nil)

			if tC.wantContainerBack {
				assert.Equal(t, tC.wantContainerBackValue, gotContainerBack.Value)
			}
		})
	}
}

func BenchmarkDequeAppend(b *testing.B) {
	b.ReportAllocs()

	deque := NewDeque[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.Append(i)
	}
}

func TestDequePrepend(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                    string
		deque                   *Deque[int]
		prependValue            int
		wantContainerLen        uint
		wantContainerFront      bool
		wantContainerFrontValue int
	}{
		{
			desc:                    "prepend to empty deque",
			deque:                   NewDeque[int](),
			prependValue:            42,
			wantContainerLen:        1,
			wantContainerFront:      true,
			wantContainerFrontValue: 42,
		},
		{
			desc:                    "prepend inserts value at the back",
			deque:                   NewDeque([]int{43, 44}...),
			prependValue:            42,
			wantContainerLen:        3,
			wantContainerFront:      true,
			wantContainerFrontValue: 42,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			tC.deque.Prepend(tC.prependValue)
			gotContainerLen := tC.deque.container.Len()
			gotContainerFront := tC.deque.container.Front()

			assert.Equal(t, tC.wantContainerLen, gotContainerLen)
			assert.Equal(t, tC.wantContainerFront, gotContainerFront != nil)

			if tC.wantContainerFront {
				assert.Equal(t, tC.wantContainerFrontValue, gotContainerFront.Value)
			}
		})
	}
}

func BenchmarkDequePrepend(b *testing.B) {
	b.ReportAllocs()

	deque := NewDeque[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.Prepend(i)
	}
}

func TestDequePop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                   string
		deque                  *Deque[int]
		wantOk                 bool
		wantValue              int
		wantContainerLen       uint
		wantContainerBack      bool
		wantContainerBackValue int
	}{
		{
			desc:              "Pop from an empty Deque",
			deque:             NewDeque[int](),
			wantOk:            false,
			wantValue:         0,
			wantContainerLen:  0,
			wantContainerBack: false,
		},
		{
			desc:                   "Pop removes and returns the back value",
			deque:                  NewDeque([]int{40, 41, 42}...),
			wantOk:                 true,
			wantValue:              42,
			wantContainerLen:       2,
			wantContainerBack:      true,
			wantContainerBackValue: 41,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotValue, gotOk := tC.deque.Pop()
			gotContainerLen := tC.deque.container.Len()
			gotContainerBack := tC.deque.container.Back()

			assert.Equal(t, tC.wantOk, gotOk)
			assert.Equal(t, tC.wantValue, gotValue)
			assert.Equal(t, tC.wantContainerLen, gotContainerLen)
			assert.Equal(t, tC.wantContainerBack, gotContainerBack != nil)

			if tC.wantContainerBack {
				assert.Equal(t, tC.wantContainerBackValue, gotContainerBack.Value)
			}
		})
	}
}

func BenchmarkDequePop(b *testing.B) {
	b.ReportAllocs()

	deque := NewDeque[int]()

	for i := 0; i < b.N; i++ {
		deque.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.Pop()
	}
}

func TestDequeShift(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                    string
		deque                   *Deque[int]
		wantOk                  bool
		wantValue               int
		wantContainerLen        uint
		wantContainerFront      bool
		wantContainerFrontValue int
	}{
		{
			desc:               "Shift from an empty Deque",
			deque:              NewDeque[int](),
			wantOk:             false,
			wantValue:          0,
			wantContainerLen:   0,
			wantContainerFront: false,
		},
		{
			desc:                    "Shift removes and returns the front value",
			deque:                   NewDeque([]int{42, 43, 44}...),
			wantOk:                  true,
			wantValue:               42,
			wantContainerLen:        2,
			wantContainerFront:      true,
			wantContainerFrontValue: 43,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotValue, gotOk := tC.deque.Shift()
			gotContainerLen := tC.deque.container.Len()
			gotContainerBack := tC.deque.container.Front()

			assert.Equal(t, tC.wantOk, gotOk)
			assert.Equal(t, tC.wantValue, gotValue)
			assert.Equal(t, tC.wantContainerLen, gotContainerLen)
			assert.Equal(t, tC.wantContainerFront, gotContainerBack != nil)

			if tC.wantContainerFront {
				assert.Equal(t, tC.wantContainerFrontValue, gotContainerBack.Value)
			}
		})
	}
}

func BenchmarkDequeShift(b *testing.B) {
	b.ReportAllocs()

	deque := NewDeque[int]()

	for i := 0; i < b.N; i++ {
		deque.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.Shift()
	}
}

func TestDequeFirst(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                    string
		deque                   *Deque[int]
		wantOk                  bool
		wantValue               int
		wantContainerLen        uint
		wantContainerFront      bool
		wantContainerFrontValue int
	}{
		{
			desc:               "First from an empty Deque",
			deque:              NewDeque[int](),
			wantOk:             false,
			wantValue:          0,
			wantContainerLen:   0,
			wantContainerFront: false,
		},
		{
			desc:                    "First returns the front value",
			deque:                   NewDeque([]int{42, 43, 44}...),
			wantOk:                  true,
			wantValue:               42,
			wantContainerLen:        3,
			wantContainerFront:      true,
			wantContainerFrontValue: 42,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotValue, gotOk := tC.deque.First()
			gotContainerLen := tC.deque.container.Len()
			gotContainerBack := tC.deque.container.Front()

			assert.Equal(t, tC.wantOk, gotOk)
			assert.Equal(t, tC.wantValue, gotValue)
			assert.Equal(t, tC.wantContainerLen, gotContainerLen)
			assert.Equal(t, tC.wantContainerFront, gotContainerBack != nil)

			if tC.wantContainerFront {
				assert.Equal(t, tC.wantContainerFrontValue, gotContainerBack.Value)
			}
		})
	}
}

func BenchmarkDequeFirst(b *testing.B) {
	b.ReportAllocs()

	deque := NewDeque[int]()

	for i := 0; i < b.N; i++ {
		deque.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.First()
	}
}

func TestDequeLast(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                   string
		deque                  *Deque[int]
		wantOk                 bool
		wantValue              int
		wantContainerLen       uint
		wantContainerBack      bool
		wantContainerBackValue int
	}{
		{
			desc:              "Last from an empty Deque",
			deque:             NewDeque[int](),
			wantOk:            false,
			wantValue:         0,
			wantContainerLen:  0,
			wantContainerBack: false,
		},
		{
			desc:                   "Last returns the front value",
			deque:                  NewDeque([]int{40, 41, 42}...),
			wantOk:                 true,
			wantValue:              42,
			wantContainerLen:       3,
			wantContainerBack:      true,
			wantContainerBackValue: 42,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotValue, gotOk := tC.deque.Last()
			gotContainerLen := tC.deque.container.Len()
			gotContainerBack := tC.deque.container.Back()

			assert.Equal(t, tC.wantOk, gotOk)
			assert.Equal(t, tC.wantValue, gotValue)
			assert.Equal(t, tC.wantContainerLen, gotContainerLen)
			assert.Equal(t, tC.wantContainerBack, gotContainerBack != nil)

			if tC.wantContainerBack {
				assert.Equal(t, tC.wantContainerBackValue, gotContainerBack.Value)
			}
		})
	}
}

func BenchmarkDequeLast(b *testing.B) {
	b.ReportAllocs()

	deque := NewDeque[int]()

	for i := 0; i < b.N; i++ {
		deque.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.Last()
	}
}

func TestDequeSize(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc             string
		deque            *Deque[int]
		wantValue        uint
		wantContainerLen uint
	}{
		{
			desc:             "Size of an empty Deque",
			deque:            NewDeque[int](),
			wantValue:        0,
			wantContainerLen: 0,
		},
		{
			desc:             "Size of a filled Deque",
			deque:            NewDeque([]int{40, 41, 42}...),
			wantValue:        3,
			wantContainerLen: 3,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotValue := tC.deque.Size()
			gotContainerLen := tC.deque.container.Len()

			assert.Equal(t, tC.wantValue, gotValue)
			assert.Equal(t, tC.wantContainerLen, gotContainerLen)
		})
	}
}

func BenchmarkDequeSize(b *testing.B) {
	b.ReportAllocs()

	deque := NewDeque[int]()

	for i := 0; i < b.N; i++ {
		deque.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.Size()
	}
}

func TestDequeEmpty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc             string
		deque            *Deque[int]
		wantValue        bool
		wantContainerLen uint
	}{
		{
			desc:             "Empty of an empty Deque",
			deque:            NewDeque[int](),
			wantValue:        true,
			wantContainerLen: 0,
		},
		{
			desc:             "Empty of a filled Deque",
			deque:            NewDeque([]int{40, 41, 42}...),
			wantValue:        false,
			wantContainerLen: 3,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotValue := tC.deque.Empty()
			gotContainerLen := tC.deque.container.Len()

			assert.Equal(t, tC.wantValue, gotValue)
			assert.Equal(t, tC.wantContainerLen, gotContainerLen)
		})
	}
}

func BenchmarkDequeEmpty(b *testing.B) {
	b.ReportAllocs()

	deque := NewDeque[int]()

	for i := 0; i < b.N; i++ {
		deque.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.Empty()
	}
}

func TestBoundDequeFull(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		deque    *BoundDeque[int]
		wantFull bool
	}{
		{
			desc:     "Empty BoundDeque with non null capacity is not full",
			deque:    NewBoundDeque[int](1),
			wantFull: false,
		},
		{
			desc:     "Empty BoundDeque with null capacity is full",
			deque:    NewBoundDeque[int](0),
			wantFull: true,
		},
		{
			desc:     "Non empty BoundDeque with non null, capacity and available space is not full",
			deque:    NewBoundDeque(4, []int{40, 41, 42}...),
			wantFull: false,
		},
		{
			desc:     "Non empty BoundDeque with non null, capacity and no available space is full",
			deque:    NewBoundDeque(3, []int{40, 41, 42}...),
			wantFull: true,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotFull := tC.deque.Full()

			assert.Equal(t, tC.wantFull, gotFull)
		})
	}
}

func BenchmarkBoundDequeFull(b *testing.B) {
	b.ReportAllocs()

	deque := NewBoundDeque[int](1)

	for i := 0; i < b.N; i++ {
		deque.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deque.Full()
	}
}

// Considering BoundDeque embeds a Deque, no need to cover general
// cases that are not specifically related to capacity management.
func TestBoundDequeAppend(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc   string
		deque  *BoundDeque[int]
		wantOk bool
	}{
		{
			desc:   "Append to BoundDeque with non null capacity and available space",
			deque:  NewBoundDeque[int](1),
			wantOk: true,
		},
		{
			desc:   "Append to BoundDeque with non null capacity and available space",
			deque:  NewBoundDeque(1, []int{42}...),
			wantOk: false,
		},
		{
			desc:   "Append to BoundDeque with null capacity",
			deque:  NewBoundDeque[int](0),
			wantOk: false,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotOk := tC.deque.Append(42)

			assert.Equal(t, tC.wantOk, gotOk)
		})
	}
}

func BenchmarkBoundDequeAppend(b *testing.B) {
	b.ReportAllocs()

	deque := NewBoundDeque[int](1)

	for i := 0; i < b.N; i++ {
		deque.Append(i)
	}
}

// Considering BoundDeque embeds a Deque, no need to cover general
// cases that are not specifically related to capacity management.
func TestBoundDequePrepend(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc   string
		deque  *BoundDeque[int]
		wantOk bool
	}{
		{
			desc:   "Prepend to BoundDeque with non null capacity and available space",
			deque:  NewBoundDeque[int](1),
			wantOk: true,
		},
		{
			desc:   "Prepend to BoundDeque with non null capacity and available space",
			deque:  NewBoundDeque(1, []int{42}...),
			wantOk: false,
		},
		{
			desc:   "Prepend to BoundDeque with null capacity",
			deque:  NewBoundDeque[int](0),
			wantOk: false,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotOk := tC.deque.Prepend(42)

			assert.Equal(t, tC.wantOk, gotOk)
		})
	}
}

func BenchmarkBoundDeque(b *testing.B) {
	b.ReportAllocs()

	deque := NewBoundDeque[int](1)

	for i := 0; i < b.N; i++ {
		deque.Prepend(i)
	}
}
