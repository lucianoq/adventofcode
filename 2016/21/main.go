package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const input = "abcdefgh"

func main() {
	str := String(input)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Fields(line)

		switch ff[0] {
		case "swap":
			if ff[1] == "position" {
				x, _ := strconv.Atoi(ff[2])
				y, _ := strconv.Atoi(ff[5])
				str.SwapPosition(x, y)
			} else {
				str.SwapLetters([]rune(ff[2])[0], []rune(ff[5])[0])
			}

		case "rotate":
			if ff[1] == "based" {
				str.RotateBasedOnLetter(ff[6])
			} else {
				x, _ := strconv.Atoi(ff[2])
				if ff[1] == "left" {
					str.Rotate(-x)
				} else {
					str.Rotate(x)
				}
			}
		case "reverse":
			x, _ := strconv.Atoi(ff[2])
			y, _ := strconv.Atoi(ff[4])
			str.Reverse(x, y)
		case "move":
			x, _ := strconv.Atoi(ff[2])
			y, _ := strconv.Atoi(ff[5])
			str.Move(x, y)
		}
	}
}

type String string

func (s *String) SwapPosition(x, y int) {
	str := []byte(*s)
	str[x], str[y] = str[y], str[x]
	*s = String(str)
}

func (s *String) SwapLetters(x, y rune) {
	str := ""
	for _, c := range string(*s) {
		switch c {
		case x:
			str += string(x)
		case y:
			str += string(y)
		default:
			str += string(c)
		}
	}
	*s = String(str)
}

func (s *String) Move(x, y int) {
	newString := "pippo"

	*s = String(newString)
}

func (s *String) Reverse(x, y int) {
	newString := "pippo"

	*s = String(newString)
}

func (s *String) RotateBasedOnLetter(x string) {
	newString := "pippo"

	*s = String(newString)
}

func (s *String) Rotate(x int) {
	newString := "pippo"

	*s = String(newString)
}
