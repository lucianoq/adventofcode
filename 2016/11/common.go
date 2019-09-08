package main

import (
	"fmt"
	"log"
)

type Atom uint

const (
	PoloniumChip Atom = iota
	ThuliumChip
	PromethiumChip
	RutheniumChip
	CobaltChip
	EleriumChip
	DilithiumChip
)

const (
	PoloniumRTG Atom = 8 + iota
	ThuliumRTG
	PromethiumRTG
	RutheniumRTG
	CobaltRTG
	EleriumRTG
	DilithiumRTG
)

type Building struct {
	Floors   [4]Floor
	Elevator int
	Parent   *Building
}

func (b Building) Valid() bool {
	for _, f := range b.Floors {
		if !f.Valid() {
			return false
		}
	}
	return true
}

func (b *Building) String() string {
	s := fmt.Sprintf("Elevator:%d Floors:", b.Elevator)
	for i := 0; i < 4; i++ {
		atoms := b.Floors[i].Atoms()
		s += fmt.Sprintf("%+v\t", atoms)
	}
	return s + "\n"
}

func (b Building) Norm() Building {
	b.Parent = nil
	return b
}

func (b Building) AllCombinations() []Building {
	buildings := make([]Building, 0)

	pairs := Pairify(b.Floors[b.Elevator].Atoms())
	for _, p := range pairs {
		if b.Elevator != 3 {
			newB := b.MoveUp(p)
			if newB.Valid() {
				buildings = append(buildings, newB)
			}
		}

		if b.Elevator != 0 {
			newB := b.MoveDown(p)
			if newB.Valid() {
				buildings = append(buildings, newB)
			}
		}
	}

	return buildings
}

func (b Building) MoveUp(atoms []Atom) Building {
	for _, a := range atoms {
		b.Floors[b.Elevator].Remove(a)
		b.Floors[b.Elevator+1].Add(a)
	}
	b.Elevator++
	return b
}

func (b Building) MoveDown(atoms []Atom) Building {
	for _, a := range atoms {
		b.Floors[b.Elevator].Remove(a)
		b.Floors[b.Elevator-1].Add(a)
	}
	b.Elevator--
	return b
}

type Floor uint16

func (f Floor) Valid() bool {
	// shift to right
	rtgs := f >> 8

	// take only interesting 8 bits
	chips := f & 0xFF

	chipsAlone := chips &^ rtgs

	// if there are no chips alone, is valid
	if chipsAlone == 0 {
		return true
	}

	// from here there are chips alone, so
	// if there is at least 1 RTG, is not valid
	if rtgs != 0 {
		return false
	}

	return true
}

func (f *Floor) Add(atoms ...Atom) {
	for _, a := range atoms {
		*f |= 1 << a
	}
}

func (f *Floor) Remove(atoms ...Atom) {
	for _, a := range atoms {
		*f &^= 1 << a
	}
}

func (f Floor) Atoms() []Atom {
	var atoms []Atom

	var i Atom
	for i = 0; i < 16; i++ {
		if f&(1<<i) != 0 {
			atoms = append(atoms, i)
		}
	}
	return atoms
}

func Pairify(atoms []Atom) [][]Atom {
	pairs := make([][]Atom, 0)

	for _, a := range atoms {
		pairs = append(pairs, []Atom{a})
	}

	for i := 0; i < len(atoms)-1; i++ {
		for j := i + 1; j < len(atoms); j++ {
			pairs = append(pairs, []Atom{atoms[i], atoms[j]})
		}
	}
	return pairs
}

func bfs(start, end Building) *Building {
	var todo = make([]Building, 0)
	var visited = make(map[Building]bool)

	todo = append(todo, start)

	for len(todo) > 0 {
		var v Building
		v, todo = todo[0], todo[1:]

		if v.Floors == end.Floors {
			return &v
		}

		for _, b := range v.AllCombinations() {
			if !visited[b.Norm()] {
				visited[b.Norm()] = true
				b.Parent = &v
				todo = append(todo, b)
			}
		}
	}

	log.Fatal("not found")
	return &Building{}
}

func countPath(b *Building) int {
	count := 0
	for b != nil {
		log.Printf("%s", b)
		count++
		b = b.Parent
	}

	// First status shouldn't be counted
	count--

	return count
}
