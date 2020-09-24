package maximumincreasingsubsequence

func MaxSumIncreasingSubsequence(array []int) []interface{} {
	// Write your code here.

	seq := make([]int, len(array))
	sum := make([]int, len(array))
	seq[0] = 0
	sum[0] = array[0]
	ms := sum[0]
	finalSeq := 0

	for i := 1; i < len(array); i++ {
		maxSum := array[i]
		s := i
		for j := i - 1; j >= 0; j-- {
			if array[j] >= array[i] {
				continue
			}
			if maxSum < array[i]+sum[j] {
				maxSum = array[i] + sum[j]
				s = j
			}
		}
		sum[i] = maxSum
		seq[i] = s
		if maxSum > ms {
			ms = maxSum
			finalSeq = i
		}
	}
	result := []int{}
	for finalSeq != seq[finalSeq] {
		result = append(result, finalSeq)
		finalSeq = seq[finalSeq]
	}
	result = append(result, finalSeq)
	resultSeq := []int{}
	for i := len(result) - 1; i >= 0; i-- {
		resultSeq = append(resultSeq, array[result[i]])
	}

	return []interface{}{
		ms,        // Example max sum
		resultSeq, // Example max sequence
	}
}
