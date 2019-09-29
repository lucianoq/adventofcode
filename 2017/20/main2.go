package main

import "fmt"

func main() {
	particles := parse()

	pp := make(map[int]Particle)
	for _, p := range particles {
		pp[p.ID] = p
	}

	oldSize, count := len(pp), 0
	for t := 0; count < 100; t++ {
		space := make(map[Triple][]int)

		for k, v := range pp {
			pos := v.PositionAfter(t)
			space[pos] = append(space[pos], k)
		}

		for _, list := range space {
			if len(list) > 1 {
				for _, id := range list {
					delete(pp, id)
				}
			}
		}

		if len(pp) == oldSize {
			count++
		} else {
			oldSize = len(pp)
			count = 0
		}
	}

	fmt.Println(oldSize)
}
