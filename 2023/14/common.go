package main

import (
	"bufio"
	"os"
)

type Map [][]byte

func parse() Map {
	scanner := bufio.NewScanner(os.Stdin)
	var m [][]byte
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		var row []byte
		for c := 0; c < len(line); c++ {
			row = append(row, line[c])
		}
		m = append(m, row)
	}
	return m
}

func (m Map) North() {
	for change := true; change; {
		change = false
		for r := 1; r < len(m); r++ {
			for c := 0; c < len(m[r]); c++ {
				if m[r][c] == 'O' {
					if m[r-1][c] == '.' {
						m[r][c], m[r-1][c] = '.', 'O'
						change = true
					}
				}
			}
		}
	}
}

func (m Map) Load() int {
	sum := 0
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[r]); c++ {
			if m[r][c] == 'O' {
				sum += len(m) - r
			}
		}
	}
	return sum
}
