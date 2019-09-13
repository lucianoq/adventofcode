package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)
	s := &Stream{buf: buf}
	count := s.ExtractGroup(0)
	fmt.Println(count)
}

func (s *Stream) ExtractGroup(level int) int {
	count := level
	for c := s.Next(); c != 0; c = s.Next() {
		switch c {

		case '}':
			return count

		case '{':
			count += s.ExtractGroup(level + 1)

		case '<':
			s.ExtractGarbage()

		case ',':
			fallthrough

		case ' ', '\n', '\r', '\t':
			// ignore spaces
			continue
		}
	}

	return count
}

func (s *Stream) ExtractGarbage() {
	for c := s.Next(); c != 0; c = s.Next() {
		switch c {
		case '>':
			return
		case '!':
			s.Next()
		}
	}
}
