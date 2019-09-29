package main

import "fmt"

func main() {
	t := Turing{
		Tape:   make(map[int]bool),
		Cursor: 0,
	}
	t.NextState = t.A

	for i := 0; i < 12134527; i++ {
		t.NextState()
	}

	count := 0
	for _, v := range t.Tape {
		if v {
			count++
		}
	}
	fmt.Println(count)
}

type Turing struct {
	Tape      map[int]bool
	Cursor    int
	NextState func()
}

func (t *Turing) A() {
	if !t.Tape[t.Cursor] {
		t.state(true, 1, t.B)
	} else {
		t.state(false, -1, t.C)
	}
}

func (t *Turing) B() {
	if !t.Tape[t.Cursor] {
		t.state(true, -1, t.A)
	} else {
		t.state(true, 1, t.C)
	}
}

func (t *Turing) C() {
	if !t.Tape[t.Cursor] {
		t.state(true, 1, t.A)
	} else {
		t.state(false, -1, t.D)
	}
}

func (t *Turing) D() {
	if !t.Tape[t.Cursor] {
		t.state(true, -1, t.E)
	} else {
		t.state(true, -1, t.C)
	}
}

func (t *Turing) E() {
	if !t.Tape[t.Cursor] {
		t.state(true, 1, t.F)
	} else {
		t.state(true, 1, t.A)
	}
}

func (t *Turing) F() {
	if !t.Tape[t.Cursor] {
		t.state(true, 1, t.A)
	} else {
		t.state(true, 1, t.E)
	}
}

func (t *Turing) state(val bool, direction int, nextState func()) {
	t.Tape[t.Cursor] = val
	t.Cursor += direction
	t.NextState = nextState
}
