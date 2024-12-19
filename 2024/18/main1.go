package main

import "fmt"

func main() {
	bytes := parseInput()

	corrupted := map[P]struct{}{}
	for _, p := range bytes[:1024] {
		corrupted[p] = struct{}{}
	}

	fmt.Println(len(BFS(corrupted)))
}
