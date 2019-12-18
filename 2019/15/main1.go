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

	fmt.Println(bfs(grid))
}

func bfs(grid map[C]int) int {
	var null = C{-100000, -100000}
	var curr C
	toDo := []C{{}}
	parent := map[C]C{{}: null}

	for len(toDo) > 0 {
		curr, toDo = toDo[0], toDo[1:]

		if grid[curr] == Oxygen {
			break
		}

		for _, dir := range []int{North, South, East, West} {
			nextPos := C{curr.x + offset[dir].x, curr.y + offset[dir].y}
			if grid[nextPos] != Wall {
				if _, ok := parent[nextPos]; !ok {
					parent[nextPos] = curr
					toDo = append(toDo, nextPos)
				}
			}
		}
	}

	count := 0
	for curr = parent[curr]; curr != null; curr = parent[curr] {
		count++
	}
	return count
}
