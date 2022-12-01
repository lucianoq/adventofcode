package main

import "fmt"

func main() {
	elves := parse()

	max := 0
	for _, cal := range elves {
		if cal > max {
			max = cal
		}
	}

	fmt.Println(max)
}
