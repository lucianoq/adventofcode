package main

import (
	"bufio"
	"os"
	"strconv"
)

type P struct{ x, y int }

func adjacent(grid [][]int, p P) []P {
	rows, cols := len(grid), len(grid[0])

	adj := []P{}
	if p.x > 0 {
		adj = append(adj, P{p.x - 1, p.y})
	}
	if p.y > 0 {
		adj = append(adj, P{p.x, p.y - 1})
	}
	if p.x < rows-1 {
		adj = append(adj, P{p.x + 1, p.y})
	}
	if p.y < cols-1 {
		adj = append(adj, P{p.x, p.y + 1})
	}
	return adj
}

func parse() [][]int {
	grid := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		grid = append(grid, []int{})
		for _, c := range line {
			num, _ := strconv.Atoi(string(c))
			grid[r] = append(grid[r], num)
		}
	}
	return grid
}
