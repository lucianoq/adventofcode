package main

import (
	"bufio"
	"errors"
	"os"
)

var (
	Width, Height = 50, 50
	NotFound      = errors.New("not found")
)

func parseInput() (map[P]byte, []byte, P) {
	var robot P
	m := map[P]byte{}

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if line == "" {
			break
		}

		for j, c := range line {
			if c == '@' {
				robot = P{i, j}
				m[P{i, j}] = '.'
			} else {
				m[P{i, j}] = byte(c)
			}
		}
	}

	moves := []byte{}
	for scanner.Scan() {
		moves = append(moves, []byte(scanner.Text())...)
	}
	return m, moves, robot
}

type P struct{ x, y int }

func getDelta(move byte) P {
	switch move {
	case '<':
		return P{0, -1}
	case '>':
		return P{0, 1}
	case '^':
		return P{-1, 0}
	case 'v':
		return P{1, 0}
	}
	panic("invalid move")
}

func findNextEmpty(m map[P]byte, p, delta P) (P, error) {
	for {
		p = P{p.x + delta.x, p.y + delta.y}
		switch m[p] {
		case '.':
			return p, nil
		case '#':
			return P{}, NotFound
		}
	}
}

func gps(m map[P]byte, char byte) int {
	sum := 0
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if m[P{i, j}] == char {
				sum += 100*i + j
			}
		}
	}
	return sum
}
