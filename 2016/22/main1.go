package main

import (
	"fmt"
)

func main() {
	devs := parse()

	count := 0
	for _, d1 := range devs {
		for _, d2 := range devs {
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
