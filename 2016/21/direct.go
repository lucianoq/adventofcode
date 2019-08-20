package main

import "bytes"

func SwapPosition(s []byte, x, y int) []byte {
	newBuf := make([]byte, len(s))
	copy(newBuf, s)
	newBuf[x], newBuf[y] = s[y], s[x]
	return newBuf
}

func SwapLetters(s []byte, x, y byte) []byte {
	newBuf := make([]byte, len(s))
	for i := range s {
		switch s[i] {
		case x:
			newBuf[i] = y
		case y:
			newBuf[i] = x
		default:
			newBuf[i] = s[i]
		}
	}
	return newBuf
}

func Move(s []byte, x, y int) []byte {
	newBuf := make([]byte, 0, len(s))
	if x < y {
		newBuf = append(newBuf, s[:x]...)
		newBuf = append(newBuf, s[x+1:y+1]...)
		newBuf = append(newBuf, s[x])
		newBuf = append(newBuf, s[y+1:]...)
	} else {
		newBuf = append(newBuf, s[:y]...)
		newBuf = append(newBuf, s[x])
		newBuf = append(newBuf, s[y:x]...)
		newBuf = append(newBuf, s[x+1:]...)
	}
	return newBuf
}

func reverse(a []byte) []byte {
	b := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		opp := len(a) - i - 1
		b[i] = a[opp]
	}
	return b
}

func Reverse(s []byte, x, y int) []byte {
	newBuf := make([]byte, 0, len(s))
	newBuf = append(newBuf, s[:x]...)
	newBuf = append(newBuf, reverse(s[x:y+1])...)
	newBuf = append(newBuf, s[y+1:]...)
	return newBuf
}

func RotateBasedOnLetter(s []byte, x byte) []byte {
	idx := bytes.IndexByte(s, x)
	if idx >= 4 {
		idx++
	}
	return RotateLeft(s, -(1 + idx))
}

func RotateLeft(s []byte, x int) []byte {
	for x < 0 {
		x += 10 * len(s)
	}
	newBuf := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		newBuf[i] = s[(i+x)%len(s)]
	}

	return newBuf
}
