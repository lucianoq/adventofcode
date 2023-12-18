package main

import "fmt"

func main() {
	curr := P{0, 0}

	borders := map[P]int{curr: 1}

	for _, op := range parse() {
		for i := 0; i < op.Steps; i++ {
			curr.R += Delta[op.Dir].R
			curr.C += Delta[op.Dir].C
			borders[curr] += 1
		}
	}

	fmt.Println(len(borders) + flood(borders, P{0, 1}))
}

func flood(borders map[P]int, p P) int {
	visited := map[P]struct{}{}
	toDo := []P{p}
	var curr P
	for len(toDo) > 0 {
		curr, toDo = toDo[0], toDo[1:]

		for _, d := range Delta {
			next := P{curr.R + d.R, curr.C + d.C}

			if borders[next] > 0 {
				continue
			}

			if _, ok := visited[next]; ok {
				continue
			}

			visited[next] = struct{}{}
			toDo = append(toDo, next)
		}
	}

	return len(visited)
}
