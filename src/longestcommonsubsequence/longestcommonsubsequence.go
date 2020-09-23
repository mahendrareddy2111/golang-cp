package longestcommonsubsequence

import "math"

func LongestCommonSubsequence(s1 string, s2 string) string {
	matrix := make([][]int, len(s2)+1)

	for i := range matrix {
		matrix[i] = make([]int, len(s1)+1)
	}

	for i := 1; i < len(matrix); i++ {
		s2r := s2[i-1]
		for j := 1; j < len(matrix[i]); j++ {
			s1r := s1[j-1]
			if s1r != s2r {

				matrix[i][j] = max(matrix[i][j-1], matrix[i-1][j], matrix[i-1][j-1])
			} else {
				matrix[i][j] = max(matrix[i][j-1], matrix[i-1][j], matrix[i-1][j-1]) + 1
			}
		}
	}
	maxSeqLength := matrix[len(s2)][len(s1)]
	if maxSeqLength == 0 {
		return ""
	}
	//temp := 1
	result := ""


	
	return result
}

func max(val ...int) int {
	max := math.MinInt32
	for _, v := range val {
		if v > max {
			max = v
		}
	}
	return max
}
