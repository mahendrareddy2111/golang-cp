package subarrayproductlessthank

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	input := []int{10,5,2,6}
	require.Equal(t, SubArrayProductLessThanK(input,100), 8)
}

func  TestCase2(t *testing.T) {
	input := []int{1,5,2,6}
	require.Equal(t, SubArrayProductLessThanK(input,100), 10)
}

func  TestCase3(t *testing.T) {
	input := []int{1,5,100,6}
	require.Equal(t, SubArrayProductLessThanK(input,100), 4)
}

func  TestCase4(t *testing.T) {
	input := []int{1,5,100,3,100,6}
	require.Equal(t, SubArrayProductLessThanK(input,100), 5)
}

func  TestCase5(t *testing.T) {
	input := []int{101,100,100,3,100,106}
	require.Equal(t, SubArrayProductLessThanK(input,100), 1)
}

func  TestCase6(t *testing.T) {
	input := []int{101,100,100,108,100,106}
	require.Equal(t, SubArrayProductLessThanK(input,100), 0)
}

func  TestCase7(t *testing.T) {
	input := []int{}
	require.Equal(t, SubArrayProductLessThanK(input,100), 0)
}

func  TestCase8(t *testing.T) {
	input := []int{100}
	require.Equal(t, SubArrayProductLessThanK(input,100), 0)
}

func  TestCase9(t *testing.T) {
	input := []int{100,1}
	require.Equal(t, SubArrayProductLessThanK(input,100), 1)
}

func  TestCase10(t *testing.T) {
	input := []int{100,1,0}
	require.Equal(t, SubArrayProductLessThanK(input,100), 3)
}