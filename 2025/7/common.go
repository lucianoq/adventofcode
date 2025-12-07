package main

import (
	"bufio"
	"os"
)

type P struct{ x, y int }

func parseInput() (map[P]struct{}, P, int) {
	var splitters = map[P]struct{}{}
	var start P

	scanner := bufio.NewScanner(os.Stdin)
	var i int
	for i = 0; scanner.Scan(); i++ {
		line := scanner.Text()

		for j, c := range line {
			switch c {
			case 'S':
				start = P{i, j}
			case '^':
				splitters[P{i, j}] = struct{}{}
			}
		}
	}
	return splitters, start, i
}
