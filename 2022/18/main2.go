package main

import "fmt"

func main() {
	lava := parse()
	min, max := minMax(lava)

	// Enlarge the space by 1 in order to avoid
	// to be blocked by the borders
	min = min.Add(P{-1, -1, -1})
	max = max.Add(P{1, 1, 1})

	outerSpace := outerSpace(lava, min, max)

	sides := 0
	for p := range outerSpace {
		for _, adj := range Adj {
			neighbour := p.Add(adj)
			if _, ok := lava[neighbour]; ok {
				sides++
			}
		}
	}

	fmt.Println(sides)
}

func minMax(space map[P]struct{}) (P, P) {
	min := P{1 << 60, 1 << 60, 1 << 60}
	max := P{0, 0, 0}
	for p := range space {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Z < min.Z {
			min.Z = p.Z
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
		if p.Z > max.Z {
			max.Z = p.Z
		}
	}
	return min, max
}

// BFS
func outerSpace(lava map[P]struct{}, min, max P) map[P]struct{} {
	start := min
	toDo := []P{start}
	visited := map[P]struct{}{start: {}}

	var curr P
	for len(toDo) > 0 {
		curr, toDo = toDo[0], toDo[1:]

		for _, adj := range Adj {
			next := curr.Add(adj)

			// Skip if visited
			if _, ok := visited[next]; ok {
				continue
			}

			// Skip if out of borders
			if next.X < min.X || next.Y < min.Y || next.Z < min.Z ||
				next.X > max.X || next.Y > max.Y || next.Z > max.Z {
				continue
			}

			// Skip if lava
			if _, ok := lava[next]; ok {
				continue
			}

			visited[next] = struct{}{}
			toDo = append(toDo, next)
		}
	}

	return visited
}
