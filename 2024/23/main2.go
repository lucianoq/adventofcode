package main

import "fmt"

func main() {
	g := parseInput()

	var lastNotEmpty Set
	for res := g.Keys(); len(res) > 0; res = g.Expand(res) {
		lastNotEmpty = res
	}

	for k := range lastNotEmpty {
		fmt.Println(k)
		return
	}
}
