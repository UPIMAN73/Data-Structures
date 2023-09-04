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

	// Push items into the queue
	Push(&queue, 1)
	Push(&queue, 2)
	Push(&queue, 4)

	// Print out the various items available
	fmt.Println(Peek[int](&queue))
	fmt.Println(fmt.Sprintf("Length of the queue: %d elements", len(queue.Bucket)))
	fmt.Println(Pop(&queue))
}
