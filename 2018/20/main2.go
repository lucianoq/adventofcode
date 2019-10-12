package main

import "fmt"

func main() {
	start := C{0, 0}

	grid := map[C]*Cell{start: {}}

	BuildGrid(grid, start, input)

	// Print(grid)

	fmt.Println(DFS(grid, map[C]struct{}{}, start, 0))
}

func DFS(grid map[C]*Cell, discovered map[C]struct{}, c C, layer int) int {
	numFarRooms := 0

	discovered[c] = struct{}{}

	if layer >= 1000 {
		numFarRooms++
	}

	check := func(neighbor C, dir int) {
		if grid[c].Doors[dir] {
			if _, ok := discovered[neighbor]; !ok {
				numFarRooms += DFS(grid, discovered, neighbor, layer+1)
			}
		}
	}

	check(c.N(), N)
	check(c.E(), E)
	check(c.S(), S)
	check(c.W(), W)

	return numFarRooms
}
