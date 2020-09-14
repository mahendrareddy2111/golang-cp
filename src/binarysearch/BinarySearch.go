package binarysearch

func BinarySearch(array []int, target int) int {
	first, last := 0, len(array)-1

	for first <= last {

		if first == last && array[first] != target {
			return -1
		}

		mid := (first + last) / 2

		if target == array[mid] {
			return mid
		}

		if target > array[mid] {
			first = mid+1
		} else {
			last = mid-1 
		}
	}

	return -1
}
