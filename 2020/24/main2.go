package main

import "fmt"

func main() {
	floor := parse()

	for t := 0; t < 100; t++ {
		floor = move(floor)
	}

	fmt.Println(countBlacks(floor))
}

func move(floor map[P]bool) map[P]bool {
	newFloor := map[P]bool{}

	toConsider := map[P]struct{}{}
	for k := range floor {
		toConsider[k] = struct{}{}
		for _, p := range neighbours(k) {
			toConsider[p] = struct{}{}
		}
	}

	for p := range toConsider {
		countBlacks := 0
		for _, n := range neighbours(p) {
			if floor[n] {
				countBlacks++
			}
		}

		if floor[p] {
			newFloor[p] = countBlacks > 0 && countBlacks <= 2
		} else {
			newFloor[p] = countBlacks == 2
		}
	}
	return newFloor
}

func neighbours(p P) []P {
	return []P{
		{p.X - 1, p.Y + 1}, // nw
		{p.X, p.Y + 1},     // ne
		{p.X + 1, p.Y},     // e
		{p.X + 1, p.Y - 1}, // se
		{p.X, p.Y - 1},     // sw
		{p.X - 1, p.Y},     // w
	}
}
