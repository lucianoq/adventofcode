package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const NumSantas = 2

type Point struct{ X, Y int }

type Santa struct {
	space map[Point]bool
	ptr   Point
}

func (s *Santa) North() {
	s.ptr.Y++
	s.Visit()
}

func (s *Santa) South() {
	s.ptr.Y--
	s.Visit()
}

func (s *Santa) West() {
	s.ptr.X--
	s.Visit()
}

func (s *Santa) East() {
	s.ptr.X++
	s.Visit()
}

func (s *Santa) Visit() {
	s.space[Point{s.ptr.X, s.ptr.Y}] = true
}

func visited(space map[Point]bool) int {
	count := 0
	for _, v := range space {
		if v {
			count++
		}
	}
	return count
}

func NewSanta(space map[Point]bool) *Santa {
	return &Santa{
		space: space,
		ptr:   Point{0, 0},
	}
}

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(buf))

	space := make(map[Point]bool)
	var santas [NumSantas]*Santa

	for i := 0; i < NumSantas; i++ {
		santas[i] = NewSanta(space)
		santas[i].Visit()
	}

	for i, c := range input {
		s := santas[i%NumSantas]
		switch c {
		case '^':
			s.North()
		case '<':
			s.West()
		case '>':
			s.East()
		case 'v':
			s.South()
		default:
			log.Printf("unknown character %+v", c)
		}
	}

	fmt.Println(visited(space))
}
