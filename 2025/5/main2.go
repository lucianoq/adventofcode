package main

import (
	"fmt"
	"sort"
)

func main() {
	ranges, _ := parseInput()

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].Start == ranges[j].Start {
			// when they have the same start, prefer bigger ranges on the left,
			// so `contains` will eat them at the beginning of the next loop
			return ranges[i].End > ranges[j].End
		}
		return ranges[i].Start < ranges[j].Start
	})

	// index of the Work in Progress range
	wipIdx := 0

	for i := 1; i < len(ranges); i++ {
		wipRange := ranges[wipIdx]

		// if the WIP contains the current, we ignore the current
		if wipRange.contains(ranges[i]) {
			continue
		}

		// if they touch, we merge them inside the WIP
		if wipRange.overlaps(ranges[i]) {
			ranges[wipIdx] = wipRange.append(ranges[i])
			continue
		}

		// disjoint ranges
		// freeze the current WIP by focusing on the next one
		// that is the current range
		wipIdx++
		ranges[wipIdx] = ranges[i]
	}

	ranges = ranges[:wipIdx+1]

	sum := 0
	for _, r := range ranges {
		sum += r.Len()
	}

	fmt.Println(sum)
}

func (r Range) overlaps(r2 Range) bool {
	return r.Start <= r2.Start && r.End >= r2.Start
}

func (r Range) contains(r2 Range) bool {
	return r.Start <= r2.Start && r.End >= r2.End
}

func (r Range) append(r2 Range) Range {
	return Range{r.Start, r2.End}
}

func (r Range) Len() int {
	return r.End - r.Start + 1
}
