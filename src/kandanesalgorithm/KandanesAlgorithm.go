package kandanesalgorithm

func KadanesAlgorithm(array []int) int {
	// Write your code here.
	maxSum := array[0]
	maxSumSoFar := array[0]
	for i:= 1 ; i < len(array) ;i++{
		maxSumSoFar= max(array[i],maxSumSoFar+array[i])
		maxSum = max(maxSum,maxSumSoFar)
	}
	return maxSum
}

func max (a,b int) int{
	if a > b {
		return a
	}
	return b
}