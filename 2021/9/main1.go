package main

import "fmt"

func main() {
	grid := parse()

	rows, cols := len(grid), len(grid[0])

	risk := 0
	for i := 0; i < rows; i++ {
	Cell:
		for j := 0; j < cols; j++ {
			for _, adj := range adjacent(grid, P{i, j}) {
				if grid[i][j] >= grid[adj.x][adj.y] {
					continue Cell
				}
			}

			risk += 1 + grid[i][j]
		}
	}

	fmt.Println(risk)
}
