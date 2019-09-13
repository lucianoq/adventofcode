package main

type Stream struct {
	buf []byte
	pos int
}

func (s *Stream) Next() byte {
	if s.pos >= len(s.buf) {
		return 0
	}
	char := s.buf[s.pos]
	s.pos++
	return char
}
