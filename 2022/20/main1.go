package main

import "fmt"

func main() {
	index, zero := parse()
	mix(index)
	fmt.Println(groveCoordinates(zero))
}
