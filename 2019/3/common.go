package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type C struct{ X, Y int }

func run(fn func(C, int, int) int) int {
	buf, _ := ioutil.ReadAll(os.Stdin)
	wires := strings.SplitN(string(buf), "\n", 2)
	cells0, cells1 := lineToCells(wires[0]), lineToCells(wires[1])

	distances := make([]int, 0)

	// for all cells where wire0 passed through
	for k := range cells0 {

		// if wire1 passed through too
		if cells1[k] > 0 {

			distances = append(distances, fn(k, cells0[k], cells1[k]))

		}
	}

	return minInt(distances)
}

func lineToCells(line string) map[C]int {
	movements := strings.Split(line, ",")

	wire := make(map[C]int)
	stepCounter := 0

	pos := C{} // start from 0

	for _, mov := range movements {
		var offset C
		switch mov[0] {
		case 'U':
			offset = C{0, 1}
		case 'R':
			offset = C{1, 0}
		case 'D':
			offset = C{0, -1}
		case 'L':
			offset = C{-1, 0}
		}

		steps, _ := strconv.Atoi(mov[1:])
		for i := 1; i <= steps; i++ {
			pos.X += offset.X
			pos.Y += offset.Y
			stepCounter++
			wire[pos] = stepCounter
		}
	}

	return wire
}

func minInt(list []int) int {
	minDist := 1<<63 - 1
	for _, d := range list {
		if d < minDist {
			minDist = d
		}
	}

	return minDist
}
