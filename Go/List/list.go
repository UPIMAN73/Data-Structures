package List

import (
	"fmt"
	"reflect"
)

// List
type List[T any] struct {
	Next  *List[T]
	Value T
}

// Print all items within the list
func ListPrintAll[T any](list *List[T]) {
	fmt.Println("\nList Elements:")
	fmt.Print("\t")
	for currentNode := list; currentNode != nil; currentNode = currentNode.Next {
		fmt.Print(" | ")
		fmt.Print(currentNode.Value)
	}
	fmt.Println(" |")
}

// Get the size of the list
func ListSize[T any](list *List[T]) uint64 {
	if list == nil {
		return 0
	} else if list.Next == nil {
		return 1
	} else {
		return ListSize[T](list.Next) + 1
	}
}

// Add a node to the list
func ListAppend[T any](list *List[T], value T) *List[T] {
	// Recursive Implementation
	if list.Next == nil {
		// Create & set the node instance and properties
		list.Next = new(List[T])
		list.Next.Value = value
		return list
	} else {
		// Go to the next element in the list
		return ListAppend[T](list.Next, value)
	}
}

// Search for the object from the list (by value)
func ListSearch[T any](list *List[T], value T) *List[T] {
	// Recursive implementation
	if list == nil {
		return nil
	} else if reflect.DeepEqual(list.Value, value) {
		return list
	} else {
		return ListSearch[T](list.Next, value)
	}
}

// Remove element from list
func ListRemove[T any](list *List[T], value T) *List[T] {
	// Recursive implementation
	if list == nil {
		return nil
	} else if reflect.DeepEqual(list.Value, value) {
		// Sliding Tape Method (essentially moving the list down by replacing the previous and removing the last object because it is duplicated)
		if list.Next != nil {
			// Setting next element value to this one (sliding tape method)
			list.Value = list.Next.Value

			// Check if the next object is the last
			if list.Next.Next == nil {
				// If it was, then just remove it
				list.Next = nil
				return nil
			}

			// Move to the next object
			return ListRemove[T](list.Next, list.Next.Value)
		} else {
			// Return nil since we want no callback
			return nil
		}
	} else {
		// Continue to the next element
		return ListRemove[T](list.Next, value)
	}
}
