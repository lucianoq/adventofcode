package main

import "fmt"

func main() {
	list := parse()

	min, max := minMax(list)

	minFuel := 1<<63 - 1

	for i := min; i <= max; i++ {
		fuel := 0
		for _, x := range list {
			fuel += abs(i - x)
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)
}
