package numberofwaystomakechange


func NumberOfWaysToMakeChange(n int, denoms []int) int {
	// Write your code here.
	ways := make([]int,n+1)
	ways[0] = 1
	for _,d := range denoms{
		for i:= d ; i < n+1 ; i++{
			ways[i] += ways[i-d]
		}
	}
	return ways[n]
}