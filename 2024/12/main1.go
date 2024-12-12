package main

// BFS
func explore(m map[P]uint8, p P, visited map[P]struct{}) (int, int) {

	visited[p] = struct{}{}

	perimeter, area := 0, 0

	var c P
	toDo := []P{p}
	for len(toDo) > 0 {
		c, toDo = toDo[0], toDo[1:]

		area++

		for _, d := range []P{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := P{c.x + d.x, c.y + d.y}
			val, ok := m[next]
			if ok && val == m[p] {
				if _, ok := visited[next]; !ok {
					toDo = append(toDo, next)
					visited[next] = struct{}{}
				}
			} else {
				perimeter++
			}
		}

	}

	return area, perimeter
}
