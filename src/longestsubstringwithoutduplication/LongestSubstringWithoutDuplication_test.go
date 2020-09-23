package longestsubstringwithoutduplication

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	expected := "mentisac"
	output := LongestSubstringWithoutDuplication("clementisacap")
	require.Equal(t, expected, output)
}

func  TestCase2(t *testing.T) {
	expected := "a"
	output := LongestSubstringWithoutDuplication("a")
	require.Equal(t, expected, output)
}

func  TestCase3(t *testing.T) {
	expected := "abc"
	output := LongestSubstringWithoutDuplication("abc")
	require.Equal(t, expected, output)
}

func  TestCase4(t *testing.T) {
	expected := ""
	output := LongestSubstringWithoutDuplication("")
	require.Equal(t, expected, output)
}

func  TestCase5(t *testing.T) {
	expected := "a"
	output := LongestSubstringWithoutDuplication("aaaaaaaaaa")
	require.Equal(t, expected, output)
}
