package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Star struct {
	X    int
	Y    int
	VelX int
	VelY int
}

func (s *Star) Current(time int) (x, y int) {
	return s.X + time*s.VelX, s.Y + time*s.VelY
}

var stars []*Star

func main() {
	stars = make([]*Star, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			line = strings.Replace(line, " ", "", -1)
			var x, y, velX, velY int
			n, err := fmt.Sscanf(line, "position=<%d,%d>velocity=<%d,%d>", &x, &y, &velX, &velY)
			if err != nil || n != 4 {
				log.Fatal(err)
			}
			stars = append(stars, &Star{x, y, velX, velY})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; ; i++ {
		Sky(i)
	}

}

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
		fmt.Println(time)
		os.Exit(0)
	}
}
