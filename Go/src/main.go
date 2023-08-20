package main

import (
	"List"
	"fmt"
)

// Test the List datastructure
func testList() {
	var test List.List[string]
	List.InitList[string](&test, "Hello World!")
	test.Elements.Next = List.InitListNode("How are you Today?")
	fmt.Println(test)
	fmt.Println(test.Elements.Value)
	fmt.Println(test.Elements.Next.Value)
}

func main() {
	testList()
}
