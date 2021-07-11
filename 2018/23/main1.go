package main

import "fmt"

func main() {
	nanobots := parse()

	var strongest Nanobot
	for _, nano := range nanobots {
		if nano.R > strongest.R {
			strongest = nano
		}
	}

	count := 0
	for _, n := range nanobots {
		if strongest.InRange(P{n.C.X, n.C.Y, n.C.Z}) {
			count++
		}
	}
	fmt.Println(count)
}
