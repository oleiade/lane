# Lane

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Build Status](https://github.com/oleiade/lane/actions/workflows/go.yml/badge.svg)](https://github.com/oleiade/lane/actions/workflows/go.yml)
[![Go Documentation](https://pkg.go.dev/badge/github.com/oleiade/lane/v2#pkg-types.svg)](https://pkg.go.dev/github.com/oleiade/lane/v2#pkg-types)
[![Go Report Card](https://goreportcard.com/badge/github.com/oleiade/lane)](https://goreportcard.com/report/github.com/oleiade/lane)
![Go Version](https://img.shields.io/github/go-mod/go-version/oleiade/lane)

The Lane package provides textbook implementations of generic `Queue`, `PriorityQueue`, `Stack`, and `Deque` data structures. Its design focuses on simplicity, performance, and concurrent usage.

<!-- toc -->

- [Lane](#lane)
  - [Installation](#installation)
    - [Using `v2` releases](#using-v2-releases)
    - [Using `v1` releases](#using-v1-releases)
  - [Usage/Examples](#usageexamples)
    - [Priority queue](#priority-queue)
      - [Example](#example)
    - [Deque](#deque)
      - [Deque example](#deque-example)
    - [Queue](#queue)
      - [Queue example](#queue-example)
    - [Stack](#stack)
      - [Stack example](#stack-example)
  - [Performance](#performance)
  - [Documentation](#documentation)
  - [License](#license)

## Installation

Using this package requires a working Go environment. [See the install instructions for Go](http://golang.org/doc/install.html).

This package requires a modern version of Go supporting modules: [see the go blog guide on using Go Modules](https://blog.golang.org/using-go-modules).

### Using `v2` releases

```bash
go get github.com/oleiade/lane/v2
```

```go
...
import (
 "github.com/oleiade/lane/v2"
)
...
```

### Using `v1` releases

```bash
go get github.com/oleiade/lane
```

```go
...
import (
 "github.com/oleiade/lane"
)
...
```

## Usage/Examples

### Priority queue

`PriorityQueue` implements a _heap priority queue_ data structure. It can be either max (descending) or min (ascending) ordered. Every operation on a `PriorityQueue` is  goroutine-safe. It performs `Push` and `Pop` operations in *O(log N)* time.

#### Example

```go
// Create a new max ordered priority queue
priorityQueue := NewMaxPriorityQueue[string, int]()

// And push some prioritized content into it
priorityQueue.Push("easy as", 3)
priorityQueue.Push("123", 2)
priorityQueue.Push("do re mi", 4)
priorityQueue.Push("abc", 1)

// Take a look at the min element in
// the priority queue
headValue, headPriority, ok := priorityQueue.Head()
if ok {
    fmt.Println(headValue)    // "abc"
    fmt.Println(headPriority) // 1
}

// The operations seem to preserve the song order
jacksonFive := make([]string, priorityQueue.Size())

for i := 0; i < len(jacksonFive); i++ {
    value, _, _ := priorityQueue.Pop()
    jacksonFive[i] = value
}

fmt.Println(strings.Join(jacksonFive, " "))
```

### Deque

Deque implements a _head-tail linked list data_ structure. Built upon a doubly linked list container, every operation performed on a `Deque` happen in *O(1)* time complexity. Every operation on a `Deque` are goroutine-safe.

Users have the option to instantiate Deques with a limited capacity using the dedicated `NewBoundDeque` constructor. When a bound Deque is full, the `Append` and `Prepend` operations fail.

#### Deque example

```go
// Create a new Deque data structure
deque := NewDeque[string]()

// And push some content into it using the Append
// and Prepend methods
deque.Append("easy as")
deque.Prepend("123")
deque.Append("do re mi")
deque.Prepend("abc")

// Take a look at what the first and
// last element stored in the Deque are.
firstValue, ok := deque.First()
if ok {
    fmt.Println(firstValue) // "abc"
}

lastValue, ok := deque.Last()
if ok {
    fmt.Println(lastValue) // 1
}

// Use the `Pop` and `Shift`
// methods to bring the song words together
jacksonFive := make([]string, deque.Size())

for i := 0; i < len(jacksonFive); i++ {
    value, ok := deque.Shift()
    if ok {
        jacksonFive[i] = value
    }
}

// abc 123 easy as do re mi
fmt.Println(strings.Join(jacksonFive, " "))
```

### Queue

`Queue` is a **FIFO** (_First In First Out_) data structure implementation. Built upon a `Deque` container, it focuses its API on the following core functionalities: `Enqueue`, `Dequeue`, `Head`. Every operation on a Queue has a time complexity of *O(1)*. Every operation on a `Queue` is goroutine-safe.

#### Queue example

```go
// Create a new queue and pretend to handle Starbucks clients
queue := NewQueue[string]()

// Add the incoming clients to the queue
queue.Enqueue("grumpyClient")
queue.Enqueue("happyClient")
queue.Enqueue("ecstaticClient")

fmt.Println(queue.Head()) // grumpyClient

// Handle the clients asynchronously
for {
    client, ok := queue.Dequeue()
    if !ok {
        break
    }

    fmt.Println(client)
}
```

### Stack

`Stack` implements a **Last In First Out** data structure. Built upon a `Deque` container, its API focuses on the following core functionalities: `Push`, `Pop`, `Head`. Every operation on a Stack has a time complexity of *O(1)*. Every operation on a `Stack` is goroutine-safe.

#### Stack example

```go
// Create a new stack and put some plates over it
stack := NewStack[string]()

// Put some plates on the stack
stack.Push("redPlate")
stack.Push("bluePlate")
stack.Push("greenPlate")

fmt.Println(stack.Head()) // greenPlate

// Check the top of the stack
value, ok := stack.Pop()
if ok {
    fmt.Println(value) // greenPlate
}

stack.Push("yellowPlate")

value, ok = stack.Pop()
if ok {
    fmt.Println(value) // yellowPlate
}

// Check the top of the stack
value, ok = stack.Pop()
if ok {
    fmt.Println(value) // bluePlate
}

// Check the top of the stack
value, ok = stack.Pop()
if ok {
    fmt.Println(value) // redPlate
}
```

## Performance

```bash
go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/oleiade/lane/v2
BenchmarkDequeAppend-8          22954384        54.44 ns/op      32 B/op       1 allocs/op
BenchmarkDequePrepend-8         25097098        44.59 ns/op      32 B/op       1 allocs/op
BenchmarkDequePop-8             63403720        18.99 ns/op       0 B/op       0 allocs/op
BenchmarkDequeShift-8           63390186        20.88 ns/op       0 B/op       0 allocs/op
BenchmarkDequeFirst-8           86662152        13.76 ns/op       0 B/op       0 allocs/op
BenchmarkDequeLast-8            85955014        13.76 ns/op       0 B/op       0 allocs/op
BenchmarkDequeSize-8            86614975        13.77 ns/op       0 B/op       0 allocs/op
BenchmarkDequeEmpty-8           86893297        14.22 ns/op       0 B/op       0 allocs/op
BenchmarkBoundDequeFull-8       590152324         2.039 ns/op       0 B/op       0 allocs/op
BenchmarkBoundDequeAppend-8     64457216        18.50 ns/op       0 B/op       0 allocs/op
BenchmarkBoundDeque-8           64631377        18.48 ns/op       0 B/op       0 allocs/op
BenchmarkPriorityQueuePush-8    19994029        65.67 ns/op      72 B/op       1 allocs/op
BenchmarkPriorityQueuePop-8     62751249        18.52 ns/op       0 B/op       0 allocs/op
BenchmarkPriorityQueueHead-8    86090420        13.77 ns/op       0 B/op       0 allocs/op
BenchmarkPriorityQueueSize-8    86768415        13.77 ns/op       0 B/op       0 allocs/op
BenchmarkPriorityQueueEmpty-8   87036146        13.76 ns/op       0 B/op       0 allocs/op
BenchmarkNewQueue-8             19570003        60.36 ns/op
BenchmarkQueueEnqueue-8         25482283        46.63 ns/op      32 B/op       1 allocs/op
BenchmarkQueueDequeue-8         63715965        18.50 ns/op       0 B/op       0 allocs/op
BenchmarkQueueHead-8            85664312        13.79 ns/op       0 B/op       0 allocs/op
BenchmarkNewStack-8             19925473        59.57 ns/op
BenchmarkStackPop-8             64704993        18.80 ns/op       0 B/op       0 allocs/op
BenchmarkStackHead-8            86119761        13.76 ns/op       0 B/op       0 allocs/op
```

## Documentation

For a more detailed overview of lane, please refer to [Documentation](https://pkg.go.dev/github.com/oleiade/lane/v2)

## License

[MIT](https://choosealicense.com/licenses/mit/)
