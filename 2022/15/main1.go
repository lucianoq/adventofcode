package main

import "fmt"

func main() {
	sensors := parse()

	// find boundaries to reduce the search space
	minX, maxX := 1<<63-1, 0
	for _, s := range sensors {
		x := s.Pos.X - s.Dist
		if x < minX {
			minX = x
		}
		x = s.Pos.X + s.Dist
		if x > maxX {
			maxX = x
		}
	}

	y := 2000000

	count := 0
Point:
	for x := minX; x < maxX; x++ {
		p := P{x, y}
		for _, s := range sensors {
			if manhattan(p, s.Pos) <= s.Dist && p != s.Beacon {
				count++
				continue Point
			}
		}
	}
	fmt.Println(count)
}
