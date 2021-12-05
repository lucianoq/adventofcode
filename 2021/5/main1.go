package main

import "fmt"

func main() {
	segments := parse()

	visited := map[Point]int{}

	for _, s := range segments {

		switch {

		// vertical
		case s.P1.X == s.P2.X:
			minY, maxY := minMax(s.P1.Y, s.P2.Y)
			for y := minY; y <= maxY; y++ {
				visited[Point{s.P1.X, y}]++
			}

		// horizontal
		case s.P1.Y == s.P2.Y:
			minX, maxY := minMax(s.P1.X, s.P2.X)
			for x := minX; x <= maxY; x++ {
				visited[Point{x, s.P1.Y}]++
			}
		}
	}

	count := 0
	for _, v := range visited {
		if v >= 2 {
			count++
		}
	}

	fmt.Println(count)
}
