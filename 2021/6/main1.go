package main

import "fmt"

func main() {
	timers := parse()

	for i := 0; i < 80; i++ {
		evolve(timers)
	}

	sum := 0
	for _, v := range timers {
		sum += v
	}
	fmt.Println(sum)
}
