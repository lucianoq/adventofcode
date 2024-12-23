package main

import (
	"fmt"
	"strings"
)

func main() {
	g := parseInput()

	res := g.Keys()     // single nodes
	res = g.Expand(res) // pairs
	res = g.Expand(res) // triplets

	count := 0
	for c := range res {
		if c[0] == 't' || strings.Contains(c, ",t") {
			count++
		}
	}
	fmt.Println(count)
}
