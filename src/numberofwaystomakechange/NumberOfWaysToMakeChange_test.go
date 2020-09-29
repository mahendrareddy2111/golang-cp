package numberofwaystomakechange

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	output := NumberOfWaysToMakeChange(6, []int{1, 5})
	require.Equal(t, 2, output)
}

func  TestCase2(t *testing.T) {
	output := NumberOfWaysToMakeChange(0, []int{2,3,4,7})
	require.Equal(t, 1, output)
}

func  TestCase3(t *testing.T) {
	output := NumberOfWaysToMakeChange(9, []int{5})
	require.Equal(t, 0, output)
}

func  TestCase4(t *testing.T) {
	output := NumberOfWaysToMakeChange(7, []int{2,4})
	require.Equal(t, 0, output)
}

func  TestCase5(t *testing.T) {
	output := NumberOfWaysToMakeChange(4, []int{1,5,10,25})
	require.Equal(t, 1, output)
}

func  TestCase6(t *testing.T) {
	output := NumberOfWaysToMakeChange(5, []int{1,5,10,25})
	require.Equal(t, 2, output)
}

func  TestCase7(t *testing.T) {
	output := NumberOfWaysToMakeChange(10, []int{1,5,10,25})
	require.Equal(t, 4, output)
}

func  TestCase8(t *testing.T) {
	output := NumberOfWaysToMakeChange(25, []int{1,5,10,25})
	require.Equal(t, 13, output)
}

func  TestCase9(t *testing.T) {
	output := NumberOfWaysToMakeChange(12, []int{2,3,7})
	require.Equal(t, 4, output)
}

func  TestCase10(t *testing.T) {
	output := NumberOfWaysToMakeChange(7, []int{2,3,4,7})
	require.Equal(t, 3, output)
}
