package main

import "fmt"

func main() {
	db := parse()

	var sum int
	for game, exs := range db {
		if possible(exs) {
			sum += game
		}
	}

	fmt.Println(sum)
}

func possible(exs []Extraction) bool {
	for _, ex := range exs {
		if ex.Blue > 14 || ex.Red > 12 || ex.Green > 13 {
			return false
		}
	}
	return true
}
