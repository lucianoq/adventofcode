package main

import (
	"fmt"
	"strings"
)

func main() {
	m := parse()

	equivalent := findEquivalent(m, 1000000000)

	for i := 0; i < equivalent; i++ {
		m.Cycle()
	}

	fmt.Println(m.Load())
}

func findEquivalent(m Map, goal int) int {
	m = m.Copy()
	var start, end int

	visited := map[string]int{}
	for i := 1; i <= goal; i++ {
		m.Cycle()

		str := m.String()

		if round, ok := visited[str]; ok {
			start, end = round, i
			break
		}

		visited[str] = i
	}

	return (goal-start)%(end-start) + start
}

func (m Map) String() string {
	s := strings.Builder{}
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[r]); c++ {
			s.WriteByte(m[r][c])
		}
	}
	return s.String()
}

func (m Map) Cycle() {
	m.North()
	m.West()
	m.South()
	m.East()
}

func (m Map) Copy() Map {
	cp := make(Map, len(m))
	for i := range m {
		cp[i] = make([]byte, len(m[i]))
		copy(cp[i], m[i])
	}
	return cp
}

func (m Map) South() {
	for change := true; change; {
		change = false
		for r := len(m) - 2; r >= 0; r-- {
			for c := 0; c < len(m[r]); c++ {
				if m[r][c] == 'O' {
					if m[r+1][c] == '.' {
						m[r][c], m[r+1][c] = '.', 'O'
						change = true
					}
				}
			}
		}
	}
}

func (m Map) West() {
	for change := true; change; {
		change = false
		for r := 0; r < len(m); r++ {
			for c := 1; c < len(m[r]); c++ {
				if m[r][c] == 'O' {
					if m[r][c-1] == '.' {
						m[r][c], m[r][c-1] = '.', 'O'
						change = true
					}
				}
			}
		}
	}
}

func (m Map) East() {
	for change := true; change; {
		change = false
		for r := 0; r < len(m); r++ {
			for c := len(m[r]) - 2; c >= 0; c-- {
				if m[r][c] == 'O' {
					if m[r][c+1] == '.' {
						m[r][c], m[r][c+1] = '.', 'O'
						change = true
					}
				}
			}
		}
	}
}
