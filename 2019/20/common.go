package main

import (
	"bufio"
	"os"
)

const (
	Wall = '#'
	Free = '.'
)

type char uint8

func (c char) Letter() bool {
	return c >= 'A' && c <= 'Z'
}

type P struct{ X, Y uint16 }

func (p P) Neighbours() []P {
	return []P{
		{p.X, p.Y + 1}, // N
		{p.X + 1, p.Y}, // E
		{p.X, p.Y - 1}, // S
		{p.X - 1, p.Y}, // W
	}
}

type Map struct {
	XMax       uint16
	YMax       uint16
	Grid       map[P]char
	AA         P
	ZZ         P
	Teleport   map[P]P
	PortalName map[P]string
	IsOuter    map[P]bool
}

func parse() *Map {
	grid := map[P]char{}
	xMax, yMax := uint16(0), uint16(0)

	scanner := bufio.NewScanner(os.Stdin)

	i := uint16(0)
	for ; scanner.Scan(); i++ {
		line := scanner.Text()

		if uint16(len(line)) > yMax {
			yMax = uint16(len(line))
		}

		for j := uint16(0); j < uint16(len(line)); j++ {
			grid[P{i, j}] = char(line[j])
		}
	}
	xMax = i

	var aa, zz P
	isOuter := map[P]bool{}
	portalName := map[P]string{}
	teleport := map[P]P{}

	cache := map[string]P{}

	for i := uint16(0); i < xMax; i++ {
		for j := uint16(0); j < yMax; j++ {
			c := grid[P{i, j}]

			if !c.Letter() {
				continue
			}

			pName, pPoint, ok := extractPortal(grid, P{i, j})

			if !ok {
				// not a first letter of a portal
				continue
			}

			portalName[pPoint] = pName

			if pName == "AA" {
				aa = pPoint
				isOuter[pPoint] = true
				continue
			}

			if pName == "ZZ" {
				zz = pPoint
				isOuter[pPoint] = true
				continue
			}

			if target, ok := cache[pName]; ok {
				teleport[pPoint] = target
				teleport[target] = pPoint
			} else {
				cache[pName] = pPoint
			}

			switch {
			case j == 0 || i == 0, // top and left outer
				i == xMax-2 || j == yMax-2: // bottom and right outer

				isOuter[pPoint] = true
			default:
				// inner
				isOuter[pPoint] = false
			}
		}
	}

	return &Map{
		XMax:       xMax,
		YMax:       yMax,
		Grid:       grid,
		AA:         aa,
		ZZ:         zz,
		Teleport:   teleport,
		PortalName: portalName,
		IsOuter:    isOuter,
	}
}

func extractPortal(grid map[P]char, p P) (string, P, bool) {
	c1 := grid[p]

	// try vertical
	if c2 := grid[P{p.X + 1, p.Y}]; c2.Letter() {
		portalName := string(c1) + string(c2)

		portalPoint := P{p.X + 2, p.Y}
		if grid[portalPoint] == '.' {
			return portalName, portalPoint, true
		}

		portalPoint = P{p.X - 1, p.Y}
		if grid[portalPoint] == '.' {
			return portalName, portalPoint, true
		}
	}

	// try horizontal
	if c2 := grid[P{p.X, p.Y + 1}]; c2.Letter() {
		portalName := string(c1) + string(c2)

		portalPoint := P{p.X, p.Y + 2}
		if grid[portalPoint] == '.' {
			return portalName, portalPoint, true
		}

		portalPoint = P{p.X, p.Y - 1}
		if grid[portalPoint] == '.' {
			return portalName, portalPoint, true
		}
	}

	return "", P{}, false
}
