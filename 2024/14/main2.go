package main

import "fmt"

func main() {
	robots := parseInput()

	for t := 0; ; t++ {
		for i := 0; i < len(robots); i++ {
			robots[i].Pos.x = (robots[i].Pos.x + robots[i].Vel.x + max.x) % max.x
			robots[i].Pos.y = (robots[i].Pos.y + robots[i].Vel.y + max.y) % max.y
		}

		// robots gather vertically at t=3
		// robots gather horizontally at t=75
		// period is known

		if ((t+max.x-3)%max.x == 0) && (t+max.y-75)%max.y == 0 {
			Print(robots)
			fmt.Println(t + 1)
			return
		}
	}
}

func Print(robots []Robot) {
	m := map[P]struct{}{}
	for _, r := range robots {
		m[r.Pos] = struct{}{}
	}
	for y := 50; y < max.y-16; y++ {
		for x := 25; x < max.x-39; x++ {
			if _, ok := m[P{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
