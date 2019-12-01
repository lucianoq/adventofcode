package main

import "fmt"

func main() {
	list := parse()
	mass := massForModules(list)
	fmt.Println(mass)
}

func massForModules(list []int) int {
	for i := range list {
		list[i] = list[i]/3 - 2
	}

	sum := 0
	for _, i := range list {
		sum += i
	}
	return sum
}
