package main

import "fmt"

func main() {
	g := makeGridEL(0, 0)
	// Print(g, 0, 0)
	fmt.Println(RiskLevel(g))
}

func RiskLevel(g map[C]int) int {
	var sum int
	for y := 0; y <= TargetRow; y++ {
		for x := 0; x <= TargetCol; x++ {
			p := C{x, y}
			sum += g[p] % 3
		}
	}
	return sum
}
