package main

import "fmt"

func main() {
	grid := parseGrid()

	count := 0
	for _, d1 := range grid {
		for _, d2 := range grid {
			if viable(d1, d2) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func viable(d1, d2 Device) bool {
	return d1 != d2 && d1.Used > 0 && d1.Used <= d2.Avail
}
