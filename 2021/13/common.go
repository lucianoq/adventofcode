package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type C struct{ x, y int }

type Grid struct {
	grid          map[C]struct{}
	width, height int
}

type Fold struct {
	along string
	line  int
}

func parse() (*Grid, []Fold) {
	folds := []Fold{}
	grid := &Grid{
		grid: map[C]struct{}{},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold") {
			ff := strings.Split(line, " ")
			ff = strings.Split(ff[2], "=")
			num, _ := strconv.Atoi(ff[1])
			folds = append(folds, Fold{
				along: ff[0],
				line:  num,
			})
			continue
		}

		ff := strings.Split(line, ",")
		x, _ := strconv.Atoi(ff[0])
		y, _ := strconv.Atoi(ff[1])
		grid.grid[C{x, y}] = struct{}{}
		if x >= grid.width {
			grid.width = x + 1
		}
		if y >= grid.height {
			grid.height = y + 1
		}
	}

	return grid, folds
}

func fold(grid *Grid, fold Fold) {
	switch fold.along {
	case "x":
		grid.width = fold.line
	case "y":
		grid.height = fold.line
	}

	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {

			var newX, newY = x, y
			switch fold.along {
			case "x":
				newX = fold.line*2 - x
			case "y":
				newY = fold.line*2 - y
			}

			if _, ok := grid.grid[C{newX, newY}]; ok {
				grid.grid[C{x, y}] = struct{}{}
				delete(grid.grid, C{newX, newY})
			}
		}
	}
}

func countDots(grid *Grid) int {
	count := 0
	for k := range grid.grid {
		if k.x < grid.width && k.y < grid.height {
			count++
		}
	}
	return count
}
