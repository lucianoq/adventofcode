package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

func (s *Santa) Visited() int {
	count := 0
	for _, v := range s.space {
		if v {
			count++
		}
	}
	return count
}

func NewSanta() *Santa {
	return &Santa{
		space: make(map[Point]bool),
		ptr:   Point{0, 0},
	}
}

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(buf))

	santa := NewSanta()
	santa.Visit()

	for _, c := range input {
		switch c {
		case '^':
			santa.North()
		case '<':
			santa.West()
		case '>':
			santa.East()
		case 'v':
			santa.South()
		default:
			log.Printf("unknown character %+v", c)
		}
	}

	fmt.Println(santa.Visited())
}
