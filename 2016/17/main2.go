package main

import "fmt"

func main() {
	var paths []string

	var p = Point{0, 0, ""}
	toVisit := []Point{p}

	for len(toVisit) > 0 {
		p, toVisit = toVisit[0], toVisit[1:]

		if p.X == 3 && p.Y == 3 {
			paths = append(paths, p.Path)
			continue
		}

		for _, adj := range adjacents(p) {
			toVisit = append(toVisit, adj)
		}
	}

	fmt.Println(findMax(paths))
}

func findMax(paths []string) int {
	maxL := 0
	for _, p := range paths {
		if len(p) > maxL {
			maxL = len(p)
		}
	}
	return maxL
}
