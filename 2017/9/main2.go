package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)
	s := &Stream{buf: buf}
	count := s.ExtractGroup()
	fmt.Println(count)
}

func (s *Stream) ExtractGarbage() int {
	deleted := 0
	for c := s.Next(); c != 0; c = s.Next() {
		switch c {
		case '>':
			return deleted
		case '!':
			s.Next()
		default:
			deleted++
		}
	}
	log.Fatal("garbage must be closed")
	return deleted
}

func (s *Stream) ExtractGroup() int {
	deleted := 0
	for c := s.Next(); c != 0; c = s.Next() {
		switch c {

		case '}':
			return deleted

		case '{':
			deleted += s.ExtractGroup()

		case '<':
			deleted += s.ExtractGarbage()

		case ',':
			fallthrough

		case ' ', '\n', '\r', '\t':
			// ignore spaces
			continue
		}
	}

	return deleted
}
