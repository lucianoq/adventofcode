package main

import "fmt"

func main() {
	splitters, start, last := parseInput()
	beams := map[P]struct{}{P{start.x + 1, start.y}: {}}
	splits := 0
	for i := 2; i < last; i++ {
		nextBeams := map[P]struct{}{}
		for b := range beams {
			if _, ok := splitters[P{i, b.y}]; ok {
				splits++
				nextBeams[P{i, b.y - 1}] = struct{}{}
				nextBeams[P{i, b.y + 1}] = struct{}{}
			} else {
				nextBeams[P{i, b.y}] = struct{}{}
			}
		}
		beams = nextBeams
	}
	fmt.Println(splits)
}
