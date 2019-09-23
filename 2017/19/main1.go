package main

import (
	"fmt"
	"os"
)

func main() {
	grid := NewGrid()

	for {
		grid.Run(func() {
			fmt.Println(grid.Word)
			os.Exit(0)
		})
	}
}
