package main

import (
	"fmt"
	"maps"
)

func main() {
	ranges := parseInput()

	invalids := map[int]struct{}{}

	for _, r := range ranges {
		for chunkLen := 1; chunkLen <= r.digits/2; chunkLen++ {

			// only divisors
			if r.digits%chunkLen != 0 {
				continue
			}
			
			maps.Copy(invalids, r.InvalidIDs(chunkLen))
		}
	}

	sum := 0
	for n := range invalids {
		sum += n
	}
	fmt.Println(sum)
}
