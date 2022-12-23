package main

import "fmt"

func main() {
	parse()

	for i := 0; i < 10; i++ {
		run()
	}

	min, max := minMax()

	count := 0
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			if _, ok := Map[P{x, y}]; !ok {
				count++
			}
		}
	}

	fmt.Println(count)
}
