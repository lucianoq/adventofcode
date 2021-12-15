package main

import (
	"bufio"
	"os"
	"strconv"
)

func parse() map[C]int {
	grid := map[C]int{}

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < Size; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j := 0; j < Size; j++ {

			num, _ := strconv.Atoi(line[j : j+1])

			for ri := 0; ri < 5; ri++ {
				for rj := 0; rj < 5; rj++ {

					p := C{ri*Size + i, rj*Size + j}
					grid[p] = num + ri + rj
					if grid[p] > 9 {
						grid[p] -= 9
					}
				}
			}
		}
	}

	Size *= 5
	return grid
}
