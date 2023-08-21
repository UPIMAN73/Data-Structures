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

// Main
func main() {
	testList()
}
