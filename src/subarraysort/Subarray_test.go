package subarraysort


import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestSubArrraySortCase1(t *testing.T) {
	expected := []int{3, 9}
	output := SubarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19})
	require.Equal(t, expected, output)
}

func  TestSubArrraySortCase2(t *testing.T) {
	expected := []int{-1,-1}
	output := SubarraySort([]int{1, 2})
	require.Equal(t, expected, output)
}

func  TestSubArrraySortCase3(t *testing.T) {
	expected := []int{0,1}
	output := SubarraySort([]int{2,1})
	require.Equal(t, expected, output)
}

func  TestSubArrraySortCase4(t *testing.T) {
	expected := []int{4, 9}
	output := SubarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 7, 7, 16, 18, 19})
	require.Equal(t, expected, output)
}

func  TestSubArrraySortCase5(t *testing.T) {
	expected := []int{0,2}
	output := SubarraySort([]int{3,2,1})
	require.Equal(t, expected, output)
}

func  TestSubArrraySortCase6(t *testing.T) {
	expected := []int{4, 6}
	output := SubarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 13, 14, 16, 18, 19})
	require.Equal(t, expected, output)
}