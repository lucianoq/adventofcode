package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type P struct{ X, Y, Z int }

func (p P) Add(p2 P) P {
	return P{p.X + p2.X, p.Y + p2.Y, p.Z + p2.Z}
}

var Adj = []P{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func parse() map[P]struct{} {
	scanner := bufio.NewScanner(os.Stdin)
	space := map[P]struct{}{}
	for scanner.Scan() {
		ff := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(ff[0])
		y, _ := strconv.Atoi(ff[1])
		z, _ := strconv.Atoi(ff[2])
		space[P{x, y, z}] = struct{}{}
	}
	return space
}
