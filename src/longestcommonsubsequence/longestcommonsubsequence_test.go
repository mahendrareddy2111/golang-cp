package longestcommonsubsequence


import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLongestCommonSubSeqCase1(t *testing.T) {
	expected := "XYZW"
	output := LongestCommonSubsequence("ZXVVYZW", "XKYKZPW")
	require.Equal(t, expected, output)
}

func TestLongestCommonSubSeqCase2(t *testing.T) {
	expected := "AE"
	output := LongestCommonSubsequence("ABCDEFG", "APPLES")
	require.Equal(t, expected, output)
}

func TestLongestCommonSubSeqCase3(t *testing.T) {
	expected := "NT"
	output := LongestCommonSubsequence("CLEMENT", "ANTONIE")
	require.Equal(t, expected, output)
}

func TestLongestCommonSubSeqCase4(t *testing.T) {
	expected := "842"
	output := LongestCommonSubsequence("8111111111111111142", "222222222822222222222222222222433333333332")
	require.Equal(t, expected, output)
}