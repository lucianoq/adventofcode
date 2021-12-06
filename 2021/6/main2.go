package main

import "fmt"

func main() {
	timers := parse()

	for i := 0; i < 256; i++ {
		timers = evolve(timers)
	}

	sum := 0
	for _, v := range timers {
		sum += v
	}
	fmt.Println(sum)
}
