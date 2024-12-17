package main

import (
	"fmt"
	"slices"
)

func main() {
	_, b, c, code := parseInput()

	a := 0
	for i := len(code) - 1; i >= 0; i-- {
		// make room for 3 bits by shifting to the left
		a <<= 3

		// check incrementally only the latest bits,
		// until we find the right value
		for !slices.Equal(run(a, b, c, code), code[i:]) {
			a++
		}
	}
	fmt.Println(a)
}
