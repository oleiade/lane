# Lane

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Build Status](https://github.com/oleiade/lane/actions/workflows/go.yml/badge.svg)](https://github.com/oleiade/lane/actions/workflows/go.yml)
[![Go Documentation](https://pkg.go.dev/badge/github.com/oleiade/lane#pkg-types.svg)](https://pkg.go.dev/github.com/oleiade/lane#pkg-types)
[![Go Report Card](https://goreportcard.com/badge/github.com/oleiade/lane)](https://goreportcard.com/report/github.com/oleiade/lane)
![Go Version](https://img.shields.io/github/go-mod/go-version/oleiade/lane)

The Lane package provides implementations of generic `Queue`, `PriorityQueue`, `Stack`, and `Deque` data structures. It was designed with simplicity, performance, and concurrent usage in mind.

## Installation

Using this package requires a working Go environment. [See the install instructions for Go](http://golang.org/doc/install.html).

Go Modules are required when using this package. [See the go blog guide on using Go Modules](https://blog.golang.org/using-go-modules).

### Using `v2` releases

```bash
go get github.com/oleiade/lane/v2
```

```go
...
import (
 "github.com/oleiade/lane/v2" // imports as package "cli"
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

### Priority Queue

`PriorityQueue` is a _heap priority queue_ data structure implementation. It can be either maximum (descending) or minimum (ascending) oriented (ordered). Every operation on a `PriorityQueue` are synchronized and goroutine-safe. It performs `Push` and `Pop` operations in `O(log N)` time.

#### Example

```go
// Let's create a new max ordered priority queue
priorityQueue := NewMaxPriorityQueue[string]()

// And push some prioritized content into it
priorityQueue.Push("easy as", 3)
priorityQueue.Push("123", 2)
priorityQueue.Push("do re mi", 4)
priorityQueue.Push("abc", 1)

// Now let's take a look at the min element in
// the priority queue
headValue, headPriority, ok := priorityQueue.Head()
if okay {
    fmt.Println(headValue) // "abc"
    fmt.Println(headPriority) // 1
}

// okay, the song order seems to be preserved; let's
// roll
var jacksonFive []string = make([]string, priorityQueue.Size())

for i := 0; i < len(jacksonFive); i++ {
    value, _, _ := priorityQueue.Pop()
    jacksonFive[i] = value
}

fmt.Println(strings.Join(jacksonFive, " "))
```

### Deque

Deque is a _head-tail linked list data_ structure implementation. It is built upon a doubly-linked list container, and every operation on a `Deque` are performed in `O(1)` time complexity. Every operation on a `Deque` is synchronized and goroutine-safe.

Deques can optionally be instantiated with a limited capacity using the dedicated `NewBoundDeque` constructor, whereby the return value of the `Append` and `Prepend` return false if the Deque was full and the item was not added.

#### Deque Example

```go
// Let's create a new deque data structure
deque := NewDeque[string]()

// And push some content into it using the Append
// and Prepend methods
deque.Append("easy as")
deque.Prepend("123")
deque.Append("do re mi")
deque.Prepend("abc")

// Now let's take a look at what are the first and
// last element stored in the Deque
firstValue, ok := deque.First()
if okay {
    fmt.Println(firstValue) // "abc"
}

lastValue, ok := deque.Last()
if ok {
    fmt.Println(lastValue) // 1
}

//okay, now let's play with the Pop and Shift
// methods to bring the song words together
var jacksonFive []string = make([]string, deque.Size())

for i := 0; i < len(jacksonFive); i++ {
    value, ok := deque.Shift()
    if okay {
        jacksonFive[i] = value
    }
}

// abc 123 easy as do re mi
fmt.Println(strings.Join(jacksonFive, " "))
```

### Queue

`Queue` is a **FIFO** (_First In First Out_) data structure implementation. It is built upon a `Deque` container and focuses its API on core functionalities: `Enqueue`, `Dequeue`, `Head`. Every operation's time complexity is O(1). Every operation on a `Queue` is synchronized and goroutine-safe.

#### Queue Example

```go
import (
"fmt"
"github.com/oleiade/lane/v2"
"sync"
)

func worker(item interface{}, wg *sync.WaitGroup) {
    fmt.Println(item)
    wg.Done()
}


func main() {
    // Create a new queue and pretend we're handling starbucks
    // clients
    queue := NewQueue[string]()

    // Let's add the incoming clients to the queue
    queue.Enqueue("grumpyClient")
    queue.Enqueue("happyClient")
    queue.Enqueue("ecstaticClient")

    fmt.Println(queue.Head()) // grumpyClient

    // Let's handle the clients asynchronously
    for client, okay:= queue.Dequeue(); ok; {
        go fmt.Println(client)
    }

    // Wait until everything is printed
    wg.Wait()
}
```

### Stack

`Stack` is a **LIFO** ( _Last in first out_ ) data structure implementation. It is built upon a `Deque` container and focuses its API on core functionalities: `Push`, `Pop`, `Head`. Every operation time complexity is O(1). Every operation on a `Stack` is synchronized and goroutine-safe.

#### Stack Example

```go
// Create a new stack and put some plates over it
stack := NewStack[string]()

// Let's put some plates on the stack
stack.Push("redPlate")
stack.Push("bluePlate")
stack.Push("greenPlate")

fmt.Println(stack.Head()) // greenPlate

// What's on top of the stack?
value, okay:= stack.Pop()
if okay {
    fmt.Println(value) // greenPlate
}

stack.Push("yellowPlate")

value, ok = stack.Pop()
if okay {
    fmt.Println(value) // yellowPlate
}

// What's on top of the stack?
value, ok = stack.Pop()
if okay {
    fmt.Println(value) // bluePlate
}

// What's on top of the stack?
value, ok = stack.Pop()
if okay {
    fmt.Println(value) // redPlate
}
```

## Documentation

For a more detailed overview of lane, please refer to [Documentation](http://godoc.org/github.com/oleiade/lane)

## License

[MIT](https://choosealicense.com/licenses/mit/)
