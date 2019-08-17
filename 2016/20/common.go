package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func merge(ranges []Range, newR Range) []Range {
	for i, r := range ranges {

		// Already included. No need to change the list
		//newR     |-----|
		//r     |-----------|
		if newR.Min >= r.Min && newR.Max <= r.Max {
			//log.Printf("%d-%d is included in %d-%d. Ignoring", newR.Min, newR.Max, r.Min, r.Max)
			return ranges
		}

		// It wraps an old one. Can take its place
		//newR  |-----------|
		//r        |-----|
		if newR.Min <= r.Min && newR.Max >= r.Max {
			//log.Printf("%d-%d is included in %d-%d. Replacing", r.Min, r.Max, newR.Min, newR.Max)
			ranges[i] = newR
			return ranges
		}

		// overlap. Merge them.
		//newR  |-----------|
		//r            |---------|
		if newR.Max >= r.Min-1 && newR.Max <= r.Max {
			//log.Printf("%d-%d is overlapping %d-%d from left. Merging into %d-%d", newR.Min, newR.Max, r.Min, r.Max, newR.Min, r.Max)
			ranges[i] = Range{newR.Min, r.Max}
			return ranges
		}

		// overlap. Merge them.
		//newR        |-----------|
		//r     |---------|
		if newR.Min >= r.Min && newR.Min <= r.Max+1 {
			//log.Printf("%d-%d is overlapping %d-%d from right. Merging into %d-%d", newR.Min, newR.Max, r.Min, r.Max, r.Min, newR.Max)
			ranges[i] = Range{r.Min, newR.Max}
			return ranges
		}
	}
	//log.Println("Didn't found any chance to merge. Appending.")
	return append(ranges, newR)
}

func parse() []Range {
	scanner := bufio.NewScanner(os.Stdin)
	var ranges []Range

	for scanner.Scan() {
		line := scanner.Text()
		xs := strings.Split(line, "-")
		min, _ := strconv.Atoi(xs[0])
		max, _ := strconv.Atoi(xs[1])
		ranges = append(ranges, Range{min, max})
	}
	return ranges
}

func simplify(ranges []Range) []Range {
	tmp := make([]Range, 0)
	changed := true
	for changed {
		changed = false

		for _, r := range ranges {
			tmp = merge(tmp, r)
		}

		changed = len(tmp) != len(ranges)
		ranges = tmp
		tmp = make([]Range, 0)
	}
	return ranges
}
