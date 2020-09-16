package bsttraversal

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(array []int) []int {
	// Write your code here.
	if tree == nil {
		return array
	}
	left := tree.Left
	if left != nil {
		array = left.InOrderTraverse(array)
	}
	array = append(array, tree.Value)
	right := tree.Right
	if right != nil {
		array = right.InOrderTraverse(array)
	}
	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	// Write your code here.
	if tree == nil {
		return array
	}
	array = append(array, tree.Value)
	left := tree.Left
	if left != nil {
		array = left.PreOrderTraverse(array)
	}
	right := tree.Right
	if right != nil {
		array = right.PreOrderTraverse(array)
	}
	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	// Write your code here.
	if tree == nil {
		return array
	}

	left := tree.Left
	if left != nil {
		array = left.PostOrderTraverse(array)
	}
	right := tree.Right
	if right != nil {
		array = right.PostOrderTraverse(array)
	}
	array = append(array, tree.Value)
	return array
}
