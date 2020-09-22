package levenshteindistance

import (
	"math"
)

func LevenshteinDistance(a, b string) int {
matrix := make([][]int,len(a)+1)
for i:= 0 ; i < len(matrix) ; i++{
	matrix[i] = make([]int,len(b)+1)
	matrix[i][0] = i
}

for  j := 0 ; j < len(b)+1; j++{
	matrix[0][j] = j
}

for i := 1 ; i < len(a)+1 ; i++{
	for j := 1 ; j < len(b)+1 ; j++{
		if a[i-1] == b[j-1]{
			matrix[i][j] = matrix[i-1][j-1]
		}else{
			matrix[i][j] = 1+min(matrix[i-1][j],matrix[i][j-1],matrix[i-1][j-1])
		}
		
	}
}

return matrix[len(a)][len(b)]
}

func min(val ...int)int{
min := math.MaxInt32
for _,v := range val{
	if ( v < min ){
		min = v
	}
}
return min
}


