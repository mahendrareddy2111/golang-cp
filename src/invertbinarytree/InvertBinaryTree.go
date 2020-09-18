package invertbinarytree

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	// Write your code here.
	if tree == nil {
		return 
	}
	tree.Right,tree.Left = tree.Left,tree.Right

	tree.Left.InvertBinaryTree()
	tree.Right.InvertBinaryTree()
}