package main

import "strings"

type Grid [Rows][Cols]Cell

func FromString(s string) Grid {
	grid := Grid{}

	lines := strings.Split(s, "\n")

	for row := 0; row < Rows; row++ {
		line := lines[row]
		for col := 0; col < Cols; col++ {
			var c byte
			if col < len(line) {
				c = line[col]
			} else {
				c = ' '
			}
			switch c {
			case '#':
				grid[row][col] = Wall
			case '.':
				grid[row][col] = Empty
			case 'A':
				grid[row][col] = Amber
			case 'B':
				grid[row][col] = Bronze
			case 'C':
				grid[row][col] = Copper
			case 'D':
				grid[row][col] = Desert
			case ' ':
				grid[row][col] = Out
			}
		}
	}
	return grid
}

func (g Grid) Goal() bool {
	for col := 3; col <= 9; col += 2 {
		for row := 2; row <= Rows-2; row++ {
			if g[row][col] != AmphipodByColumn[col] {
				return false
			}
		}
	}
	return true
}

func (g Grid) Copy() Grid {
	var newG Grid
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			newG[r][c] = g[r][c]
		}
	}
	return newG
}

func (g Grid) Hash() string {
	var s strings.Builder
	for col := 1; col <= Cols-2; col++ {
		s.WriteString(g[1][col].String())
	}
	for col := 3; col <= 9; col += 2 {
		for row := 2; row <= Rows-2; row++ {
			s.WriteString(g[row][col].String())
		}
	}
	return s.String()
}

func (g Grid) String() string {
	var s strings.Builder
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			s.WriteString(g[r][c].String())
		}
		s.WriteRune('\n')
	}
	return s.String()
}

func outsideRoom(r, c int) bool {
	return r == 1 && (c == 3 || c == 5 || c == 7 || c == 9)
}

func (g Grid) goOut(r, c int) ([]*Node, bool) {
	cell := g[r][c]
	energyToAdd := Energy[cell]

	energy := 0
	// bubble up until the hallway is reached
	for row := r - 1; row >= 1; row-- {
		// if the path is free
		if g[row][c] != Empty {
			return nil, false
		}
		energy += energyToAdd
	}

	nodes := []*Node{}

	// move sideways while there are free cell, and add all of them
	// to the list of possible moves
	// try both directions: left=-1, right=+1
	for _, dir := range []int{1, -1} {
		// sum up from parent energy without changing it
		energy := energy

		// walk step by step
		// adding energy
		for s := c + dir; g[1][s] == Empty; s += dir {
			energy += energyToAdd

			// Amphipods will never stop on the space
			// immediately outside any room.
			if !outsideRoom(1, s) {
				newG := g.Copy()
				newG[1][s], newG[r][c] = newG[r][c], newG[1][s]
				nodes = append(nodes, &Node{
					Grid:   newG,
					Energy: energy,
				})
			}
		}
	}
	return nodes, true
}

func (g Grid) goHome(c int) (*Node, bool) {
	cell := g[1][c]
	energyToAdd := Energy[cell]
	goalColumn := ColumnByAmphipod[cell]

	goalRow, room := g.roomAtHome(goalColumn)
	if !room {
		return nil, false
	}

	energy := 0

	// move sideways until the right column is reached
	step := 1 // to right
	arrived := func(curr int) bool {
		return curr <= goalColumn
	}
	// if destination column is on the left of current column
	// we need to go backward (-1) while current >= destination
	if c > goalColumn {
		step = -1 // to left
		arrived = func(curr int) bool {
			return curr >= goalColumn
		}
	}
	for currCol := c + step; arrived(currCol); currCol += step {
		// if the path is free
		if g[1][currCol] != Empty {
			return nil, false
		}
		energy += energyToAdd
	}

	energy += energyToAdd * (goalRow - 1)

	g[1][c], g[goalRow][goalColumn] = Empty, cell

	return &Node{
		Grid:   g,
		Energy: energy,
	}, true
}

// return true if that room is ready to host their amphipods
// (no foreign amphipods in), false otherwise.
// If there is room, return the row index of the highest free cell.
func (g Grid) roomAtHome(c int) (int, bool) {

	// check if there are only the right amphipods and empty cells
	for r := 2; r <= Rows-2; r++ {
		if g[r][c] != AmphipodByColumn[c] && g[r][c] != Empty {
			return 0, false
		}
	}

	// from bottom, return the index of the first Empty
	for r := Rows - 2; r >= 2; r-- {
		if g[r][c] == Empty {
			return r, true
		}
	}

	// no room, because all amphipods are in place
	return 0, false
}
