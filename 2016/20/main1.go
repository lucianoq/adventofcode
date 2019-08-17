package main

import (
	"fmt"
	"math"
)

func main() {
	ranges := parse()

	// After simplify I'll know that all the ranges are not overlapping
	// and also they're not consecutive (because I merge consecutive ranges already)
	ranges = simplify(ranges)

	// There is always at least 1 free ip between ranges
	// So, I just need to find the smallest Max and take the next ip.
	smallestMax := math.MaxInt32
	for _, r := range ranges {
		if r.Max < smallestMax {
			smallestMax = r.Max
		}
	}

	fmt.Println(smallestMax + 1)
}
