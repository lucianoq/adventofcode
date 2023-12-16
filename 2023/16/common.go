package main

import (
	"bufio"
	"github.com/lucianoq/container/set"
	"os"
)

type P struct{ R, C int }

type Direction uint8

const (
	N Direction = iota
	E
	S
	W
)

type Map struct {
	M          map[P]byte
	MaxR, MaxC int
}

type Beam struct {
	R, C int
	Dir  Direction
}

func (b *Beam) Right() { b.Dir = (b.Dir + 1) % 4 }
func (b *Beam) Left()  { b.Dir = (b.Dir + 3) % 4 }

func (b *Beam) Move(m Map) ([]Beam, set.Set[P]) {
	energized := set.New[P]()

	for {
		switch b.Dir {
		case N:
			b.R--
		case S:
			b.R++
		case E:
			b.C++
		case W:
			b.C--
		}

		// out of bounds, stop here
		if b.R < 0 || b.R >= m.MaxR || b.C < 0 || b.C >= m.MaxC {
			return nil, energized
		}

		pt := P{b.R, b.C}

		energized.Add(pt)

		switch m.M[pt] {
		case '|':
			if b.Dir == E || b.Dir == W {
				return []Beam{
					{b.R, b.C, N},
					{b.R, b.C, S},
				}, energized
			}

		case '-':
			if b.Dir == N || b.Dir == S {
				return []Beam{
					{b.R, b.C, W},
					{b.R, b.C, E},
				}, energized
			}

		case '/':
			switch b.Dir {
			case N, S:
				b.Right()
			case E, W:
				b.Left()
			}

		case '\\':
			switch b.Dir {
			case N, S:
				b.Left()
			case E, W:
				b.Right()
			}
		}
	}
}

func energy(m Map, start Beam) int {
	var (
		visited   = set.Set[Beam]{}
		energized = set.New[P]()
		beams     = []Beam{start}
		beam      Beam
	)

	for len(beams) > 0 {
		beam, beams = beams[0], beams[1:]

		if visited.Contains(beam) {
			continue
		}
		visited.Add(beam)

		newBeams, pointsWentThrough := beam.Move(m)
		energized.AddSets(pointsWentThrough)
		beams = append(beams, newBeams...)
	}

	return len(energized)
}

func parse() Map {
	m := map[P]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	var r, c int
	for r = 0; scanner.Scan(); r++ {
		line := scanner.Text()
		for c = 0; c < len(line); c++ {
			if line[c] != '.' {
				m[P{r, c}] = line[c]
			}
		}
	}
	return Map{m, r, c}
}
