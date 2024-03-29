package main

type Status struct {
	Point P
	Steps int
}

var Dirs = [4]P{
	{0, 1},  // Right
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
}

func bfs(start P, goal func(p P) bool, forbidden func(from, to P) bool) int {
	todo := []Status{{start, 0}}
	visited := map[P]struct{}{start: {}}

	var curr Status

	for len(todo) > 0 {
		curr, todo = todo[0], todo[1:]

		if goal(curr.Point) {
			return curr.Steps
		}

		for _, dir := range Dirs {
			neighbour, ok := curr.Point.Add(dir)

			// out of bound
			if !ok {
				continue
			}

			if forbidden(curr.Point, neighbour) {
				continue
			}

			// already visited
			if _, seen := visited[neighbour]; seen {
				continue
			}

			visited[neighbour] = struct{}{}
			todo = append(todo, Status{neighbour, curr.Steps + 1})
		}
	}
	return 1<<63 - 1
}
