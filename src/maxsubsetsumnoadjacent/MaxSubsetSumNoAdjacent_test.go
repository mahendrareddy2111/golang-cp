package maxsubsetsumnoadjacent

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestMaxSubSetCase1(t *testing.T) {
	res := MaxSubsetSumNoAdjacent([]int{75, 105, 120, 75, 90, 135})
	require.Equal(t, res, 330)
}

func  TestMaxSubSetCase2(t *testing.T) {
	res := MaxSubsetSumNoAdjacent([]int{})
	require.Equal(t, res, 0)
}

func  TestMaxSubSetCase3(t *testing.T) {
	res := MaxSubsetSumNoAdjacent([]int{1})
	require.Equal(t, res, 1)
}

func  TestMaxSubSetCase4(t *testing.T) {
	res := MaxSubsetSumNoAdjacent([]int{1,2})
	require.Equal(t, res, 2)
}

func  TestMaxSubSetCase5(t *testing.T) {
	res := MaxSubsetSumNoAdjacent([]int{1,2,3})
	require.Equal(t, res, 4)
}

func  TestMaxSubSetCase6(t *testing.T) {
	res := MaxSubsetSumNoAdjacent([]int{1,15,3})
	require.Equal(t, res, 15)
}

func  TestMaxSubSetCase7(t *testing.T) {
	res := MaxSubsetSumNoAdjacent([]int{7,10,12,7,9,14})
	require.Equal(t, res, 33)
}

