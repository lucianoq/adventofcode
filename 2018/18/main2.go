package main

import "fmt"

const minutes = 1000000000

func main() {
	area := parse()
	// Print(area)

	visited := map[[size][size]rune]int{}

	for i := 0; i < minutes; i++ {
		area = Minute(area)

		if visited[area] != 0 {
			// Found an already seen configuration

			// first occurrence of that configuration
			first := visited[area]
			// second occurrence of that configuration
			second := i
			// period
			period := second - first

			// fast forward i:
			// - skip `first` items (because they were not periodic)
			// - find how many times the period is included in the remaining
			// - minutes
			// - multiply by period to fill the whole span with full iterations
			// - re-add the beginning offset
			i = ((minutes-first)/period)*period + first

			// reset `visited` in order to skip this branch again
			// we are not supposed to find duplicates in the remaining
			// iterations
			visited = map[[size][size]rune]int{}
		}
		visited[area] = i
	}

	fmt.Println(CountResources(area))
}
