package main

import "fmt"

func main() {
	parse()
	fmt.Println(countRest())
}

func empty(p P) bool {
	_, ok := Map[p]
	return !ok
}
