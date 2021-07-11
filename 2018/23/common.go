package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type P struct {
	X, Y, Z int
}

type Nanobot struct {
	C P
	R int
}

func (n Nanobot) InRange(p P) bool {
	return manhattan(n.C, p) <= n.R
}

func parse() []Nanobot {
	nanobots := make([]Nanobot, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, z, r int
		n, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		if err != nil || n != 4 {
			log.Fatal(err)
		}
		nanobots = append(nanobots, Nanobot{
			C: P{x, y, z},
			R: r,
		})
	}
	return nanobots
}

func manhattan(p1, p2 P) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y) + abs(p1.Z-p2.Z)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
