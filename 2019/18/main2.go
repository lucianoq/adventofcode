package main

import "fmt"

func main() {
	vault := parse()
	fmt.Println(BFS(vault))
}

func BFS(vault *Vault) int {

	const numBots = 4

	graph := toGraph(vault)

	type Status struct {
		IDs    [numBots]char
		Keybag int32
	}

	discovered := map[Status]int{}
	toDo := []Status{} // let Q be a queue

	root := Status{
		IDs:    [numBots]char{'0', '1', '2', '3'},
		Keybag: 0,
	}
	discovered[root] = 0      // label root as discovered
	toDo = append(toDo, root) // Q.enqueue(root)

	bestSolutionSoFar := 1<<63 - 1

	for len(toDo) > 0 {
		var curr Status

		curr, toDo = toDo[0], toDo[1:]

		// if all keybag bits are turned on
		if curr.Keybag&(1<<vault.NumKeys-1) == 1<<vault.NumKeys-1 {
			if dist := discovered[curr]; dist < bestSolutionSoFar {
				bestSolutionSoFar = dist
			}
		}

		for bot := 0; bot < numBots; bot++ {

			for cell, weight := range graph.GetNext(curr.IDs[bot]) {

				// if it is a door
				if cell >= 'A' && cell <= 'Z' {
					// skip if we have no key for it
					if curr.Keybag&(1<<(cell-'A')) == 0 {
						continue
					}
				}

				keyBag := curr.Keybag
				// if it is a key
				if cell >= 'a' && cell <= 'z' {
					// add it to the bitset
					keyBag |= 1 << (cell - 'a')
				}

				newDistance := discovered[curr] + weight
				// optimization to cut off worse solutions
				if newDistance > bestSolutionSoFar {
					continue
				}

				newIDs := curr.IDs
				newIDs[bot] = cell

				nextStatus := Status{newIDs, keyBag}

				if oldDist, found := discovered[nextStatus]; !found || newDistance < oldDist {
					discovered[nextStatus] = newDistance
					toDo = append(toDo, nextStatus)
				}
			}
		}
	}

	return bestSolutionSoFar
}
