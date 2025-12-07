package main

import "fmt"

func main() {
	splitters, start, last := parseInput()
	beams := map[P]int{P{start.x + 1, start.y}: 1}
	for i := 2; i < last; i++ {
		nextBeams := map[P]int{}
		for b, journeys := range beams {
			if _, ok := splitters[P{i, b.y}]; ok {
				nextBeams[P{i, b.y - 1}] += journeys
				nextBeams[P{i, b.y + 1}] += journeys
			} else {
				nextBeams[P{i, b.y}] += journeys
			}
		}
		beams = nextBeams
	}

	paths := 0
	for _, journeys := range beams {
		paths += journeys
	}

	fmt.Println(paths)
}
