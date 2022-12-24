package main

import (
	"bufio"
	"os"
)

const (
	Width  = 102
	Height = 37
	LCM    = 700
)

const (
	N = iota
	E
	S
	W
)

type P struct{ r, c int }

var (
	Map      = map[P]int{}
	Blizzard = [LCM]map[P]struct{}{}
	Entrance = P{0, 1}
	Exit     = P{Height - 1, Width - 2}
	Dirs     = [5]P{
		{+1, 0}, // Move S
		{0, +1}, // Move E
		{-1, 0}, // Move N
		{0, -1}, // Move W
		{0, 0},  // Stay in place
	}
)

type Status struct {
	Min int
	Pos P
}

func bfs(from, to P, startMinute int) int {
	start := Status{
		Min: startMinute,
		Pos: from,
	}

	toDo := []Status{start}
	visited := map[Status]struct{}{start: {}}

	var curr Status
	for len(toDo) > 0 {
		curr, toDo = toDo[0], toDo[1:]

		for _, dir := range Dirs {

			adj := Status{
				Pos: P{curr.Pos.r + dir.r, curr.Pos.c + dir.c},
				Min: curr.Min + 1,
			}

			if adj.Pos == to {
				return adj.Min
			}

			if adj.Pos != Entrance && adj.Pos != Exit {
				if adj.Pos.r <= 0 || adj.Pos.r >= Height-1 {
					continue
				}
				if adj.Pos.c <= 0 || adj.Pos.c >= Width-1 {
					continue
				}

				if _, ok := Blizzard[adj.Min%LCM][adj.Pos]; ok {
					continue
				}
			}

			if _, ok := visited[adj]; !ok {
				visited[adj] = struct{}{}
				toDo = append(toDo, adj)
			}
		}
	}
	panic("not found")
}

func precomputeStates() {
	m := map[P][]int{}

	Blizzard[0] = map[P]struct{}{}
	for r := 0; r < Height; r++ {
		for c := 0; c < Width; c++ {
			p := P{r, c}
			val, ok := Map[p]
			if ok {
				m[p] = append(m[p], val)
				Blizzard[0][p] = struct{}{}
			}
		}
	}

	for t := 1; t < LCM; t++ {
		m = next(m)
		blizzard := map[P]struct{}{}
		for r := 1; r < Height-1; r++ {
			for c := 1; c < Width-1; c++ {
				p := P{r, c}
				if len(m[p]) > 0 {
					blizzard[p] = struct{}{}
				}
			}
		}
		Blizzard[t] = blizzard
	}
}

func next(old map[P][]int) map[P][]int {
	new := map[P][]int{}
	for r := 0; r < Height; r++ {
		for c := 0; c < Width; c++ {
			p := P{r, c}
			for _, item := range old[p] {

				var n P
				switch item {
				case N:
					n = P{p.r - 1, p.c}
					if n.r == 0 {
						n.r = Height - 2
					}
				case E:
					n = P{p.r, p.c + 1}
					if n.c == Width-1 {
						n.c = 1
					}
				case S:
					n = P{p.r + 1, p.c}
					if n.r == Height-1 {
						n.r = 1
					}
				case W:
					n = P{p.r, p.c - 1}
					if n.c == 0 {
						n.c = Width - 2
					}
				}
				new[n] = append(new[n], item)
			}
		}
	}
	return new
}

func parse() {
	scanner := bufio.NewScanner(os.Stdin)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		for col, char := range line {
			switch char {
			case '^':
				Map[P{row, col}] = N
			case '>':
				Map[P{row, col}] = E
			case 'v':
				Map[P{row, col}] = S
			case '<':
				Map[P{row, col}] = W
			}
		}
	}
}
