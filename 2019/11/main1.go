package main

import (
	"fmt"
	"os"
)

func main() {
	grid := make(map[C]int)

	run(grid)
}

// if output channel is closed,
// program stopped to work
func exit(closed bool, grid map[C]int) {
	if closed {
		fmt.Println(len(grid))
		os.Exit(0)
	}
}
