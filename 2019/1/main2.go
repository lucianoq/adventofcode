package main

import "fmt"

func main() {
	list := parse()
	mass := massAndFuel(list)
	fmt.Println(mass)
}

func massAndFuel(list []int) int {
	for i := range list {
		list[i] = list[i]/3 - 2

		for toAdd := list[i]/3 - 2; toAdd > 0; toAdd = toAdd/3 - 2 {
			list[i] += toAdd
		}
	}

	sum := 0
	for _, i := range list {
		sum += i
	}
	return sum
}
