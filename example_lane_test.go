package lane

import (
	"fmt"
	"strings"
)

func ExamplePriorityQueue() {
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
	// Output:
	// do re mi
	// 4
	// do re mi easy as 123 abc
}

func ExampleDeque() {
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
	// Output:
	// abc
	// do re mi
	// abc 123 easy as do re mi
}

func ExampleQueue() {
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

	// Output:
	// grumpyClient true
	// grumpyClient
	// happyClient
	// ecstaticClient
}

func ExampleStack() {
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

	// Output:
	// greenPlate true
	// greenPlate
	// yellowPlate
	// bluePlate
	// redPlate
}
