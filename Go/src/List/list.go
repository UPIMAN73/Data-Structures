package List

import (
	"reflect"
)

// List
type List[T any] struct {
	Next  *List[T]
	Value T
}

// Get the size of the list
func ListSize[T any](list *List[T]) uint64 {
	var size uint64
	for currentNode := list; currentNode != nil; currentNode = currentNode.Next {
		size += 1
	}
	return size
}

// Add a node to the list
func AddNode[T any](list *List[T], value T) {
	// Set node to the list elements
	currentNode := list

	// Iterate through the list to obtain the last one
	for i := ListSize(list); i > 1; i-- {
		currentNode = currentNode.Next
	}

	// Create & set the node instance and properties
	currentNode.Next = new(List[T])
	currentNode.Next.Value = value
}

// Search for the object from the list (by value)
func ListSearch[T any](list *List[T], value T) *List[T] {
	// Iterate through the list to obtain the last one
	for currentNode := list; ; /* Do Nothing */ currentNode = currentNode.Next {
		if currentNode == nil {
			return nil
		} else if reflect.DeepEqual(currentNode.Value, value) {
			return currentNode
		} else {
			continue
		}
	}
}
