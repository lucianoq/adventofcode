package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	updates, cmpFn := parseInput()

	sum := 0
	for _, u := range updates {
		if slices.IsSortedFunc(u, cmpFn) {
			n, _ := strconv.Atoi(u[len(u)/2])
			sum += n
		}
	}

	fmt.Println(sum)
}
