package validatesubsequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	require.True(t, IsValidSubsequence(array, sequence))
}

func  TestCase2(t *testing.T) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{5, 1, 22, 25, 6, -1, 8, 10}
	require.True(t, IsValidSubsequence(array, sequence))
}

func  TestCase3(t *testing.T) {
	array := []int{5, 1, 22, 25, -1, 8, 10}
	sequence := []int{5, 1, 22, 25, 6, -1, 8, 10}
	require.False(t, IsValidSubsequence(array, sequence))
}

func  TestCase4(t *testing.T) {
	array := []int{5, 1, 22, 25,6,9, -1, 8, 10}
	sequence := []int{5, 1, 22, 25, 6, -1, 8, 10,11}
	require.False(t, IsValidSubsequence(array, sequence))
}

func  TestCase5(t *testing.T) {
	array := []int{}
	sequence := []int{}
	require.True(t, IsValidSubsequence(array, sequence))
}
