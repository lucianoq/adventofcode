package main

import (
	"bufio"
	"os"
)

var (
	Size = 0
	Map  = make(map[P]uint8)
)

type P struct{ X, Y int }

func (p P) Add(delta P) (P, bool) {
	next := P{p.X + delta.X, p.Y + delta.Y}
	return next, next.X >= 0 && next.X < Size && next.Y >= 0 && next.Y < Size
}

var (
	Left  = P{0, -1}
	Right = P{0, 1}
	Up    = P{-1, 0}
	Down  = P{1, 0}
)

func init() {
	scanner := bufio.NewScanner(os.Stdin)
	var r int
	for r = 0; scanner.Scan(); r++ {
		line := scanner.Text()
		for c := 0; c < len(line); c++ {
			Map[P{r, c}] = line[c] - '0'
		}
	}
	Size = r
}
