package riversizes

func RiverSizes(matrix [][]int) []int {
	// Write your code here.
	result := []int{}
	for i:= 0 ; i < len(matrix) ; i++{
		for j := 0 ; j < len(matrix[i]) ;j++{
			if matrix[i][j] == 0 {
				continue
			}
			result = append(result,RiverSizesHelper(matrix,i,j))
		}
	}

	return result
}

func RiverSizesHelper(matrix [][]int,i,j int) int{
	value := 1 
	if  i == len(matrix) ||  i < 0 || j == len(matrix[i])  || j < 0 || matrix[i][j] == 0 {
		return 0
	}
	matrix[i][j] = 0
	value +=  RiverSizesHelper(matrix,i+1,j)
	value +=  RiverSizesHelper(matrix,i,j+1)
	value +=  RiverSizesHelper(matrix,i-1,j)
	value +=  RiverSizesHelper(matrix,i,j-1)
	return value
}