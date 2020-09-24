package minnumberofcoinsforchange

import "math"

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	coinCount := make([]int, n+1)
	for i := 0; i < len(coinCount); i++ {
		coinCount[i] = math.MaxInt32
	}
	coinCount[0] = 0
	for _, d := range denoms {
		for i := d; i < n+1; i++ {
			coinCount[i] = min(coinCount[i], coinCount[i-d]+1)
		}
	}
	if coinCount[n] == math.MaxInt32 {
		return -1
	}
	return coinCount[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
