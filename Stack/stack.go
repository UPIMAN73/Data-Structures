package Stack

import (
	"fmt"
)

// A datastructure used to stack items into a "stack"
type Stack[T any] struct {
	Bucket []T
}

// Adds an item to the stack
func Push[T any](stack *Stack[T], value T) T {
	stack.Bucket = append(stack.Bucket, value)
	return value
}

// Viewes the last item in the stack
func Peek[T any](stack *Stack[T]) T {
	if len(stack.Bucket) > 0 {
		return stack.Bucket[len(stack.Bucket)-1]
	} else {
		var noop T
		return noop
	}
}

// Removes the last item from the stack
func Pop[T any](stack *Stack[T]) T {
	if len(stack.Bucket) > 0 {
		// Grab the last value & remove the last item from the bucket
		output := stack.Bucket[len(stack.Bucket)-1]
		stack.Bucket = stack.Bucket[:len(stack.Bucket)-1]
		return output
	} else {
		var noop T
		return noop
	}
}

// Test function that demonstrates the use of the datastructure
func TestStack() {
	// Initilize the stack
	var stack Stack[int]

	// Push items into the stack
	Push(&stack, 1)
	Push(&stack, 2)
	Push(&stack, 4)

	// Print out the various items available
	fmt.Println(Peek[int](&stack))
	fmt.Println(Pop(&stack))
}
