package twonumbersum

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	result := []int{}
	resultMap := make(map[int]int)
	for i := 0; i < len(array); i++ {
		temp := target - array[i]
		_, ok := resultMap[array[i]]
		if ok {
			result = make([]int, 2)
			result[0] = array[i]
			result[1] = target - array[i]
			return result
		} else {
			resultMap[temp] = array[i]
		}
	}
	return result
}
