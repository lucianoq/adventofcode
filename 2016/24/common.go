package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	wall  = -1
	empty = -2
)

type Grid [][]int

type Status struct {
	X, Y int

	// bitmap of Points of Interest found
	PoICollected uint64

	Steps int
}

func (s Status) Norm() Status {
	s.Steps = 0
	return s
}

func main() {
	g, _, _, startX, startY, numPoI := parse()
	//g.Print(rows, cols)

	result := BFS(g, startX, startY, numPoI)

	fmt.Println(result)
}

func parse() (g Grid, rows, cols, startX, startY, numPoI int) {
	var i, j int
	g = make(Grid, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for i = 0; scanner.Scan(); i++ {
		line := scanner.Text()

		g = append(g, make([]int, len(line)))

		for j = 0; j < len(line); j++ {
			c := []byte(line)[j]
			switch {
			case c == '#':
				g[i][j] = wall
			case c == '.':
				g[i][j] = empty
			case c >= '0' && c <= '9':
				x, _ := strconv.Atoi(string(c))
				if x > numPoI {
					numPoI = x
				}
				if x == 0 {
					startX, startY = i, j
				}
				g[i][j] = x
			}
		}
	}
	return g, i, j, startX, startY, numPoI
}

func (g Grid) Print(rows, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch g[i][j] {
			case wall:
				fmt.Print("#")
			case empty:
				fmt.Print(" ")
			default:
				fmt.Print(g[i][j])
			}
		}
		fmt.Println()
	}
}

func (g Grid) Adjacent(s Status) []Status {
	ss := []Status{}

	for _, offset := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		offX, offY := offset[0], offset[1]
		cell := g[s.X+offX][s.Y+offY]

		switch {
		case cell == wall:
		case cell == empty, cell == 0:
			ss = append(ss, Status{s.X + offX, s.Y + offY, s.PoICollected, s.Steps + 1})
		default:
			// set the bit at cell-th position to 1
			newPoICollected := s.PoICollected | (1 << uint(cell))

			ss = append(ss, Status{s.X + offX, s.Y + offY, newPoICollected, s.Steps + 1})
		}
	}
	return ss
}

func BFS(grid Grid, startX, startY, numPoI int) int {
	start := Status{startX, startY, 0, 0}

	visited := map[Status]bool{start: true}
	toDoList := []Status{start}

	var curr = start
	for len(toDoList) > 0 {
		curr, toDoList = toDoList[0], toDoList[1:]

		if curr.Complete(numPoI, startX, startY) {
			return curr.Steps
		}

		for _, s := range grid.Adjacent(curr) {
			if !visited[s.Norm()] {
				visited[s.Norm()] = true
				toDoList = append(toDoList, s)
			}
		}
	}

	return -1
}
