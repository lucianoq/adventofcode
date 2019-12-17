package main

import "fmt"

func main() {
	input := make(chan int, 0)
	output := make(chan int, 0)

	go func() {
		NewVM("input", input, output).Run()
		close(output)
	}()

	grid := buildGrid(input, output)

	fmt.Println(dfs(C{-12, -12}, grid, make(map[C]struct{}), 0))
}

func dfs(node C, grid map[C]int, visited map[C]struct{}, depth int) int {
	visited[node] = struct{}{}

	maxDepth := depth
	for _, dir := range []int{North, South, East, West} {
		nextPos := C{node.x + offset[dir].x, node.y + offset[dir].y}
		if grid[nextPos] != Wall {
			if _, ok := visited[nextPos]; !ok {
				d := dfs(nextPos, grid, visited, depth+1)
				if d > maxDepth {
					maxDepth = d
				}
			}
		}
	}
	return maxDepth
}
