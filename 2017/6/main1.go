package main

import "fmt"

func main() {
	var input = [size]int{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6}

	visited := make(map[[size]int]struct{})

	for time := 1; ; time++ {
		idx := maxIdx(input) // find the max
		blocks := input[idx] // num of blocks to reallocate
		input[idx] = 0       // empty the bank

		// reallocate
		for blocks > 0 {
			idx = (idx + 1) % size // move index forward

			// assign the block
			input[idx]++
			blocks--
		}

		if _, ok := visited[input]; ok {
			fmt.Println(time)
			return
		}

		visited[input] = struct{}{}
	}
}
