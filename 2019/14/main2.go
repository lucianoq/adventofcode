package main

import "fmt"

const ore = 1000000000000

func main() {
	reactions := parse()

	left, right := 0, 10000000

	for right-left > 1 {

		mid := (left + right) / 2
		if costPerFuel(reactions, mid) <= ore {
			left = mid
		} else {
			right = mid
		}
	}

	fmt.Println(left)
}
