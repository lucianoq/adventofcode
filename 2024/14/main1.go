package main

import "fmt"

func main() {
	robots := parseInput()

	for t := 0; t < 100; t++ {
		for i := 0; i < len(robots); i++ {
			robots[i].Pos.x = (robots[i].Pos.x + robots[i].Vel.x + max.x) % max.x
			robots[i].Pos.y = (robots[i].Pos.y + robots[i].Vel.y + max.y) % max.y
		}
	}

	q := [4]int{}
	for _, r := range robots {
		if r.Pos.x < max.x/2 {
			if r.Pos.y < max.y/2 {
				q[0]++
			}
			if r.Pos.y > max.y/2 {
				q[1]++
			}
		}
		if r.Pos.x > max.x/2 {
			if r.Pos.y < max.y/2 {
				q[2]++
			}
			if r.Pos.y > max.y/2 {
				q[3]++
			}
		}
	}

	fmt.Println(q[0] * q[1] * q[2] * q[3])
}
