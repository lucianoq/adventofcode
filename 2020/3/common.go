package main

import (
	"bufio"
	"os"
	"strings"
)

// P point type
type P struct{ R, C int }

// Map type
type Map struct {
	maxR  int
	maxC  int
	trees map[P]bool
}

func NewMap() *Map {
	m := &Map{
		trees: make(map[P]bool),
	}

	scanner := bufio.NewScanner(os.Stdin)
	r := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		m.maxC = len(line)

		for c := 0; c < m.maxC; c++ {
			if line[c] == '#' {
				m.trees[P{r, c}] = true
			}
		}
		r++
	}
	m.maxR = r
	return m
}

func (m Map) Tree(c P) bool {
	return m.trees[P{c.R % m.maxR, c.C % m.maxC}]
}

func (m Map) End(c P) bool {
	return c.R >= m.maxR
}

// descent returns the number of encountered trees given a map and x/y steps.
func descent(m *Map, r, c int) int {
	x := P{0, 0}
	trees := 0

	for !m.End(x) {
		if m.Tree(x) {
			trees++
		}

		x.R += r
		x.C += c
	}
	return trees
}
