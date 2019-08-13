package main

import "fmt"

const maxSteps = 50

func main() {
	var p = Point{1, 1, 0}

	found := make(map[int]bool)
	found[toInt(p)] = true

	var toVisit []Point
	toVisit = append(toVisit, p)

	for len(toVisit) > 0 {
		p, toVisit = toVisit[0], toVisit[1:]

		// do not continue after 50 steps
		if p.Steps == maxSteps {
			continue
		}

		for _, adj := range adjacents(p) {
			if !found[toInt(adj)] {
				found[toInt(adj)] = true
				toVisit = append(toVisit, adj)
			}
		}
	}
	fmt.Println(len(found))
}
