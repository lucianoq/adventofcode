package main

import (
	"bufio"
	"os"
)

var (
	Width  int
	Height int
	Map    = make(map[P]int8)
	Start  P
	End    P
)

type P struct{ X, Y int }

func (p P) Add(delta P) (P, bool) {
	next := P{p.X + delta.X, p.Y + delta.Y}
	return next, next.X >= 0 && next.X < Height && next.Y >= 0 && next.Y < Width
}

func init() {
	scanner := bufio.NewScanner(os.Stdin)
	var r int
	for r = 0; scanner.Scan(); r++ {
		line := scanner.Text()
		for c := 0; c < len(line); c++ {
			switch line[c] {
			case 'S':
				Start = P{r, c}
				Map[P{r, c}] = 'a'
			case 'E':
				End = P{r, c}
				Map[P{r, c}] = 'z'
			default:
				Map[P{r, c}] = int8(line[c])
			}
		}
		Width = len(line)
	}
	Height = r
}
