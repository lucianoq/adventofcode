package main

import (
	"bufio"
	"os"
)

const Size = 130

const (
	N uint8 = iota
	E
	S
	W
)

type P struct{ x, y int }

func parseMap() (map[P]any, P, uint8) {
	var pos P
	scanner := bufio.NewScanner(os.Stdin)
	m := map[P]any{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			switch c {
			case '#':
				m[P{i, j}] = struct{}{}
			case '^':
				pos = P{i, j}
			}
		}
	}
	return m, pos, N
}
