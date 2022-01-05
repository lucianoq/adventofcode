package main

import "fmt"

func main() {
	numbers := parse()

	acc := numbers[0]

	for i := 1; i < len(numbers); i++ {
		acc = Add(acc, numbers[i])
	}

	fmt.Println(acc.Magnitude())
}
