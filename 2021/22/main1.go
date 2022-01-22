package main

import "fmt"

type P struct{ x, y, z int }

func (c Cuboid) contains(p P) bool {
	return c.minX <= p.x && p.x <= c.maxX &&
		c.minY <= p.y && p.y <= c.maxY &&
		c.minZ <= p.z && p.z <= c.maxZ
}

func main() {
	rules := parse()

	space := map[P]bool{}

	for _, rule := range rules {
		for x := -50; x <= 50; x++ {
			for y := -50; y <= 50; y++ {
				for z := -50; z <= 50; z++ {
					p := P{x, y, z}
					if rule.contains(p) {
						space[p] = rule.state
					}
				}
			}
		}
	}

	count := 0
	for _, c := range space {
		if c {
			count++
		}
	}
	fmt.Println(count)
}
