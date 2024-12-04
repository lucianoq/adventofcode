package main

import (
	"bufio"
	"os"
)

const Size = 140

type P struct{ x, y int }

func (p P) Add(d P, times int) P { return P{p.x + times*d.x, p.y + times*d.y} }

func input() map[P]byte {
	scanner := bufio.NewScanner(os.Stdin)

	out := map[P]byte{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			out[P{i, j}] = line[j]
		}
	}
	return out
}
