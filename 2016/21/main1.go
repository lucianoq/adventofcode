package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = []byte("abcdefgh")

func main() {
	str := input

	for _, line := range parse() {
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

	fmt.Println(string(str))
}
