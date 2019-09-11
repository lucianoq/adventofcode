package main

import "fmt"

const input = 289326

func main() {
	grid := make(map[int]Coord)

	pos := Coord{0, 0}
	dir := East
	edgeLength := 1

	grid[1] = pos

	// Stop when we reach our goal
	for i := 1; i < input; {

		// every 2 sides increase edgeLength by 1
		for j := 0; j < 2; j++ {

			// for all the edge length
			for k := 0; k < edgeLength; k++ {
				//1 step forward in the right direction,
				pos.Move(dir)

				// increase the counter,
				i++

				// draw the cell.
				grid[i] = pos
			}
			dir.RotateLeft90()
		}
		edgeLength++
	}

	goal := grid[input]

	fmt.Println(abs(goal.X) + abs(goal.Y))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
