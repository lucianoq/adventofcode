package main

import (
	"bufio"
	"os"
)

const Size = 130

type Direction uint8

const (
	N Direction = iota
	E
	S
	W
)

type P struct{ x, y int }

func (p P) Next(d Direction) P {
	switch d {
	case N:
		return P{p.x - 1, p.y}
	case E:
		return P{p.x, p.y + 1}
	case S:
		return P{p.x + 1, p.y}
	case W:
		return P{p.x, p.y - 1}
	}
	return P{0, 0}
}

func parseMap() (map[P]struct{}, P, Direction) {
	var pos P
	scanner := bufio.NewScanner(os.Stdin)
	m := map[P]struct{}{}
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

func run(m map[P]struct{}, pos P, dir Direction) map[P]struct{} {
	visited := map[P]struct{}{
		pos: {},
	}

	for {
		next := pos.Next(dir)

		if next.x < 0 || next.x >= Size || next.y < 0 || next.y >= Size {
			return visited
		}

		if _, ok := m[next]; ok {
			dir = (dir + 1) % 4
		} else {
			pos = next
			visited[pos] = struct{}{}
		}
	}
}
