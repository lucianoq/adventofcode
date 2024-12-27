package main

import (
	"bufio"
	"os"
)

type Dir uint8

const (
	N Dir = iota
	E
	S
	W
)

func (d Dir) Clock() Dir {
	return (d + 1) % 4
}

func (d Dir) CounterClock() Dir {
	return (d - 1 + 4) % 4
}

func parseInput() map[P]int {
	m := map[P]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, ch := range line {
			m[P{i, j}] = int(ch - '0')
		}
	}
	return m
}
