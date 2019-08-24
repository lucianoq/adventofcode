package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	startX, startY int
	numPoI         int
)

const (
	wall  = -1
	empty = -2
)

type Status struct {
	X, Y         int
	PoICollected uint64
	Steps        int
}

func (s Status) Norm() Status {
	s.Steps = 0
	return s
}

type Grid [][]int

func main() {
	g, rows, cols := parse()
	g.Print(rows, cols)

	fmt.Println(BFS(g))
}

func parse() (Grid, int, int) {
	var i, j int
	g := make(Grid, 0)

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
	return g, i, j
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

	for offsetX := range []int{-1, 0, +1} {
		for offsetY := range []int{-1, 0, +1} {
			if (offsetX+offsetY)*(offsetX+offsetY) == 1 {
				cell := g[s.X+offsetX][s.Y+offsetY]

				switch {
				case cell == wall:
				case cell == empty, cell == 0:
					ss = append(ss, Status{s.X + offsetX, s.Y + offsetY, s.PoICollected, s.Steps + 1})
				default:
					s.PoICollected = s.PoICollected|(1<<)
					ss = append(ss, Status{s.X + offsetX, s.Y + offsetY, s.PoICollected + 1, s.Steps + 1})
				}
			}
		}
	}
	fmt.Printf("from %d,%d I can go to %+v\n", s.X, s.Y, ss)
	return ss
}




func BFS(grid Grid) int {
	//successMap := bits.

	start := Status{startX, startY, 0, 0}

	visited := map[Status]bool{start: true}
	toDoList := []Status{start}

	var curr = start
	for len(toDoList) > 0 {
		curr, toDoList = toDoList[0], toDoList[1:]

		fmt.Printf("Visiting (%d,%d) with PoI collected=%d and steps=%d\n", curr.X, curr.Y, curr.PoICollected, curr.Steps)

		if curr.PoICollected == (numPoI<<8)-1 {
			fmt.Println("Caught'em'all!")
			return curr.Steps
		}

		for _, s := range grid.Adjacent(curr) {
			if !visited[s.Norm()] {
				visited[s.Norm()] = true
				toDoList = append(toDoList, s)
			}
		}
	}

	for k := range visited {
		fmt.Println(k)
	}

	return -1
}
