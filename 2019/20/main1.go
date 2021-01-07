package main

import "fmt"

func main() {
	m := parse()

	fmt.Println(BFS(m))
}

func BFS(m *Map) int {
	discovered := map[P]struct{}{}
	toDo := []P{}

	discovered[m.AA] = struct{}{}
	toDo = append(toDo, m.AA)

	depth := 0

	for len(toDo) > 0 {
		for levelSize := len(toDo); levelSize > 0; levelSize-- {
			var curr P
			curr, toDo = toDo[0], toDo[1:]

			if curr == m.ZZ {
				return depth
			}

			for _, n := range curr.Neighbours() {
				dest := m.Grid[n]

				switch {

				case dest == Wall:
					continue

				case dest == Free:

					if _, found := discovered[n]; !found {
						discovered[n] = struct{}{}
						toDo = append(toDo, n)
					}

				case dest.Letter():
					next := m.Teleport[curr]
					if _, found := discovered[next]; !found {
						discovered[next] = struct{}{}
						toDo = append(toDo, next)
					}
				}
			}
		}
		depth++
	}

	return -1
}
