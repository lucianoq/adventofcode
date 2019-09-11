package main

import (
	"fmt"
)

const input = 289326

func main() {
	grid := make(map[Coord]int)

	pos := Coord{0, 0}
	dir := East
	edgeLength := 1

	grid[pos] = 1

	i := 1
	for {
		// every 2 sides increase edgeLength by 1
		for j := 0; j < 2; j++ {

			// for all the edge length
			for k := 0; k < edgeLength; k++ {
				//1 step forward in the right direction,
				pos.Move(dir)

				// compute value
				i = sumAdjacents(grid, pos)

				if i > input {
					fmt.Println(i)
					return
				}

				// draw the cell.
				grid[pos] = i
			}
			dir.RotateLeft90()
		}
		edgeLength++
	}
}

func sumAdjacents(grid map[Coord]int, pos Coord) int {
	sum := 0
	for _, i := range []int{-1, 0, +1} {
		for _, j := range []int{-1, 0, +1} {
			if i|j != 0 {
				sum += grid[Coord{pos.X + i, pos.Y + j}]
			}
		}
	}
	return sum
}
