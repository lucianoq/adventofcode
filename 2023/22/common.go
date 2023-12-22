package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type P struct{ X, Y, Z int }

type Brick struct {
	ID         int
	Start, End P
}

type Space struct {
	Map    map[P]int
	Bricks map[int]*Brick
}

func parse() *Space {
	s := NewSpace()
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		ff := strings.Split(scanner.Text(), "~")
		var start, end P
		_, _ = fmt.Sscanf(ff[0], "%d,%d,%d", &start.X, &start.Y, &start.Z)
		_, _ = fmt.Sscanf(ff[1], "%d,%d,%d", &end.X, &end.Y, &end.Z)
		s.Add(&Brick{i, start, end})
	}
	return s
}

func NewSpace() *Space {
	return &Space{
		Map:    make(map[P]int),
		Bricks: make(map[int]*Brick),
	}
}

func (s *Space) Add(b *Brick) {
	s.Bricks[b.ID] = b
	for x := b.Start.X; x <= b.End.X; x++ {
		for y := b.Start.Y; y <= b.End.Y; y++ {
			for z := b.Start.Z; z <= b.End.Z; z++ {
				s.Map[P{x, y, z}] = b.ID
			}
		}
	}
}

func (s *Space) Fall() int {
	fallen := map[int]struct{}{}
	for changed := true; changed; {
		changed = false
		for _, b := range s.Bricks {
			for s.EmptyBelow(b) {
				s.Move(b, P{0, 0, -1})
				changed = true
				fallen[b.ID] = struct{}{}
			}
		}
	}
	return len(fallen)
}

func (s *Space) EmptyBelow(b *Brick) bool {
	if b.Start.Z == 1 {
		return false
	}
	for x := b.Start.X; x <= b.End.X; x++ {
		for y := b.Start.Y; y <= b.End.Y; y++ {
			if _, full := s.Map[P{x, y, b.Start.Z - 1}]; full {
				return false
			}
		}
	}
	return true
}

func (s *Space) Move(b *Brick, p P) {
	for x := b.Start.X; x <= b.End.X; x++ {
		for y := b.Start.Y; y <= b.End.Y; y++ {
			for z := b.Start.Z; z <= b.End.Z; z++ {
				delete(s.Map, P{x, y, z})
			}
		}
	}
	for x := b.Start.X + p.X; x <= b.End.X+p.X; x++ {
		for y := b.Start.Y + p.Y; y <= b.End.Y+p.Y; y++ {
			for z := b.Start.Z + p.Z; z <= b.End.Z+p.Z; z++ {
				s.Map[P{x, y, z}] = b.ID
			}
		}
	}

	s.Bricks[b.ID].Start.X += p.X
	s.Bricks[b.ID].Start.Y += p.Y
	s.Bricks[b.ID].Start.Z += p.Z
	s.Bricks[b.ID].End.X += p.X
	s.Bricks[b.ID].End.Y += p.Y
	s.Bricks[b.ID].End.Z += p.Z
}

func (s *Space) TopOf(b *Brick) []*Brick {
	var onTop []*Brick
	for x := b.Start.X; x <= b.End.X; x++ {
		for y := b.Start.Y; y <= b.End.Y; y++ {
			if id, ok := s.Map[P{x, y, b.End.Z + 1}]; ok {
				onTop = append(onTop, s.Bricks[id])
			}
		}
	}
	return onTop
}

func (s *Space) SafeToDisintegrate(b *Brick) bool {
	for _, topB := range s.TopOf(b) {
		if s.DependsOn(topB, b) {
			return false
		}
	}
	return true
}

func (s *Space) DependsOn(topB, botB *Brick) bool {
	for x := topB.Start.X; x <= topB.End.X; x++ {
		for y := topB.Start.Y; y <= topB.End.Y; y++ {
			if id, full := s.Map[P{x, y, topB.Start.Z - 1}]; full && id != botB.ID {
				return false
			}
		}
	}
	return true
}
