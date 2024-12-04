package main

import "fmt"

func main() {
	m := input()

	count := 0
	for x := 0; x < Size; x++ {
		for y := 0; y < Size; y++ {
			if findCrossMas(m, P{x, y}) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func findCrossMas(m map[P]byte, start P) bool {
	if m[start] != 'A' {
		return false
	}
	s := string([]byte{
		m[P{start.x - 1, start.y - 1}],
		m[P{start.x - 1, start.y + 1}],
		m[P{start.x + 1, start.y + 1}],
		m[P{start.x + 1, start.y - 1}],
	})
	return s == "MMSS" || s == "MSSM" || s == "SSMM" || s == "SMMS"
}
