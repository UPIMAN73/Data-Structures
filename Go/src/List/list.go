package List

// List Node
type ListNode[T any] struct {
	Next  *ListNode[T]
	Value T
}

// Generic List
type List[T any] struct {
	Size     uint64
	Elements *ListNode[T]
}

// Initilize a List
func InitList[T any](list *List[T], value T) {
	list.Size = 1
	list.Elements = InitListNode[T](value)
}

// Initilize a ListNode type
func InitListNode[T any](value T) *ListNode[T] {
	newElement := ListNode[T]{Value: value, Next: nil}
	return &newElement
}
