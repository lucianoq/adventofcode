package main

import "fmt"

func main() {
	m := input()
	count := 0
	for x := 0; x < Size; x++ {
		for y := 0; y < Size; y++ {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx != 0 || dy != 0 {
						if findXMas(m, P{x, y}, P{dx, dy}) {
							count++
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}

func findXMas(m map[P]byte, start, delta P) bool {
	return m[start] == 'X' &&
		m[start.Add(delta, 1)] == 'M' &&
		m[start.Add(delta, 2)] == 'A' &&
		m[start.Add(delta, 3)] == 'S'
}
