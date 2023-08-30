package List

import (
	"fmt"
	"reflect"
)

// LinkedList
type LinkedList[T any] struct {
	Next     *LinkedList[T]
	Previous *LinkedList[T]
	Value    T
}

// Print all items within the LinkedList
func LinkedListPrintAll[T any](list *LinkedList[T]) {
	fmt.Println("\nLinkedList Elements:")
	fmt.Print("\t")
	for currentNode := list; currentNode != nil; currentNode = currentNode.Next {
		fmt.Print(" | ")
		fmt.Print(currentNode.Value)
	}
	fmt.Println(" |")
}

// Get the size of the LinkedList
func LinkedListSize[T any](list *LinkedList[T]) uint64 {
	if list == nil {
		return 0
	} else if list.Next == nil {
		return 1
	} else {
		return LinkedListSize[T](list.Next) + 1
	}
}

// Add a node to the LinkedList
func LinkedListAppend[T any](list *LinkedList[T], value T) *LinkedList[T] {
	// Recursive Implementation
	if list.Next == nil {
		// Create & set the node instance and properties
		list.Next = new(LinkedList[T])
		list.Next.Value = value
		return list
	} else {
		// Go to the next element in the LinkedList
		return LinkedListAppend[T](list.Next, value)
	}
}

// Add a node to the LinkedList
func LinkedListAppendPrevious[T any](list *LinkedList[T], value T) *LinkedList[T] {
	// Recursive Implementation
	if list.Previous == nil {
		// Create & set the node instance and properties
		list.Previous = new(LinkedList[T])
		list.Previous.Value = value
		return list
	} else {
		// Go to the next element in the LinkedList
		return LinkedListAppend[T](list.Previous, value)
	}
}

// Search for the object from the LinkedList (by value)
func LinkedListSearch[T any](list *LinkedList[T], value T) *LinkedList[T] {
	// Recursive implementation
	if list == nil {
		return nil
	} else if reflect.DeepEqual(list.Value, value) {
		return list
	} else {
		return LinkedListSearch[T](list.Next, value)
	}
}

// Remove element from LinkedList
func LinkedListRemove[T any](list *LinkedList[T], value T) *LinkedList[T] {
	// Recursive implementation
	if list == nil {
		return nil
	} else if reflect.DeepEqual(list.Value, value) {
		// Sliding Tape Method (essentially moving the LinkedList down by replacing the previous and removing the last object because it is duplicated)
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
			return LinkedListRemove[T](list.Next, list.Next.Value)
		} else {
			// Return nil since we want no callback
			return nil
		}
	} else {
		// Continue to the next element
		return LinkedListRemove[T](list.Next, value)
	}
}

// Linked List Test
func TestLinkedList() {
	// List definition
	var test LinkedList[string]
	var test2 LinkedList[string]

	// List value declaration
	test.Value = "Hello World!"

	// Appending elements to the list
	LinkedListAppend[string](&test, "How are you Today?")
	LinkedListAppend[string](&test, "I am good, how are you?")
	LinkedListAppend[string](&test, "Hi!")
	LinkedListAppendPrevious[string](&test, "Hola!")
	LinkedListAppend[string](&test, "ni hao!")

	// Print the structure and the size of the LinkedList
	fmt.Println(test)
	fmt.Println(LinkedListSize(&test))
	fmt.Println(LinkedListSize(&test2)) // Has one empty element in there

	// LinkedList Searching
	fmt.Println(LinkedListSearch(&test, "I am good, how are you?"))
	fmt.Println(LinkedListSearch(&test, "Hiiiii!"))

	// Removing Items in the LinkedList
	LinkedListPrintAll(&test)
	fmt.Println(LinkedListSize(&test))
	LinkedListRemove(&test, "Hi!")
	fmt.Println(LinkedListSize(&test))
	LinkedListPrintAll(&test)
}
