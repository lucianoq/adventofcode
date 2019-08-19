package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var input = []byte("abcdefgh")

func main() {
	str := input
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Fields(line)

		switch ff[0] {
		case "swap":
			if ff[1] == "position" {
				x, _ := strconv.Atoi(ff[2])
				y, _ := strconv.Atoi(ff[5])
				str = SwapPosition(str, x, y)
			} else {
				str = SwapLetters(str, []byte(ff[2])[0], []byte(ff[5])[0])
			}

		case "rotate":
			if ff[1] == "based" {
				x := []byte(ff[6])
				str = RotateBasedOnLetter(str, x[0])
			} else {
				x, _ := strconv.Atoi(ff[2])
				if ff[1] == "left" {
					str = RotateLeft(str, x)
				} else {
					str = RotateLeft(str, -x)
				}
			}
		case "reverse":
			x, _ := strconv.Atoi(ff[2])
			y, _ := strconv.Atoi(ff[4])
			str = Reverse(str, x, y)
		case "move":
			x, _ := strconv.Atoi(ff[2])
			y, _ := strconv.Atoi(ff[5])
			str = Move(str, x, y)
		}
	}
}

func SwapPosition(s []byte, x, y int) []byte {
	newBuf := make([]byte, len(s))
	copy(newBuf, s)
	newBuf[x], newBuf[y] = s[y], s[x]
	return s
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
	val := s[x]

	for i:=0; i<len(s); i++ {
		newBuf = append(newBuf, )
	}

	if x < y {
		newBuf = append(newBuf, s[:x]...)
		newBuf = append(newBuf, s[x+1:]...)
		newBuf = append(newBuf, ...)
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
	return s
}

func RotateBasedOnLetter(s []byte, x byte) []byte {
	return s
}

func RotateLeft(s []byte, x int) []byte {
	return s
}
