package main

import "fmt"

func main() {
	m := parse()

	fmt.Println(BFSNested(m))
}

func BFSNested(m *Map) int {
	type Status struct {
		P     P
		Depth int
	}

	discovered := map[Status]struct{}{}
	toDo := []Status{}

	root := Status{m.AA, 0}

	discovered[root] = struct{}{}
	toDo = append(toDo, root)

	steps := 0

	for len(toDo) > 0 {

		for levelSize := len(toDo); levelSize > 0; levelSize-- {

			var curr Status

			curr, toDo = toDo[0], toDo[1:]

			for _, n := range curr.P.Neighbours() {
				dest := m.Grid[n]

				switch {

				case dest == Wall:
					continue

				case dest == Free:

					target := Status{n, curr.Depth}

					if _, found := discovered[target]; !found {
						discovered[target] = struct{}{}
						toDo = append(toDo, target)
					}

				case dest.Letter():

					var target Status

					isOuter := m.IsOuter[curr.P]

					if !isOuter {
						// this is a inner portal
						// go deep

						target = Status{
							m.Teleport[curr.P],
							curr.Depth + 1,
						}

					} else {

						portalName := m.PortalName[curr.P]

						if curr.Depth == 0 {

							// our goal
							if portalName == "ZZ" {
								return steps
							}

							// any other outer portal is blocked at level 0
							continue

						}

						// curr.Dept > 0
						// that means for all inner levels

						// invalidate AA and ZZ portals
						if portalName == "AA" || portalName == "ZZ" {
							continue
						}

						// return to parent level
						target = Status{
							m.Teleport[curr.P],
							curr.Depth - 1,
						}
					}

					if _, found := discovered[target]; !found {
						discovered[target] = struct{}{}
						toDo = append(toDo, target)
					}
				}
			}
		}

		steps++
	}

	return -1
}
