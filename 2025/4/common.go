package main

import (
	"bufio"
	"os"
)

const Size = 136

type P struct{ x, y int }

func parseMap() map[P]struct{} {
	m := map[P]struct{}{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		for j, c := range scanner.Text() {
			if c == '@' {
				m[P{i, j}] = struct{}{}
			}
		}
	}
	return m
}

func adjacentRolls(m map[P]struct{}, x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx, ny := x+dx, y+dy
			if nx < 0 || nx >= Size ||
				ny < 0 || ny >= Size {
				continue
			}

			if _, ok := m[P{nx, ny}]; ok {
				count++
			}
		}
	}
	return count
}
