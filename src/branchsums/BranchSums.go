package branchsums

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	return BranchSumsHelper(root, root.Value, []int{})
}

func BranchSumsHelper(root *BinaryTree, sum int, array []int) []int {
	if root == nil {
		return array
	}
	left := root.Left
	right := root.Right
	if left == nil && right == nil {
		array = append(array, sum)
		return array
	}

	if left != nil {
		array = BranchSumsHelper(left, sum+left.Value, array)
	}

	if right != nil {
		array = BranchSumsHelper(right, sum+right.Value, array)
	}

	return array
}
