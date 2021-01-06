package main

import (
	"bufio"
	"fmt"
	"os"
)

type char uint8

type P struct{ X, Y uint16 }

func (p P) Neighbours() []P {
	return []P{
		{p.X - 1, p.Y}, // W
		{p.X, p.Y - 1}, // S
		{p.X + 1, p.Y}, // E
		{p.X, p.Y + 1}, // N
	}
}

const (
	Wall = '#'
	Free = '.'
)

type Map struct {
	XMax    uint16
	YMax    uint16
	Grid    map[P]char
	Portals map[P]P
	Start   P
	End     P
}

func parse() *Map {
	grid := map[P]char{}

	scanner := bufio.NewScanner(os.Stdin)
	i, j := uint16(0), uint16(0)

	for ; scanner.Scan(); i++ {
		line := scanner.Text()
		for j = 0; j < uint16(len(line)); j++ {
			grid[P{i, j}] = char(line[j])
		}
	}

	portals := map[P]P{}
	cache := map[string]P{}

	xMax, yMax := i, j

	for i := uint16(0); i < xMax; i++ {
		for j := uint16(0); j < yMax; j++ {
			c := grid[P{i, j}]

			if !isLetter(c) {
				continue
			}

			// try south
			if c2 := grid[P{i + 1, j}]; isLetter(c2) {
				portalName := string(c) + string(c2)

				portalPoint := P{i + 2, j}
				if grid[portalPoint] != '.' {
					portalPoint = P{i - 1, j}
				}

				if p, ok := cache[portalName]; ok {
					portals[p] = portalPoint
					portals[portalPoint] = p
				} else {
					cache[portalName] = portalPoint
				}
			}

			// try east
			if c2 := grid[P{i, j + 1}]; isLetter(c2) {
				portalName := string(c) + string(c2)

				portalPoint := P{i, j + 2}
				if grid[portalPoint] != '.' {
					portalPoint = P{i, j - 1}
				}

				if p, ok := cache[portalName]; ok {
					portals[p] = portalPoint
					portals[portalPoint] = p
				} else {
					cache[portalName] = portalPoint
				}
			}
		}
	}

	return &Map{
		XMax:    xMax,
		YMax:    yMax,
		Grid:    grid,
		Portals: portals,
		Start:   cache["AA"],
		End:     cache["ZZ"],
	}
}

func isLetter(c char) bool {
	return c >= 'A' && c <= 'Z'
}

func main() {
	fmt.Println(parse().BFS())
}

func (m *Map) BFS() int {
	discovered := map[P]struct{}{}
	toDo := []P{}

	discovered[m.Start] = struct{}{}
	toDo = append(toDo, m.Start)

	depth := 0

	for len(toDo) > 0 {
		for levelSize := len(toDo); levelSize > 0; levelSize-- {
			var curr P
			curr, toDo = toDo[0], toDo[1:]

			if curr == m.End {
				return depth
			}

			for _, n := range curr.Neighbours() {
				dest := m.Grid[n]

				switch {
				case dest == Wall:
					continue

				case dest == Free:

					if _, found := discovered[n]; !found {
						discovered[n] = struct{}{}
						toDo = append(toDo, n)
					}

				case isLetter(dest):
					stargate := m.Portals[curr]
					if _, found := discovered[stargate]; !found {
						discovered[stargate] = struct{}{}
						toDo = append(toDo, stargate)
					}
				}
			}
		}
		depth++
	}

	return -1
}
