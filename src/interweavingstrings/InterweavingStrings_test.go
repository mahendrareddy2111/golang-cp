package interweavingstrings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestInterWeavingStringsCase1(t *testing.T) {
	one := "algoexpert"
	two := "your-dream-job"
	three := "your-algodream-expertjob"
	require.Equal(t, InterweavingStrings(one, two, three), true)
}

func  TestInterWeavingStringsCase2(t *testing.T) {
	one := "a"
	two := "b"
	three := "ab"
	require.Equal(t, InterweavingStrings(one, two, three), true)
}

func  TestInterWeavingStringsCase3(t *testing.T) {
	one := "a"
	two := "b"
	three := "ba"
	require.Equal(t, InterweavingStrings(one, two, three), true)
}

func  TestInterWeavingStringsCase4(t *testing.T) {
	one := "a"
	two := "b"
	three := "ac"
	require.Equal(t, InterweavingStrings(one, two, three), false)
}

func  TestInterWeavingStringsCase5(t *testing.T) {
	one := "abc"
	two := "def"
	three := "abcdef"
	require.Equal(t, InterweavingStrings(one, two, three), true)
}

func  TestInterWeavingStringsCase6(t *testing.T) {
	one := "abc"
	two := "def"
	three := "deabcf"
	require.Equal(t, InterweavingStrings(one, two, three), true)
}

func  TestInterWeavingStringsCase7(t *testing.T) {
	one := "aabcc"
	two := "dbbca"
	three := "aadbbcbcac"
	require.Equal(t, InterweavingStrings(one, two, three), true)
}

func  TestInterWeavingStringsCase8(t *testing.T) {
	one := "aabcc"
	two := "dbbca"
	three := "aadbbcbcac"
	require.Equal(t, InterweavingStrings(one, two, three), true)
}