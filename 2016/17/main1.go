package main

import (
	"fmt"
)

func main() {
	var p = Point{0, 0, ""}
	toVisit := []Point{p}

	for len(toVisit) > 0 {
		p, toVisit = toVisit[0], toVisit[1:]

		if p.X == 3 && p.Y == 3 {
			fmt.Println(p.Path)
			return
		}

		for _, adj := range adjacents(p) {
			toVisit = append(toVisit, adj)
		}
	}
}
