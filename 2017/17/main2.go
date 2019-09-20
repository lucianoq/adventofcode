package main

import "fmt"

const stepsForward = 303

func main() {
	currentPos := 0
	res := -1

	for i := 1; i <= 50000000; i++ {
		currentPos = (currentPos + stepsForward) % i
		if currentPos == 0 {
			res = i
		}
		currentPos++
	}
	fmt.Println(res)
}
