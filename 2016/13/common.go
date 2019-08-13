package main

import "math/bits"

const input = 1358

type Point struct{ X, Y, Steps int }

func adjacents(p Point) []Point {
	points := []Point{
		{p.X + 1, p.Y, p.Steps + 1},
		{p.X, p.Y + 1, p.Steps + 1},
		{p.X - 1, p.Y, p.Steps + 1},
		{p.X, p.Y - 1, p.Steps + 1},
	}

	var filteredPoints []Point
	for _, p := range points {
		if p.X >= 0 && p.Y >= 0 {
			if open(uint(p.X), uint(p.Y)) {
				filteredPoints = append(filteredPoints, p)
			}
		}
	}
	return filteredPoints
}

//Find x*x + 3*x + 2*x*y + y + y*y.
//Add the office designer's favorite number (your puzzle input).
//Find the binary representation of that sum; count the number of bits that are 1.
//If the number of bits that are 1 is even, it's an open space.
//If the number of bits that are 1 is odd, it's a wall.
func open(x, y uint) bool {
	return bits.OnesCount(x*x+3*x+2*x*y+y+y*y+input)%2 == 0
}

func toInt(p Point) int {
	return p.X*1000000 + p.Y
}
