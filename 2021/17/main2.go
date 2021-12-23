package main

import "fmt"

func main() {
	xMin, xMax, yMin, yMax := parse()

	count := 0

	// if xv > xMax,
	// x would overshoot at the first step
	// and will never recover because xv->0 so x never decrease
	for xv := 0; xv <= xMax; xv++ {
	V:
		// if yv < yMin,
		// y would overshoot at the first step
		// and will never recover because yv > -inf so y always decrease
		for yv := yMin; yv < 500; yv++ {

			xv, yv := xv, yv

			x, y := 0, 0
			for step := 0; step < 500; step++ {
				x += xv
				y += yv

				if x > xMax {
					continue V
				}

				if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
					count++
					continue V
				}

				// new xv
				switch {
				case xv > 0:
					xv--
				case xv < 0:
					xv++
				}

				// new yv
				yv--
			}
		}
	}

	fmt.Println(count)
}
