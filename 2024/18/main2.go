package main

import "fmt"

func main() {
	bytes := parseInput()

	corrupted := map[P]struct{}{}
	for _, p := range bytes {
		corrupted[p] = struct{}{}
		if path := BFS(corrupted); path == nil {
			fmt.Printf("%d,%d\n", p.x, p.y)
			return
		}
	}
}
