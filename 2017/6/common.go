package main

const size = 16

func maxIdx(arr [size]int) int {
	maxIdx, maxVal := -1, -1
	for i := 0; i < size; i++ {

		// looping from 0 and not including `=` ensure the first is taken
		// that is the tie solving strategy
		if arr[i] > maxVal {
			maxIdx, maxVal = i, arr[i]
		}
	}
	return maxIdx
}
