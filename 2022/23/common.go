package main

import (
	"bufio"
	"os"
)

type P struct{ X, Y int }

type Elf struct {
	Pos     P
	Moving  bool
	NextPos P
}

const (
	N int = iota + 1
	E     = 2*iota + 1
	S
	W
)

var (
	Map     = map[P]struct{}{}
	Elves   []*Elf
	Order   = [4]int{N, S, W, E}
	CurrDir = 0
	Dirs    = [8]P{
		{-1, -1}, // NW
		{-1, 0},  // N
		{-1, +1}, // NE
		{0, +1},  // E
		{+1, +1}, // SE
		{+1, 0},  // S
		{+1, -1}, // SW
		{0, -1},  // W
	}
)

func (e *Elf) AroundAllEmpty() bool {
	for _, d := range Dirs {
		adj := P{e.Pos.X + d.X, e.Pos.Y + d.Y}
		if _, ok := Map[adj]; ok {
			// at least one is not empty
			return false
		}
	}
	return true
}

func (e *Elf) ElfInDirection(wannaGo int) bool {
	for j := -1; j <= 1; j++ {
		dxy := Dirs[(wannaGo+j+8)%8]
		adj := P{e.Pos.X + dxy.X, e.Pos.Y + dxy.Y}
		if _, ok := Map[adj]; ok {
			return true
		}
	}
	return false
}

func run() bool {
	proposes := map[P]int{}

	for _, e := range Elves {
		if e.AroundAllEmpty() {
			continue
		}

		for i := 0; i < 4; i++ {
			dir := Order[(CurrDir+i)%4]

			if e.ElfInDirection(dir) {
				continue
			}

			dxy := Dirs[dir]
			dest := P{e.Pos.X + dxy.X, e.Pos.Y + dxy.Y}
			proposes[dest]++
			e.NextPos = dest
			e.Moving = true
			break
		}
	}

	someoneMoved := false
	for _, e := range Elves {

		if !e.Moving {
			continue
		}

		if proposes[e.NextPos] > 1 {
			e.Moving = false
			continue
		}

		someoneMoved = true
		delete(Map, e.Pos)
		Map[e.NextPos] = struct{}{}
		e.Pos = e.NextPos
		e.Moving = false
	}

	CurrDir = (CurrDir + 1) % 4

	return someoneMoved
}

func minMax() (P, P) {
	min := P{1 << 30, 1 << 30}
	max := P{-(1 << 30), -(1 << 30)}
	for p := range Map {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}
	return min, max
}

func parse() {
	scanner := bufio.NewScanner(os.Stdin)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		for col, char := range line {
			if char == '#' {
				p := P{row, col}
				Map[p] = struct{}{}
				Elves = append(Elves, &Elf{Pos: p})
			}
		}
	}
}
