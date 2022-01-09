package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ r, c int }

type Type uint8

type Trench struct {
	grid       map[P]Type
	rows, cols int
}

const (
	Empty Type = iota
	East
	South
)

func (t *Trench) Move() bool {
	moved := false
	for _, direction := range []Type{East, South} {
		newTrenchGrid := map[P]Type{}
		for r := 0; r < t.rows; r++ {
			for c := 0; c < t.cols; c++ {
				p := P{r, c}

				switch t.grid[p] {
				case East:
					if direction == East {
						next := P{r, (c + 1 + t.cols) % t.cols}
						if t.grid[next] == Empty {
							moved = true
							newTrenchGrid[p] = Empty
							newTrenchGrid[next] = East
						} else {
							newTrenchGrid[p] = East
						}
					} else {
						newTrenchGrid[p] = East
					}

				case South:
					if direction == South {
						next := P{(r + 1 + t.rows) % t.rows, c}
						if t.grid[next] == Empty {
							moved = true
							newTrenchGrid[p] = Empty
							newTrenchGrid[next] = South
						} else {
							newTrenchGrid[p] = South
						}
					} else {
						newTrenchGrid[p] = South
					}

				case Empty:
					if _, ok := newTrenchGrid[p]; !ok {
						newTrenchGrid[p] = Empty
					}
				}
			}
		}
		t.grid = newTrenchGrid
	}
	return moved
}

func main() {
	trench := parse()
	for step := 1; ; step++ {
		moved := trench.Move()
		if !moved {
			fmt.Println(step)
			return
		}
	}
}

func parse() Trench {
	grid := map[P]Type{}
	scanner := bufio.NewScanner(os.Stdin)

	row, col := 0, 0
	for row = 0; scanner.Scan(); row++ {
		line := scanner.Text()

		if line == "" {
			continue
		}

		var c rune
		for col, c = range line {
			switch c {
			case 'v':
				grid[P{row, col}] = South
			case '>':
				grid[P{row, col}] = East
			case '.':
				grid[P{row, col}] = Empty
			default:
				panic("wrong char")
			}
		}
	}

	return Trench{
		grid: grid,
		rows: row,
		cols: col + 1,
	}
}
