package main

import (
	"List"
	"fmt"
)

// Test the List datastructure
func testList() {
	// List definition
	var test List.List[string]
	var test2 List.List[string]

	// List value declaration
	test.Value = "Hello World!"

	// Appending elements to the list
	List.ListAppend[string](&test, "How are you Today?")
	List.ListAppend[string](&test, "I am good, how are you?")
	List.ListAppend[string](&test, "Hi!")
	List.ListAppend[string](&test, "Hola!")
	List.ListAppend[string](&test, "ni hao!")

	// Print the structure and the size of the list
	fmt.Println(test)
	fmt.Println(List.ListSize(&test))
	fmt.Println(List.ListSize(&test2)) // Has one empty element in there

	// List Searching
	fmt.Println(List.ListSearch(&test, "I am good, how are you?"))
	fmt.Println(List.ListSearch(&test, "Hiiiii!"))

	// Removing Items in the List
	List.ListPrintAll(&test)
	fmt.Println(List.ListSize(&test))
	List.ListRemove(&test, "Hi!")
	fmt.Println(List.ListSize(&test))
	List.ListPrintAll(&test)
}

// Linked List Test
func testLinkedList() {
	// List definition
	var test List.LinkedList[string]
	var test2 List.LinkedList[string]

	// List value declaration
	test.Value = "Hello World!"

	// Appending elements to the list
	List.LinkedListAppend[string](&test, "How are you Today?")
	List.LinkedListAppend[string](&test, "I am good, how are you?")
	List.LinkedListAppend[string](&test, "Hi!")
	List.LinkedListAppendPrevious[string](&test, "Hola!")
	List.LinkedListAppend[string](&test, "ni hao!")

	// Print the structure and the size of the LinkedList
	fmt.Println(test)
	fmt.Println(List.LinkedListSize(&test))
	fmt.Println(List.LinkedListSize(&test2)) // Has one empty element in there

	// LinkedList Searching
	fmt.Println(List.LinkedListSearch(&test, "I am good, how are you?"))
	fmt.Println(List.LinkedListSearch(&test, "Hiiiii!"))

	// Removing Items in the LinkedList
	List.LinkedListPrintAll(&test)
	fmt.Println(List.LinkedListSize(&test))
	List.LinkedListRemove(&test, "Hi!")
	fmt.Println(List.LinkedListSize(&test))
	List.LinkedListPrintAll(&test)
}

// Main
func main() {
	testLinkedList()
}
