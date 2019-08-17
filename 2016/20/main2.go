package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	ranges := parse()

	// After simplify I'll know that all the ranges are not overlapping
	// and also they're not consecutive (because I merge consecutive ranges already)
	ranges = simplify(ranges)

	// Now they are all non-overlapping. So I just need to order the slice
	// And count the differences between a Max and the following Min.
	// Last Max is MaxUint32.

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	count := 0
	for i := 0; i < len(ranges)-1; i++ {
		count += ranges[i+1].Min - ranges[i].Max - 1
	}
	count += math.MaxUint32 - ranges[len(ranges)-1].Max

	fmt.Println(count)
}
