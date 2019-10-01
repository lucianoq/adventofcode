package main

import (
	"bufio"
	"io"
)

func parse(reader io.Reader) (map[C]bool, map[C]Unit) {
	world := make(map[C]bool)
	units := make(map[C]Unit)

	scanner := bufio.NewScanner(reader)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			world[C{i, j}] = line[j] != '#'

			switch line[j] {
			case 'E':
				units[C{i, j}] = Unit{
					Type: Elf,
					AP:   3,
					HP:   200,
					Pos:  C{i, j},
				}
			case 'G':
				units[C{i, j}] = Unit{
					Type: Goblin,
					AP:   3,
					HP:   200,
					Pos:  C{i, j},
				}
			}
		}
	}

	return world, units
}
