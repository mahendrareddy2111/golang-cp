package binarysearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinarySearchCase1(t *testing.T) {
	expected := 3
	output := BinarySearch([]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 33)
	require.Equal(t, expected, output)
}

func TestBinarySearchCase2(t *testing.T) {
	expected := 1
	output := BinarySearch([]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 1)
	require.Equal(t, expected, output)
}

func TestBinarySearchCase3(t *testing.T) {
	expected := 3
	output := BinarySearch([]int{1,5,23,111}, 111)
	require.Equal(t, expected, output)
}

func TestBinarySearchCase4(t *testing.T) {
	expected := 1
	output := BinarySearch([]int{1,5,23,111}, 5)
	require.Equal(t, expected, output)
}

func TestBinarySearchCase5(t *testing.T) {
	expected := -1
	output := BinarySearch([]int{1,5,23,111}, 35)
	require.Equal(t, expected, output)
}

func TestBinarySearchCase6(t *testing.T) {
	expected := 0
	output := BinarySearch([]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 0)
	require.Equal(t, expected, output)
}
