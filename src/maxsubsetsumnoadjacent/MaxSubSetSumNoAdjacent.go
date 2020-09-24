package maxsubsetsumnoadjacent

import "math"

func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if len(array) == 1 {
		return array[0]
	}
	maxSum := math.MinInt32
	
	if array[0] < array[1] {
		maxSum = array[1]

	} else {
		maxSum = array[0]
		array[1] = array[0]
	}
	for i := 2; i < len(array); i++ {
		tempSum := array[i] + array[i-2]
		if tempSum > maxSum{
			maxSum = tempSum
		}
		array[i] = maxSum
	}
	return maxSum
}
