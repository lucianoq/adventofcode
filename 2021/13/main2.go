package main

import "fmt"

func main() {
	grid, folds := parse()

	for _, f := range folds {
		fold(grid, f)
	}

	fmt.Println(toString(grid.grid, grid.width, grid.height, 5, 0))
}
