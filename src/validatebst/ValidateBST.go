package validatebst

import (
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Validate() bool {

	return tree.ValidateHelper(math.MinInt32, math.MaxInt32)

}

func (tree *BST) ValidateHelper(min, max int) bool {
	// Write your code here.
	if tree.Value < min || tree.Value >= max {
		return false
	}

	if tree.Left != nil && !tree.Left.ValidateHelper(min, tree.Value) {
		return false
	}

	if tree.Right != nil && !tree.Right.ValidateHelper(tree.Value, max) {
		return false
	}

	return true

}
