package main

import (
	"fmt"
)

const (
	input int32 = 34000000
	
	// Elf #[3.4M] alone will deliver 34M presents to house #[3.4M]
	// that means that house #[3.4M] is a solution, maybe not the best.
	// Looking for something smaller, we should iterate no more than
	// input/10 houses and input/10 elves.
	limit = input / 10
)

func main() {

	// This slice is (4 * 3.4M) B = 13.6 MB =~ 13 MiB. We can handle it
	houses := make([]int32, limit)

	var e, h int32 = 1, 1
	for e = 1; e < limit; e++ {
		for h = e; h < limit; h += e {
			houses[h] += 10 * e
		}
	}

	for i, h := range houses {
		if h >= input {
			fmt.Println(i)
			return
		}
	}
}
