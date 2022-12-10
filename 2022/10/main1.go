package main

import "fmt"

func main() {
	signal := 0

	updateSignal := func(cycle, x int) {
		if cycle%40 == 20 && cycle <= 220 {
			signal += x * cycle
		}
	}

	run(updateSignal)

	fmt.Println(signal)
}
