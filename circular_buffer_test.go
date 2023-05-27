package lane

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularBufferPut_EmptyWithoutCapacity(t *testing.T) {
	t.Parallel()

	rb := NewCircularBuffer[int](0)
	gotOk := rb.Put(1)

	assert.False(t, gotOk)
	assert.True(t, rb.empty())
}

func TestCircularBufferPut_EmptyWithCapacity(t *testing.T) {
	t.Parallel()

	rb := NewCircularBuffer[int](3)
	gotOk := rb.Put(1)

	assert.True(t, gotOk)
	assert.False(t, rb.empty())
	assert.Equal(t, 1, rb.data[0])
	assert.Equal(t, 0, rb.readPos)
	assert.Equal(t, 1, rb.writePos)
}

func TestCircularBufferPut_FillUpToCapacity(t *testing.T) {
	t.Parallel()

	rb := NewCircularBuffer[int](3)
	rb.Put(1)
	rb.Put(2)
	rb.Put(3)

	assert.True(t, rb.full)
	assert.Equal(t, 1, rb.data[0])
	assert.Equal(t, 2, rb.data[1])
	assert.Equal(t, 3, rb.data[2])
	assert.Equal(t, 0, rb.readPos)
	assert.Equal(t, 0, rb.writePos)
}

func TestCircularBufferPut_Overwrite(t *testing.T) {
	t.Parallel()

	rb := NewCircularBuffer[int](3)
	rb.Put(1)
	rb.Put(2)
	rb.Put(3)
	rb.Put(4)

	assert.True(t, rb.full)
	assert.Equal(t, 4, rb.data[0])
	assert.Equal(t, 2, rb.data[1])
	assert.Equal(t, 3, rb.data[2])
	assert.Equal(t, 1, rb.readPos)
	assert.Equal(t, 1, rb.writePos)
}

func BenchmarkCircularBufferPut(b *testing.B) {
	b.ReportAllocs()

	cb := NewCircularBuffer[int](b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.Put(i)
	}
}

func TestCircularBufferPop_EmptyWithoutCapacity(t *testing.T) {
	t.Parallel()

	rb := NewCircularBuffer[int](0)
	_, ok := rb.Pop()
	assert.False(t, ok)
	assert.True(t, rb.empty())
}

func TestCircularBufferPop_EmptyWithCapacity(t *testing.T) {
	t.Parallel()

	rb := NewCircularBuffer[int](3)
	_, ok := rb.Pop()
	assert.False(t, ok)
	assert.True(t, rb.empty())
}

func TestCircularBufferPop_LessElementsThanCapacity(t *testing.T) {
	t.Parallel()

	rb := NewCircularBuffer[int](3)
	rb.Put(1)
	rb.Put(2)

	item, ok := rb.Pop()

	assert.True(t, ok)
	assert.Equal(t, 1, item)
	assert.False(t, rb.empty())
	assert.Equal(t, 1, rb.readPos)
	assert.Equal(t, 2, rb.writePos)
}

func TestCircularBufferPop_PopFilledBuffer(t *testing.T) {
	t.Parallel()

	rb := NewCircularBuffer[int](3)
	rb.Put(1)
	rb.Put(2)
	rb.Put(3)

	item, ok := rb.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, item)
	assert.False(t, rb.empty())
	assert.False(t, rb.full)
	assert.Equal(t, 1, rb.readPos)
	assert.Equal(t, 0, rb.writePos)

	item, ok = rb.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, item)
	assert.False(t, rb.empty())
	assert.False(t, rb.full)
	assert.Equal(t, 2, rb.readPos)
	assert.Equal(t, 0, rb.writePos)

	item, ok = rb.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, item)
	assert.True(t, rb.empty())
	assert.False(t, rb.full)
	assert.Equal(t, 0, rb.readPos)
	assert.Equal(t, 0, rb.writePos)
}

func BenchmarkCircularBufferPop(b *testing.B) {
	b.ReportAllocs()

	cb := NewCircularBuffer[int](b.N)

	for i := 0; i < b.N; i++ {
		cb.Put(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.Pop()
	}
}

func TestCircularBuffer_PutAfterPopAll(t *testing.T) {
	t.Parallel()

	rb := NewCircularBuffer[int](3)
	rb.Put(1)
	rb.Put(2)

	_, ok := rb.Pop()
	assert.True(t, ok)

	_, ok = rb.Pop()
	assert.True(t, ok)

	rb.Put(3)
	rb.Put(4)

	item, ok := rb.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, item)

	item, ok = rb.Pop()
	assert.True(t, ok)
	assert.Equal(t, 4, item)
}

func TestCircularBufferPeek_EmptyWithCapacity(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	_, ok := cb.Peek()
	assert.False(t, ok)
	assert.True(t, cb.empty())
}

func TestCircularBufferPeek_LessElementsThanCapacity(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	cb.Put(1)
	cb.Put(2)

	item, ok := cb.Peek()

	assert.True(t, ok)
	assert.Equal(t, 1, item)
	assert.False(t, cb.empty())
	assert.Equal(t, 0, cb.readPos)
	assert.Equal(t, 2, cb.writePos)
}

func TestCircularBufferPeek_FilledBuffer(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	cb.Put(1)
	cb.Put(2)
	cb.Put(3)

	item, ok := cb.Peek()

	assert.True(t, ok)
	assert.Equal(t, 1, item)
	assert.False(t, cb.empty())
	assert.True(t, cb.full)
	assert.Equal(t, 0, cb.readPos)
	assert.Equal(t, 0, cb.writePos)
}

func BenchmarkCircularBufferPeek(b *testing.B) {
	b.ReportAllocs()

	cb := NewCircularBuffer[int](b.N)
	cb.Put(203895836)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.Peek()
	}
}

func TestCircularBufferSlice_EmptyBuffer(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	slice := cb.View()

	assert.Nil(t, slice)
}

func TestCircularBufferSlice_FullBuffer(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	cb.Put(1)
	cb.Put(2)
	cb.Put(3)

	slice := cb.View()

	assert.Equal(t, []int{1, 2, 3}, slice)
}

func TestCircularBufferSlice_PartialBuffer(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	cb.Put(1)
	cb.Put(2)

	slice := cb.View()

	assert.Equal(t, []int{1, 2}, slice)
}

func TestCircularBufferSlice_BufferWithWrapAround(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	cb.Put(1)
	cb.Put(2)
	cb.Put(3)
	cb.Put(4) // causes a wrap around, overwriting 1

	slice := cb.View()

	assert.Equal(t, []int{2, 3, 4}, slice)
}

func TestCircularBufferSize_EmptyWithCapacity(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	size := cb.Size()
	assert.Equal(t, 0, size)
}

func TestCircularBufferSize_LessElementsThanCapacity(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	cb.Put(1)
	cb.Put(2)

	size := cb.Size()
	assert.Equal(t, 2, size)
}

func TestCircularBufferSize_FilledBuffer(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	cb.Put(1)
	cb.Put(2)
	cb.Put(3)

	size := cb.Size()
	assert.Equal(t, 3, size)
}

func TestCircularBufferSize_OverfilledBuffer(t *testing.T) {
	t.Parallel()

	cb := NewCircularBuffer[int](3)
	cb.Put(1)
	cb.Put(2)
	cb.Put(3)
	cb.Put(4)

	size := cb.Size()
	assert.Equal(t, 3, size)
}
