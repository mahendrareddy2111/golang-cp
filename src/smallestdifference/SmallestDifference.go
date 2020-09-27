package smallestdifference

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)
	minDiff := math.MaxInt32
	result := make([]int, 2)

	for i := 0; i < len(array1); i++ {
		j := 0
		for ; j < len(array2) && array2[j] <= array1[i]; j++ {
			minDiff,result = getMinDiff(minDiff, array1[i], array2[j], result)
		}
		if j != len(array2) {
			minDiff,result = getMinDiff(minDiff, array1[i], array2[j], result)
		}
	}
	return result
}

func getMinDiff(minDiff, i, j int, result []int) (int,[]int) {
	tempDiff := int(math.Abs(float64(i - j)))
	if minDiff > tempDiff {
		minDiff = tempDiff
		result[0] = i
		result[1] = j
	}
	return minDiff,result
}
