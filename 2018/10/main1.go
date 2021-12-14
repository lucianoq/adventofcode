package main

import (
	"fmt"
	"math"
	"os"
)

func Sky(time int) {
	interpolatedStars := make([]*Star, 0)
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64
	for _, s := range stars {
		x, y := s.Current(time)
		interpolatedStars = append(interpolatedStars, &Star{x, y, 0, 0})
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	// Normalize
	for _, s := range interpolatedStars {
		s.X -= minX
		s.Y -= minY
	}
	maxX -= minX
	maxY -= minY
	maxX++
	maxY++

	if maxX < 64 {
		// Initialize
		sky := make([][]bool, maxX+1)
		for i := 0; i < maxX+1; i++ {
			sky[i] = make([]bool, maxY+1)
		}

		// Populate
		for _, s := range interpolatedStars {
			sky[s.X][s.Y] = true
		}

		fmt.Println(toString(sky, maxX, maxY, 6, 2))
		os.Exit(0)
	}
}
