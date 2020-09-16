package diskstacking

import (
	"sort"
)

type Disk []int
type Disks []Disk

func (d Disks) Len() int{
	return len(d)
}

func (d Disks) Swap(i,j int){
	d[i],d[j] = d[j],d[i]
}

func (d Disks) Less(i,j int) bool{
	return d[i][2] < d[j][2]
}



func DiskStacking(input [][]int) [][]int {
	// Write your code here.
	disks := make(Disks,len(input))
	for i,value := range input{
		disks[i] = value
	}
	sort.Sort(disks)
	height := make([]int,len(input))
	height[0] = disks[0][2]
	sequence := make([]int,len(input))
	sequence[0] = 0
	highestStack := height[0]
	highestSeq := 0
	for i := 1 ; i < len(disks) ; i++{
		len := disks[i][2]
		prevLen := len
		seq := i
		for j := i-1 ; j >= 0 ; j--{
			if checkForDiskValidity(i,j,disks){
				temp := len+height[j]
				if prevLen < temp{
					prevLen = temp
					seq = j
				}
			}
		}
		height[i] = prevLen
		sequence[i] = seq

		if highestStack < height[i]{
			highestStack = height[i]
			highestSeq = i
		}
	}

	result := []int{highestSeq}
	for highestSeq != sequence[highestSeq]{
		highestSeq = sequence[highestSeq]
		result = append(result,highestSeq)
	}

	resultDiskStack := [][]int{}

	for i := len(result)-1 ; i >= 0 ; i--{
		resultDiskStack = append(resultDiskStack,disks[result[i]])
	}


	return resultDiskStack
}

func checkForDiskValidity(i,j int,disks Disks) bool{
	if disks[i][0] <= disks[j][0]{
		return false
	}
	if disks[i][1] <= disks[j][1]{
		return false
	}
	if disks[i][2] <= disks[j][2]{
		return false
	}
	return true
}