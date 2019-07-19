package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const size = 100

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := &[size][size]bool{}
	var row, col int

	for scanner.Scan() {
		col = 0
		line := strings.TrimSpace(scanner.Text())
		if len(line) != size {
			log.Fatal("line not size chars")
		}
		for _, c := range line {
			if c == '#' {
				grid[row][col] = true
			}
			col++
		}
		row++
	}

	for i := 0; i < size; i++ {
		grid = step(grid)

		// Uncomment following to see the animation in the terminal
		//fmt.Print("\033[H\033[2J") //clear the screen
		//printGrid(grid)
		//time.Sleep(800 * time.Millisecond)
	}

	fmt.Println(countTrue(grid))
}

// Conway's Game of Life rules.
func step(in *[size][size]bool) *[size][size]bool {
	out := &[size][size]bool{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			neighborsOn := neighborsOn(in, i, j)

			if in[i][j] {
				switch neighborsOn {
				case 2, 3:
					out[i][j] = true
				default:
					out[i][j] = false
				}
			} else {
				if neighborsOn == 3 {
					out[i][j] = true
				} else {
					out[i][j] = false
				}
			}
		}
	}
	out[0][0] = true
	out[0][size-1] = true
	out[size-1][0] = true
	out[size-1][size-1] = true
	return out
}

// count the neighbors that are turned on
func neighborsOn(in *[size][size]bool, x int, y int) int {
	trues := 0

	// -1, 0, +1 in both axis
	for δx := -1; δx < 2; δx++ {
		x2 := x + δx

		// if out of bounds, count as off.
		if x2 < 0 || x2 >= size {
			continue
		}

		for δy := -1; δy < 2; δy++ {
			y2 := y + δy

			// if out of bounds, count as off.
			if y2 < 0 || y2 >= size {
				continue
			}

			// ignore the cell itself
			if x == x2 && y == y2 {
				continue
			}

			// count if light is on
			if in[x2][y2] {
				trues++
			}
		}
	}

	return trues
}

func countTrue(in *[size][size]bool) int {
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if in[i][j] {
				count++
			}
		}
	}
	return count
}

func printGrid(in *[size][size]bool) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if in[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
