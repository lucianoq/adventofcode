package main

import "fmt"

func main() {
	g := parse()

	for changed := true; changed; {
		g, changed = g.Step(g.Adjacent, 4)
	}

	fmt.Println(g.Occupied())
}

func (g Grid) Adjacent(p P) []P {
	res := make([]P, 0, 8)
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr != 0 || dc != 0 {
				newP := p.Add(dr, dc)
				if !g.OutOfBounds(newP) {
					res = append(res, newP)
				}
			}
		}
	}
	return res
}
