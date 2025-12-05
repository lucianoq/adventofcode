package main

import "fmt"

func main() {
	ranges, ids := parseInput()

	count := 0
	for _, id := range ids {
		if fresh(ranges, id) {
			count++
		}
	}

	fmt.Println(count)
}

func fresh(ranges []Range, id int) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}
	return false
}
