package subarrayproductlessthank

func SubArrayProductLessThanK(array []int, k int)int{
	i , j := 0,0
	product := 1
	result := 0
	for j < len(array){
		if array[j] >= k {
			n := j-i
			result += ((n * (n+1))/2)
			// for product >= k {
			// 	product /= array[i]
			// 	i++
			// }
			j++
			i = j
			product = 1
			continue
		}
		product *= array[j]
		if product >= k {
			
			n := j-i
			result += ((n * (n+1))/2)-1
			for product >= k {
				product /= array[i]
				i++
			}
			//product = 1
		}
		j++
	}
	n := j-i
	result += (n * (n+1))/2
	return result
}