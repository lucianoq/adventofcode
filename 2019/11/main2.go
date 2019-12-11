package main

import (
	"fmt"
	"os"
)

func main() {
	grid := make(map[C]int)

	// first position must be white
	grid[C{0, 0}] = White

	run(grid)
}

// if output channel is closed,
// program stopped to work
func exit(closed bool, grid map[C]int) {
	if closed {
		print(grid)
		os.Exit(0)
	}
}

func print(grid map[C]int) {
	// image is mirrored, so print upside down
	for j := 1; j > -6; j-- {

		// left to right
		for i := 0; i < 41; i++ {

			switch grid[C{i, j}] {
			case Black:
				fmt.Print(" ")
			case White:
				fmt.Print("█")
			}

		}
		fmt.Println()
	}
}
