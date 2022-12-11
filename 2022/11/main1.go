package main

import "fmt"

func main() {
	monkeys := parse()

	manage := func(worryLevel int) int {
		return worryLevel / 3
	}

	fmt.Println(monkeyBusiness(monkeys, 20, manage))
}
