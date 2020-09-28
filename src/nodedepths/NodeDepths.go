package nodedepths

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}



func NodeDepths(root *BinaryTree) int {
	// Write your code here.
	
	return NodeDepthsHelper(root, 0)
}

func NodeDepthsHelper(root *BinaryTree, depths int) int {
	if root == nil {
		return 0
	}
	return depths + NodeDepthsHelper(root.Left,depths+1)+NodeDepthsHelper(root.Right,depths+1)
}
