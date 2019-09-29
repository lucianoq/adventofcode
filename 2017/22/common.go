package main

import (
	"bufio"
	"log"
	"os"
)

type (
	C struct{ X, Y int }

	Grid map[C]Status

	Direction int

	Status int

	Virus struct {
		Pos                   C
		Dir                   Direction
		G                     Grid
		BurstsCausedInfection int
	}
)

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

const (
	CLEAN Status = iota
	WEAKENED
	INFECTED
	FLAGGED
)

func (v *Virus) Burst() {
	switch v.G[v.Pos] {
	case CLEAN:
		v.Left()
	case INFECTED:
		v.Right()
	case FLAGGED:
		v.Reverse()
	}

	v.Toggle()

	if v.G[v.Pos] == INFECTED {
		v.BurstsCausedInfection++
	}

	v.Forward()
}

func (v *Virus) Toggle() {
	v.G[v.Pos] = v.G[v.Pos].Touch()
}

func (v *Virus) Right() {
	v.Dir = (v.Dir + 1) % 4
}
func (v *Virus) Reverse() {
	v.Dir = (v.Dir + 2) % 4
}

func (v *Virus) Left() {
	v.Dir = (v.Dir + 3) % 4
}

func (v *Virus) Forward() {
	switch v.Dir {
	case UP:
		v.Pos.X--
	case DOWN:
		v.Pos.X++
	case RIGHT:
		v.Pos.Y++
	case LEFT:
		v.Pos.Y--
	default:
		log.Fatal("wrong dir")
	}

	v.G[v.Pos] = v.G[v.Pos]
}

func parse() Grid {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sizeX, sizeY := len(lines), len(lines[0])
	cX, cY := sizeX/2, sizeY/2

	grid := make(Grid)

	for i := 0; i < sizeX; i++ {
		for j, r := range lines[i] {
			switch r {
			case '.':
				grid[C{i - cX, j - cY}] = CLEAN
			case '#':
				grid[C{i - cX, j - cY}] = INFECTED
			default:
				log.Fatal("wrong char")
			}
		}
	}

	return grid
}
