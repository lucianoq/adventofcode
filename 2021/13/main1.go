package main

import "fmt"

func main() {
	grid, folds := parse()

	fold(grid, folds[0])

	fmt.Println(countDots(grid))
}
