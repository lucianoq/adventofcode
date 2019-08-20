package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = []byte("fbgdceah")

func main() {
	str := input

	lines := parse()

	for i := len(lines) - 1; i >= 0; i-- {

		line := lines[i]

		ff := strings.Fields(line)

		switch ff[0] {
		case "swap":
			if ff[1] == "position" {
				x, _ := strconv.Atoi(ff[2])
				y, _ := strconv.Atoi(ff[5])
				str = InverseSwapPosition(str, x, y)
			} else {
				str = InverseSwapLetters(str, []byte(ff[2])[0], []byte(ff[5])[0])
			}

		case "rotate":
			if ff[1] == "based" {
				x := []byte(ff[6])
				str = InverseRotateBasedOnLetter(str, x[0])
			} else {
				x, _ := strconv.Atoi(ff[2])
				if ff[1] == "left" {
					str = InverseRotateLeft(str, x)
				} else {
					str = InverseRotateLeft(str, -x)
				}
			}
		case "reverse":
			x, _ := strconv.Atoi(ff[2])
			y, _ := strconv.Atoi(ff[4])
			str = InverseReverse(str, x, y)
		case "move":
			x, _ := strconv.Atoi(ff[2])
			y, _ := strconv.Atoi(ff[5])
			str = InverseMove(str, x, y)
		}
	}

	fmt.Println(string(str))
}
