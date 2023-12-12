package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ r, c int }

func parse() []P {
	var universe []P
	scanner := bufio.NewScanner(os.Stdin)
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		for c, ch := range line {
			if ch == '#' {
				universe = append(universe, P{r, c})
			}
		}
	}
	return universe
}

func findEmpty(galaxies []P) (map[int]struct{}, map[int]struct{}) {
	eRows := map[int]struct{}{}
	eCols := map[int]struct{}{}

	for i := 0; i < 140; i++ {

		foundRow, foundCol := false, false

		for _, g := range galaxies {
			if g.r == i {
				foundRow = true
			}
			if g.c == i {
				foundCol = true
			}
		}

		if !foundRow {
			eRows[i] = struct{}{}
		}
		if !foundCol {
			eCols[i] = struct{}{}

		}
	}
	return eRows, eCols
}

func manhattan(p1, p2 P) int {
	return abs(p1.r-p2.r) + abs(p1.c-p2.c)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sort(x1, x2 int) (int, int) {
	if x1 <= x2 {
		return x1, x2
	}
	return x2, x1
}

func main() {
	gxs := parse()

	eRows, eCols := findEmpty(gxs)

	sum := 0
	var min, max int
	for i := 0; i < len(gxs)-1; i++ {
		for j := i + 1; j < len(gxs); j++ {
			dist := manhattan(gxs[i], gxs[j])

			min, max = sort(gxs[i].r, gxs[j].r)
			for dr := min + 1; dr < max; dr++ {
				if _, ok := eRows[dr]; ok {
					dist += Expansion - 1
				}
			}

			min, max = sort(gxs[i].c, gxs[j].c)
			for dc := min + 1; dc < max; dc++ {
				if _, ok := eCols[dc]; ok {
					dist += Expansion - 1
				}
			}

			sum += dist
		}
	}

	fmt.Println(sum)
}
