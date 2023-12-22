package main

import "fmt"

func main() {
	space := parse()
	_ = space.Fall()

	count := 0
	for _, b := range space.Bricks {
		if space.SafeToDisintegrate(b) {
			count++
		}
	}
	fmt.Println(count)
}
