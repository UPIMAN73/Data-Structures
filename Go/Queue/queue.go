package Queue

import (
	"fmt"
)

// A datastructure used to queue items into a "queue"
type Queue[T any] struct {
	Bucket []T
}

// Adds an item to the queue
func Push[T any](queue *Queue[T], value T) T {
	queue.Bucket = append(queue.Bucket, value)
	return value
}

// Viewes the first item in the queue
func Peek[T any](queue *Queue[T]) T {
	if len(queue.Bucket) > 0 {
		return queue.Bucket[0]
	} else {
		var noop T
		return noop
	}
}

// Removes the last item from the queue
func Pop[T any](queue *Queue[T]) T {
	if len(queue.Bucket) > 1 {
		// Grab the last value & remove the last item from the bucket
		output := queue.Bucket[0]
		queue.Bucket = queue.Bucket[1:]
		return output
	} else if len(queue.Bucket) == 1 {
		// Setting the first element in the bucket to empty so that removal is easier
		var noop T
		output := queue.Bucket[0]
		queue.Bucket[0] = noop
		queue.Bucket = queue.Bucket[:0]
		return output
	} else {
		var noop T
		return noop
	}
}

// Test function that demonstrates the use of the datastructure
func TestQueue() {
	// Initilize the queue
	var queue Queue[int]
	numberOfItems := 7

	// Push items into the queue
	fmt.Println("Adding items into the queue")
	fmt.Println(fmt.Sprintf("Initial length of the queue: %d elements", len(queue.Bucket)))
	for i := 1; i <= numberOfItems; i++ {
		Push(&queue, i)
	}
	fmt.Println(fmt.Sprintf("Final length of the queue: %d elements", len(queue.Bucket)))
	fmt.Println("\n")

	// Print out the various items available
	for i := 0; i < numberOfItems; i++ {
		fmt.Println(fmt.Sprintf("\tItem Value: %+v", Pop(&queue)))
		fmt.Println(fmt.Sprintf("\t\tLength of the queue: %d elements", len(queue.Bucket)))
	}

	// Error checking
	fmt.Println("\nSince the previous items in the queue have been removed,\n   we handle the error case by providing an empty variable\n   which in this case is a 0.\n")
	fmt.Println(fmt.Sprintf("\tItem Value: %+v", Pop(&queue)))
	fmt.Println(fmt.Sprintf("Length of the queue: %d elements", len(queue.Bucket)))
}
