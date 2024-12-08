package main

import "fmt"

func main() {
	freqs := parseInput()

	antiNodes := map[P]struct{}{}
	for _, locs := range freqs {
		for a := 0; a < len(locs)-1; a++ {
			for b := a + 1; b < len(locs); b++ {
				delta := P{locs[b].X - locs[a].X, locs[b].Y - locs[a].Y}

				anti1 := P{locs[a].X - delta.X, locs[a].Y - delta.Y}
				if anti1.X >= 0 && anti1.X < Size && anti1.Y >= 0 && anti1.Y < Size {
					antiNodes[anti1] = struct{}{}
				}

				anti2 := P{locs[b].X + delta.X, locs[b].Y + delta.Y}
				if anti2.X >= 0 && anti2.X < Size && anti2.Y >= 0 && anti2.Y < Size {
					antiNodes[anti2] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(antiNodes))
}
