package main

import "fmt"

func main() {
	start := C{0, 0}

	grid := map[C]*Cell{start: {}}

	BuildGrid(grid, start, input)

	// Print(grid)

	fmt.Println(BFS(grid, start))
}

func BFS(grid map[C]*Cell, start C) int {
	nilC := C{1<<63 - 1, 1<<63 - 1}
	end := nilC

	previous := map[C]C{start: nilC}
	discovered := map[C]struct{}{start: {}}
	queue := []C{start}

	check := func(c, neighbor C, dir int) {
		if grid[c].Doors[dir] {
			if _, ok := discovered[neighbor]; !ok {
				discovered[neighbor] = struct{}{}

				// if this is the last one I could discover
				if len(discovered) == len(grid) {
					end = neighbor
				}
				previous[neighbor] = c
				queue = append(queue, neighbor)
			}
		}
	}

	var curr C
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]

		if curr == end {
			i := 0
			for previous[curr] != nilC {
				curr = previous[curr]
				i++
			}
			return i
		}

		check(curr, curr.N(), N)
		check(curr, curr.E(), E)
		check(curr, curr.S(), S)
		check(curr, curr.W(), W)
	}

	return 0
}
