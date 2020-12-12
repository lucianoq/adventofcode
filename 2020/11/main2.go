package main

import "fmt"

func main() {
	g := parse()

	for changed := true; changed; {
		g, changed = g.Step(g.InLine, 5)
	}

	fmt.Println(g.Occupied())
}

func (g Grid) InLine(p P) []P {
	res := make([]P, 0, 8)

	for dr := -1; dr <= 1; dr++ {

	NextDirection:
		for dc := -1; dc <= 1; dc++ {

			if dr != 0 || dc != 0 {

				for p := p.Add(dr, dc); !g.OutOfBounds(p); p = p.Add(dr, dc) {
					if g.grid[p] == Empty || g.grid[p] == Occupied {
						res = append(res, p)
						continue NextDirection
					}

				}
			}
		}
	}
	return res

}
