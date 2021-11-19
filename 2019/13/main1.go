package main

import "fmt"

func main() {
	input := make(chan int, 0)
	output := make(chan int, 0)

	go func() {
		NewVM("input", input, output).Run()
		close(output)
	}()

	grid := make(map[C]Tile)

	var blocks int

	for {
		x, open := <-output

		if !open {
			fmt.Println(blocks)
			return
		}

		y := <-output
		tile := Tile(<-output)

		if tile == Block {
			blocks++
		}

		grid[C{x, y}] = tile
	}
}
