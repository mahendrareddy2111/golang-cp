package maximumincreasingsubsequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	output := MaxSumIncreasingSubsequence([]int{10, 70, 20, 30, 50, 11, 30})
	require.Equal(t, []interface{}{110, []int{10, 20, 30, 50}}, output)
}

func  TestCase2(t *testing.T) {
	output := MaxSumIncreasingSubsequence([]int{1})
	require.Equal(t, []interface{}{1, []int{1}}, output)
}	

func  TestCase3(t *testing.T) {
	output := MaxSumIncreasingSubsequence([]int{-1})
	require.Equal(t, []interface{}{-1, []int{-1}}, output)
}	

func  TestCase4(t *testing.T) {
	output := MaxSumIncreasingSubsequence([]int{5,4,3,2,1})
	require.Equal(t, []interface{}{5, []int{5}}, output)
}