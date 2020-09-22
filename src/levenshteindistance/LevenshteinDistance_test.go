package levenshteindistance

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	require.Equal(t, LevenshteinDistance("abc", "yabd"), 2)
}

func  TestCase2(t *testing.T) {
	require.Equal(t, LevenshteinDistance("", ""), 0)
}

func  TestCase3(t *testing.T) {
	require.Equal(t, LevenshteinDistance("", "abc"), 3)
}

func  TestCase4(t *testing.T) {
	require.Equal(t, LevenshteinDistance("abc", "abcx"), 1)
}

func  TestCase5(t *testing.T) {
	require.Equal(t, LevenshteinDistance("algoexpert", "algozexpert"), 1)
}

func  TestCase6(t *testing.T) {
	require.Equal(t, LevenshteinDistance("gumbo", "gambo1"), 2)
}

func  TestCase7(t *testing.T) {
	require.Equal(t, LevenshteinDistance("table", "bal"), 3)
}