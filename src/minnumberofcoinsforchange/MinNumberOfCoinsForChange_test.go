package minnumberofcoinsforchange

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	input := []int{1, 5, 10}
	actual := MinNumberOfCoinsForChange(7, input)
	require.Equal(t, 3, actual)
}

func  TestCase2(t *testing.T) {
	input := []int{1,2,3}
	actual := MinNumberOfCoinsForChange(0, input)
	require.Equal(t, 0, actual)
}

func  TestCase3(t *testing.T) {
	input := []int{2,1}
	actual := MinNumberOfCoinsForChange(3, input)
	require.Equal(t, 2, actual)
}

func  TestCase4(t *testing.T) {
	input := []int{1,5,10}
	actual := MinNumberOfCoinsForChange(4, input)
	require.Equal(t, 4, actual)
}

func  TestCase5(t *testing.T) {
	input := []int{1,5,10}
	actual := MinNumberOfCoinsForChange(10, input)
	require.Equal(t, 1, actual)
}

func  TestCase6(t *testing.T) {
	input := []int{1,5,10}
	actual := MinNumberOfCoinsForChange(11, input)
	require.Equal(t, 2, actual)
}

func  TestCase7(t *testing.T) {
	input := []int{1,5,10}
	actual := MinNumberOfCoinsForChange(24, input)
	require.Equal(t, 6, actual)
}

func  TestCase8(t *testing.T) {
	input := []int{1,5,10}
	actual := MinNumberOfCoinsForChange(25, input)
	require.Equal(t, 3, actual)
}