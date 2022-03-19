package main

import (
	"fmt"

	"github.com/lucianoq/container/set"
)

func main() {
	grid := parse()

	for step := 1; ; step++ {

		// First, the energy level of each octopus increases by 1.
		for i := 0; i < Size; i++ {
			for j := 0; j < Size; j++ {
				grid[i][j]++
			}
		}

		// Then, any octopus with an energy level greater than 9 flashes.

		flashed := set.New[pair]()

		for {
			changed := false
			for i := 0; i < Size; i++ {
				for j := 0; j < Size; j++ {
					if grid[i][j] > 9 {

						if !flashed.Contains(pair{i, j}) {
							flashed.Add(pair{i, j})
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

		if flashed.Len() == Size*Size {
			fmt.Println(step)
			return
		}

		// Finally, any octopus that flashed during this step has its energy
		// level set to 0, as it used all of its energy to flash.
		for f := range flashed {
			grid[f.i][f.j] = 0
		}
	}
}
