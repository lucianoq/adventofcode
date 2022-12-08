package main

import "fmt"

func main() {
	v := visibleFromEdge()

	count := 0
	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if v[P{r, c}] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func visibleFromEdge() map[P]bool {
	v := map[P]bool{}
	for i := 0; i < Size; i++ {
		walk(v, P{i, 0}, Right)       // left to right
		walk(v, P{i, Size - 1}, Left) // right to left
		walk(v, P{0, i}, Down)        // top down
		walk(v, P{Size - 1, i}, Up)   // bottom up
	}
	return v
}

func walk(visible map[P]bool, start, dir P) {
	visible[start] = true
	tallest := Map[start]

	for curr, valid := start, true; valid; curr, valid = curr.Add(dir) {
		if Map[curr] > tallest {
			visible[curr] = true
			tallest = Map[curr]
		}
		// Optimization: nothing else can be seen after a 9
		if Map[curr] == 9 {
			break
		}
	}
}
