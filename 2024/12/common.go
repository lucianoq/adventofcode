package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ x, y int }

func parseInput() (map[P]uint8, int) {
	scanner := bufio.NewScanner(os.Stdin)
	m := map[P]uint8{}
	var i int
	for i = 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			m[P{i, j}] = uint8(c)
		}
	}
	return m, i
}

func main() {
	m, size := parseInput()
	price := 0
	visited := map[P]struct{}{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if _, ok := visited[P{i, j}]; ok {
				continue
			}
			area, fence := explore(m, P{i, j}, visited)
			price += area * fence
		}
	}
	fmt.Println(price)
}
