package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Nanobot struct {
	X, Y, Z int
	R       int
}

func main() {
	nanobots, strongest := parse()

	count := 0
	for _, n := range nanobots {
		if strongest.InRange(n) {
			count++
		}
	}
	fmt.Println(count)
}

func parse() ([]Nanobot, Nanobot) {
	var strongest Nanobot
	nanobots := make([]Nanobot, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, z, r int
		n, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		if err != nil || n != 4 {
			log.Fatal(err)
		}

		nano := Nanobot{x, y, z, r}
		if nano.R > strongest.R {
			strongest = nano
		}
		nanobots = append(nanobots, nano)
	}

	return nanobots, strongest
}

func (n Nanobot) manhattan(n2 Nanobot) int {
	return abs(n.X-n2.X) + abs(n.Y-n2.Y) + abs(n.Z-n2.Z)
}

func (n Nanobot) InRange(n2 Nanobot) bool {
	return n.manhattan(n2) <= n.R
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

