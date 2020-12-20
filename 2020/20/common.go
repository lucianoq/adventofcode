package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Tile [10][10]bool

type Side uint

const (
	Up Side = iota
	Right
	Down
	Left
	UpFlip
	RightFlip
	DownFlip
	LeftFlip
)

func parse() map[int]*Tile {
	tiles := map[int]*Tile{}
	tile := &Tile{}
	latestID := -1
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if line == "" {
			tiles[latestID] = tile
		}

		if strings.HasPrefix(line, "Tile ") {
			tile = &Tile{}
			line := strings.Trim(line, "Tile :")
			latestID, _ = strconv.Atoi(line)
			i = -1
		}

		for j, c := range line {
			if c == '#' {
				tile[i][j] = true
			}
		}
	}

	return tiles
}

func (t Tile) Edge(s Side) [10]bool {
	edge := [10]bool{}

	switch s {
	case Up:
		for j := 0; j < 10; j++ {
			edge[j] = t[0][j]
		}
	case Right:
		for i := 0; i < 10; i++ {
			edge[i] = t[i][10-1]
		}
	case Down:
		for j := 0; j < 10; j++ {
			edge[j] = t[10-1][10-j-1]
		}
	case Left:
		for i := 0; i < 10; i++ {
			edge[i] = t[10-i-1][0]
		}
	case UpFlip:
		for j := 0; j < 10; j++ {
			edge[j] = t[0][10-j-1]
		}
	case RightFlip:
		for i := 0; i < 10; i++ {
			edge[i] = t[10-i-1][10-1]
		}
	case DownFlip:
		for j := 0; j < 10; j++ {
			edge[j] = t[10-1][j]
		}
	case LeftFlip:
		for i := 0; i < 10; i++ {
			edge[i] = t[i][0]
		}
	}

	return edge
}

func Overlap(edge1, edge2 [10]bool) bool {
	for i := 0; i < 10; i++ {
		if edge1[i] != edge2[10-i-1] {
			return false
		}
	}
	return true
}
