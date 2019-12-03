package main

import "fmt"

func main() {
	fmt.Println(run(manhattan))
}

func manhattan(c C, _, _ int) int {
	return abs(c.X) + abs(c.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
