package main

import (
	"bufio"
	"os"
	"strings"
)

type C struct{ x, y int }

func parse() map[C]struct{} {
	res := make(map[C]struct{})

	scanner := bufio.NewScanner(os.Stdin)

	var x, y int
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), "")
		for _, l := range line {
			if l == "#" {
				res[C{x, y}] = struct{}{}
			}
			x++
		}
		x = 0
		y++
	}

	return res
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
