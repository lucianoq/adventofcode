package main

import (
	"fmt"
	"strconv"
)

func main() {
	curr := P{0, 0}

	var edges int
	var vertices = []P{curr}

	for _, op := range parse() {
		dir, steps := extract(op.Color)

		curr = P{
			curr.R + Delta[dir].R*steps,
			curr.C + Delta[dir].C*steps,
		}
		vertices = append(vertices, curr)

		edges += steps
	}

	fmt.Println(area(vertices) + edges/2 + 1)
}

func extract(s string) (string, int) {
	var dir string
	switch s[5] {
	case '0':
		dir = "R"
	case '1':
		dir = "D"
	case '2':
		dir = "L"
	case '3':
		dir = "U"
	}
	steps, _ := strconv.ParseInt(s[:5], 16, 64)
	return dir, int(steps)
}

// Shoelace formula
func area(vertices []P) int {
	var a int
	for i := 0; i < len(vertices); i++ {
		j := (i + 1) % len(vertices)
		a += vertices[i].R * vertices[j].C
		a -= vertices[i].C * vertices[j].R
	}
	return abs(a) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
