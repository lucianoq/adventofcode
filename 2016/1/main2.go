package main

import (
	"fmt"
	"log"
	"strconv"
)

type Point struct{ X, Y int }

func main() {
	items := parse()
	visited := map[Point]bool{}

	x, y := 0, 0
	visited[Point{x, y}] = true

	face := 0
	for _, i := range items {
		switch i[0] {
		case 'R':
			face = (face + 1 + 4) % 4
		case 'L':
			face = (face - 1 + 4) % 4
		}
		steps, err := strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
		}

		for j := 0; j < steps; j++ {
			switch dirs[face] {
			case 'N':
				y++
			case 'E':
				x++
			case 'S':
				y--
			case 'W':
				x--
			}
			if visited[Point{x, y}] {
				fmt.Println(x + y)
				return
			}

			visited[Point{x, y}] = true
		}
	}
}
