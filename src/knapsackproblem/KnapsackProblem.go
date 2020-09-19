package knapsackproblem


func KnapsackProblem(items [][]int, capacity int) []interface{} {
	// Write your code here.

	valueWeightIndexedMap := make(map[int][]int)

	for idx, value := range items {
		valueWeightIndexedMap[idx+1] = value
	}

	matrix := make([][]int, len(items)+1)
	for i := range matrix {
		matrix[i] = make([]int, capacity+1)
	}

	for i := 1; i <= len(items); i++ {
		for j := 0; j <= capacity; j++ {
			valueWeight := valueWeightIndexedMap[i]
			value := valueWeight[0]
			weight := valueWeight[1]
			if j < weight {
				matrix[i][j] = matrix[i-1][j]
				continue
			}
			preMaxValue := matrix[i-1][j]
			value += matrix[i-1][j-weight]
			if value > preMaxValue {
				matrix[i][j] = value
			} else {
				matrix[i][j] = preMaxValue
			}

		}
	}

	// for i := 0; i <= len(items); i++ {
	// 	for j := 0; j <= capacity; j++ {
	// 		fmt.Printf("%d\t", matrix[i][j])
	// 	}
	// 	fmt.Println()
	// }

	sequence := findSequence(matrix, items)

	// Replace return below.
	return []interface{}{
		matrix[len(items)][capacity],          // total value
		sequence, // item indices
	}
}

func findSequence(matrix, items [][]int) []int {

	sequence := []int{}

	i, c := len(matrix)-1, len(matrix[0])-1

	for i > 0 {
		if matrix[i][c] == matrix[i-1][c] {
			i--
		} else {
			sequence = append(sequence, i-1)
			c -= items[i-1][1]
			i--
		}
		if c == 0 {
			break
		}
	}
	reverse(sequence)

	return sequence

}

func reverse(numbers []int){
	for i,j := 0,len(numbers)-1; i < j ; i ,j = i+1,j-1{
		numbers[i],numbers[j] = numbers[j],numbers[i]
	}
}
