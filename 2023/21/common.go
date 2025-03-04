package main

import (
	"bufio"
	"os"
	"strings"
)

type P struct{ x, y int }

type Map struct {
	m     []string
	Start P
	Size  int
}

var Delta = []P{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (m Map) Get(p P) byte {
	var mod = P{
		x: ((p.x % m.Size) + m.Size) % m.Size,
		y: ((p.y % m.Size) + m.Size) % m.Size,
	}
	return m.m[mod.x][mod.y]
}

func parse() Map {
	scanner := bufio.NewScanner(os.Stdin)
	var m Map
	var start P
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		if idx := strings.IndexByte(line, 'S'); idx > -1 {
			start.x = r
			start.y = idx
			line = line[:idx] + "." + line[idx+1:]
		}
		m.m = append(m.m, line)
	}
	m.Size = len(m.m)
	m.Start = start
	return m
}

func Reach(m Map, steps int) int {
	reachable := map[int]map[P]struct{}{}
	reachable[0] = map[P]struct{}{m.Start: {}}

	for step := 1; step <= steps; step++ {
		newSteps := map[P]struct{}{}

		for p := range reachable[step-1] {
			for _, delta := range Delta {
				next := P{p.x + delta.x, p.y + delta.y}
				if m.Get(next) == '#' {
					continue
				}
				newSteps[next] = struct{}{}
			}
		}

		reachable[step] = newSteps
		delete(reachable, step-1)
	}

	return len(reachable[steps])
}
