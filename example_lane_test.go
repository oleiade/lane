package lane

import (
	"fmt"
	"strings"
)

func ExamplePriorityQueue() {
	// Let's create a new max ordered priority queue
	var priorityQueue *PQueue = NewPQueue(MINPQ)

	// And push some prioritized content into it
	priorityQueue.Push("easy as", 3)
	priorityQueue.Push("123", 2)
	priorityQueue.Push("do re mi", 4)
	priorityQueue.Push("abc", 1)

	// Now let's take a look at the min element in
	// the priority queue
	headValue, headPriority := priorityQueue.Head()
	fmt.Println(headValue)    // "abc"
	fmt.Println(headPriority) // 1

	// Okay the song order seems to be preserved, let's
	// roll
	var jacksonFive []string = make([]string, priorityQueue.Size())

	for i := 0; i < len(jacksonFive); i++ {
		value, _ := priorityQueue.Pop()

		jacksonFive[i] = value.(string)
	}

	fmt.Println(strings.Join(jacksonFive, " "))
}

func ExampleDeque() {

}

func ExampleQueue() {

}

func ExampleStack() {

}
