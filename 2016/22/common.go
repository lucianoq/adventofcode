package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Device struct {
	P
	Size, Used, Avail int
}

type P struct {
	X, Y int
}

type Grid map[P]Device

func (g Grid) Copy() Grid {
	targetMap := make(Grid)
	for k, v := range g {
		targetMap[k] = v
	}
	return targetMap
}

func parseGrid() Grid {
	g := make(Grid)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "/dev/grid") {
			ff := strings.Fields(line)

			x, y := extractCoordinates(ff[0])
			size, _ := strconv.Atoi(strings.TrimRight(ff[1], "T"))
			used, _ := strconv.Atoi(strings.TrimRight(ff[2], "T"))
			avail, _ := strconv.Atoi(strings.TrimRight(ff[3], "T"))

			g[P{x, y}] = Device{
				P:     P{x, y},
				Size:  size,
				Used:  used,
				Avail: avail,
			}
		}
	}
	return g
}

func extractCoordinates(s string) (int, int) {
	ff := strings.Split(s, "-")
	x, _ := strconv.Atoi(strings.TrimLeft(ff[1], "x"))
	y, _ := strconv.Atoi(strings.TrimLeft(ff[2], "y"))
	return x, y
}
