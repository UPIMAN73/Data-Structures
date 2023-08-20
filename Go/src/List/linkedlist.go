package List

// Linked List Node
type LinkedListNode[T any] struct {
	Next     *LinkedListNode[T]
	Previous *LinkedListNode[T]
	Value    T
}

// Linked List Object
type LinkedList[T any] struct {
	Size uint64
	Root *LinkedListNode[T]
}

// Initilize a List
func InitLinkedList[T any](list *LinkedList[T], value T) {
	list.Size = 1
	list.Root = InitLinkedListNode[T](value)
}

// Initilize a ListNode type
func InitLinkedListNode[T any](value T) *LinkedListNode[T] {
	newElement := LinkedListNode[T]{Value: value, Next: nil}
	return &newElement
}
