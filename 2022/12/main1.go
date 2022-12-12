package main

import "fmt"

func main() {
	// Stop when we find the end
	goal := func(p P) bool {
		return p == End
	}

	// avoid going uphill by 2 or more
	forbidden := func(from, to P) bool {
		return Map[to]-Map[from] >= 2
	}

	// find the End
	fmt.Println(bfs(Start, goal, forbidden))
}
