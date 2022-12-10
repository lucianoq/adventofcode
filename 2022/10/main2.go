package main

import "fmt"

func main() {
	crt := [6][40]bool{}

	updateCRT := func(cycle, x int) {
		crtCol := (cycle - 1) % 40
		ctrRow := (cycle - 1) / 40
		if x-1 <= crtCol && crtCol <= x+1 {
			crt[ctrRow][crtCol] = true
		}
	}

	run(updateCRT)

	for r := 0; r < 6; r++ {
		for c := 0; c < 40; c++ {
			if crt[r][c] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
