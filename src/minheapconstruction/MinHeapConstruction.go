package minheapconstruction

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
	// Write your code here.

	length := len(array)

	for i := (length / 2) - 1; i >= 0; i-- {

		idx := getMinValue(i, array)
		if idx != i {
			swap(i, idx, array)
		}

	}
}

func swap(i, idx int, array []int) {
	array[i], array[idx] = array[idx], array[i]
	if idx <= (len(array)/2)-1 {
		idx1 := getMinValue(idx, array)
		if idx1 != idx {
			swap(idx, idx1, array)
		}
	}
}

func getMinValue(idx int, array []int) int {

	idx1 := 2*idx + 1
	idx2 := 2*idx + 2

	if idx2 >= len(array) {
		if array[idx1] >= array[idx] {
			return idx
		}
		return idx1
	} else if array[idx1] <= array[idx2] {
		if array[idx1] >= array[idx] {
			return idx
		}
		return idx1
	}
	if array[idx2] >= array[idx] {
		return idx
	}
	return idx2
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {
	// Write your code here.
}

func (h *MinHeap) siftUp() {
	// Write your code here.
}

func (h MinHeap) Peek() int {
	// Write your code here.
	return h[0]
}

func (h MinHeap) Last() int {
	// Write your code here.
	return h[len(h)-1]
}

func (h *MinHeap) Remove() int {
	// Write your code here.
	value := (*h)[0]
	last := (*h)[len(*h)-1]
	(*h)[0] = last
	idx := getMinValue(0, *h)
	if idx != 0 {
		swap(0, idx, *h)
	}
	return value
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	idx := (len(*h) / 2) - 1
	// idx1 := getMinValue(idx, *h)

	// if idx1 == idx {
	// 	return
	// }
	for idx >= 0 {
		idx1 := getMinValue(idx, *h)

		if idx1 == idx {
			return
		}
		swap(idx, idx1, *h)
		idx1 = idx
		idx = (idx-1)/ 2 
	}
}
