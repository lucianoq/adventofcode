package main

import (
	"bufio"
	"fmt"
	"os"
)

const Size = 60

type P struct{ x, y int }

func (p P) Neighbors() []P {
	return []P{
		{p.x - 1, p.y}, //N
		{p.x + 1, p.y}, //S
		{p.x, p.y - 1}, //W
		{p.x, p.y + 1}, //E
	}
}

func inMap(p P) bool {
	return p.x >= 0 && p.x < Size && p.y >= 0 && p.y < Size
}

func main() {
	h := parseMap()

	sum := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if h[P{i, j}] == 0 {
				sum += score(h, P{i, j})
			}
		}
	}

	fmt.Println(sum)
}

func parseMap() map[P]uint8 {
	scanner := bufio.NewScanner(os.Stdin)

	h := map[P]uint8{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			h[P{i, j}] = line[j] - '0'
		}
	}
	return h
}
