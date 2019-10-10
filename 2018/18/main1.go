package main

import "fmt"

const minutes = 10

func main() {
	area := parse()
	// Print(area)

	for i := 0; i < minutes; i++ {
		area = Minute(area)
	}

	fmt.Println(CountResources(area))
}
