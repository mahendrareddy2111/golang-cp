package twonumbersum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{-1, 11}
	output := TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
	require.ElementsMatch(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := []int{4,6}
	output := TwoNumberSum([]int{4,6}, 10)
	require.ElementsMatch(t, expected, output)
}


func TestCase3(t *testing.T) {
	expected := []int{4,1}
	output := TwoNumberSum([]int{4,6,1}, 5)
	require.ElementsMatch(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := []int{6,-3}
	output := TwoNumberSum([]int{4,6,1,-3}, 3)
	require.ElementsMatch(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := []int{}
	output := TwoNumberSum([]int{4,6,1,-4}, 3)
	require.ElementsMatch(t, expected, output)
}