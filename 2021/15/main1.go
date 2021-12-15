package main

import (
	"bufio"
	"os"
	"strconv"
)

func parse() map[C]int {
	scanner := bufio.NewScanner(os.Stdin)
	grid := map[C]int{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j := 0; j < Size; j++ {
			num, _ := strconv.Atoi(line[j : j+1])
			grid[C{i, j}] = num
		}
	}
	return grid
}
