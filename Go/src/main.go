package main

import (
	"List"
	"fmt"
)

// Test the List datastructure
func testList() {
	// List definition
	var test List.List[string]

	// List value declaration
	test.Value = "Hello World!"

	// Adding nodes to the list
	List.AddNode[string](&test, "How are you Today?")
	List.AddNode[string](&test, "I am good, how are you?")

	// Print the structure and the size of the list
	fmt.Println(test)
	fmt.Println(List.ListSize(&test))

	// List Searching
	fmt.Println(List.ListSearch(&test, "I am good, how are you?"))
	fmt.Println(List.ListSearch(&test, "Hi!"))
}

// Main
func main() {
	testList()
}
