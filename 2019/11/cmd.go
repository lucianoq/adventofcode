package main

import "math"

type Mode int

const (
	Position Mode = iota
	Immediate
	Relative
)

type Cmd int

func (c Cmd) OpCode() int {
	// last 2 digits
	return int(c) % 100
}

func (c Cmd) Modes(arity int) []Mode {
	// remove last two digits
	modeSection := int(c) / 100

	modes := make([]Mode, arity)
	for i := 0; i < arity; i++ {
		modes[i] = Mode(modeSection / int(math.Pow10(i)) % 10)
	}
	return modes
}
