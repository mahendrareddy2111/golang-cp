package minheapconstruction

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func isMinHeapPropertySatisfied(heap MinHeap) bool {
	for i := 1; i < len(heap); i++ {
		parentIdx := (i - 1) / 2
		if parentIdx < 0 {
			return true
		}

		if heap[parentIdx] > heap[i] {
			return false
		}
	}
	return true
}

func TestMinHeapConstructionCase1(t *testing.T) {
	var minHeap = NewMinHeap([]int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41})
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	minHeap.Insert(76)
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	require.Equal(t, -5, minHeap.Peek())
	 require.Equal(t, -5, minHeap.Remove())
	 require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	 require.Equal(t, 2, minHeap.Peek())
	 require.Equal(t, 2, minHeap.Remove())
	 require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	 require.Equal(t, 6, minHeap.Peek())
	 minHeap.Insert(87)
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
}


func TestMinHeapConstructionCase2(t *testing.T) {
	var minHeap = NewMinHeap([]int{-7,2,3,8,-10,4,-6,-10,-2,-7,10,5,2,9,-9,-5,3,8})
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	require.Equal(t, -10, minHeap.Remove())
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	 require.Equal(t, -10, minHeap.Peek())
	 require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	 minHeap.Insert(-8)
	 require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	 require.Equal(t, -10, minHeap.Peek())
	 require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	 require.Equal(t, -10, minHeap.Remove())
	require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
	 require.Equal(t, -9, minHeap.Peek())
	 minHeap.Insert(8)
	 require.Equal(t, -9, minHeap.Peek())

}
