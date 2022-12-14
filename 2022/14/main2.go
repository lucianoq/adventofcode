package main

import "fmt"

func main() {
	parse()
	Floor += 2
	fmt.Println(countRest())
}

func empty(p P) bool {
	if p.X == Floor {
		return false
	}

	_, ok := Map[p]
	return !ok
}
