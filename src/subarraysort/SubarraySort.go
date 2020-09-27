package subarraysort

import "math"

func SubarraySort(array []int) []int {
	result := []int{-1, -1}
	smallestOOO := math.MaxInt32
	largestOOO := math.MinInt32
	outOfRange := false

	for i := 1; i < len(array); i++ {

		if array[i-1] > array[i] {
			outOfRange = true
			if smallestOOO > array[i] {
				smallestOOO = array[i]
			}

			if largestOOO < array[i-1] {
				largestOOO = array[i-1]
			}
		}
	}
	sIdx := -1
	lIdx := -1
	if outOfRange {
		for i := 0; i < len(array); i++ {
			if sIdx == -1 && array[i] > smallestOOO {
				sIdx = i
			}

			if lIdx == -1 && array[i] > largestOOO {
				lIdx = i-1
			}
		}

		if sIdx != -1 && lIdx == -1{
			lIdx = len(array)-1
		}
	}
	result[0] = sIdx
	result[1] = lIdx
	return result
}
