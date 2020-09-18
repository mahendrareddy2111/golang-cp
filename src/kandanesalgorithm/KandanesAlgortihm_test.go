package kandanesalgorithm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	expected := 19
	output := KadanesAlgorithm([]int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4})
	require.Equal(t, expected, output)
}

func  TestCase2(t *testing.T) {
	expected := 8
	output := KadanesAlgorithm([]int{3, 5, -9})
	require.Equal(t, expected, output)
}

func  TestCase3(t *testing.T) {
	expected := -3
	output := KadanesAlgorithm([]int{-3, -5, -9})
	require.Equal(t, expected, output)
}
