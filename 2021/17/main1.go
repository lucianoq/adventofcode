package main

import "fmt"

func main() {
	_, _, yMin, yMax := parse()

	// going backward, the first found would be the highest
	for yv := 1000; yv > yMin; yv-- {
		if ok, yReached := findMaxYReached(yv, yMin, yMax); ok {
			fmt.Println(yReached)
			return
		}
	}
}

func findMaxYReached(yv, yMin, yMax int) (bool, int) {
	maxYReached := 0
	y := 0
	for step := 0; step < 10000; step++ {
		y += yv

		if y > maxYReached {
			maxYReached = y
		}

		if y >= yMin && y <= yMax {
			return true, maxYReached
		}

		yv--
	}
	return false, 0
}
