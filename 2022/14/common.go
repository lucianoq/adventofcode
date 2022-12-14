package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type P struct{ X, Y int }

var (
	Map   = map[P]struct{}{}
	Floor int
)

func parsePoint(s string) P {
	pp := strings.Split(s, ",")
	pC, _ := strconv.Atoi(pp[0])
	pR, _ := strconv.Atoi(pp[1])
	return P{pR, pC}
}

func parse() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		segments := strings.Split(scanner.Text(), " -> ")
		start := parsePoint(segments[0])
		for i := 1; i < len(segments); i++ {
			end := parsePoint(segments[i])
			for x := min(start.X, end.X); x <= max(start.X, end.X); x++ {
				for y := min(start.Y, end.Y); y <= max(start.Y, end.Y); y++ {
					Map[P{x, y}] = struct{}{}
				}
			}

			maxX := max(start.X, end.X)
			if maxX > Floor {
				Floor = maxX
			}

			start = end
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func countRest() int {
	count := 0
	for {
		settled := pourSand()
		if !settled {
			return count
		}
		count++
	}
}

func pourSand() bool {
	sand := P{0, 500}

	// If start tile is not empty,
	// we filled the map
	if !empty(sand) {
		return false
	}

Fall:
	for sand.X < Floor {
		for _, y := range [3]int{0, -1, 1} {
			next := P{sand.X + 1, sand.Y + y}
			if empty(next) {
				sand = next
				continue Fall
			}
		}

		// settle there
		Map[sand] = struct{}{}
		return true
	}

	// into the abyss
	return false
}
