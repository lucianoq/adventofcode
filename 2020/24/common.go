package main

import (
	"bufio"
	"os"
)

type P struct{ X, Y int }

func popToken(input string) (string, string) {
	first := string(input[0])
	if first == "w" || first == "e" {
		return first, input[1:]
	}
	return input[:2], input[2:]
}

func parse() map[P]bool {
	floor := map[P]bool{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		x, y := 0, 0

		var token string
		for len(input) > 0 {
			token, input = popToken(input)

			switch token {
			case "ne":
				y++
			case "se":
				y--
				x++
			case "nw":
				y++
				x--
			case "sw":
				y--
			case "w":
				x--
			case "e":
				x++
			}
		}

		floor[P{x, y}] = !floor[P{x, y}]
	}

	return floor
}

func countBlacks(floor map[P]bool) int {
	count := 0
	for _, v := range floor {
		if v {
			count++
		}
	}
	return count
}
