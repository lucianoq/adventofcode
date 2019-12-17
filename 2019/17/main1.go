package main

import "fmt"

type C struct{ x, y int }

func main() {
	input := make(chan int, 0)
	output := make(chan int, 0)

	go func() {
		NewVM("input", input, output).Run()
		close(output)
	}()

	pos := C{0, 0}
	grid := make(map[C]int)
	var maxX, maxY int
	for x := range output {
		// fmt.Print(string(x))

		grid[pos] = x

		if x == '\n' {
			pos.y++
			pos.x = 0
		} else {
			pos.x++
		}

		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}

	sum := 0
	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			if intersection(C{i, j}, grid) {
				sum += i * j
			}
		}
	}

	fmt.Println(sum)
}

func intersection(c C, grid map[C]int) bool {
	return grid[C{c.x, c.y}] == '#' &&
		grid[C{c.x + 1, c.y}] == '#' &&
		grid[C{c.x - 1, c.y}] == '#' &&
		grid[C{c.x, c.y + 1}] == '#' &&
		grid[C{c.x, c.y - 1}] == '#'
}
