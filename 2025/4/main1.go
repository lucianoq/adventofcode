package main

import "fmt"

func main() {
	m := parseMap()
	count := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if _, ok := m[P{i, j}]; ok && adjacentRolls(m, i, j) < 4 {
				count++
			}
		}
	}
	fmt.Println(count)
}
