package main

import "fmt"

func main() {
	// Stop when we find an 'a'
	goal := func(p P) bool {
		return Map[p] == 'a'
	}

	// avoid going downhill by 2 or more
	forbidden := func(from, to P) bool {
		return Map[from]-Map[to] >= 2
	}

	// Find the first 'a', starting from End
	fmt.Println(bfs(End, goal, forbidden))
}
