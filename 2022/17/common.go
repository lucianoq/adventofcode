package main

import (
	"bufio"
	"os"
)

var (
	Map = map[P]struct{}{}

	Rocks = [][]P{
		// |..@@@@.|
		{{2, 0}, {3, 0}, {4, 0}, {5, 0}},

		// |...@...|
		// |..@@@..|
		// |...@...|
		{{3, 0}, {2, 1}, {3, 1}, {4, 1}, {3, 2}},

		// |....@..|
		// |....@..|
		// |..@@@..|
		{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}},

		// |..@....|
		// |..@....|
		// |..@....|
		// |..@....|
		{{2, 0}, {2, 1}, {2, 2}, {2, 3}},

		// |..@@...|
		// |..@@...|
		{{2, 0}, {3, 0}, {2, 1}, {3, 1}},
	}

	Directions = map[byte]P{
		'<': {-1, 0},
		'>': {1, 0},
		'v': {0, -1},
	}
)

type P struct{ X, Y int }

func (p P) Add(delta P) P {
	p.X += delta.X
	p.Y += delta.Y
	return p
}

func parse() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getDir(input string) <-chan byte {
	ch := make(chan byte)
	go func() {
		for {
			for i := 0; i < len(input); i++ {
				ch <- input[i]
			}
		}
	}()
	return ch
}

func getRock(round, height int) []P {
	rock := Rocks[round%len(Rocks)]

	newRock := make([]P, len(rock))
	for i := 0; i < len(rock); i++ {
		newRock[i].X = rock[i].X
		newRock[i].Y = rock[i].Y + height + 3
	}

	return newRock
}

func runRound(round, peak int, dir <-chan byte) int {
	rock := getRock(round, peak)
	for {
		// Pushed by jet
		direction := <-dir
		if newRock, ok := move(rock, direction); ok {
			rock = newRock
		}

		// Down by 1
		if newRock, ok := move(rock, 'v'); ok {
			rock = newRock
			continue
		}

		// we stopped going down. Freeze
		for _, p := range rock {
			Map[p] = struct{}{}

			if p.Y+1 > peak {
				peak = p.Y + 1
			}
		}

		return peak
	}
}

func move(rock []P, direction byte) ([]P, bool) {
	newRock := make([]P, len(rock))
	for i, p := range rock {
		p = p.Add(Directions[direction])

		// out of borders
		if p.X < 0 || p.X > 6 || p.Y < 0 {
			return nil, false
		}

		// clash with existing
		if _, ok := Map[p]; ok {
			return nil, false
		}

		newRock[i] = p
	}

	return newRock, true
}
