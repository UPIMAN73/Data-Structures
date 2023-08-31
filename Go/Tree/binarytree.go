package Tree

import "fmt"

// Binary Tree Datatype
type BinaryTree[T any] struct {
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
	Value T
}

// Add left leaf to binary tree
func BinaryTreeAddLeft[T any](node *BinaryTree[T], value T) *BinaryTree[T] {
	if node.Left == nil {
		node.Left = new(BinaryTree[T])
		node.Left.Value = value
		return node
	}
	return nil
}

// Add right leaf to binary tree
func BinaryTreeAddRight[T any](node *BinaryTree[T], value T) *BinaryTree[T] {
	if node.Right == nil {
		node.Right = new(BinaryTree[T])
		node.Right.Value = value
		return node
	}
	return nil
}

// Print out tree based on values
func BinaryTreePrintTree[T any](node *BinaryTree[T], depth uint64) {
	// Spacing
	for i := uint64(0); i < depth*10; i++ {
		fmt.Print(" ")
	}

	// Print value (at the 'parent')
	fmt.Println(node.Value)

	// Spacing
	for i := uint64(0); i < depth*8; i++ {
		fmt.Print(" ")
	}

	// Print left branch visual
	fmt.Print("/")

	// Spacing
	for i := uint64(0); i < depth*7; i++ {
		fmt.Print(" ")
	}

	// Print right branch visual
	fmt.Println("\\")

	// Spacing
	for i := uint64(0); i < depth*5; i++ {
		fmt.Print(" ")
	}

	// Print right node value
	fmt.Print(node.Left.Value)

	// Spacing
	for i := uint64(0); i < depth*3; i++ {
		fmt.Print(" ")
	}

	// Print right node value
	fmt.Println(node.Right.Value)
}

// Recursive binary tree inversion
func BinaryTreeInvert[T any](node *BinaryTree[T]) *BinaryTree[T] {
	// If the node is empty return nil (in this case the node)
	if node == nil {
		return node
	} else {
		// Recursively invert the nodes, starting with left, then the right
		leftNode := BinaryTreeInvert[T](node.Left)
		node.Left = BinaryTreeInvert[T](node.Right)
		node.Right = leftNode
	}
	// Return the node afterwards
	return node
}

// Test function that demonstrates the use of the datastructure
func TestBinaryTree() {
	// Create a binary tree variable
	var root BinaryTree[string]

	// set the value of the node
	root.Value = "Hello"

	// print out current state
	fmt.Println(root)

	// fmt.Println("Binary Tree Components: ")
	// Add left node
	BinaryTreeAddLeft[string](&root, "World!")
	fmt.Println(root)

	// Add a node to the right
	BinaryTreeAddRight[string](&root, "Joshua!")
	fmt.Println(root)

	// Print out contents of the tree
	fmt.Println("Regular Binary Tree:")
	BinaryTreePrintTree[string](&root, 0)

	// Add values to some of the leaf nodes
	BinaryTreeAddRight[string](root.Right, "How are you today?")
	BinaryTreeAddLeft[string](root.Right, "How you doing?")
	BinaryTreePrintTree[string](root.Right, 2)
	BinaryTreePrintTree[string](&root, 1)

	// Invert the tree
	fmt.Println("\n\nInverted Binary Tree:")
	BinaryTreeInvert[string](&root)
	BinaryTreePrintTree[string](&root, 1)
	BinaryTreePrintTree[string](root.Left, 2)
}
