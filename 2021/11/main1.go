package main

import "fmt"

func main() {
	grid := parse()

	flashes := 0

	for step := 0; step < 100; step++ {

		// First, the energy level of each octopus increases by 1.
		for i := 0; i < Size; i++ {
			for j := 0; j < Size; j++ {
				grid[i][j]++
			}
		}

		// Then, any octopus with an energy level greater than 9 flashes.
		flashed := map[pair]struct{}{}

		for {
			changed := false
			for i := 0; i < Size; i++ {
				for j := 0; j < Size; j++ {

					if grid[i][j] > 9 {

						if _, ok := flashed[pair{i, j}]; !ok {

							flashed[pair{i, j}] = struct{}{}
							changed = true

							// This increases the energy level of all adjacent octopuses by 1,
							// including octopuses that are diagonally adjacent.
							for _, di := range []int{-1, 0, +1} {
								for _, dj := range []int{-1, 0, +1} {
									if di|dj != 0 {
										if i+di >= 0 && i+di < Size {
											if j+dj >= 0 && j+dj < Size {
												grid[i+di][j+dj]++
											}
										}
									}
								}
							}
						}
					}
				}
			}
			if !changed {
				break
			}
		}

		flashes += len(flashed)

		// Finally, any octopus that flashed during this step has its energy
		// level set to 0, as it used all of its energy to flash.
		for f, _ := range flashed {
			grid[f.i][f.j] = 0
		}
	}

	fmt.Println(flashes)
}
