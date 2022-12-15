package main

import "fmt"

func main() {
	sensors := parse()

	for y := 0; y <= 4000000; y++ {
	Point:
		for x := 0; x <= 4000000; x++ {
			p := P{x, y}

			for _, s := range sensors {
				if manhattan(p, s.Pos) < s.Dist {

					// No reason to check all points inside the diamond
					// When we find ourselves inside one: we can skip many Xs
					// in order to reach the closest border
					//
					// we skip an amount of Xs that is calculated as follows:
					// - take the full distance between the sensors and its
					//   closest beacon
					// - remove the y component of the distance between p and s
					// - if px < sx ==>    add the x component
					// - if px > sx ==> remove the x component
					//
					// last two points can be condensed in adding the difference
					// (positive in the first case, negative in the second)
					x += s.Dist - abs(s.Pos.Y-p.Y) + s.Pos.X - p.X
					continue Point
				}
			}

			// no sensor in range
			fmt.Println(4000000*x + y)
			return
		}
	}
}
