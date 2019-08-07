package main

import (
	"fmt"
)

const (
	input int32 = 34000000

	// Elf #limit alone will deliver 34M+ presents to house #limit.
	// that means that house #limit is a solution, maybe not the best.
	// Looking for something smaller, we should iterate no more than
	// `limit` houses and `limit` elves.
	limit = input/11 + 1
)

func main() {

	// This slice is (4 * 3090910 ) bytes = 12363640 bytes =~ 11.8 MiB
	// We can handle it
	houses := make([]int32, limit)

	var e, h int32 = 1, 1
	for e = 1; e < limit; e++ {
		c := 0
		for h = e; c < 50 && h < limit; h += e {
			houses[h] += 11 * e
			c++
		}
	}

	for i, h := range houses {
		if h >= input {
			fmt.Println(i)
			return
		}
	}
}
