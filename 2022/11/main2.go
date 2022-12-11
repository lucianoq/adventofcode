package main

import "fmt"

func main() {
	monkeys := parse()

	lcm := 1
	for _, monkey := range monkeys {
		lcm *= monkey.Div
	}

	manage := func(worryLevel int) int {
		return worryLevel % lcm
	}

	fmt.Println(monkeyBusiness(monkeys, 10000, manage))
}
