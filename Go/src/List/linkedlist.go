package List

// Linked List Node
type LinkedListNode struct {
	Next     *LinkedListNode
	Previous *LinkedListNode
	Value    string
}

// Linked List Object
type LinkedList struct {
	Size uint64
	Root *LinkedListNode
}
