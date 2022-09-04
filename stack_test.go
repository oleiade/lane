package lane

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                    string
		items                   []int
		wantContainerFirstOk    bool
		wantContainerFirstValue int
		wantContainerSize       uint
	}{
		{
			desc:                 "NewStack initializes a Stack",
			wantContainerFirstOk: false,
			wantContainerSize:    0,
		},
		{
			desc:                    "NewStack with initializer produces FIFO ordering",
			items:                   []int{1, 2, 3},
			wantContainerFirstOk:    true,
			wantContainerFirstValue: 1,
			wantContainerSize:       3,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotStack := NewStack(tC.items...)
			gotContainerFirstValue, gotContainerFirstOk := gotStack.container.First()
			gotContainerSize := gotStack.container.Size()

			assert.Equal(t, tC.wantContainerFirstOk, gotContainerFirstOk)
			assert.Equal(t, tC.wantContainerFirstValue, gotContainerFirstValue)
			assert.Equal(t, tC.wantContainerSize, gotContainerSize)
		})
	}
}

func TestStackPush(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                    string
		stack                   *Stack[int]
		pushValue               int
		wantContainerSize       uint
		wantContainerFirst      bool
		wantContainerFirstValue int
	}{
		{
			desc:                    "Push to an empty Stack inserts value",
			stack:                   NewStack[int](),
			pushValue:               42,
			wantContainerSize:       1,
			wantContainerFirst:      true,
			wantContainerFirstValue: 42,
		},
		{
			desc:                    "Push inserts value at the head",
			stack:                   NewStack([]int{41, 40}...),
			pushValue:               42,
			wantContainerSize:       3,
			wantContainerFirst:      true,
			wantContainerFirstValue: 42,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			tC.stack.Push(tC.pushValue)
			gotContainerLen := tC.stack.container.Size()
			gotContainerFirstValue, gotContainerFirstOk := tC.stack.container.First()

			assert.Equal(t, tC.wantContainerSize, gotContainerLen)
			assert.Equal(t, tC.wantContainerFirst, gotContainerFirstOk)

			if tC.wantContainerFirst {
				assert.Equal(t, tC.wantContainerFirstValue, gotContainerFirstValue)
			}
		})
	}
}

func BenchmarkNewStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewStack[int]()
	}
}

func TestStackPop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                    string
		stack                   *Stack[int]
		wantPopOk               bool
		wantPopValue            int
		wantContainerSize       uint
		wantContainerFirst      bool
		wantContainerFirstValue int
	}{
		{
			desc:               "Pop on an empty Stack",
			stack:              NewStack[int](),
			wantPopOk:          false,
			wantContainerSize:  0,
			wantContainerFirst: false,
		},
		{
			desc:                    "Pop removes and returns the head value",
			stack:                   NewStack([]int{42, 41, 40}...),
			wantPopOk:               true,
			wantPopValue:            42,
			wantContainerSize:       2,
			wantContainerFirst:      true,
			wantContainerFirstValue: 41,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotPopValue, gotPopOk := tC.stack.Pop()
			gotContainerLen := tC.stack.container.Size()
			gotContainerFirstValue, gotContainerFirstOk := tC.stack.container.First()

			assert.Equal(t, tC.wantPopOk, gotPopOk)
			assert.Equal(t, tC.wantPopValue, gotPopValue)
			assert.Equal(t, tC.wantContainerSize, gotContainerLen)
			assert.Equal(t, tC.wantContainerFirst, gotContainerFirstOk)

			if tC.wantContainerFirst {
				assert.Equal(t, tC.wantContainerFirstValue, gotContainerFirstValue)
			}
		})
	}
}

func BenchmarkStackPop(b *testing.B) {
	b.ReportAllocs()

	stack := NewStack[int]()

	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func TestStackHead(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc                    string
		stack                   *Stack[int]
		wantHeadOk              bool
		wantHeadValue           int
		wantContainerSize       uint
		wantContainerFirst      bool
		wantContainerFirstValue int
	}{
		{
			desc:               "Head on an empty Stack",
			stack:              NewStack[int](),
			wantHeadOk:         false,
			wantContainerSize:  0,
			wantContainerFirst: false,
		},
		{
			desc:                    "Head returns the head value and leaves the Stack untouched",
			stack:                   NewStack([]int{42, 41, 40}...),
			wantHeadOk:              true,
			wantHeadValue:           42,
			wantContainerSize:       3,
			wantContainerFirst:      true,
			wantContainerFirstValue: 42,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			gotHeadValue, gotHeadOk := tC.stack.Head()
			gotContainerLen := tC.stack.container.Size()
			gotContainerFirstValue, gotContainerFirstOk := tC.stack.container.First()

			assert.Equal(t, tC.wantHeadOk, gotHeadOk)
			assert.Equal(t, tC.wantHeadValue, gotHeadValue)
			assert.Equal(t, tC.wantContainerSize, gotContainerLen)
			assert.Equal(t, tC.wantContainerFirst, gotContainerFirstOk)

			if tC.wantContainerFirst {
				assert.Equal(t, tC.wantContainerFirstValue, gotContainerFirstValue)
			}
		})
	}
}

func BenchmarkStackHead(b *testing.B) {
	b.ReportAllocs()

	stack := NewStack[int]()

	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Head()
	}
}
