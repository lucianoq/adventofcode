package main

import (
	"bufio"
	"log"
	"os"
)

const (
	N = iota
	E
	S
	W
)

type P struct{ r, c int }

func (p P) N() P { return P{p.r - 1, p.c} }
func (p P) S() P { return P{p.r + 1, p.c} }
func (p P) W() P { return P{p.r, p.c - 1} }
func (p P) E() P { return P{p.r, p.c + 1} }

func parse() (map[P]byte, P) {
	scanner := bufio.NewScanner(os.Stdin)

	m := map[P]byte{}
	var start P

	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()

		for c := 0; c < len(line); c++ {

			p := P{r, c}

			switch line[c] {
			case 'S':
				start = p
				m[p] = '-'
			case '.':
				continue
			default:
				m[p] = line[c]
			}
		}
	}
	return m, start
}

type State struct {
	Current P
	From    int
}

func (s State) Next(m map[P]byte) State {
	switch m[s.Current] {
	case '|':
		switch s.From {
		case N:
			return State{s.Current.S(), N}
		case S:
			return State{s.Current.N(), S}
		}
	case '-':
		switch s.From {
		case W:
			return State{s.Current.E(), W}
		case E:
			return State{s.Current.W(), E}
		}
	case 'L':
		switch s.From {
		case N:
			return State{s.Current.E(), W}
		case E:
			return State{s.Current.N(), S}
		}
	case 'J':
		switch s.From {
		case N:
			return State{s.Current.W(), E}
		case W:
			return State{s.Current.N(), S}
		}
	case '7':
		switch s.From {
		case S:
			return State{s.Current.W(), E}
		case W:
			return State{s.Current.S(), N}
		}
	case 'F':
		switch s.From {
		case S:
			return State{s.Current.E(), W}
		case E:
			return State{s.Current.S(), N}
		}
	}
	log.Fatal("should never happen")
	return State{}
}
