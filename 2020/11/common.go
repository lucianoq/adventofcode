package main

import (
	"bufio"
	"os"
)

const (
	Floor uint8 = iota
	Empty
	Occupied
)

type P struct{ R, C int }

func (p P) Add(r, c int) P {
	return P{p.R + r, p.C + c}
}

type Grid struct {
	grid map[P]uint8
	maxR int
	maxC int
}

func (g Grid) Step(neighbors func(P) []P, limitOccupied int) (Grid, bool) {
	newGrid := NewGrid(g.maxR, g.maxC)
	changed := false

	for r := 0; r < g.maxR; r++ {

	NextSeat:
		for c := 0; c < g.maxC; c++ {

			// by default it doesn't change
			newGrid.grid[P{r, c}] = g.grid[P{r, c}]

			switch g.grid[P{r, c}] {

			case Empty:

				for _, p := range neighbors(P{r, c}) {
					if g.grid[p] == Occupied {
						// if we find just one Occupied, no need to continue
						// the loop
						continue NextSeat
					}
				}
				// no Occupied found, we can set it
				newGrid.grid[P{r, c}] = Occupied
				changed = true

			case Occupied:
				countOcc := 0
				for _, p := range neighbors(P{r, c}) {
					if g.grid[p] == Occupied {
						countOcc++
					}

					if countOcc >= limitOccupied {
						newGrid.grid[P{r, c}] = Empty
						changed = true
						continue NextSeat
					}
				}
			}
		}
	}
	return newGrid, changed
}

// Occupied returns the number of occupied seats in a Grid
func (g Grid) Occupied() int {
	count := 0
	for i := 0; i < g.maxR; i++ {
		for j := 0; j < g.maxC; j++ {
			if g.grid[P{i, j}] == Occupied {
				count++
			}
		}
	}
	return count
}

// OutOfBounds returns true if that point is outside the grid boundaries.
func (g Grid) OutOfBounds(p P) bool {
	if p.R < 0 || p.R >= g.maxR {
		return true
	}

	if p.C < 0 || p.C >= g.maxC {
		return true
	}

	return false
}

func NewGrid(r, c int) Grid {
	return Grid{
		grid: map[P]uint8{},
		maxR: r,
		maxC: c,
	}
}

func parse() Grid {
	g := NewGrid(0, 0)

	r := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		c := 0

		for _, char := range line {
			switch char {
			case 'L':
				g.grid[P{r, c}] = Empty
			case '#':
				g.grid[P{r, c}] = Occupied
			case '.':
				g.grid[P{r, c}] = Floor
			}
			c++

			if c > g.maxC {
				g.maxC = c
			}
		}

		r++
	}
	g.maxR = r

	return g
}
