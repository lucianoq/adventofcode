package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var DfsOrBfs bool

func init() {
	rand.Seed(time.Now().UnixNano())
	DfsOrBfs = rand.Intn(2) == 0
}

func main() {
	grid := parse()

	rows, cols := len(grid), len(grid[0])
	var basins []int
	for i := 0; i < rows; i++ {
	Cell:
		for j := 0; j < cols; j++ {
			for _, adj := range adjacent(grid, P{i, j}) {
				if grid[i][j] >= grid[adj.x][adj.y] {
					continue Cell
				}
			}

			if DfsOrBfs {
				basins = append(basins, findBasinSizeBFS(grid, P{i, j}))
			} else {
				discovered := map[P]struct{}{}
				basins = append(basins, findBasinSizeDFS(grid, P{i, j}, discovered))
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	fmt.Println(basins[0] * basins[1] * basins[2])
}

// BFS
func findBasinSizeBFS(grid [][]int, root P) int {
	discovered := map[P]struct{}{root: {}}
	queue := []P{root}

	var curr P
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		for _, p := range adjacent(grid, curr) {

			if grid[p.x][p.y] == 9 {
				continue
			}

			if _, ok := discovered[p]; !ok {
				discovered[p] = struct{}{}
				queue = append(queue, p)
			}
		}
	}

	return len(discovered)
}

// DFS
func findBasinSizeDFS(grid [][]int, point P, discovered map[P]struct{}) int {
	discovered[point] = struct{}{}

	sum := 0
	for _, p := range adjacent(grid, point) {

		if grid[p.x][p.y] == 9 {
			continue
		}

		if _, ok := discovered[p]; !ok {
			sum += findBasinSizeDFS(grid, p, discovered)
		}
	}
	return 1 + sum
}
