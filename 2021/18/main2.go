package main

import "fmt"

func main() {
	numbers := parse()

	max := 0
	for i := range numbers {
		for j := range numbers {
			if i == j {
				continue
			}

			res := Add(numbers[i].copy(), numbers[j].copy())
			magnitude := res.Magnitude()
			if magnitude > max {
				max = magnitude
			}
		}
	}
	fmt.Println(max)
}
