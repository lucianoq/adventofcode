package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct{ Start, End int }

func parseInput() ([]Range, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	// Scan ranges
	var ranges []Range
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		ff := strings.Split(line, "-")
		start, _ := strconv.Atoi(ff[0])
		end, _ := strconv.Atoi(ff[1])
		ranges = append(ranges, Range{start, end})
	}

	// Scan IDs
	var ids []int
	for scanner.Scan() {
		line := scanner.Text()
		id, _ := strconv.Atoi(line)
		ids = append(ids, id)
	}

	return ranges, ids
}
