package main

import "fmt"

func main() {
	m := parseMap()
	count := 0
	for {
		toDelete := make([]P, 0)
		for i := 0; i < Size; i++ {
			for j := 0; j < Size; j++ {
				if _, ok := m[P{i, j}]; ok && adjacentRolls(m, i, j) < 4 {
					toDelete = append(toDelete, P{i, j})
				}
			}
		}

		if len(toDelete) == 0 {
			break
		}

		for _, p := range toDelete {
			delete(m, P{p.x, p.y})
			count++
		}
	}

	fmt.Println(count)
}
