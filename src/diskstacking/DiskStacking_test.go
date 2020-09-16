package diskstacking

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiskStackingCase1(t *testing.T) {
	expected := [][]int{{2, 1, 2}, {3, 2, 3}, {4, 4, 5}}
	input := [][]int{{2, 1, 2}, {3, 2, 3}, {2, 2, 8}, {2, 3, 4}, {2, 2, 1}, {4, 4, 5}}
	output := DiskStacking(input)
	require.Equal(t, expected, output)
}

func TestDiskStackingCase2(t *testing.T) {
	expected := [][]int{{2, 1, 2}}
	input := [][]int{{2, 1, 2}}
	output := DiskStacking(input)
	require.Equal(t, expected, output)
}

func TestDiskStackingCase3(t *testing.T) {
	expected := [][]int{{2, 1, 2}, {3, 2, 3}}
	input := [][]int{{2, 1, 2}, {3, 2, 3}}
	output := DiskStacking(input)
	require.Equal(t, expected, output)
}

func TestDiskStackingCase4(t *testing.T) {
	expected := [][]int{{2,2,8}}
	input := [][]int{{2, 1, 2}, {3, 2, 3},{2,2,8}}
	output := DiskStacking(input)
	require.Equal(t, expected, output)
}

func TestDiskStackingCase5(t *testing.T) {
	expected := [][]int{{2, 1, 2}, {3, 2, 3}}
	input := [][]int{{2, 1, 2}, {3, 2, 3},{2,3,4}}
	output := DiskStacking(input)
	require.Equal(t, expected, output)
}
