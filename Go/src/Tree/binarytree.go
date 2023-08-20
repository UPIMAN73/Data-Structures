package Tree

type BinaryTreeNode[T any] struct {
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
	Value T
}

type BinaryTree[T any] struct {
	Size uint64
	Root *BinaryTreeNode[T]
}
