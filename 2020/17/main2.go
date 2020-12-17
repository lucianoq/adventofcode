package main

import (
	"bufio"
	"os"
)

type Point struct {
	X, Y, Z, W int
}

func (p Point) Neighbours() []Point {
	n := []Point{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x|y|z|w != 0 {
						n = append(n, Point{p.X + x, p.Y + y, p.Z + z, p.W + w})
					}
				}
			}
		}
	}
	return n
}

func parse() Dimension {
	dim := Dimension{}
	scanner := bufio.NewScanner(os.Stdin)
	x := 0
	for scanner.Scan() {
		line := scanner.Text()

		for y, c := range line {
			if c == '#' {
				dim[Point{x, y, 0, 0}] = struct{}{}
			}
		}
		x++
	}
	return dim
}
