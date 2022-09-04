package lane

// List represents a doubly linked list.
type List[T any] struct {
	root Element[T]
	len  uint
}

// Init initializes or clears list l.
func (l *List[T]) Init() *List[T] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0

	return l
}

// New returns an initialized list.
func New[T any]() *List[T] {
	return new(List[T]).Init()
}

// Len returns the number of elements of list l.
func (l *List[T]) Len() uint {
	return l.len
}

// Front returns the first element of list l or nil.
func (l *List[T]) Front() *Element[T] {
	if l.len == 0 {
		return nil
	}

	return l.root.next
}

// Back returns the last element of list l or nil.
func (l *List[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}

	return l.root.prev
}

// PushFront inserts a new element e with value v at
// the front of list l and returns e.
func (l *List[T]) PushFront(v T) *Element[T] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at
// the back of list l and returns e.
func (l *List[T]) PushBack(v T) *Element[T] {
	l.lazyInit()

	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v
// immediately before mark and returns e.
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}

	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v
// immediately after mark and returns e.
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}

	return l.insertValue(v, mark)
}

// Remove removes e from l if e is an element of list l.
func (l *List[T]) Remove(e *Element[T]) T {
	if e.list == l {
		l.remove(e)
	}

	return e.Value
}

// MoveToFront moves element e to the front of list l.
func (l *List[T]) MoveToFront(e *Element[T]) {
	if e.list != l || l.root.next == e {
		return
	}

	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
func (l *List[T]) MoveToBack(e *Element[T]) {
	if e.list != l || l.root.prev == e {
		return
	}

	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}

	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}

	l.move(e, mark)
}

// PushBackList inserts a copy of an other list at the back of list l.
func (l *List[T]) PushBackList(other *List[T]) {
	l.lazyInit()

	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of list l.
func (l *List[T]) PushFrontList(other *List[T]) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

func (l *List[T]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List[T]) insert(e, at *Element[T]) *Element[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++

	return e
}

func (l *List[T]) insertValue(v T, at *Element[T]) *Element[T] {
	return l.insert(&Element[T]{Value: v}, at)
}

func (l *List[T]) remove(e *Element[T]) *Element[T] {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}

func (l *List[T]) move(e, at *Element[T]) *Element[T] {
	if e == at {
		return e
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

// Element is a node of a linked list.
type Element[T any] struct {
	next, prev *Element[T]

	list *List[T]

	Value T
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}

	return nil
}

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}

	return nil
}
