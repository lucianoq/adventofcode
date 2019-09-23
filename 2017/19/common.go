package main

import (
	"bufio"
	"os"
)

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

type Grid struct {
	Map       [][]byte
	X, Y      int
	Direction Direction
	Word      string
	Steps     int
}

func (g *Grid) Run(terminateFn func()) {
	g.Steps++

	switch g.Direction {
	case N:
		g.X--
	case S:
		g.X++
	case E:
		g.Y++
	case W:
		g.Y--
	}

	c := g.Map[g.X][g.Y]
	switch c {
	case '|', '-':
		return
	case '+':
		g.turn()
	case ' ':
		terminateFn()
	default:
		g.Word += string(c)
	}
}

func (g *Grid) turn() {
	switch g.Direction {
	case N, S:
		if g.Map[g.X][g.Y+1] != ' ' {
			g.Direction = E
		}
		if g.Map[g.X][g.Y-1] != ' ' {
			g.Direction = W
		}
	case E, W:
		if g.Map[g.X+1][g.Y] != ' ' {
			g.Direction = S
		}
		if g.Map[g.X-1][g.Y] != ' ' {
			g.Direction = N
		}
	}
}

func NewGrid() *Grid {
	grid := parse()
	x, y := start(grid)
	return &Grid{
		Map:       grid,
		X:         x,
		Y:         y,
		Direction: S,
		Word:      "",
		Steps:     0,
	}
}

func parse() [][]byte {
	grid := make([][]byte, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	return grid
}

func start(g [][]byte) (int, int) {
	for j := 0; j < len(g[0]); j++ {
		if g[0][j] != ' ' {
			return 0, j
		}
	}
	return 0, 0
}
