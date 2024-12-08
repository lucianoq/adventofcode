package main

import "fmt"

func main() {
	freqs := parseInput()

	antiNodes := map[P]struct{}{}
	for _, locs := range freqs {
		for a := 0; a < len(locs)-1; a++ {
			for b := a + 1; b < len(locs); b++ {
				delta := P{locs[b].X - locs[a].X, locs[b].Y - locs[a].Y}

				outOfBound := 0
				for period := 0; outOfBound < 2; period++ {
					outOfBound = 0

					anti1 := P{locs[a].X - period*delta.X, locs[a].Y - period*delta.Y}
					if anti1.Valid() {
						antiNodes[anti1] = struct{}{}
					} else {
						outOfBound++
					}

					anti2 := P{locs[b].X + period*delta.X, locs[b].Y + period*delta.Y}
					if anti2.Valid() {
						antiNodes[anti2] = struct{}{}
					} else {
						outOfBound++
					}
				}
			}
		}
	}

	fmt.Println(len(antiNodes))
}
