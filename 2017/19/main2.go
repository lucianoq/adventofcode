package main

import (
	"fmt"
	"os"
)

func main() {
	grid := NewGrid()

	for {
		grid.Run(func() {
			fmt.Println(grid.Steps)
			os.Exit(0)
		})
	}
}
