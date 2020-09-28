package validatesubsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	
	j := 0 
	for i := 0 ; i < len(array) ; i++{
		if array[i] == sequence[j] {
			j++
		}
		if j == len(sequence){
			return true
		}
	} 
	return j == len(sequence)
}