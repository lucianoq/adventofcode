package main

import "fmt"

const (
	destX = 31
	destY = 39
)

func main() {
	var p = Point{1, 1, 0}

	found := make(map[int]bool)
	found[toInt(p)] = true

	var toVisit []Point
	toVisit = append(toVisit, p)

	for len(toVisit) > 0 {
		p, toVisit = toVisit[0], toVisit[1:]

		// check final node
		if p.X == destX && p.Y == destY {
			fmt.Println(p.Steps)
			return
		}

		for _, adj := range adjacents(p) {
			if !found[toInt(adj)] {
				found[toInt(adj)] = true
				toVisit = append(toVisit, adj)
			}
		}
	}
}
